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

	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/notifier"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

type StartupMods struct {
	cfg        *StartupIO
	meteo      *meteo.MeteoStation
	dynCfg     *configs.DynamicConfigs
	meteoLCD   *meteo.MeteoDisplays
	notify     *notifier.Notifier
	devMonitor *monitoring.DeviceMonitor
}

// NewStartupMods make new struct
func NewStartupMods(cfg *StartupIO, dc *configs.DynamicConfigs,
	meteo *meteo.MeteoStation, mlcd *meteo.MeteoDisplays,
	ntf *notifier.Notifier, mon *monitoring.DeviceMonitor) *StartupMods {
	return &StartupMods{
		cfg:        cfg,
		dynCfg:     dc,
		meteo:      meteo,
		meteoLCD:   mlcd,
		notify:     ntf,
		devMonitor: mon,
	}
}

// LoadFromFile loading configs from file
func (s *StartupMods) LoadFromFile(fileName string) error {
	err := s.cfg.LoadCommands(fileName, func(cmd *StartupCmd) {
		s.applyConfigs(cmd.Module, cmd.Command, cmd.Device, cmd.Args)
	})
	return err
}

// ExecModCommand exec module command
func (s *StartupMods) ExecModCommand(fileName string, module string, cmd string, dev string, args []string) error {

	// Add command to command list
	s.cfg.AddCommand(module, cmd, dev, args)

	// Apply current command to application
	s.applyConfigs(module, cmd, dev, args)

	return nil
}

// SaveCommands save all commands to startup-configs file
func (s *StartupMods) SaveCommands(fileName string) error {
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

	return nil
}

// applyConfigs general function for applying configs
func (s *StartupMods) applyConfigs(module string, cmd string, dev string, args []string) {
	// Apply command to application
	switch module {
	case "meteo":
		s.applyMeteoCfg(cmd, dev, args)
		break

	case "display":
		s.applyMeteoLCDCfg(cmd, dev, args)
		break

	case "db":
		s.applyDBCfg(cmd, dev, args)
		break

	case "timer":
		s.applyTimerCfg(cmd, dev, args)
		break

	case "notify":
		s.applyNotifyCfg(cmd, dev, args)
		break

	case "monitor":
		s.applyMonitorCfg(cmd, dev, args)
		break
	}
}

// applyMeteoCfg apply commands for meteo module
func (s *StartupMods) applyMeteoCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		sensor := s.meteo.NewMeteoSensor(dev, "", "", 0)
		s.meteo.AddSensor(dev, sensor)
		logger.Infof("Meteo add new device \"%s\"", dev)
		break

	case "ip":
		sensor := s.meteo.Sensor(dev)
		sensor.IP = args[0]
		logger.Infof("Meteo set ip address \"%s\" for device \"%s\"", sensor.IP, dev)
		break

	case "type":
		sensor := s.meteo.Sensor(dev)
		sensor.Type = args[0]
		logger.Infof("Meteo set sensor type \"%s\" for device \"%s\"", sensor.Type, dev)
		break

	case "channel":
		ch, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		sensor := s.meteo.Sensor(dev)
		sensor.Channel = ch
		logger.Infof("Meteo set sensor channel \"%d\" for device \"%s\"", sensor.Channel, dev)
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
		logger.Infof("Display add new device \"%s\"", dev)
		break

	case "ip":
		lcd := s.meteoLCD.Display(dev)
		lcd.IP = args[0]
		logger.Infof("Display set ip address \"%s\" for device \"%s\"", lcd.IP, dev)
		break

	case "sensors":
		lcd := s.meteoLCD.Display(dev)

		for _, sensor := range args {
			lcd.AddDisplayingSensor(sensor)
			logger.Infof("Display add displaying sensor \"%s\" for device \"%s\"", sensor, dev)
		}
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
			s.dynCfg.Settings().Timers.MeteoSensorsDelay = delay
			break

		case "meteo-lcd":
			s.dynCfg.Settings().Timers.MeteoDisplayDelay = delay
			break

		case "meteo-db":
			s.dynCfg.Settings().Timers.MeteoDBDelay = delay
			break
		}
	} else {
		return errors.New("command not found")
	}

	logger.Infof("Global apply timer delay \"%d\" for timer \"%s\"", delay, dev)

	return nil
}

// applyNotifyCfg apply commands for notifier
func (s *StartupMods) applyNotifyCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-server":
		s.notify.SetName(dev)
		logger.Infof("Notifier add new server \"%s\"", dev)
		break

	case "api-key":
		s.notify.SetApiKey(args[0])
		logger.Infof("Notifier set api-key \"%s\" for server \"%s\"", args[0], dev)
		break

	case "chat-id":
		for _, chatID := range args {
			s.notify.AddChatID(chatID)
			logger.Infof("Notifier add ChatID \"%s\" for server \"%s\"", chatID, dev)
		}
		break
	}

	return nil
}

// applyMonitorCfg apply commands for devices monitor
func (s *StartupMods) applyMonitorCfg(cmd string, dev string, args []string) error {
	devs := ""

	switch cmd {
	case "add-monitor":
		s.devMonitor.SetName(dev)
		logger.Infof("Monitor add new monitor \"%s\"", dev)
		break

	case "device":
		switch args[0] {
		case "meteo":
			for i, device := range args {
				if i == 0 {
					continue
				}
				devs += device + " "
				sensor := s.meteo.Sensor(device)
				s.devMonitor.AddDevice(sensor.Name, "meteo", sensor.IP)
			}
			break

		case "display":
			for i, device := range args {
				if i == 0 {
					continue
				}
				devs += device + " "
				lcd := s.meteoLCD.Display(device)
				s.devMonitor.AddDevice(lcd.Name, "display", lcd.IP)
			}
			break
		}

		logger.Infof("Monitor add new device from module \"%s\" : \"%s\" for monitor \"%s\"", args[0], devs, dev)
		break
	}

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
			s.dynCfg.Settings().MeteoDB.IP = args[0]
			s.dynCfg.Settings().MeteoDB.Port = port
			s.dynCfg.Settings().MeteoDB.User = args[2]
			s.dynCfg.Settings().MeteoDB.Passwd = args[3]
			s.dynCfg.Settings().MeteoDB.Base = args[4]
			break
		}
	} else {
		return errors.New("command not found")
	}

	logger.Infof("Global apply database configs for base \"%s\"", dev)

	return nil
}
