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

package meteo

import (
	"time"

	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

const (
	taskDelay = 1
)

// meteoTask meteo task struct
type MeteoTask struct {
	meteo           *MeteoStation
	meteoDB         *MeteoDatabase
	meteoLCD        *MeteoDisplays
	dynCfg          *configs.DynamicConfigs
	reqTimer        *time.Timer
	displaysCounter int
	sensorsCounter  int
	databaseCounter int
	lastHour        int
}

// NewMeteoTask make new struct
func NewMeteoTask(meteo *MeteoStation, mdb *MeteoDatabase, dc *configs.DynamicConfigs,
	mlcd *MeteoDisplays) *MeteoTask {
	return &MeteoTask{
		meteo:           meteo,
		meteoDB:         mdb,
		dynCfg:          dc,
		meteoLCD:        mlcd,
		displaysCounter: 0,
		sensorsCounter:  0,
		databaseCounter: 0,
		lastHour:        -1,
	}
}

// TaskHandler process timer loop
func (m *MeteoTask) TaskHandler() {
	for {
		<-m.reqTimer.C
		m.displaysCounter++
		m.sensorsCounter++
		m.databaseCounter++

		// Get actual data from controller
		if m.sensorsCounter == m.dynCfg.Settings().Timers.MeteoSensorsDelay {
			m.sensorsCounter = 0
			sensors := m.meteo.AllSensors()

			// Get actual data from controllers
			for _, sensor := range sensors {
				err := sensor.SyncMeteoData()
				if err != nil {
					continue
				}
			}
		}

		// Display actual data on LCDs
		if m.displaysCounter == m.dynCfg.Settings().Timers.MeteoDisplayDelay {
			m.displaysCounter = 0
			for _, display := range m.meteoLCD.Displays() {
				for _, sensorName := range *display.Sensors() {
					sensor := m.meteo.Sensor(sensorName)
					data := sensor.MeteoData()
					err := display.DisplayMeteo(sensorName, data.Temp, data.Humidity, data.Pressure)
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
				err := m.meteoDB.Load(m.dynCfg.Settings().MeteoDB.FileName)
				if err != nil {
					logger.Errorf("Fail to load %s database", "meteo")
					logger.Error(err.Error())
					continue
				}

				for _, sensor := range m.meteo.AllSensors() {
					mdata := sensor.MeteoData()

					data := &MeteoDBData{
						Sensor:   sensor.Name,
						Temp:     mdata.Temp,
						Humidity: mdata.Humidity,
						Pressure: mdata.Pressure,
					}

					err = m.meteoDB.AddMeteoData(data)
					if err != nil {
						logger.Errorf("Fail to add to database data from sensor %s",
							sensor.Name)
						logger.Error(sensor.Name, " ", err.Error())
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
