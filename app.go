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
	"github.com/futcamp/controller/meteo"
	"github.com/futcamp/controller/net"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"
	"github.com/google/logger"
)

type Application struct {
	Log      *utils.Logger
	Cfg      *configs.Configs
	MeteoCfg *configs.MeteoConfigs
	Meteo    *meteo.MeteoStation
	Server   *net.WebServer
}

// NewApplication make new struct
func NewApplication(log *utils.Logger, cfg *configs.Configs, mCfg *configs.MeteoConfigs,
	meteo *meteo.MeteoStation, srv *net.WebServer) *Application {
	return &Application{
		Log:      log,
		Cfg:      cfg,
		MeteoCfg: mCfg,
		Meteo:    meteo,
		Server:   srv,
	}
}

// Start run init functions of all modules
func (a *Application) Start() {
	a.Log.Init(utils.LogPath)

	// Load app configs
	err := a.Cfg.LoadFromFile(utils.ConfigsPath)
	if err != nil {
		logger.Errorf("Fail to load [%s] configs", "main")
		return
	}
	logger.Infof("Configs [%s] was loaded", "main")

	if a.Cfg.Settings().Modules.Meteo {
		// Load meteo configs
		err = a.MeteoCfg.LoadFromFile(utils.MeteoConfigsPath)
		if err != nil {
			logger.Infof("Fail to load [%s] configs", "meteo")
			return
		}
		logger.Infof("Configs [%s] was loaded", "meteo")

		// Add meteo sensors
		for _, sensor := range a.MeteoCfg.Settings().Sensors {
			a.Meteo.AddSensor(sensor.Name, sensor.Type, sensor.IP, sensor.Channel)
			logger.Infof("New sensor [%s] type [%s] IP [%s] channel [%d]",
				sensor.Name, sensor.Type, sensor.IP, sensor.Channel)
		}
	}

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
