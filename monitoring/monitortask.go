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

package monitoring

import (
	"github.com/futcamp/controller/utils/configs"
	"time"
)

const (
	taskDelay = 5
)

// meteoTask meteo task struct
type MonitorTask struct {
	monitor    *DeviceMonitor
	dynCfg     *configs.DynamicConfigs
	reqTimer   *time.Timer
	lastHour   int
	monCounter int
}

// NewMonitorTask make new struct
func NewMonitorTask(monitor *DeviceMonitor, dc *configs.DynamicConfigs) *MonitorTask {
	return &MonitorTask{
		monitor:    monitor,
		dynCfg:     dc,
		monCounter: 0,
	}
}

// TaskHandler process timer loop
func (m *MonitorTask) TaskHandler() {
	for {
		<-m.reqTimer.C
		m.monCounter++

		if m.monCounter >= m.dynCfg.Settings().Timers.MonitorDelay {
			m.monCounter = 0
			m.monitor.CheckDevices()
		}

		m.reqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (m *MonitorTask) Start() {
	m.reqTimer = time.NewTimer(taskDelay * time.Second)
	m.TaskHandler()
}
