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
	dynCfg          *configs.DynamicConfigs
	reqTimer        *time.Timer
	displays        *MeteoDisplays
	sensorsCounter  int
	databaseCounter int
	displayCounter  int
	lastHour        int
}

// NewMeteoTask make new struct
func NewMeteoTask(meteo *MeteoStation, mdb *MeteoDatabase, dc *configs.DynamicConfigs,
	disp *MeteoDisplays) *MeteoTask {
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
			sensors := m.meteo.AllSensors()

			// Get actual data from controllers
			for _, sensor := range sensors {
				err := sensor.SyncMeteoData()
				if err != nil {
					sensor.Errors++

					if sensor.Errors == 1 {
						logger.Errorf("Fail to read meteo data from sensor \"%s\"", sensor.Name)
						logger.Error(err.Error())
					}

					if sensor.Errors > MaxReadDelay {
						errData := MeteoData{
							Temp:     SensorErrorValue,
							Humidity: SensorErrorValue,
							Pressure: SensorErrorValue,
						}

						sensor.SetMeteoData(&errData)
						sensor.Errors = 0
					}
					continue
				} else {
					sensor.Errors = 0
				}
			}
		}

		// Display actual data on LCDs
		if m.displayCounter == m.dynCfg.Settings().Timers.DisplayDelay {
			m.displayCounter = 0
			for _, display := range m.displays.Displays() {
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
				db := m.dynCfg.Settings().MeteoDB
				err := m.meteoDB.Connect(db.IP, db.User, db.Passwd, db.Base)
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
