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

package modules

import (
	"fmt"
	"sync"

	"github.com/futcamp/controller/devices/data"
	"github.com/futcamp/controller/devices/hardware"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

type Humidifier interface {
	Update() bool
	Name() string
	Sensor() string
	IP() string
	Error() bool
	SetError(err bool)
	SetSensor(sensor string)
	SetIP(ip string)
	SetUpdate(state bool)
	Status() bool
	Threshold() int
	Humidity() int
	Humidifier() bool
	SetHumidity(value int)
	SetHumidifier(status bool)
	SetThreshold(value int)
	SetStatus(status bool)
	SwitchStatus()
	SyncData() error
}

type HumCtrlModule struct {
	name      string
	ip        string
	sensor    string
	error     bool
	update    bool
	mtxUpdate sync.Mutex
	data      data.HumCtrlData
	dynCfg    *configs.DynamicConfigs
}

// NewHumCtrlModule make new struct
func NewHumCtrlModule(name string, dc *configs.DynamicConfigs) *HumCtrlModule {
	return &HumCtrlModule{
		name:   name,
		dynCfg: dc,
		update: true,
	}
}

//
// Simple data getters and setters
//

// Update get current update state
func (h *HumCtrlModule) Update() bool {
	var value bool

	h.mtxUpdate.Lock()
	value = h.update
	h.mtxUpdate.Unlock()

	return value
}

// Name get mod name
func (h *HumCtrlModule) Name() string {
	return h.name
}

// Sensor get humidity sensor name
func (h *HumCtrlModule) Sensor() string {
	return h.sensor
}

// IP get current mod ip
func (h *HumCtrlModule) IP() string {
	return h.ip
}

// Error get current error state
func (h *HumCtrlModule) Error() bool {
	return h.error
}

// SetError set new error state
func (h *HumCtrlModule) SetError(err bool) {
	h.error = err
}

// SetSensor set new humidity sensor name
func (h *HumCtrlModule) SetSensor(sensor string) {
	h.sensor = sensor
}

// SetIP set new ip address
func (h *HumCtrlModule) SetIP(ip string) {
	h.ip = ip
}

// SetUpdate set new update state
func (h *HumCtrlModule) SetUpdate(state bool) {
	h.mtxUpdate.Lock()
	h.update = state
	h.mtxUpdate.Unlock()
}

//
// Nested getters and setters
//

// Status get current humctrl status
func (h *HumCtrlModule) Status() bool {
	return h.data.Status()
}

// Threshold get current threshold value
func (h *HumCtrlModule) Threshold() int {
	return h.data.Threshold()
}

// Humidity get current humidity value
func (h *HumCtrlModule) Humidity() int {
	return h.data.Humidity()
}

// Humidifier get current humidifier status
func (h *HumCtrlModule) Humidifier() bool {
	return h.data.Humidifier()
}

// SetHumidity set new humidity value
func (h *HumCtrlModule) SetHumidity(value int) {
	h.data.SetHumidity(value)
}

// SetHumidifier set new humidifier status
func (h *HumCtrlModule) SetHumidifier(status bool) {
	h.data.SetHumidifier(status)
}

//
// Setters with additional functional
//

// SetThreshold set new threshold value
func (h *HumCtrlModule) SetThreshold(value int) {
	// Save value to local storage
	h.data.SetThreshold(value)

	logger.Infof("HumControl set new threshold \"%d\" for device \"%s\"", value, h.Name())

	// Save threshold to configs
	h.dynCfg.AddCommand(fmt.Sprintf("humctrl threshold %s %d", h.Name(), value))
	h.dynCfg.SaveConfigs()
}

// SetStatus set new humctrl status
func (h *HumCtrlModule) SetStatus(status bool) {
	var stat string

	if status {
		stat = "on"
	} else {
		stat = "off"
	}

	// Save value to local storage
	h.data.SetStatus(status)

	logger.Infof("HumControl set new status \"%s\" for device \"%s\"", stat, h.Name())

	// Save status to configs
	h.dynCfg.AddCommand(fmt.Sprintf("humctrl status %s %s", h.Name(), stat))
	h.dynCfg.SaveConfigs()
}

// SwitchStatus switch humctrl status
func (h *HumCtrlModule) SwitchStatus() {
	status := h.Status()
	h.SetStatus(!status)
}

//
// Other functional
//

// SyncData sync data with module
func (h *HumCtrlModule) SyncData() error {
	_, err := hardware.HdkSyncHumCtrlData(h.ip, h.data.Status(), h.data.Humidifier())
	if err != nil {
		return err
	}

	return nil
}
