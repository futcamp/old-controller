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

package airctrl

import "sync"

type AirCtrlModule struct {
	Name      string
	IP        string
	Sensor    string
	Threshold int
	state     bool
	exchError bool
	mtxState  sync.Mutex
	mtxErr    sync.Mutex
}

type AirControl struct {
	status  bool
	mtxStat sync.Mutex
	modules map[string]*AirCtrlModule
}

// NewAirCtrlModule make new AirCtrlModule struct
func NewAirCtrlModule(name string, ip string, sensor string, thr int) *AirCtrlModule {
	return &AirCtrlModule{
		Name:      name,
		IP:        ip,
		Sensor:    sensor,
		Threshold: thr,
		state:     false,
		exchError: false,
	}
}

// SwitchRelay send switch relay command to controller and switch oper state
func (m *AirCtrlModule) SwitchRelay(state bool) error {
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

// GetRelayState get current relay oper state
func (m *AirCtrlModule) GetRelayState() bool {
	var state bool

	m.mtxState.Lock()
	state = m.state
	m.mtxState.Unlock()

	return state
}

// SetError set error flag
func (m *AirCtrlModule) SetError() {
	m.mtxState.Lock()
	m.exchError = true
	m.mtxState.Unlock()
}

// GetError get error flag state
func (m *AirCtrlModule) GetError() bool {
	var err bool

	m.mtxState.Lock()
	err = m.exchError
	m.mtxState.Unlock()

	return err
}

// ClearError clear error flag
func (m *AirCtrlModule) ClearError() {
	m.mtxState.Lock()
	m.exchError = false
	m.mtxState.Unlock()
}

// SyncModule relay sync state with controller
func (m *AirCtrlModule) SyncModule() error {
	ctrl := NewWiFiController(m.IP)
	data, err := ctrl.SyncAirData()
	if err != nil {
		return err
	}

	if !data.Synced || m.GetError() {
		curState := m.GetRelayState()

		err = ctrl.SwitchRelay(curState)
		if err != nil {
			return err
		}
	}

	return nil
}

// NewAirControl make new AirControl struct
func NewAirControl() *AirControl {
	modules := make(map[string]*AirCtrlModule)

	return &AirControl{
		status:  false,
		modules: modules,
	}
}

// AddModule add new air control module
func (a *AirControl) AddModule(name string, module *AirCtrlModule) {
	a.modules[name] = module
}

// setHumidityControl set new state to humidity control
func (a *AirControl) setHumidityControl(state bool) error {
	a.mtxStat.Lock()
	a.status = state
	a.mtxStat.Unlock()

	// TODO: add save state to database

	return nil
}

// GetHumidityControl get current state of humidity control
func (a *AirControl) GetHumidityControl() bool {
	var state bool

	a.mtxStat.Lock()
	state = a.status
	a.mtxStat.Unlock()

	return state
}

// HumidityControlOn switch on control
func (a *AirControl) HumidityControlOn() error {
	return a.setHumidityControl(true)
}

// HumidityControlOff switch off control
func (a *AirControl) HumidityControlOff() error {
	return a.setHumidityControl(false)
}

// GetModules get all modules array
func (a *AirControl) GetModules() []*AirCtrlModule {
	var modules []*AirCtrlModule

	for _, mod := range a.modules {
		modules = append(modules, mod)
	}

	return modules
}
