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

package devices

import (
	"fmt"
	"github.com/futcamp/controller/devices/data"
	"github.com/futcamp/controller/devices/modules"
	"github.com/futcamp/controller/utils/configs"
	"github.com/google/logger"
)

type UserKey struct {
	User string
	Key  string
}

type Security struct {
	modules map[string]modules.SecurityController
	dynCfg  *configs.DynamicConfigs
	data    data.SecurityData
	keys    []UserKey
}

// NewSecurity make new struct
func NewSecurity(dc *configs.DynamicConfigs) *Security {
	mods := make(map[string]modules.SecurityController)
	return &Security{
		modules: mods,
		dynCfg:  dc,
	}
}

//
// Getters and setters
//

// Status get current security status
func (s *Security) Status() bool {
	return s.data.Status()
}

// Status get current security alarm
func (s *Security) Alarm() bool {
	return s.data.Alarm()
}

// Status get current security status
func (s *Security) SetStatus(stat bool) {
	var state string

	// Set new security status
	s.data.SetStatus(stat)

	logger.Infof("Security set new status \"%t\"", stat)

	if stat {
		state = "on"
	} else {
		state = "off"
	}

	// Save status to configs
	s.dynCfg.AddCommand(fmt.Sprintf("security-status add-state %s", state))
	s.dynCfg.SaveConfigs()
}

// SetAlarm set new security alarm
func (s *Security) SetAlarm(alarm bool) {
	s.data.SetAlarm(alarm)
}

// AddKey add new user key
func (s *Security) AddKey(user string, key string) error {
	uKey := UserKey{
		User: user,
		Key:  key,
	}

	s.keys = append(s.keys, uKey)

	return nil
}

// Keys get all users keys
func (s *Security) Keys() *[]UserKey {
	return &s.keys
}

//
// Main functional
//

// AddModule add new security mod
func (s *Security) AddModule(name string, mod modules.SecurityController) {
	s.modules[name] = mod
}

// DeleteModule delete mod from storage
func (s *Security) DeleteModule(name string) {
	delete(s.modules, name)
}

// Module get mod by name
func (s *Security) Module(name string) modules.SecurityController {
	return s.modules[name]
}

// AllModules get all devices list
func (s *Security) AllModules() []modules.SecurityController {
	var mods []modules.SecurityController

	for _, mod := range s.modules {
		mods = append(mods, mod)
	}

	return mods
}
