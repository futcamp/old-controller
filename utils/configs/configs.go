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

package configs

import (
	"bufio"
	"os"

	"go.uber.org/config"
)

const (
	AppName    = "futcamp"
	ApiVersion = "v2"
)

type ModCfg struct {
	Meteo   bool
	Humctrl bool
}

type ServerCfg struct {
	IP   string
	Port int
}

type AppSettings struct {
	Server     ServerCfg
	RCliServer ServerCfg
	Modules    ModCfg
}

type Configs struct {
	settings AppSettings
}

// NewConfigs make new struct
func NewConfigs() *Configs {
	return &Configs{}
}

// LoadFromFile loading configs from cfg file
func (c *Configs) LoadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	provider, err := config.NewYAML(config.Source(reader))
	if err != nil {
		return err
	}

	err = provider.Get(AppName).Populate(&c.settings)
	if err != nil {
		return err
	}

	return nil
}

// GetSettings get pointer of application settings
func (c *Configs) Settings() *AppSettings {
	return &c.settings
}
