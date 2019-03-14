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
	"github.com/futcamp/controller/modules/humctrl/mod"
	"github.com/futcamp/controller/utils/configs"
)

type HumControl struct {
	modules map[string]*mod.HumCtrlModule
	dynCfg  *configs.DynamicConfigs
}

// NewHumidityControl make new struct
func NewHumidityControl(dc *configs.DynamicConfigs) *HumControl {
	mods := make(map[string]*mod.HumCtrlModule)
	return &HumControl{
		modules: mods,
		dynCfg:  dc,
	}
}

// AddModule add new humidity control mod
func (h *HumControl) AddModule(name string, mod *mod.HumCtrlModule) {
	h.modules[name] = mod
}

// DeleteModule delete mod from storage
func (h *HumControl) DeleteModule(name string) {
	delete(h.modules, name)
}

// Module get mod by name
func (h *HumControl) Module(name string) *mod.HumCtrlModule {
	return h.modules[name]
}

// AllModules get all modules list
func (h *HumControl) AllModules() []*mod.HumCtrlModule {
	var mods []*mod.HumCtrlModule

	for _, mod := range h.modules {
		mods = append(mods, mod)
	}

	return mods
}
