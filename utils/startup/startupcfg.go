/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2018 Sergey Denisov.
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

type StartupCfg struct {
	mods     *StartupMods
	fileName string
}

// NewStartupCfg make new struct
func NewStartupCfg(mods *StartupMods) *StartupCfg {
	return &StartupCfg{
		mods: mods,
	}
}

// LoadFromFile loading startup-configs from file
func (s *StartupCfg) LoadFromFile(fileName string) error {
	s.fileName = fileName
	return s.mods.LoadFromFile(fileName)
}

// SaveAirCtrlThreshold save air control threshold
func (s *StartupCfg) SaveAirCtrlThreshold(device string, value int) error {
	return s.mods.SaveModCommand(s.fileName, "airctrl", "threshold", device, "int", value)
}

// DeleteAirCtrlThreshold delete air control threshold
func (s *StartupCfg) DeleteAirCtrlThreshold(device string) error {
	return s.mods.DeleteModCommand(s.fileName, "airctrl", "threshold", device)
}

// SaveAirCtrlStatus save air control status
func (s *StartupCfg) SaveAirCtrlStatus(device string, value bool) error {
	return s.mods.SaveModCommand(s.fileName, "airctrl", "status", device, "int", value)
}

// DeleteAirCtrlThreshold delete air control status
func (s *StartupCfg) DeleteAirCtrlStatus(device string) error {
	return s.mods.DeleteModCommand(s.fileName, "airctrl", "status", device)
}
