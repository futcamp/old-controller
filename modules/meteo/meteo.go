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
	"github.com/pkg/errors"
	"sync"

	"github.com/google/logger"
)

type DisplayedSensor struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Temp     int    `json:"temp"`
	Humidity int    `json:"humidity"`
	Pressure int    `json:"pressure"`
}

type MeteoData struct {
	Temp     int
	Humidity int
	Pressure int
}

type Sensor struct {
	Name    string
	Type    string
	IP      string
	Channel int
	Mtx     sync.Mutex
	Data    MeteoData
	Online  bool
}

type MeteoStation struct {
	Sensors map[string]*Sensor
}

// SetMeteoData set new meteo data to sensor
func (s *Sensor) SetMeteoData(data *MeteoData) {
	s.Mtx.Lock()
	s.Data.Temp = data.Temp
	s.Data.Humidity = data.Humidity
	s.Data.Pressure = data.Pressure
	s.Mtx.Unlock()
}

// MeteoData get meteo data from sensor
func (s *Sensor) MeteoData() MeteoData {
	var data MeteoData

	s.Mtx.Lock()
	data.Temp = s.Data.Temp
	data.Humidity = s.Data.Humidity
	data.Pressure = s.Data.Pressure
	s.Mtx.Unlock()

	return data
}

// NewMeteoStation make new struct
func NewMeteoStation() *MeteoStation {
	m := &MeteoStation{}

	m.Sensors = make(map[string]*Sensor)

	return m
}

// AddSensor add new meteo sensor
func (m *MeteoStation) AddSensor(name string, sType string, ip string, ch int) {
	sensor := &Sensor{
		Name:    name,
		Type:    sType,
		IP:      ip,
		Channel: ch,
		Online:  true,
	}

	m.Sensors[name] = sensor
}

// Sensor get meteo data from sensor
func (m *MeteoStation) Sensor(name string) (DisplayedSensor, error) {
	var dSensor DisplayedSensor

	sensor := m.Sensors[name]
	if sensor == nil {
		return dSensor, errors.New("sensor not found")
	}
	data := sensor.MeteoData()

	dSensor.Name = sensor.Name
	dSensor.Type = sensor.Type
	dSensor.Temp = data.Temp
	dSensor.Humidity = data.Humidity
	dSensor.Pressure = data.Pressure

	return dSensor, nil
}

// AllSensors get all sensors list
func (m *MeteoStation) AllSensors() []DisplayedSensor {
	var sensors []DisplayedSensor

	for _, sensor := range m.Sensors {
		data := sensor.MeteoData()

		dSensor := DisplayedSensor{
			Name:     sensor.Name,
			Type:     sensor.Type,
			Temp:     data.Temp,
			Humidity: data.Humidity,
			Pressure: data.Pressure,
		}

		sensors = append(sensors, dSensor)
	}

	return sensors
}

// SyncData get actual data from all wi-fi sensors
func (m *MeteoStation) SyncData() {
	for _, sensor := range m.Sensors {
		ctrl := NewWiFiController(sensor.Type, sensor.IP, sensor.Channel)
		data, err := ctrl.SyncMeteoData()
		if err != nil {
			if sensor.Online {
				sensor.Online = false
				logger.Errorf("Fail to sync meteo data with sensor %s", sensor.Name)
			}
			continue
		}

		if !sensor.Online {
			sensor.Online = true
			logger.Errorf("Sensor %s is now online", sensor.Name)
		}

		sensor.SetMeteoData(&MeteoData{
			Temp:     data.Temp,
			Humidity: data.Humidity,
			Pressure: data.Pressure,
		})
	}
}
