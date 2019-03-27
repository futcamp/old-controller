/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2019 Sergey Denisov.
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
	"github.com/futcamp/controller/devices/modules"
	"github.com/futcamp/controller/utils/configs"
)

const (
	motionTaskDelay = 1
)

type MotionTask struct {
	dynCfg   *configs.DynamicConfigs
	reqTimer *time.Timer
	motion   *devices.Motion
	light    *devices.Light
}

// NewMotionTask make new struct
func NewMotionTask(dc *configs.DynamicConfigs, mot *devices.Motion, lgh *devices.Light) *MotionTask {
	return &MotionTask{
		dynCfg: dc,
		motion: mot,
		light:  lgh,
	}
}

// ActivityRun start activity
func (m *MotionTask) ActivityRun(motCtrl modules.MotionController) {
	for _, lamp := range *motCtrl.Lamps() {
		module := m.light.Module(lamp)
		if !module.Status() {
			motCtrl.SetAlreadyOn(false)
			module.SetStatus(true)
			module.SetUpdate(true)
		} else {
			motCtrl.SetAlreadyOn(true)
		}
	}
}

// ActivityStop stop activity
func (m *MotionTask) ActivityStop(motCtrl modules.MotionController) {
	for _, lamp := range *motCtrl.Lamps() {
		module := m.light.Module(lamp)
		// If lamp was switch on before motion not switch off
		if !motCtrl.AlreadyOn() {
			module.SetStatus(false)
			module.SetUpdate(false)
		}
	}
}

// TaskHandler process timer loop
func (m *MotionTask) TaskHandler() {
	for {
		<-m.reqTimer.C

		for _, mod := range m.motion.AllModules() {
			if mod.Activity() {
				// First activity run
				if mod.CurDelay() == mod.Delay() {
					m.ActivityRun(mod)
				}

				if mod.CurDelay() != 0 {
					mod.SetCurDelay(mod.CurDelay() - 1)
				} else {
					mod.SetActivity(false)
					mod.SetCurDelay(mod.Delay())

					// Activity stop
					m.ActivityStop(mod)
				}
			}
		}

		m.reqTimer.Reset(motionTaskDelay * time.Second)
	}
}

// Start start new timer
func (m *MotionTask) Start() {
	m.reqTimer = time.NewTimer(motionTaskDelay * time.Second)
	m.TaskHandler()
}
