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

type Illuminator interface {
	Update() bool
	Name() string
	IP() string
	Error() bool
	SetError(err bool)
	SetChannel(ch int)
	SetIP(ip string)
	SetUpdate(state bool)
	Status() bool
	SetStatus(status bool)
	SwitchStatus()
	SyncData() error
}

type LightModule struct {
	name      string
	ip        string
	channel   int
	error     bool
	update    bool
	mtxUpdate sync.Mutex
	data      data.LightData
	dynCfg    *configs.DynamicConfigs
}

// NewLightModule make new struct
func NewLightModule(name string, dc *configs.DynamicConfigs) *LightModule {
	return &LightModule{
		name:   name,
		dynCfg: dc,
		update: true,
	}
}

//
// Simple data getters and setters
//

// Update get current update state
func (l *LightModule) Update() bool {
	var value bool

	l.mtxUpdate.Lock()
	value = l.update
	l.mtxUpdate.Unlock()

	return value
}

// Name get mod name
func (l *LightModule) Name() string {
	return l.name
}

// IP get current mod ip
func (l *LightModule) IP() string {
	return l.ip
}

// Channel get current mod channel
func (l *LightModule) Channel() int {
	return l.channel
}

// Error get current error state
func (l *LightModule) Error() bool {
	return l.error
}

// SetError set new error state
func (l *LightModule) SetError(err bool) {
	l.error = err
}

// SetChannel set new channel
func (l *LightModule) SetChannel(ch int) {
	l.channel = ch
}

// SetIP set new ip address
func (l *LightModule) SetIP(ip string) {
	l.ip = ip
}

// SetUpdate set new update state
func (l *LightModule) SetUpdate(state bool) {
	l.mtxUpdate.Lock()
	l.update = state
	l.mtxUpdate.Unlock()
}

//
// Nested getters and setters
//

// Status get current status
func (l *LightModule) Status() bool {
	return l.data.Status()
}

//
// Setters with additional functional
//

// SetStatus set new status
func (l *LightModule) SetStatus(status bool) {
	var stat string

	if status {
		stat = "on"
	} else {
		stat = "off"
	}

	// Save value to local storage
	l.data.SetStatus(status)

	logger.Infof("Light set new status \"%s\" for device \"%s\"", stat, l.Name())

	// Save status to configs
	l.dynCfg.AddCommand(fmt.Sprintf("light status %s %s", l.Name(), stat))
	l.dynCfg.SaveConfigs()
}

// SwitchStatus switch status
func (l *LightModule) SwitchStatus() {
	status := l.Status()
	l.SetStatus(!status)
}

//
// Other functional
//

// SyncData sync data with module
func (l *LightModule) SyncData() error {
	_, err := hardware.HdkSyncLightData(l.IP(), l.Channel(), l.data.Status())
	if err != nil {
		return err
	}

	return nil
}
