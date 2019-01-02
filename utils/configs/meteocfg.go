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

type MeteoDatabase struct {
	IP       string
	User     string
	Password string
	Base     string
}

type MeteoDelays struct {
	Sensors  int
	Displays int
	Database int
}

type MeteoDisplay struct {
	Name    string
	IP      string
	Enable  bool
	Sensors []string
}

type MeteoSensor struct {
	Name    string
	Type    string
	IP      string
	Channel int
	Enable  bool
}

type MeteoSettings struct {
	Delays   MeteoDelays
	Sensors  []MeteoSensor
	Displays []MeteoDisplay
	Database MeteoDatabase
}

type MeteoConfigs struct {
	settings MeteoSettings
}

// NewMeteoConfigs make new struct
func NewMeteoConfigs() *MeteoConfigs {
	return &MeteoConfigs{}
}

// LoadFromFile loading configs from cfg file
func (m *MeteoConfigs) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	provider, err := config.NewYAML(config.Source(reader))
	if err != nil {
		return err
	}

	err = provider.Get("meteo").Populate(&m.settings)
	if err != nil {
		return err
	}

	return nil
}

// GetSettings get pointer of application settings
func (m *MeteoConfigs) Settings() *MeteoSettings {
	return &m.settings
}
