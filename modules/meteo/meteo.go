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
	"sync"
)


type MeteoData struct {
	Temp     int
	Humidity int
	Pressure int
}

type MeteoSensor struct {
	Name    string
	Type    string
	IP      string
	Channel int
	Mtx     sync.Mutex
	Data    MeteoData
}

type MeteoStation struct {
	Sensors map[string]*MeteoSensor
}

// SetMeteoData set new meteo data to sensor
func (s *MeteoSensor) SetMeteoData(data *MeteoData) {
	s.Mtx.Lock()
	s.Data.Temp = data.Temp
	s.Data.Humidity = data.Humidity
	s.Data.Pressure = data.Pressure
	s.Mtx.Unlock()
}

// MeteoData get meteo data from sensor
func (s *MeteoSensor) MeteoData() MeteoData {
	var data MeteoData

	s.Mtx.Lock()
	data.Temp = s.Data.Temp
	data.Humidity = s.Data.Humidity
	data.Pressure = s.Data.Pressure
	s.Mtx.Unlock()

	return data
}

// SyncMeteoData get meteo data from controller
func (s *MeteoSensor) SyncMeteoData() error {
	var mData MeteoData

	ctrl := NewWiFiController(s.Type, s.IP, s.Channel)
	data, err := ctrl.SyncMeteoData()
	if err != nil {
		return err
	}

	mData.Temp = data.Temp
	mData.Humidity = data.Humidity
	mData.Pressure = data.Pressure

	s.SetMeteoData(&mData)

	return nil
}

// NewMeteoStation make new struct
func NewMeteoStation() *MeteoStation {
	sensors := make(map[string]*MeteoSensor)
	return &MeteoStation{
		Sensors: sensors,
	}
}

// NewMeteoSensor make new meteo sensor
func (m *MeteoStation) NewMeteoSensor(name string, sType string, ip string, ch int) *MeteoSensor {
	return &MeteoSensor{
		Name:    name,
		Type:    sType,
		IP:      ip,
		Channel: ch,
	}
}

// AddSensor add new meteo sensor
func (m *MeteoStation) AddSensor(name string, sensor *MeteoSensor) {
	m.Sensors[name] = sensor
}

// MeteoSensor get meteo sensor
func (m *MeteoStation) Sensor(name string) *MeteoSensor {
	return m.Sensors[name]
}

// AllSensors get all sensors list
func (m *MeteoStation) AllSensors() []*MeteoSensor {
	var sensors []*MeteoSensor

	for _, sensor := range m.Sensors {
		sensors = append(sensors, sensor)
	}

	return sensors
}