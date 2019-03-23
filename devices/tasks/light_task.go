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
	lightTaskDelay = 500
)

// HumControlTask humidity control task struct
type LightlTask struct {
	dynCfg   *configs.DynamicConfigs
	light    *devices.Light
	reqTimer *time.Timer
}

// NewLightlTask make new struct
func NewLightlTask(lgh *devices.Light, dc *configs.DynamicConfigs) *LightlTask {
	return &LightlTask{
		light:  lgh,
		dynCfg: dc,
	}
}

// TaskHandler process timer loop
func (l *LightlTask) TaskHandler() {
	for {
		<-l.reqTimer.C

		// SyncData data with remote devices
		for _, module := range l.light.AllModules() {
			// Sync states with module
			if module.Update() {
				err := module.SyncData()
				if err != nil {
					if !module.Error() {
						module.SetError(true)
						logger.Errorf("Light fail to sync data with \"%s\" module!", module.Name())
						logger.Error(err.Error())
					}
					continue
				}
				if module.Error() {
					module.SetError(false)
					logger.Errorf("Light module \"%s\" was synced", module.Name())
				}
				module.SetUpdate(false)
			}
		}

		l.reqTimer.Reset(lightTaskDelay * time.Millisecond)
	}
}

// Start start new timer
func (l *LightlTask) Start() {
	l.reqTimer = time.NewTimer(lightTaskDelay * time.Millisecond)
	l.TaskHandler()
}
