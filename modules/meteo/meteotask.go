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
	MeteoCfg        *configs.MeteoConfigs
	ReqTimer        *time.Timer
	DisplaysCounter int
	SensorsCounter  int
	DatabaseCounter int
	LastHour        int
}

// NewMeteoTask make new struct
func NewMeteoTask(conf *configs.MeteoConfigs, meteo *MeteoStation,
	mdb *MeteoDatabase, mcfg *configs.MeteoConfigs) *MeteoTask {
	return &MeteoTask{
		Cfg:             conf,
		Meteo:           meteo,
		MeteoDB:         mdb,
		MeteoCfg:        mcfg,
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
			sensors := m.Meteo.AllSensors()

			for _, sensor := range sensors {
				ctrl := NewWiFiController(sensor.Type, sensor.IP, sensor.Channel)
				data, err := ctrl.SyncMeteoData()
				if err != nil {
					continue
				}

				sensor.SetMeteoData(&MeteoData{
					Temp:     data.Temp,
					Humidity: data.Humidity,
					Pressure: data.Pressure,
				})
			}
		}

		// Display actual data on LCDs
		if m.DisplaysCounter == m.Cfg.Settings().Delays.Displays {
			m.DisplaysCounter = 0
			for _, display := range m.Cfg.Settings().Displays {
				if display.Enable {
					for _, sensorName := range display.Sensors {
						sensor, err := m.Meteo.Sensor(sensorName)
						if err != nil {
							logger.Errorf("Sensor %s not found", sensorName)
							continue
						}
						data := sensor.MeteoData()

						ctrlSensor := &CtrlMeteoData{
							Temp:     data.Temp,
							Humidity: data.Humidity,
							Pressure: data.Pressure,
						}

						// Send data to controller
						ctrl := NewWiFiControllerDisplay(display.IP)
						err = ctrl.DisplayMeteoData(sensorName, ctrlSensor)
						if err != nil {
							continue
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
				dbCfg := m.MeteoCfg.Settings().Database
				err := m.MeteoDB.Connect(dbCfg.IP, dbCfg.User, dbCfg.Password, dbCfg.Base)
				if err != nil {
					logger.Errorf("Fail to load %s database", "Meteo")
					logger.Error(err.Error())
					continue
				}

				for _, sensor := range m.Meteo.AllSensors() {
					mdata := sensor.MeteoData()

					data := &MeteoDBData{
						Sensor:   sensor.Name,
						Temp:     mdata.Temp,
						Humidity: mdata.Humidity,
						Pressure: mdata.Pressure,
						Altitude: mdata.Pressure,
					}
					err = m.MeteoDB.AddMeteoData(data)
					if err != nil {
						logger.Errorf("Fail to add to database data from sensor %s",
							sensor.Name)
						logger.Error(sensor.Name, err.Error())
					}
				}

				m.MeteoDB.Close()
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
