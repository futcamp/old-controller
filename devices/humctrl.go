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


type HumControl struct {
	modules map[string]modules.Humidifier
	dynCfg  *configs.DynamicConfigs
}

// NewHumidityControl make new struct
func NewHumidityControl(dc *configs.DynamicConfigs) *HumControl {
	mods := make(map[string]modules.Humidifier)
	return &HumControl{
		modules: mods,
		dynCfg:  dc,
	}
}

// AddModule add new humidity control mod
func (h *HumControl) AddModule(name string, mod *modules.HumCtrlModule) {
	h.modules[name] = mod
}

// DeleteModule delete mod from storage
func (h *HumControl) DeleteModule(name string) {
	delete(h.modules, name)
}

// Module get mod by name
func (h *HumControl) Module(name string) modules.Humidifier {
	return h.modules[name]
}

// AllModules get all devices list
func (h *HumControl) AllModules() []modules.Humidifier {
	var mods []modules.Humidifier

	for _, mod := range h.modules {
		mods = append(mods, mod)
	}

	return mods
}
