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

package tasks

import (
	"time"

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/utils/configs"
	"github.com/google/logger"
)

const (
	humTaskDelay = 500
)

// HumControlTask humidity control task struct
type HumControlTask struct {
	dynCfg   *configs.DynamicConfigs
	humCtrl  *devices.HumControl
	reqTimer *time.Timer
}

// NewHumControlTask make new struct
func NewHumControlTask(hctrl *devices.HumControl, dc *configs.DynamicConfigs) *HumControlTask {
	return &HumControlTask{
		humCtrl: hctrl,
		dynCfg:  dc,
	}
}

// TaskHandler process timer loop
func (h *HumControlTask) TaskHandler() {
	for {
		<-h.reqTimer.C

		// SyncData data with remote devices
		for _, module := range h.humCtrl.AllModules() {
			// Process data
			if module.Status() {
				if module.Humidity() < module.Threshold() {
					if !module.Humidifier() {
						logger.Infof("HumControl current humidity \"%d\" from sensor \"%s\" less then \"%s\" threshold value \"%d\"",
							module.Humidity(), module.Sensor(), module.Name(), module.Threshold())
						module.SetUpdate(true)
					}
					module.SetHumidifier(true)
				} else {
					if module.Humidifier() {
						logger.Infof("HumControl current humidity \"%d\" from sensor \"%s\" more then \"%s\" threshold value \"%d\"",
							module.Humidity(), module.Sensor(), module.Name(), module.Threshold())
						module.SetUpdate(true)
					}
					module.SetHumidifier(false)
				}
			} else {
				module.SetHumidifier(false)
			}

			// Sync states with module
			if module.Update() {
				err := module.SyncData()
				if err != nil {
					if !module.Error() {
						module.SetError(true)
						logger.Errorf("HumControl fail to sync data with \"%s\" module!", module.Name())
						logger.Error(err.Error())
					}
					continue
				}
				if module.Error() {
					module.SetError(false)
					logger.Errorf("HumControl module \"%s\" was synced", module.Name())
				}
				module.SetUpdate(false)
			}
		}

		h.reqTimer.Reset(humTaskDelay * time.Millisecond)
	}
}

// Start start new timer
func (h *HumControlTask) Start() {
	time.Sleep(time.Duration(h.dynCfg.Settings().Timers.MeteoDBDelay+5) * time.Second)
	h.reqTimer = time.NewTimer(humTaskDelay * time.Millisecond)
	h.TaskHandler()
}
