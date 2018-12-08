/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2018 Sergey Denisov.
/*
/* Written by Sergey Denisov aka LittleBuster (DenisovS21@gmail.com)
/* Github: https://github.com/LittleBuster
/*	       https://github.com/futcamp
/*
/* This library is free software; you can redistribute it and/or
/* modify it under the terms of the GNU General Public Licence 3
/* as published by the Free Software Foundation; either version 3
/* of the Licence, or (at your option) any later version.
/*
/*******************************************************************/

package main

import (
	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/net"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

type Application struct {
	Log         *utils.Logger
	LogTask     *utils.LogTask
	Cfg         *configs.Configs
	MeteoCfg    *configs.MeteoConfigs
	Meteo       *meteo.MeteoStation
	Server      *net.WebServer
	MeteoTask   *meteo.MeteoTask
	Locker      *utils.Locker
	MeteoDB     *meteo.MeteoDatabase
	Monitor     *monitoring.DeviceMonitor
	MonitorTask *monitoring.MonitorTask
}

// NewApplication make new struct
func NewApplication(log *utils.Logger, cfg *configs.Configs, mCfg *configs.MeteoConfigs,
	meteo *meteo.MeteoStation, srv *net.WebServer, mTask *meteo.MeteoTask,
	lTask *utils.LogTask, lck *utils.Locker, mdb *meteo.MeteoDatabase,
	monitor *monitoring.DeviceMonitor, monitorTask *monitoring.MonitorTask) *Application {
	return &Application{
		Log:         log,
		Cfg:         cfg,
		MeteoCfg:    mCfg,
		Meteo:       meteo,
		Server:      srv,
		MeteoTask:   mTask,
		LogTask:     lTask,
		Locker:      lck,
		MeteoDB:     mdb,
		Monitor:     monitor,
		MonitorTask: monitorTask,
	}
}

// Start run init functions of all modules
func (a *Application) Start() {
	a.Log.Init(utils.LogPath)

	// Load app configs
	err := a.Cfg.LoadFromFile(utils.ConfigsPath)
	if err != nil {
		logger.Errorf("Fail to load %s configs", "main")
		return
	}
	logger.Infof("Configs %s was loaded", "main")

	if a.Cfg.Settings().Modules.Meteo {
		// Load meteo configs
		err = a.MeteoCfg.LoadFromFile(utils.MeteoConfigsPath)
		if err != nil {
			logger.Infof("Fail to load %s configs", "meteo")
			return
		}
		logger.Infof("Configs %s was loaded", "meteo")

		// Add meteo sensors
		for _, sensor := range a.MeteoCfg.Settings().Sensors {
			a.Meteo.AddSensor(sensor.Name, sensor.Type, sensor.IP, sensor.Channel)
			a.Monitor.AddDevice(sensor.Name, "Sensor", sensor.IP)
			logger.Infof("New sensor %s type %s IP %s channel %d",
				sensor.Name, sensor.Type, sensor.IP, sensor.Channel)
		}

		// Add monitoring for displays
		for _, display := range a.MeteoCfg.Settings().Displays {
			if display.Enable {
				a.Monitor.AddDevice(display.Name, "Display", display.IP)
			}
		}

		// Set path to db
		a.MeteoDB.SetDBFile(utils.MeteoDBPath)

		// Add db lock
		a.Locker.AddLock(utils.MeteoDBName)

		// Start task
		go a.MeteoTask.Start()
	}

	// Start logger task
	go a.LogTask.Start()

	// Start monitoring
	go a.MonitorTask.Start()

	// Start web server
	logger.Infof("Starting Web server at %s:%d...", a.Cfg.Settings().Server.IP,
		a.Cfg.Settings().Server.Port)
	err = a.Server.Start(a.Cfg.Settings().Server.IP, a.Cfg.Settings().Server.Port)
	if err != nil {
		logger.Error("Fail to start WebServer")
	}

	logger.Info("Application was finished")
}

// Free unload all modules from memory
func (a *Application) Free() {
	a.Log.Free()
}
