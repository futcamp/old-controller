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
	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/devices/db"
	"github.com/futcamp/controller/devices/tasks"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/net/rcli"
	"github.com/futcamp/controller/net/webserver"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"
	"github.com/futcamp/controller/utils/configs/cfgtask"
	"github.com/futcamp/controller/utils/startup"

	"github.com/google/logger"
)

type Application struct {
	log         *utils.Logger
	logTask     *utils.LogTask
	cfg         *configs.Configs
	meteo       *devices.MeteoStation
	server      *webserver.WebServer
	locker      *utils.Locker
	meteoDB     *db.MeteoDatabase
	monitor     *monitoring.DeviceMonitor
	monitorTask *monitoring.MonitorTask
	startup     *startup.Startup
	rcli        *rcli.RCliServer
	dynCfg      *configs.DynamicConfigs
	dynCfgTask  *cfgtask.DynConfigsTask
	tasks       *tasks.DeviceTasks
}

// NewApplication make new struct
func NewApplication(log *utils.Logger, cfg *configs.Configs,
	meteo *devices.MeteoStation, srv *webserver.WebServer,
	lTask *utils.LogTask, lck *utils.Locker, mdb *db.MeteoDatabase,
	monitor *monitoring.DeviceMonitor, monitorTask *monitoring.MonitorTask,
	stp *startup.Startup, rc *rcli.RCliServer, dc *configs.DynamicConfigs,
	dct *cfgtask.DynConfigsTask, tsks *tasks.DeviceTasks) *Application {
	return &Application{
		log:         log,
		cfg:         cfg,
		meteo:       meteo,
		server:      srv,
		logTask:     lTask,
		locker:      lck,
		meteoDB:     mdb,
		monitor:     monitor,
		monitorTask: monitorTask,
		startup:     stp,
		rcli:        rc,
		dynCfg:      dc,
		dynCfgTask:  dct,
		tasks:       tsks,
	}
}

// Start run init functions of all devices
func (a *Application) Start() {
	a.log.Init(utils.LogPath)

	// Load general application configs
	err := a.cfg.LoadFromFile(utils.ConfigsPath)
	if err != nil {
		logger.Errorf("Application fail to load %s configs", "main")
		return
	}
	logger.Infof("Application configs %s was loaded", "main")

	// Add lock for meteo database
	a.locker.AddLock(utils.MeteoDBName)

	// Load startup-configs from file and apply to application
	err = a.startup.Load(utils.StartupCfgPath)
	if err != nil {
		logger.Error("Application fail to read startup configs")
		return
	}
	logger.Info("Application startup configs was loaded")

	// Start all application tasks
	go a.logTask.Start()
	go a.monitorTask.Start()
	go a.dynCfgTask.Start()
	a.tasks.RunTasks()

	// Start RemoteCLI server
	go func() {
		logger.Infof("Application starting RemoteCLI server at %s:%d", a.cfg.Settings().RCliServer.IP,
			a.cfg.Settings().RCliServer.Port)

		err = a.rcli.Start(a.cfg.Settings().RCliServer.IP, a.cfg.Settings().RCliServer.Port)
		if err != nil {
			logger.Error("Application fail to start RemoteCLI server")
			return
		}
	}()

	// Start web server
	logger.Infof("Application starting Web server at %s:%d", a.cfg.Settings().Server.IP,
		a.cfg.Settings().Server.Port)

	err = a.server.Start(a.cfg.Settings().Server.IP, a.cfg.Settings().Server.Port)
	if err != nil {
		logger.Error("Application fail to start WebServer")
	}

	logger.Info("Application was finished")
}

// Free unload all devices from memory
func (a *Application) Free() {
	a.log.Free()
}
