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
	"fmt"
	"sync"

	"github.com/futcamp/controller/utils/configs"
)

type HumControl struct {
	modules map[string]*Module
	dynCfg  *configs.DynamicConfigs
}

type ControlData struct {
	threshold int
	status    bool
	mtxThresh sync.Mutex
	mtxStatus sync.Mutex
}

type Module struct {
	Name       string
	IP         string
	Sensor     string
	Humidity   int
	Humidifier bool
	Data       ControlData
	Error      bool
	humCtrl    *HumControl
}

// SetThreshold set new threshold value
func (c *ControlData) SetThreshold(thresh int) {
	c.mtxThresh.Lock()
	c.threshold = thresh
	c.mtxThresh.Unlock()
}

// Threshold get threshold value
func (c *ControlData) Threshold() int {
	var thresh int

	c.mtxThresh.Lock()
	thresh = c.threshold
	c.mtxThresh.Unlock()

	return thresh
}

// SetStatus set new status state
func (c *ControlData) SetStatus(status bool) {
	c.mtxThresh.Lock()
	c.status = status
	c.mtxThresh.Unlock()
}

// Status get status state
func (c *ControlData) Status() bool {
	var status bool

	c.mtxThresh.Lock()
	status = c.status
	c.mtxThresh.Unlock()

	return status
}

// NewModule make new humidity control module
func (h *HumControl) NewModule(name string, ip string, sensor string, err bool) *Module {
	return &Module{
		Name:    name,
		IP:      ip,
		Sensor:  sensor,
		Error:   err,
		humCtrl: h,
	}
}

// SyncData get data from controller
func (m *Module) SyncData() error {
	var status bool
	var thresh int
	var stat string

	status = m.Data.Status()
	thresh = m.Data.Threshold()

	// Control logic
	if status {
		if m.Humidity < (thresh - 1) {
			m.Humidifier = true
		}
		if m.Humidity > (thresh + 1) {
			m.Humidifier = false
		}
	} else {
		m.Humidifier = false
	}

	// Send current states
	ctrl := NewWiFiController(m.IP)
	data, err := ctrl.SyncData(status, m.Humidifier)
	if err != nil {
		return err
	}

	// SAve new data
	if data.Status != status {
		m.Data.SetStatus(data.Status)
		if data.Status {
			stat = "on"
		} else {
			stat = "off"
		}
		m.humCtrl.dynCfg.AddCommand(fmt.Sprintf("humctrl status %s %s", m.Name, stat))
		m.humCtrl.dynCfg.SaveConfigs()
	}

	return nil
}

// NewHumidityControl make new struct
func NewHumidityControl(dc *configs.DynamicConfigs) *HumControl {
	mods := make(map[string]*Module)
	return &HumControl{
		modules: mods,
		dynCfg:  dc,
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
