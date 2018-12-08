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
	"errors"
	"sync"
)


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
	}

	m.Sensors[name] = sensor
}

// Sensor get meteo sensor
func (m *MeteoStation) Sensor(name string) (*Sensor, error) {
	s := m.Sensors[name]

	if s == nil {
		return nil, errors.New("sensor not found")
	}

	return s, nil
}

// AllSensors get all sensors list
func (m *MeteoStation) AllSensors() []*Sensor {
	var sensors []*Sensor

	for _, sensor := range m.Sensors {
		sensors = append(sensors, sensor)
	}

	return sensors
}
