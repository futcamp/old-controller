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
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/devices/modules"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/notifier"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

type StartupMods struct {
	cfg        *configs.Configs
	io         *StartupIO
	meteo      *devices.MeteoStation
	dynCfg     *configs.DynamicConfigs
	notify     *notifier.Notifier
	devMonitor *monitoring.DeviceMonitor
	humCtrl    *devices.HumControl
	meteoLCD   *devices.MeteoDisplay
	tempCtrl   *devices.TempControl
	light      *devices.Light
	motion     *devices.Motion
}

// NewStartupMods make new struct
func NewStartupMods(io *StartupIO, dc *configs.DynamicConfigs,
	meteo *devices.MeteoStation, ntf *notifier.Notifier,
	mon *monitoring.DeviceMonitor, cfg *configs.Configs,
	hctrl *devices.HumControl, mlcd *devices.MeteoDisplay,
	tctrl *devices.TempControl, lgh *devices.Light,
	mot *devices.Motion) *StartupMods {
	return &StartupMods{
		io:         io,
		dynCfg:     dc,
		meteo:      meteo,
		notify:     ntf,
		devMonitor: mon,
		cfg:        cfg,
		humCtrl:    hctrl,
		meteoLCD:   mlcd,
		tempCtrl:   tctrl,
		light:      lgh,
		motion:     mot,
	}
}

// LoadFromFile loading configs from file
func (s *StartupMods) LoadFromFile(fileName string) error {
	err := s.io.LoadCommands(fileName, func(cmd *StartupCmd) {
		s.applyConfigs(cmd.Module, cmd.Command, cmd.Device, cmd.Args)
	})
	return err
}

// ExecModCommand exec mod command
func (s *StartupMods) ExecModCommand(fileName string, module string, cmd string, dev string, args []string) error {

	// Add command to command list
	s.io.AddCommand(module, cmd, dev, args)

	// Apply current command to application
	s.applyConfigs(module, cmd, dev, args)

	return nil
}

// AddModCommand add new mod command
func (s *StartupMods) AddModCommand(fileName string, module string, cmd string, dev string, args []string) error {

	// Add command to command list
	s.io.AddCommand(module, cmd, dev, args)

	return nil
}

// SaveCommands save all commands to startup-configs file
func (s *StartupMods) SaveCommands(fileName string) error {
	// Save command to file
	err := s.io.SaveCommands(fileName)
	if err != nil {
		return err
	}
	return nil
}

// DeleteModCommand delete command from startup-configs
func (s *StartupMods) DeleteModCommand(fileName string, module string, cmd string, dev string) error {
	// Delete command from command list
	s.io.DeleteCommand(module, cmd, dev)

	// Delete structures from storages
	if cmd[0] == 'a' && cmd[1] == 'd' && cmd[2] == 'd' {
		switch module {
		case "meteo":
			if s.cfg.Settings().Modules.Meteo {
				s.meteo.DeleteModule(dev)
			}
			break

		case "humctrl":
			if s.cfg.Settings().Modules.Humctrl {
				s.humCtrl.DeleteModule(dev)
			}
			break

		case "tempctrl":
			if s.cfg.Settings().Modules.Tempctrl {
				s.tempCtrl.DeleteModule(dev)
			}
			break

		case "light":
			if s.cfg.Settings().Modules.Light {
				s.light.DeleteModule(dev)
			}
			break

		case "motion":
			if s.cfg.Settings().Modules.Motion {
				s.motion.DeleteModule(dev)
			}
			break

		case "display":
			s.meteoLCD.DeleteDisplay(dev)
			break

		case "monitor":
			s.devMonitor.DeleteDevice(dev)
			break
		}
	}

	return nil
}

// GenHashStr generate hash string
func (s *StartupMods) GenHashStr(login string, passwd string) string {
	hash := ""
	tmp := login + passwd
	byteData, _ := json.Marshal(tmp)

	h := sha512.New()
	h.Write(byteData)
	hOut := h.Sum(nil)

	for i := 0; i < len(hOut); i++ {
		hash += fmt.Sprintf("%x", hOut[i])
	}

	return hash
}

// applyConfigs general function for applying configs
func (s *StartupMods) applyConfigs(module string, cmd string, dev string, args []string) {
	// Apply command to application
	switch module {
	case "meteo":
		if s.cfg.Settings().Modules.Meteo {
			s.applyMeteoCfg(cmd, dev, args)
		}
		break

	case "humctrl":
		if s.cfg.Settings().Modules.Humctrl {
			s.applyHumCtrlCfg(cmd, dev, args)
		}
		break

	case "tempctrl":
		if s.cfg.Settings().Modules.Tempctrl {
			s.applyTempCtrlCfg(cmd, dev, args)
		}
		break

	case "light":
		if s.cfg.Settings().Modules.Light {
			s.applyLightCfg(cmd, dev, args)
		}
		break

	case "motion":
		if s.cfg.Settings().Modules.Motion {
			s.applyMotionCfg(cmd, dev, args)
		}
		break

	case "db":
		s.applyMeteoDBCfg(cmd, dev, args)
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

	case "display":
		s.applyMeteoLCDCfg(cmd, dev, args)
		break

	case "rcli":
		s.applyRCliCfg(cmd, dev, args)
		break
	}
}

// applyMeteoLCDCfg apply commands for meteo displays
func (s *StartupMods) applyMeteoLCDCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		lcd := modules.NewDisplayModule(dev)
		s.meteoLCD.AddDisplay(dev, lcd)
		logger.Infof("Display add new device \"%s\"", dev)
		break

	case "ip":
		lcd := s.meteoLCD.Display(dev)
		lcd.SetIP(args[0])
		logger.Infof("Display set ip address \"%s\" for device \"%s\"", args[0], dev)
		break

	case "sensors":
		lcd := s.meteoLCD.Display(dev)

		for _, sensor := range args {
			lcd.AddSensor(sensor)
			logger.Infof("Display add displaying sensor \"%s\" for device \"%s\"", sensor, dev)
		}
		break
	}

	return nil
}

// applyMeteoCfg apply commands for meteo mod
func (s *StartupMods) applyMeteoCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		mod := modules.NewMeteoModule(dev)
		s.meteo.AddModule(dev, mod)
		logger.Infof("Meteo add new device \"%s\"", dev)
		break

	case "ip":
		mod := s.meteo.Module(dev)
		mod.SetIP(args[0])
		logger.Infof("Meteo set ip address \"%s\" for device \"%s\"", args[0], dev)
		break

	case "type":
		mod := s.meteo.Module(dev)
		mod.SetType(args[0])
		logger.Infof("Meteo set sensor type \"%s\" for device \"%s\"", args[0], dev)
		break

	case "channel":
		ch, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		mod := s.meteo.Module(dev)
		mod.SetChannel(ch)
		logger.Infof("Meteo set sensor channel \"%d\" for device \"%s\"", ch, dev)
		break

	case "temp-delta":
		td, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		mod := s.meteo.Module(dev)
		mod.SetDelta(td)
		logger.Infof("Meteo set sensor temperature delta \"%d\" for device \"%s\"", td, dev)
		break
	}

	return nil
}

// applyHumCtrlCfg apply commands for humctrl mod
func (s *StartupMods) applyHumCtrlCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		mod := modules.NewHumCtrlModule(dev, s.dynCfg)
		s.humCtrl.AddModule(dev, mod)
		logger.Infof("HumControl add new device \"%s\"", dev)
		break

	case "ip":
		mod := s.humCtrl.Module(dev)
		mod.SetIP(args[0])
		logger.Infof("HumControl set ip address \"%s\" for device \"%s\"", mod.IP(), dev)
		break

	case "sensor":
		mod := s.humCtrl.Module(dev)
		mod.SetSensor(args[0])
		logger.Infof("HumControl set sensor \"%s\" for device \"%s\"", mod.Sensor(), dev)
		break

	case "status":
		mod := s.humCtrl.Module(dev)
		if args[0] == "on" {
			mod.SetStatus(true)
		} else {
			mod.SetStatus(false)
		}
		logger.Infof("HumControl set status \"%s\" for device \"%s\"", args[0], dev)
		break

	case "threshold":
		threshold, err := strconv.Atoi(args[0])
		if err != nil {
			logger.Infof("HumControl fail to convert threshold value for device \"%s\"", dev)
			break
		}

		mod := s.humCtrl.Module(dev)
		mod.SetThreshold(threshold)
		logger.Infof("HumControl set threshold \"%d\" for device \"%s\"", threshold, dev)
		break
	}

	return nil
}

// applyLightCfg apply commands for light mod
func (s *StartupMods) applyLightCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		mod := modules.NewLightModule(dev, s.dynCfg)
		s.light.AddModule(dev, mod)
		logger.Infof("Light add new device \"%s\"", dev)
		break

	case "ip":
		mod := s.light.Module(dev)
		mod.SetIP(args[0])
		logger.Infof("Light set ip address \"%s\" for device \"%s\"", mod.IP(), dev)
		break

	case "channel":
		ch, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		mod := s.light.Module(dev)
		mod.SetChannel(ch)
		logger.Infof("Light set channel \"%d\" for device \"%s\"", ch, dev)
		break

	case "status":
		mod := s.light.Module(dev)
		if args[0] == "on" {
			mod.SetStatus(true)
		} else {
			mod.SetStatus(false)
		}
		logger.Infof("Light set status \"%s\" for device \"%s\"", args[0], dev)
		break
	}

	return nil
}

// applyMotionCfg apply commands for motion mod
func (s *StartupMods) applyMotionCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		mod := modules.NewMotionModule(dev)
		s.motion.AddModule(dev, mod)
		logger.Infof("Motion add new device \"%s\"", dev)
		break

	case "ip":
		mod := s.motion.Module(dev)
		mod.SetIP(args[0])
		logger.Infof("Motion set ip address \"%s\" for device \"%s\"", mod.IP(), dev)
		break

	case "delay":
		delay, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		mod := s.motion.Module(dev)
		mod.SetDelay(delay)
		logger.Infof("Motion set delay \"%d\" for device \"%s\"", delay, dev)
		break

	case "lamps":
		mod := s.motion.Module(dev)

		for _, lamp := range args {
			mod.AddLamp(lamp)
			logger.Infof("Motion add motion lamp \"%s\" for device \"%s\"", lamp, dev)
		}
		break

	}

	return nil
}

// applyTempCtrlCfg apply commands for tempctrl mod
func (s *StartupMods) applyTempCtrlCfg(cmd string, dev string, args []string) error {
	switch cmd {
	case "add-device":
		mod := modules.NewTempCtrlModule(dev, s.dynCfg)
		s.tempCtrl.AddModule(dev, mod)
		logger.Infof("TempControl add new device \"%s\"", dev)
		break

	case "ip":
		mod := s.tempCtrl.Module(dev)
		mod.SetIP(args[0])
		logger.Infof("TempControl set ip address \"%s\" for device \"%s\"", mod.IP(), dev)
		break

	case "sensor":
		mod := s.tempCtrl.Module(dev)
		mod.SetSensor(args[0])
		logger.Infof("TempControl set sensor \"%s\" for device \"%s\"", mod.Sensor(), dev)
		break

	case "status":
		mod := s.tempCtrl.Module(dev)
		if args[0] == "on" {
			mod.SetStatus(true)
		} else {
			mod.SetStatus(false)
		}
		logger.Infof("TempControl set status \"%s\" for device \"%s\"", args[0], dev)
		break

	case "threshold":
		threshold, err := strconv.Atoi(args[0])
		if err != nil {
			logger.Infof("TempControl fail to convert threshold value for device \"%s\"", dev)
			break
		}

		mod := s.tempCtrl.Module(dev)
		mod.SetThreshold(threshold)
		logger.Infof("TempControl set threshold \"%d\" for device \"%s\"", threshold, dev)
		break
	}

	return nil
}

// applyTimerCfg apply commands for timers
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

		case "meteo-db":
			s.dynCfg.Settings().Timers.MeteoDBDelay = delay
			break

		case "display":
			s.dynCfg.Settings().Timers.DisplayDelay = delay
			break

		case "monitor":
			s.dynCfg.Settings().Timers.MonitorDelay = delay
			break
		}
	} else {
		return errors.New("command not found")
	}

	logger.Infof("Timers apply timer delay \"%d\" for timer \"%s\"", delay, dev)

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
				mod := s.meteo.Module(device)
				dev := monitoring.NewDevice(mod.Name(), "meteo", mod.IP())
				s.devMonitor.AddDevice(dev)
				logger.Infof("Monitor add new device from mod \"%s\" sensor \"%s\" for monitor \"%s\"",
					args[0], mod.Name(), dev.Name())
			}
			break

		case "humctrl":
			for i, device := range args {
				if i == 0 {
					continue
				}
				mod := s.humCtrl.Module(device)
				dev := monitoring.NewDevice(mod.Name(), "humctrl", mod.IP())
				s.devMonitor.AddDevice(dev)
				logger.Infof("Monitor add new device from mod \"%s\" mod \"%s\" for monitor \"%s\"",
					args[0], mod.Name(), dev.Name())
			}
			break

		case "tempctrl":
			for i, device := range args {
				if i == 0 {
					continue
				}
				mod := s.tempCtrl.Module(device)
				dev := monitoring.NewDevice(mod.Name(), "tempctrl", mod.IP())
				s.devMonitor.AddDevice(dev)
				logger.Infof("Monitor add new device from mod \"%s\" mod \"%s\" for monitor \"%s\"",
					args[0], mod.Name(), dev.Name())
			}
			break

		case "light":
			for i, device := range args {
				if i == 0 {
					continue
				}
				mod := s.light.Module(device)
				dev := monitoring.NewDevice(mod.Name(), "light", mod.IP())
				s.devMonitor.AddDevice(dev)
				logger.Infof("Monitor add new device from mod \"%s\" mod \"%s\" for monitor \"%s\"",
					args[0], mod.Name(), dev.Name())
			}
			break

		case "motion":
			for i, device := range args {
				if i == 0 {
					continue
				}
				mod := s.motion.Module(device)
				dev := monitoring.NewDevice(mod.Name(), "motion", mod.IP())
				s.devMonitor.AddDevice(dev)
				logger.Infof("Monitor add new device from mod \"%s\" mod \"%s\" for monitor \"%s\"",
					args[0], mod.Name(), dev.Name())
			}
			break

		case "display":
			for i, device := range args {
				if i == 0 {
					continue
				}
				lcd := s.meteoLCD.Display(device)
				dev := monitoring.NewDevice(lcd.Name(), "display", lcd.IP())
				s.devMonitor.AddDevice(dev)
				logger.Infof("Monitor add new device from mod \"%s\" display \"%s\" for monitor \"%s\"",
					args[0], lcd.Name(), dev.Name())
			}
			break
		}
		break
	}

	return nil
}

// applyRCliCfg apply commands for RemoteCLI monitor
func (s *StartupMods) applyRCliCfg(cmd string, dev string, args []string) error {
	if cmd == "add-user" {
		switch dev {
		case "main":
			if len(args) < 3 {
				return errors.New("wrong args count")
			}

			login := args[1]
			passwd := args[3]
			s.dynCfg.Settings().RCli.UserHash = s.GenHashStr(login, passwd)
			break
		}
	} else {
		return errors.New("command not found")
	}

	logger.Infof("RemoteCLI add user \"%s\" for remote CLI \"%s\"", args[1], dev)

	return nil
}

// applyMeteoDBCfg apply commands for MeteoDB mod
func (s *StartupMods) applyMeteoDBCfg(cmd string, dev string, args []string) error {
	if cmd == "add-base" {
		switch dev {
		case "meteodb":
			if len(args) < 10 {
				return errors.New("wrong args count")
			}

			port, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}

			s.dynCfg.Settings().MeteoDB.IP = args[1]
			s.dynCfg.Settings().MeteoDB.Port = port
			s.dynCfg.Settings().MeteoDB.User = args[5]
			s.dynCfg.Settings().MeteoDB.Passwd = args[7]
			s.dynCfg.Settings().MeteoDB.Base = args[9]
			break
		}
	} else {
		return errors.New("MeteoDB command not found")
	}

	logger.Infof("MeteoDB apply database configs for base \"%s\"", dev)

	return nil
}
