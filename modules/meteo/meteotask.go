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

package meteo

import (
	"time"

	"github.com/futcamp/controller/utils/configs"
	"github.com/google/logger"
)

const (
	taskDelay = 1
)

// MeteoTask meteo task struct
type MeteoTask struct {
	Cfg             *configs.MeteoConfigs
	Meteo           *MeteoStation
	MeteoDB         *MeteoDatabase
	ReqTimer        *time.Timer
	DisplaysCounter int
	SensorsCounter  int
	DatabaseCounter int
	LastHour        int
}

// NewMeteoTask make new struct
func NewMeteoTask(conf *configs.MeteoConfigs, meteo *MeteoStation,
	mdb *MeteoDatabase) *MeteoTask {
	return &MeteoTask{
		Cfg:             conf,
		Meteo:           meteo,
		MeteoDB:         mdb,
		DisplaysCounter: 0,
		SensorsCounter:  0,
		DatabaseCounter: 0,
		LastHour:        -1,
	}
}

// TaskHandler process timer loop
func (m *MeteoTask) TaskHandler() {
	for {
		<-m.ReqTimer.C
		m.DisplaysCounter++
		m.SensorsCounter++
		m.DatabaseCounter++

		// Get actual data from controller
		if m.SensorsCounter == m.Cfg.Settings().Delays.Sensors {
			m.SensorsCounter = 0
			m.Meteo.SyncData()
		}

		// Display actual data on LCDs
		if m.DisplaysCounter == m.Cfg.Settings().Delays.Displays {
			m.DisplaysCounter = 0
			for _, display := range m.Cfg.Settings().Displays {
				if display.Enable {
					for _, sensor := range display.Sensors {
						data, err := m.Meteo.Sensor(sensor)
						if err != nil {
							logger.Errorf("Sensor %s not found", sensor)
							continue
						}

						ctrlSensor := &CtrlMeteoData{
							Temp:     data.Temp,
							Humidity: data.Humidity,
							Pressure: data.Pressure,
						}

						// Send data to controller
						ctrl := NewWiFiController("", display.IP, 0)
						err = ctrl.DisplayMeteoData(sensor, ctrlSensor)
						if err != nil {
							logger.Errorf("Fail to display data on %s from sensor %s",
								display.Name, sensor)
						}
					}

				}
			}
		}

		// Save meteo data to database
		if m.DatabaseCounter == m.Cfg.Settings().Delays.Database {
			m.DatabaseCounter = 0
			hour := time.Now().Hour()

			if hour != m.LastHour {
				err := m.MeteoDB.Load()
				if err != nil {
					logger.Errorf("Fail to load %s database", "Meteo")
					continue
				}

				for _, sensor := range m.Meteo.AllSensors() {
					data := &MeteoDBData{
						Sensor:   sensor.Name,
						Temp:     sensor.Temp,
						Humidity: sensor.Humidity,
						Pressure: sensor.Pressure,
						Altitude: sensor.Pressure,
					}
					err = m.MeteoDB.AddMeteoData(data)
					if err != nil {
						logger.Errorf("Fail to add to database data from sensor %s")
					}
				}

				m.MeteoDB.Unload()
				m.LastHour = hour
			}
		}

		m.ReqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (m *MeteoTask) Start() {
	m.ReqTimer = time.NewTimer(taskDelay * time.Second)
	m.TaskHandler()
}
