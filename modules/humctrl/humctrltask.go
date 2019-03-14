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

package humctrl

import (
	"github.com/google/logger"
	"time"

	"github.com/futcamp/controller/modules/meteo"
)

const (
	taskDelay = 1
)

// HumControlTask humidity control task struct
type HumControlTask struct {
	meteo    *meteo.MeteoStation
	humCtrl  *HumControl
	reqTimer *time.Timer
}

// NewHumControlTask make new struct
func NewHumControlTask(hctrl *HumControl, meteo *meteo.MeteoStation) *HumControlTask {
	return &HumControlTask{
		humCtrl: hctrl,
		meteo:   meteo,
	}
}

// TaskHandler process timer loop
func (h *HumControlTask) TaskHandler() {
	for {
		<-h.reqTimer.C

		// Update current humidity
		for _, module := range h.humCtrl.AllModules() {
			(*module).SetHumidity(h.meteo.Sensor(module.Sensor()).MeteoData().Humidity)
		}

		// SyncData data with remote mod
		for _, module := range h.humCtrl.AllModules() {
			// Process data
			if module.Status() {
				if module.Humidity() < module.Threshold() {
					module.SetHumidifier(true)
				} else {
					module.SetHumidifier(false)
				}
			} else {
				module.SetHumidifier(false)
			}

			// Sync states with module
			err := module.SyncData()
			if err != nil {
				if !module.Error() {
					module.SetError(true)
					logger.Errorf("Fail to sync data with \"%s\" mod!", module.Name())
					logger.Error(err.Error())
				}
				continue
			}
			if module.Error() {
				module.SetError(false)
				logger.Errorf("Module \"%s\" was synced.", module.Name())
			}
		}

		h.reqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (h *HumControlTask) Start() {
	h.reqTimer = time.NewTimer(taskDelay * time.Second)
	h.TaskHandler()
}
