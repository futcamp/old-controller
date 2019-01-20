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

package thermoctrl

import "sync"

type ThermoCtrlModule struct {
	Name      string
	IP        string
	Sensor    string
	threshold int
	state     bool
	exchError bool
	status    bool
	mtxState  sync.Mutex
	mtxErr    sync.Mutex
	mtxThr    sync.Mutex
	mtxStat   sync.Mutex
}

type ThermoControl struct {
	modules map[string]*ThermoCtrlModule
}

// NewThermoCtrlModule make new ThermoCtrlModule struct
func NewThermoCtrlModule(name string, ip string, sensor string, thr int) *ThermoCtrlModule {
	return &ThermoCtrlModule{
		Name:      name,
		IP:        ip,
		Sensor:    sensor,
		threshold: thr,
		state:     false,
		exchError: false,
	}
}

// SwitchRelay send switch relay command to controller and switch oper state
func (m *ThermoCtrlModule) SwitchRelay(state bool) error {
	ctrl := NewWiFiController(m.IP)
	err := ctrl.SwitchRelay(state)
	if err != nil {
		return err
	}

	m.mtxState.Lock()
	m.state = state
	m.mtxState.Unlock()

	return nil
}

// RelayState get current relay oper state
func (m *ThermoCtrlModule) RelayState() bool {
	var state bool

	m.mtxState.Lock()
	state = m.state
	m.mtxState.Unlock()

	return state
}

// SetError set error flag
func (m *ThermoCtrlModule) SetError() {
	m.mtxState.Lock()
	m.exchError = true
	m.mtxState.Unlock()
}

// Error get error flag state
func (m *ThermoCtrlModule) Error() bool {
	var err bool

	m.mtxState.Lock()
	err = m.exchError
	m.mtxState.Unlock()

	return err
}

// SetThreshold set threshold value
func (m *ThermoCtrlModule) SetThreshold(value int) {
	m.mtxThr.Lock()
	m.threshold = value
	m.mtxThr.Unlock()
}

// Threshold get threshold value
func (m *ThermoCtrlModule) Threshold() int {
	var value int

	m.mtxThr.Lock()
	value = m.threshold
	m.mtxThr.Unlock()

	return value
}

// SwitchThermoControl switch mode of thermo control for device
func (m *ThermoCtrlModule) SwitchThermoControl(state bool) {
	m.mtxStat.Lock()
	m.status = state
	m.mtxStat.Unlock()
}

// ThermoControl get status of humidity control for device
func (m *ThermoCtrlModule) ThermoControl() bool {
	var status bool

	m.mtxThr.Lock()
	status = m.status
	m.mtxThr.Unlock()

	return status
}

// ClearError clear error flag
func (m *ThermoCtrlModule) ClearError() {
	m.mtxState.Lock()
	m.exchError = false
	m.mtxState.Unlock()
}

// SyncModule relay sync state with controller
func (m *ThermoCtrlModule) SyncModule() error {
	ctrl := NewWiFiController(m.IP)
	data, err := ctrl.SyncThermoData()
	if err != nil {
		return err
	}

	if !data.Synced || m.Error() {
		curState := m.RelayState()

		err = ctrl.SwitchRelay(curState)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewThermoControl make new ThermoControl struct
func NewThermoControl() *ThermoControl {
	modules := make(map[string]*ThermoCtrlModule)

	return &ThermoControl{
		modules: modules,
	}
}

// AddModule add new thermo control module
func (a *ThermoControl) AddModule(name string, module *ThermoCtrlModule) {
	a.modules[name] = module
}

// Modules get all modules array
func (a *ThermoControl) Modules() []*ThermoCtrlModule {
	var modules []*ThermoCtrlModule

	for _, mod := range a.modules {
		modules = append(modules, mod)
	}

	return modules
}

// Module get single module pointer
func (a *ThermoControl) Module(name string) *ThermoCtrlModule {
	return a.modules[name]
}
