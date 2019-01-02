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

package configs

import (
	"bufio"
	"os"

	"go.uber.org/config"
)

type AirCtrlModule struct {
	Name      string
	IP        string
	Sensor    string
	Threshold int
}

type AirCtrlSettings struct {
	Modules []AirCtrlModule
}

type AirCtrlConfigs struct {
	settings AirCtrlSettings
}

// NewAirCtrlConfigs make new struct
func NewAirCtrlConfigs() *AirCtrlConfigs {
	return &AirCtrlConfigs{}
}

// LoadFromFile loading configs from cfg file
func (a *AirCtrlConfigs) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	provider, err := config.NewYAML(config.Source(reader))
	if err != nil {
		return err
	}

	err = provider.Get("airctrl").Populate(&a.settings)
	if err != nil {
		return err
	}

	return nil
}

// GetSettings get pointer of application settings
func (a *AirCtrlConfigs) Settings() *AirCtrlSettings {
	return &a.settings
}
