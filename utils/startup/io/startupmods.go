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

package io

import (
	"errors"
	"strconv"

	"github.com/futcamp/controller/modules/airctrl"
	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

type StartupMods struct {
	cfg      *StartupIO
	airCtrl  *airctrl.AirControl
	meteo    *meteo.MeteoStation
	dynCfg   *configs.DynamicConfigs
	meteoLCD *meteo.MeteoDisplays
}

// NewStartupMods make new struct
func NewStartupMods(cfg *StartupIO, ac *airctrl.AirControl, dc *configs.DynamicConfigs,
	meteo *meteo.MeteoStation, mlcd *meteo.MeteoDisplays) *StartupMods {
	return &StartupMods{
		cfg:      cfg,
		airCtrl:  ac,
		dynCfg:   dc,
		meteo:    meteo,
		meteoLCD: mlcd,
	}
}

// LoadFromFile loading configs from file
func (s *StartupMods) LoadFromFile(fileName string) error {
	err := s.cfg.LoadCommands(fileName, func(cmd *StartupCmd) {
		switch cmd.Module {
		case "meteo":
			s.applyMeteoCfg(cmd.Command, cmd.Device, cmd.Args)
			break

		case "display":
			s.applyMeteoLCDCfg(cmd.Command, cmd.Device, cmd.Args)
			break

		case "airctrl":
			s.applyAirCtrlCfg(cmd.Command, cmd.Device, cmd.Args)
			break

		case "db":
			s.applyDBCfg(cmd.Command, cmd.Device, cmd.Args)
			break

		case "timer":
			s.applyTimerCfg(cmd.Command, cmd.Device, cmd.Args)
			break
		}
	})

	return err
}

// SaveModCommand save module command to startup-configs
func (s *StartupMods) SaveModCommand(fileName string, module string, cmd string, dev string, args []string) error {

	// Add command to command list
	s.cfg.AddCommand(module, cmd, dev, args)

	// Apply command to application
	switch module {
	case "meteo":
		s.applyMeteoCfg(cmd, dev, args)
		break

	case "display":
		s.applyMeteoLCDCfg(cmd, dev, args)
		break

	case "airctrl":
		s.applyAirCtrlCfg(cmd, dev, args)
		break

	case "db":
		s.applyDBCfg(cmd, dev, args)
		break

	case "timer":
		s.applyTimerCfg(cmd, dev, args)
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

	// Save command to file
	err := s.cfg.SaveCommands(fileName)
	if err != nil {
		return err
	}

	return nil
}

// applyMeteoCfg apply commands for meteo module
func (s *StartupMods) applyMeteoCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		sensor := s.meteo.NewMeteoSensor(dev, "", "", 0)
		s.meteo.AddSensor(dev, sensor)
		logger.Infof("meteo add new device \"%s\"", dev)
		break

	case "ip":
		sensor := s.meteo.Sensor(dev)
		sensor.IP = args[0]
		logger.Infof("meteo set ip address \"%s\" for device \"%s\"", sensor.IP, dev)
		break

	case "type":
		sensor := s.meteo.Sensor(dev)
		sensor.Type = args[0]
		logger.Infof("meteo set sensor type \"%s\" for device \"%s\"", sensor.Type, dev)
		break

	case "channel":
		ch, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		sensor := s.meteo.Sensor(dev)
		sensor.Channel = ch
		logger.Infof("meteo set sensor channel \"%d\" for device \"%s\"", sensor.Channel, dev)
		break
	}

	return nil
}

// applyMeteoLCDCfg apply commands for meteo displays
func (s *StartupMods) applyMeteoLCDCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		lcd := meteo.NewMeteoDisplay(dev, "")
		s.meteoLCD.AddMeteoDisplay(dev, lcd)
		logger.Infof("display add new device \"%s\"", dev)
		break

	case "ip":
		lcd := s.meteoLCD.Display(dev)
		lcd.IP = args[0]
		logger.Infof("display set ip address \"%s\" for device \"%s\"", lcd.IP, dev)
		break

	case "sensors":
		lcd := s.meteoLCD.Display(dev)

		for _, sensor := range args {
			lcd.AddDisplayingSensor(sensor)
			logger.Infof("display add displaying sensor \"%s\" for device \"%s\"", sensor, dev)
		}
		break
	}

	return nil
}

// applyAirCtrlCfg apply commands for air control module
func (s *StartupMods) applyAirCtrlCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		mod := airctrl.NewAirCtrlModule(dev, "", "", 0)
		s.airCtrl.AddModule(dev, mod)
		logger.Infof("airctrl add new device \"%s\"", dev)
		break

	case "ip":
		module := s.airCtrl.Module(dev)
		module.IP = args[0]
		logger.Infof("airctrl set ip address \"%s\" for device \"%s\"", module.IP, dev)
		break

	case "sensor":
		module := s.airCtrl.Module(dev)
		module.Sensor = args[0]
		logger.Infof("airctrl set sensor \"%s\" for device \"%s\"", module.Sensor, dev)
		break

	case "threshold":
		humidity, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		module := s.airCtrl.Module(dev)
		module.SetThreshold(humidity)
		logger.Infof("airctrl apply threshold value \"%d\" for device \"%s\"", module.Threshold(), dev)
		break

	case "status":
		status, err := strconv.ParseBool(args[0])
		if err != nil {
			return err
		}

		module := s.airCtrl.Module(dev)
		module.SwitchHumidityControl(status)
		logger.Infof("airctrl apply status value \"%t\" for device \"%s\"", module.HumidityControl(), dev)
		break
	}

	return nil
}

// applyAirCtrlCfg apply commands for air control module
func (s *StartupMods) applyTimerCfg(cmd string, dev string, args []string) error {
	delay, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	if cmd == "add-delay" {
		switch dev {
		case "meteo-sensors":
			s.dynCfg.Timers.MeteoSensorsDelay = delay
			break

		case "meteo-lcd":
			s.dynCfg.Timers.MeteoDisplayDelay = delay
			break

		case "meteo-db":
			s.dynCfg.Timers.MeteoDBDelay = delay
			break
		}
	} else {
		return errors.New("command not found")
	}

	logger.Infof("global apply timer delay \"%d\" for timer \"%s\"", delay, dev)

	return nil
}

// applyAirCtrlCfg apply commands for air control module
func (s *StartupMods) applyDBCfg(cmd string, dev string, args []string) error {
	if cmd == "add-base" {
		switch dev {
		case "meteodb":
			port, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			s.dynCfg.MeteoDB.IP = args[0]
			s.dynCfg.MeteoDB.Port = port
			s.dynCfg.MeteoDB.User = args[2]
			s.dynCfg.MeteoDB.Passwd = args[3]
			s.dynCfg.MeteoDB.Base = args[4]
			break
		}
	} else {
		return errors.New("command not found")
	}

	logger.Infof("global apply database configs for base \"%s\"", dev)

	return nil
}
