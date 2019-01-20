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

package startup

import (
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/startup/io"

	"strconv"
)

type MeteoStartupCfg struct {
	mods *io.StartupMods
}

// NewMeteoStartupCfg make new struct
func NewMeteoStartupCfg(mods *io.StartupMods) *MeteoStartupCfg {
	return &MeteoStartupCfg{
		mods: mods,
	}
}

// AddDevice add new meteo device to configs
func (m *MeteoStartupCfg) AddDevice(device string) error {
	return m.mods.SaveModCommand(utils.StartupCfgPath, "meteo", "add-device", device, nil)
}

// DeleteDevice remove device from configs
func (m *MeteoStartupCfg) DeleteDevice(device string) error {
	return m.mods.DeleteModCommand(utils.StartupCfgPath, "meteo", "add-device", device)
}

// SetSensorIP set meteo sensor IP for device
func (m *MeteoStartupCfg) SetSensorIP(device string, sensIP string) error {
	return m.mods.SaveModCommand(utils.StartupCfgPath, "meteo", "ip", device, []string{sensIP})
}

// SetSensorType set meteo sensor type for device
func (m *MeteoStartupCfg) SetSensorType(device string, sensType string) error {
	return m.mods.SaveModCommand(utils.StartupCfgPath, "meteo", "type", device, []string{sensType})
}

// SetSensorChannel set meteo sensor channel for device
func (m *MeteoStartupCfg) SetSensorChannel(device string, channel int) error {
	ch := strconv.Itoa(channel)
	return m.mods.SaveModCommand(utils.StartupCfgPath, "meteo", "channel", device, []string{ch})
}
