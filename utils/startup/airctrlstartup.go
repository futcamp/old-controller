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
	"strconv"

	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/startup/io"
)

type AirctrlStartupCfg struct {
	mods     *io.StartupMods
	fileName string
}

// NewAirctrlStartupCfg make new struct
func NewAirctrlStartupCfg(mods *io.StartupMods) *AirctrlStartupCfg {
	return &AirctrlStartupCfg{
		mods: mods,
	}
}

// AddDevice add new air control device to configs
func (a *AirctrlStartupCfg) AddDevice(device string) error {
	return a.mods.SaveModCommand(utils.StartupCfgPath, "airctrl", "add-device", device, nil)
}

// DeleteDevice remove device from configs
func (a *AirctrlStartupCfg) DeleteDevice(device string) error {
	return a.mods.DeleteModCommand(utils.StartupCfgPath, "airctrl", "add-device", device)
}

// SetSensor set new sensor name for device
func (a *AirctrlStartupCfg) SetSensor(device string, sensor string) error {
	return a.mods.SaveModCommand(utils.StartupCfgPath, "airctrl", "sensor", device, []string{sensor})
}

// SetIP set new ip address value for device
func (a *AirctrlStartupCfg) SetIP(device string, ip string) error {
	return a.mods.SaveModCommand(utils.StartupCfgPath, "airctrl", "ip", device, []string{ip})
}

// SetThreshold set new threshold value for device
func (a *AirctrlStartupCfg) SetThreshold(device string, threshold int) error {
	sThreshold := strconv.Itoa(threshold)
	return a.mods.SaveModCommand(utils.StartupCfgPath, "airctrl", "threshold", device, []string{sThreshold})
}

// SetThreshold set new status value for device
func (a *AirctrlStartupCfg) SetStatus(device string, status bool) error {
	sStatus := strconv.FormatBool(status)
	return a.mods.SaveModCommand(utils.StartupCfgPath, "airctrl", "threshold", device, []string{sStatus})
}
