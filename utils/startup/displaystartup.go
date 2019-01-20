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
)

type DisplayStartupCfg struct {
	mods *io.StartupMods
}

// NewDisplayStartupCfg make new struct
func NewDisplayStartupCfg(mods *io.StartupMods) *DisplayStartupCfg {
	return &DisplayStartupCfg{
		mods: mods,
	}
}

// AddDevice add new display device to configs
func (m *DisplayStartupCfg) AddDevice(device string) error {
	return m.mods.SaveModCommand(utils.StartupCfgPath, "display", "add-device", device, nil)
}

// DeleteDevice remove device from configs
func (m *DisplayStartupCfg) DeleteDevice(device string) error {
	return m.mods.DeleteModCommand(utils.StartupCfgPath, "display", "add-device", device)
}

// SetIP set display IP
func (m *DisplayStartupCfg) SetIP(device string, ip string) error {
	return m.mods.SaveModCommand(utils.StartupCfgPath, "display", "ip", device, []string{ip})
}

// SetIP set display IP
func (m *DisplayStartupCfg) SetSensors(device string, ip string) error {
	return m.mods.SaveModCommand(utils.StartupCfgPath, "display", "ip", device, []string{ip})
}