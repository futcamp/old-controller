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
	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/net/rcli"
	"github.com/futcamp/controller/net/webserver"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"
	"github.com/futcamp/controller/utils/startup"

	"github.com/google/logger"
)

type Application struct {
	log         *utils.Logger
	logTask     *utils.LogTask
	cfg         *configs.Configs
	meteo       *meteo.MeteoStation
	server      *webserver.WebServer
	meteoTask   *meteo.MeteoTask
	locker      *utils.Locker
	meteoDB     *meteo.MeteoDatabase
	monitor     *monitoring.DeviceMonitor
	monitorTask *monitoring.MonitorTask
	startup     *startup.Startup
	rcli        *rcli.RCliServer
	dynCfg      *configs.DynamicConfigs
}

// NewApplication make new struct
func NewApplication(log *utils.Logger, cfg *configs.Configs,
	meteo *meteo.MeteoStation, srv *webserver.WebServer, mTask *meteo.MeteoTask,
	lTask *utils.LogTask, lck *utils.Locker, mdb *meteo.MeteoDatabase,
	monitor *monitoring.DeviceMonitor, monitorTask *monitoring.MonitorTask,
	stp *startup.Startup, rc *rcli.RCliServer, dc *configs.DynamicConfigs) *Application {
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
		startup:     stp,
		rcli:        rc,
		dynCfg:      dc,
	}
}

// Start run init functions of all modules
func (a *Application) Start() {
	a.log.Init(utils.LogPath)

	// Load general application configs
	err := a.cfg.LoadFromFile(utils.ConfigsPath)
	if err != nil {
		logger.Errorf("Fail to load %s configs", "main")
		logger.Error(err.Error())
		return
	}
	logger.Infof("Configs %s was loaded", "main")

	// Add lock for meteo database
	a.locker.AddLock(utils.MeteoDBName)

	// Load startup-configs from file and apply to application
	err = a.startup.Load(utils.StartupCfgPath)
	if err != nil {
		logger.Error("Fail to read startup configs")
		logger.Error(err.Error())
		return
	}
	logger.Info("Startup configs was loaded")

	// Start all application tasks
	go a.logTask.Start()
	go a.meteoTask.Start()
	go a.monitorTask.Start()

	// Start RemoteCLI server
	go func() {
		logger.Infof("Starting RemoteCLI server at %s:%d...", a.cfg.Settings().RCliServer.IP,
			a.cfg.Settings().RCliServer.Port)

		a.rcli.SetHash(a.dynCfg.Settings().RCli.UserHash)

		err = a.rcli.Start(a.cfg.Settings().RCliServer.IP, a.cfg.Settings().RCliServer.Port)
		if err != nil {
			logger.Error("Fail to start RemoteCLI server")
			logger.Error(err.Error())
			return
		}
	}()

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
