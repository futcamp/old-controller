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
	"time"

	"github.com/futcamp/controller/modules/meteo"

	"github.com/google/logger"
)

const (
	taskDelay = 200
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

		// Update current temperatures
		for _, module := range h.humCtrl.AllModules() {
			data := module.ServerData()
			temp := h.meteo.Sensor(module.Sensor).MeteoData().Temp
			module.SetServerData(data.Status, data.Threshold, temp)
		}

		// Sync data with remote module
		for _, module := range h.humCtrl.AllModules() {
			err := module.SyncData()
			if err != nil {
				if !module.Error {
					module.Error = true
					logger.Errorf("Fail to sync data with \"%s\" module!", module.Name)
					logger.Error(err.Error())
				}
				continue
			}
			if module.Error {
				module.Error = false
				logger.Errorf("Module \"%s\" was synced.", module.Name)
			}
		}

		h.reqTimer.Reset(taskDelay * time.Millisecond)
	}
}

// Start start new timer
func (h *HumControlTask) Start() {
	h.reqTimer = time.NewTimer(taskDelay * time.Millisecond)
	h.TaskHandler()
}
