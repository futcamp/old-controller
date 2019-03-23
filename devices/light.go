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
	"github.com/futcamp/controller/devices/modules"
	"github.com/futcamp/controller/utils/configs"
)

type Light struct {
	modules map[string]*modules.LightModule
	dynCfg  *configs.DynamicConfigs
}

// NewHumidityControl make new struct
func NewLight(dc *configs.DynamicConfigs) *Light {
	mods := make(map[string]*modules.LightModule)
	return &Light{
		modules: mods,
		dynCfg:  dc,
	}
}

// AddModule add new humidity control mod
func (l *Light) AddModule(name string, mod *modules.LightModule) {
	l.modules[name] = mod
}

// DeleteModule delete mod from storage
func (l *Light) DeleteModule(name string) {
	delete(l.modules, name)
}

// Module get mod by name
func (l *Light) Module(name string) *modules.LightModule {
	return l.modules[name]
}

// AllModules get all devices list
func (l *Light) AllModules() []*modules.LightModule {
	var mods []*modules.LightModule

	for _, mod := range l.modules {
		mods = append(mods, mod)
	}

	return mods
}
