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

type TempControl struct {
	modules map[string]modules.TempController
	dynCfg  *configs.DynamicConfigs
}

// NewTemperatureControl make new struct
func NewTemperatureControl(dc *configs.DynamicConfigs) *TempControl {
	mods := make(map[string]modules.TempController)
	return &TempControl{
		modules: mods,
		dynCfg:  dc,
	}
}

// AddModule add new humidity control mod
func (t *TempControl) AddModule(name string, mod modules.TempController) {
	t.modules[name] = mod
}

// DeleteModule delete mod from storage
func (t *TempControl) DeleteModule(name string) {
	delete(t.modules, name)
}

// Module get mod by name
func (t *TempControl) Module(name string) modules.TempController {
	return t.modules[name]
}

// AllModules get all devices list
func (t *TempControl) AllModules() []modules.TempController {
	var mods []modules.TempController

	for _, mod := range t.modules {
		mods = append(mods, mod)
	}

	return mods
}
