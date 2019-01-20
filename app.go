/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2018-2019 Sergey Denisov.
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
	"github.com/futcamp/controller/modules/airctrl"
	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/net"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"
	"github.com/futcamp/controller/utils/startup/io"

	"github.com/google/logger"
)

type Application struct {
	log         *utils.Logger
	logTask     *utils.LogTask
	cfg         *configs.Configs
	meteo       *meteo.MeteoStation
	server      *net.WebServer
	meteoTask   *meteo.MeteoTask
	locker      *utils.Locker
	meteoDB     *meteo.MeteoDatabase
	monitor     *monitoring.DeviceMonitor
	monitorTask *monitoring.MonitorTask
	airCtrl     *airctrl.AirControl
	airTask     *airctrl.AirCtrlTask
	startupMods *io.StartupMods
}

// NewApplication make new struct
func NewApplication(log *utils.Logger, cfg *configs.Configs,
	meteo *meteo.MeteoStation, srv *net.WebServer, mTask *meteo.MeteoTask,
	lTask *utils.LogTask, lck *utils.Locker, mdb *meteo.MeteoDatabase,
	monitor *monitoring.DeviceMonitor, monitorTask *monitoring.MonitorTask,
	ac *airctrl.AirControl, acTask *airctrl.AirCtrlTask, sm *io.StartupMods) *Application {
	return &Application{
		log:         log,
		cfg:         cfg,
		meteo:       meteo,
		server:      srv,
		meteoTask:   mTask,
		logTask:     lTask,
		locker:      lck,
		meteoDB:     mdb,
		monitor:     monitor,
		monitorTask: monitorTask,
		airCtrl:     ac,
		airTask:     acTask,
		startupMods: sm,
	}
}

// Start run init functions of all modules
func (a *Application) Start() {
	a.log.Init(utils.LogPath)

	// Load app configs
	err := a.cfg.LoadFromFile(utils.ConfigsPath)
	if err != nil {
		logger.Errorf("Fail to load %s configs", "main")
		logger.Error(err.Error())
		return
	}
	logger.Infof("Configs %s was loaded", "main")

	err = a.startupMods.LoadFromFile(utils.StartupCfgPath)
	if err != nil {
		logger.Error("Fail to read startup configs")
		logger.Error(err.Error())
	}
	logger.Info("Startup configs was loaded")

	// Start all application tasks
	go a.logTask.Start()
	go a.meteoTask.Start()
	go a.airTask.Start()
	go a.monitorTask.Start()

	// Start web server
	logger.Infof("Starting Web server at %s:%d...", a.cfg.Settings().Server.IP,
		a.cfg.Settings().Server.Port)

	err = a.server.Start(a.cfg.Settings().Server.IP, a.cfg.Settings().Server.Port)
	if err != nil {
		logger.Error("Fail to start WebServer")
		logger.Error(err.Error())
	}

	logger.Info("Application was finished")
}

// Free unload all modules from memory
func (a *Application) Free() {
	a.log.Free()
}
