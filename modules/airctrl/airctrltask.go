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

package airctrl

import (
	"time"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/google/logger"
)

const (
	taskDelay = 1
)

// AirCtrlTask air control task struct
type AirCtrlTask struct {
	reqTimer *time.Timer
	airCtrl  *AirControl
	meteo    *meteo.MeteoStation
}

// NewAirCtrlTask make new struct
func NewAirCtrlTask(ac *AirControl, meteo *meteo.MeteoStation) *AirCtrlTask {
	return &AirCtrlTask{
		airCtrl: ac,
		meteo: meteo,
	}
}

// TaskHandler process timer loop
func (a *AirCtrlTask) TaskHandler() {
	for {
		<-a.reqTimer.C

		modules := a.airCtrl.GetModules()

		// Syncing relay state
		for _, module := range modules {
			err := module.SyncModule()
			if err != nil {
				continue
			} else {
				module.ClearError()
			}
		}

		// If control is on - get data from humidity sensor and control air humidity
		if a.airCtrl.GetHumidityControl() {
			for _, module := range modules {
				// Get humidity sensor pointer for humidity control
				sensor, err := a.meteo.Sensor(module.Sensor)
				if err != nil {
					logger.Error(err.Error())
					continue
				}

				// Get current humidity from meteo sensor and current humidity module state
				humidity := sensor.MeteoData().Humidity
				state := module.GetRelayState()

				if humidity < module.Threshold && !state {
					// If switched off send relay on command
					err := module.SwitchRelay(true)
					if err != nil {
						logger.Error(err.Error())
						module.SetError()
					}
				} else if humidity >= module.Threshold && state {
					// If switched on send relay off command
					err := module.SwitchRelay(false)
					if err != nil {
						logger.Error(err.Error())
						module.SetError()
					}
				}
			}
		} else {
			for _, module := range modules {
				if module.GetRelayState() {
					// Switch off all air humidity control modules if control is off
					err := module.SwitchRelay(false)
					if err != nil {
						logger.Error(err.Error())
					}
				}
			}
		}

		a.reqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (a *AirCtrlTask) Start() {
	a.reqTimer = time.NewTimer(taskDelay * time.Second)
	a.TaskHandler()
}
