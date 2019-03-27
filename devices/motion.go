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
)

type Motion struct {
	modules map[string]modules.MotionController
}

// NewMotion make new struct
func NewMotion() *Motion {
	mods := make(map[string]modules.MotionController)
	return &Motion{
		modules: mods,
	}
}

// AddModule add new motion control mod
func (m *Motion) AddModule(name string, mod modules.MotionController) {
	m.modules[name] = mod
}

// DeleteModule delete mod from storage
func (m *Motion) DeleteModule(name string) {
	delete(m.modules, name)
}

// Module get mod by name
func (m *Motion) Module(name string) modules.MotionController {
	return m.modules[name]
}

// AllModules get all devices list
func (m *Motion) AllModules() []modules.MotionController {
	var mods []modules.MotionController

	for _, mod := range m.modules {
		mods = append(mods, mod)
	}

	return mods
}
