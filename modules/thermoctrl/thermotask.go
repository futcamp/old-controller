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

package thermoctrl

import (
	"time"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/google/logger"
)

const (
	taskDelay = 1
)

// ThermoCtrlTask air control task struct
type ThermoCtrlTask struct {
	reqTimer   *time.Timer
	thermoCtrl *ThermoControl
	meteo      *meteo.MeteoStation
}

// NewThermoCtrlTask make new struct
func NewThermoCtrlTask(tc *ThermoControl, meteo *meteo.MeteoStation) *ThermoCtrlTask {
	return &ThermoCtrlTask{
		thermoCtrl: tc,
		meteo:      meteo,
	}
}

// TaskHandler process timer loop
func (a *ThermoCtrlTask) TaskHandler() {
	for {
		<-a.reqTimer.C

		modules := a.thermoCtrl.Modules()

		// Syncing relay state
		for _, module := range modules {
			err := module.SyncModule()
			if err != nil {
				continue
			} else {
				module.ClearError()
			}
		}

		// If control is on - get data from meteo sensor and control temperature
		for _, module := range modules {
			if module.ThermoControl() {
				// Get humidity sensor pointer for humidity control
				sensor := a.meteo.Sensor(module.Sensor)

				// Get current humidity from meteo sensor and current humidity module state
				temp := sensor.MeteoData().Temp
				state := module.RelayState()

				if temp < module.Threshold() && !state {
					// If switched off send relay on command
					err := module.SwitchRelay(true)
					if err != nil {
						logger.Error(err.Error())
						module.SetError()
					}
				} else if temp >= module.Threshold() && state {
					// If switched on send relay off command
					err := module.SwitchRelay(false)
					if err != nil {
						logger.Error(err.Error())
						module.SetError()
					}
				}
			} else {
				if module.RelayState() {
					// Switch off air humidity control module if control is off
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
func (a *ThermoCtrlTask) Start() {
	a.reqTimer = time.NewTimer(taskDelay * time.Second)
	a.TaskHandler()
}
