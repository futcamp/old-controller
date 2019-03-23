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
	tempTaskDelay = 500
)

// TempControlTask temperature control task struct
type TempControlTask struct {
	dynCfg   *configs.DynamicConfigs
	tempCtrl  *devices.TempControl
	reqTimer *time.Timer
}

// NewTempControlTask make new struct
func NewTempControlTask(tctrl *devices.TempControl, dc *configs.DynamicConfigs) *TempControlTask {
	return &TempControlTask{
		tempCtrl: tctrl,
		dynCfg:  dc,
	}
}

// TaskHandler process timer loop
func (h *TempControlTask) TaskHandler() {
	for {
		<-h.reqTimer.C

		// SyncData data with remote devices
		for _, module := range h.tempCtrl.AllModules() {
			// Process data
			if module.Status() {
				if module.Temperature() < module.Threshold() {
					if !module.Heater() {
						logger.Infof("TempControl current temp \"%d\" from sensor \"%s\" less then \"%s\" threshold value \"%d\"",
							module.Temperature(), module.Sensor(), module.Name(), module.Threshold())
						module.SetUpdate(true)
					}
					module.SetHeater(true)
				} else {
					if module.Heater() {
						logger.Infof("TempControl current temp \"%d\" from sensor \"%s\" more then \"%s\" threshold value \"%d\"",
							module.Temperature(), module.Sensor(), module.Name(), module.Threshold())
						module.SetUpdate(true)
					}
					module.SetHeater(false)
				}
			} else {
				module.SetHeater(false)
			}

			// Sync states with module
			if module.Update() {
				err := module.SyncData()
				if err != nil {
					if !module.Error() {
						module.SetError(true)
						logger.Errorf("TempControl fail to sync data with \"%s\" module!", module.Name())
						logger.Error(err.Error())
					}
					continue
				}
				if module.Error() {
					module.SetError(false)
					logger.Errorf("TempControl module \"%s\" was synced", module.Name())
				}
				module.SetUpdate(false)
			}
		}

		h.reqTimer.Reset(tempTaskDelay * time.Millisecond)
	}
}

// Start start new timer
func (h *TempControlTask) Start() {
	time.Sleep(time.Duration(h.dynCfg.Settings().Timers.MeteoDBDelay+5) * time.Second)
	h.reqTimer = time.NewTimer(tempTaskDelay * time.Millisecond)
	h.TaskHandler()
}
