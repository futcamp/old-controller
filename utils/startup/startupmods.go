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

import (
	"github.com/futcamp/controller/modules/airctrl"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
	"github.com/pkg/errors"
)

type StartupMods struct {
	cfg        *StartupIO
	airCtrl    *airctrl.AirControl
	airCtrlCfg *configs.AirCtrlConfigs
}

// NewStartupMods make new struct
func NewStartupMods(cfg *StartupIO, ac *airctrl.AirControl, acc *configs.AirCtrlConfigs) *StartupMods {
	return &StartupMods{
		cfg:        cfg,
		airCtrl:    ac,
		airCtrlCfg: acc,
	}
}

// LoadFromFile loading configs from file
func (s *StartupMods) LoadFromFile(fileName string) error {
	err := s.cfg.LoadCommands(fileName, func(module string, cmd string, dev string, value interface{}) {
		switch module {
		case "meteo":
			s.applyMeteoCfg(cmd, dev, value)
			break

		case "airctrl":
			s.applyAirCtrlCfg(cmd, dev, value)
			break
		}
	})

	return err
}

// applyMeteoCfg apply commands for meteo module
func (s *StartupMods) applyMeteoCfg(cmd string, dev string, value interface{}) {
}

// applyAirCtrlCfg apply commands for air control module
func (s *StartupMods) applyAirCtrlCfg(cmd string, dev string, value interface{}) {
	switch cmd {
	case "threshold":
		module := s.airCtrl.Module(dev)
		module.SetThreshold(value.(int))
		logger.Infof("airctrl apply threshold value %d for device %s", module.Threshold(), dev)
		break

	case "status":
		module := s.airCtrl.Module(dev)
		module.SwitchHumidityControl(value.(bool))
		logger.Infof("airctrl apply status value %t for device %s", module.HumidityControl(), dev)
		break
	}
}

// SaveModCommand save module command to startup-configs
func (s *StartupMods) SaveModCommand(fileName string, module string, cmd string, dev string, valType string,
	value interface{}) error {

	// Add command to command list
	s.cfg.AddCommand(module, cmd, dev, valType, value)

	// Apply command to application
	switch module {
	case "meteo":
		s.applyMeteoCfg(cmd, dev, value)
		break

	case "airctrl":
		s.applyAirCtrlCfg(cmd, dev, value)
		break
	}

	// Save command to file
	err := s.cfg.SaveCommands(fileName)
	if err != nil {
		return err
	}

	return nil
}

// DeleteModCommand delete command from startup-configs
func (s *StartupMods) DeleteModCommand(fileName string, module string, cmd string, dev string) error {
	// Delete command from command list
	s.cfg.DeleteCommand(module, cmd, dev)

	// Restore default value from configs file
	switch module {
	case "meteo":
		err := s.restoreMeteoCfg(cmd, dev)
		if err != nil {
			return err
		}
		break

	case "airctrl":
		err := s.restoreAirCtrlCfg(cmd, dev)
		if err != nil {
			return err
		}
		break
	}

	// Save command to file
	err := s.cfg.SaveCommands(fileName)
	if err != nil {
		return err
	}

	return nil
}

// restoreMeteoCfg restore default values from meteo cfg
func (s *StartupMods) restoreMeteoCfg(cmd string, dev string) error {
	return nil
}

// restoreAirCtrlCfg restore default values from air control cfg
func (s *StartupMods) restoreAirCtrlCfg(cmd string, dev string) error {
	switch cmd {
	case "threshold":
		for _, module := range s.airCtrlCfg.Settings().Modules {
			if module.Name == dev {
				logger.Infof("airctrl restore default threshold value for device %s", dev)
				s.applyAirCtrlCfg(cmd, dev, module.Threshold)
				return nil
			}
		}
		return errors.New("Fail to restore. Device not found")
		break

	case "status":
		for _, module := range s.airCtrlCfg.Settings().Modules {
			if module.Name == dev {
				logger.Infof("airctrl restore default status value for device %s", dev)
				s.applyAirCtrlCfg(cmd, dev, module.Status)
				return nil
			}
		}
		return errors.New("Fail to restore. Device not found")
		break
	}

	return errors.New("Unknown restore command")
}
