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

package humctrl

import (
	"sync"
)

type HumCtrlModuleData struct {
	Humidifier bool
}

type HumCtrlServerData struct {
	Status    bool
	Threshold int
	Hum       int
}

type Module struct {
	Name    string
	IP      string
	Sensor  string
	Error   bool
	ModData HumCtrlModuleData
	SrvData HumCtrlServerData
	ModMtx  sync.Mutex
	SrvMtx  sync.Mutex
}

type HumControl struct {
	modules map[string]*Module
}

// SetServerData set new server data
func (m *Module) SetServerData(status bool, thresh int, hum int) {
	m.SrvMtx.Lock()
	m.SrvData.Status = status
	m.SrvData.Threshold = thresh
	m.SrvData.Hum = hum
	m.SrvMtx.Unlock()
}

// ServerData get current server data
func (m *Module) ServerData() HumCtrlServerData {
	var data HumCtrlServerData

	m.SrvMtx.Lock()
	data.Threshold = m.SrvData.Threshold
	data.Status = m.SrvData.Status
	data.Hum = m.SrvData.Hum
	m.SrvMtx.Unlock()

	return data
}

// SetModuleData set new module data
func (m *Module) SetModuleData(hum bool) {
	m.ModMtx.Lock()
	m.ModData.Humidifier = hum
	m.ModMtx.Unlock()
}

// ServerData get current module data
func (m *Module) ModuleData() HumCtrlModuleData {
	var data HumCtrlModuleData

	m.ModMtx.Lock()
	data.Humidifier = m.ModData.Humidifier
	m.ModMtx.Unlock()

	return data
}

// SyncData get data from controller
func (s *Module) SyncData() error {
	var status bool
	var newStatus bool
	var newThreshold int
	var dataChngd bool

	sData := s.ServerData()
	newStatus = sData.Status
	newThreshold = sData.Threshold

	// Control logic
	if sData.Status {
		if sData.Hum < (sData.Threshold - 1) {
			status = true
		}
		if sData.Hum > (sData.Threshold + 1) {
			status = false
		}
	} else {
		status = false
	}

	// Send current states
	ctrl := NewWiFiController(s.IP)
	data, err := ctrl.SyncData(status, sData.Threshold, sData.Hum)
	if err != nil {
		return err
	}

	if data.Switch {
		newStatus = !sData.Status
		dataChngd = true
	}
	if data.Plus {
		newThreshold = sData.Threshold + 1
		dataChngd = true
	}
	if data.Minus {
		newThreshold = sData.Threshold - 1
		dataChngd = true
	}

	s.SetModuleData(data.Humidifier)

	if dataChngd {
		s.SetServerData(newStatus, newThreshold, sData.Hum)
		// save configs
	}

	return nil
}

// NewHumidityControl make new struct
func NewHumidityControl() *HumControl {
	mods := make(map[string]*Module)
	return &HumControl{
		modules: mods,
	}
}

// NewModule make new humidity control module
func (h *HumControl) NewModule(name string, ip string, sensor string, err bool) *Module {
	return &Module{
		Name:   name,
		IP:     ip,
		Sensor: sensor,
		Error:  err,
	}
}

// AddModule add new humidity control module
func (h *HumControl) AddModule(name string, mod *Module) {
	h.modules[name] = mod
}

// DeleteModule delete module from storage
func (h *HumControl) DeleteModule(name string) {
	delete(h.modules, name)
}

// Module get module by name
func (h *HumControl) Module(name string) *Module {
	return h.modules[name]
}

// AllModules get all modules list
func (h *HumControl) AllModules() []*Module {
	var mods []*Module

	for _, mod := range h.modules {
		mods = append(mods, mod)
	}

	return mods
}
