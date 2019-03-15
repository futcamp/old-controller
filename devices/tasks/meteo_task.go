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

package tasks

import (
	"time"

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/devices/db"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

const (
	taskDelay        = 1
	MaxReadDelay     = 10
	SensorErrorValue = -255
)

// meteoTask meteo task struct
type MeteoTask struct {
	meteo           *devices.MeteoStation
	meteoDB         *db.MeteoDatabase
	dynCfg          *configs.DynamicConfigs
	reqTimer        *time.Timer
	displays        *devices.MeteoDisplay
	sensorsCounter  int
	databaseCounter int
	displayCounter  int
	lastHour        int
}

// NewMeteoTask make new struct
func NewMeteoTask(meteo *devices.MeteoStation, mdb *db.MeteoDatabase,
	dc *configs.DynamicConfigs, disp *devices.MeteoDisplay) *MeteoTask {
	return &MeteoTask{
		meteo:           meteo,
		meteoDB:         mdb,
		dynCfg:          dc,
		displays:        disp,
		sensorsCounter:  0,
		databaseCounter: 0,
		displayCounter:  0,
		lastHour:        -1,
	}
}

// TaskHandler process timer loop
func (m *MeteoTask) TaskHandler() {
	for {
		<-m.reqTimer.C
		m.sensorsCounter++
		m.databaseCounter++
		m.displayCounter++

		// Get actual data from controller
		if m.sensorsCounter == m.dynCfg.Settings().Timers.MeteoSensorsDelay {
			m.sensorsCounter = 0
			mods := m.meteo.AllModules()

			// Get actual data from controllers
			for _, mod := range mods {
				err := mod.SyncData()
				if err != nil {
					mod.SetError()

					if mod.Errors() == 1 {
						logger.Errorf("Meteo fail to read meteo data from sensor \"%s\"", mod.Name())
					}

					if mod.Errors() > MaxReadDelay {
						mod.SetErrorValues()
						mod.ResetErrors()
					}
					continue
				} else {
					mod.ResetErrors()
				}
			}
		}

		// Display actual data on LCDs
		if m.displayCounter == m.dynCfg.Settings().Timers.DisplayDelay {
			m.displayCounter = 0
			for _, display := range m.displays.Displays() {
				for _, modName := range *display.Sensors() {
					mod := m.meteo.Module(modName)
					err := display.SyncData(modName, mod.Temp(), mod.Humidity(), mod.Pressure())
					if err != nil {
						continue
					}
				}
			}
		}

		// Save meteo data to database
		if m.databaseCounter == m.dynCfg.Settings().Timers.MeteoDBDelay {
			m.databaseCounter = 0
			hour := time.Now().Hour()

			if hour != m.lastHour {
				mdb := m.dynCfg.Settings().MeteoDB
				err := m.meteoDB.Connect(mdb.IP, mdb.User, mdb.Passwd, mdb.Base)
				if err != nil {
					logger.Errorf("Fail to load %s database", "meteo")
					logger.Error(err.Error())
					continue
				}

				for _, mod := range m.meteo.AllModules() {
					data := &db.MeteoDBData{
						Sensor:   mod.Name(),
						Temp:     mod.Temp(),
						Humidity: mod.Humidity(),
						Pressure: mod.Pressure(),
					}

					err = m.meteoDB.AddMeteoData(data)
					if err != nil {
						logger.Errorf("Fail to add to database data from sensor %s", mod.Name())
					}
				}

				m.meteoDB.Close()
				m.lastHour = hour
			}
		}

		m.reqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (m *MeteoTask) Start() {
	m.reqTimer = time.NewTimer(taskDelay * time.Second)
	m.TaskHandler()
}
