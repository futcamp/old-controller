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

type TempController interface {
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
	Temperature() int
	Heater() bool
	SetTemperature(value int)
	SetHeater(status bool)
	SetThreshold(value int)
	SetStatus(status bool)
	SwitchStatus()
	SyncData() error
}

type TempCtrlModule struct {
	name      string
	ip        string
	sensor    string
	error     bool
	update    bool
	mtxUpdate sync.Mutex
	data      data.TempCtrlData
	dynCfg    *configs.DynamicConfigs
}

// NewTempCtrlModule make new struct
func NewTempCtrlModule(name string, dc *configs.DynamicConfigs) *TempCtrlModule {
	return &TempCtrlModule{
		name:   name,
		dynCfg: dc,
		update: true,
	}
}

//
// Simple data getters and setters
//

// Update get current update state
func (t *TempCtrlModule) Update() bool {
	var value bool

	t.mtxUpdate.Lock()
	value = t.update
	t.mtxUpdate.Unlock()

	return value
}

// Name get mod name
func (t *TempCtrlModule) Name() string {
	return t.name
}

// Sensor get temperature sensor name
func (t *TempCtrlModule) Sensor() string {
	return t.sensor
}

// IP get current mod ip
func (t *TempCtrlModule) IP() string {
	return t.ip
}

// Error get current error state
func (t *TempCtrlModule) Error() bool {
	return t.error
}

// SetError set new error state
func (t *TempCtrlModule) SetError(err bool) {
	t.error = err
}

// SetSensor set new temperature sensor name
func (t *TempCtrlModule) SetSensor(sensor string) {
	t.sensor = sensor
}

// SetIP set new ip address
func (t *TempCtrlModule) SetIP(ip string) {
	t.ip = ip
}

// SetUpdate set new update state
func (t *TempCtrlModule) SetUpdate(state bool) {
	t.mtxUpdate.Lock()
	t.update = state
	t.mtxUpdate.Unlock()
}

//
// Nested getters and setters
//

// Status get current tempctrl status
func (t *TempCtrlModule) Status() bool {
	return t.data.Status()
}

// Threshold get current threshold value
func (t *TempCtrlModule) Threshold() int {
	return t.data.Threshold()
}

// Temperature get current temperature value
func (t *TempCtrlModule) Temperature() int {
	return t.data.Temperature()
}

// Heater get current heater status
func (t *TempCtrlModule) Heater() bool {
	return t.data.Heater()
}

// SetTemperature set new temperature value
func (t *TempCtrlModule) SetTemperature(value int) {
	t.data.SetTemperature(value)
}

// SetHeater set new heater status
func (t *TempCtrlModule) SetHeater(status bool) {
	t.data.SetHeater(status)
}

//
// Setters with additional functional
//

// SetThreshold set new threshold value
func (t *TempCtrlModule) SetThreshold(value int) {
	// Save value to local storage
	t.data.SetThreshold(value)

	logger.Infof("TempControl set new threshold \"%d\" for device \"%s\"", value, t.Name())

	// Save threshold to configs
	t.dynCfg.AddCommand(fmt.Sprintf("tempctrl threshold %s %d", t.Name(), value))
	t.dynCfg.SaveConfigs()
}

// SetStatus set new tempctrl status
func (t *TempCtrlModule) SetStatus(status bool) {
	var stat string

	if status {
		stat = "on"
	} else {
		stat = "off"
	}

	// Save value to local storage
	t.data.SetStatus(status)

	logger.Infof("TempControl set new status \"%s\" for device \"%s\"", stat, t.Name())

	// Save status to configs
	t.dynCfg.AddCommand(fmt.Sprintf("tempctrl status %s %s", t.Name(), stat))
	t.dynCfg.SaveConfigs()
}

// SwitchStatus switch tempctrl status
func (t *TempCtrlModule) SwitchStatus() {
	status := t.Status()
	t.SetStatus(!status)
}

//
// Other functional
//

// SyncData sync data with module
func (t *TempCtrlModule) SyncData() error {
	_, err := hardware.HdkSyncTempCtrlData(t.ip, t.data.Status(), t.data.Heater())
	if err != nil {
		return err
	}

	return nil
}
