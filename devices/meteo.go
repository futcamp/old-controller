/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2018-2019 Sergey Denisov.
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

import "github.com/futcamp/controller/devices/modules"

type MeteoStation struct {
	modules map[string]modules.MeteoController
}

// NewMeteoStation make new struct
func NewMeteoStation() *MeteoStation {
	mods := make(map[string]modules.MeteoController)
	return &MeteoStation{
		modules: mods,
	}
}

// AddModule add new meteo module
func (m *MeteoStation) AddModule(name string, module modules.MeteoController) {
	m.modules[name] = module
}

// DeleteModule delete module from storage
func (m *MeteoStation) DeleteModule(name string) {
	delete(m.modules, name)
}

// MeteoSensor get meteo module
func (m *MeteoStation) Module(name string) modules.MeteoController {
	return m.modules[name]
}

// AllSensors get all devices list
func (m *MeteoStation) AllModules() []modules.MeteoController {
	var mods []modules.MeteoController

	for _, sensor := range m.modules {
		mods = append(mods, sensor)
	}

	return mods
}
