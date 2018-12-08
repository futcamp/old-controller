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

package monitoring

import (
	"time"
)

const (
	taskDelay = 5
)

// MeteoTask meteo task struct
type MonitorTask struct {
	Monitor *DeviceMonitor
	ReqTimer        *time.Timer
	DisplaysCounter int
	SensorsCounter  int
	DatabaseCounter int
	LastHour        int
}

// NewMonitorTask make new struct
func NewMonitorTask(monitor *DeviceMonitor) *MonitorTask {
	return &MonitorTask{
		Monitor:monitor,
	}
}

// TaskHandler process timer loop
func (m *MonitorTask) TaskHandler() {
	for {
		<-m.ReqTimer.C

		m.Monitor.CheckDevices()

		m.ReqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (m *MonitorTask) Start() {
	m.ReqTimer = time.NewTimer(taskDelay * time.Second)
	m.TaskHandler()
}