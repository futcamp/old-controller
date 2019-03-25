/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2019 Sergey Denisov.
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

package modules

import (
	"github.com/futcamp/controller/devices/data"
	"github.com/futcamp/controller/devices/hardware"
)

const (
	SensorErrorValue = -255
)

type MeteoController interface {
	Name() string
	IP() string
	Type() string
	Errors() int
	SetError()
	ResetErrors()
	SetDelta(value int)
	SetType(value string)
	SetIP(value string)
	SetChannel(value int)
	Temp() int
	Humidity() int
	Pressure() int
	SetErrorValues()
	SyncData() error
}

type MeteoModule struct {
	name     string
	ip       string
	sensType string
	channel  int
	delta    int
	err      int
	data     data.MeteoData
}

// NewMeteoModule make new struct
func NewMeteoModule(name string) *MeteoModule {
	return &MeteoModule{
		name: name,
	}
}

//
// Simple data getters and setters
//

// Name get module name
func (m *MeteoModule) Name() string {
	return m.name
}

// IP get module ip
func (m *MeteoModule) IP() string {
	return m.ip
}

// Type get module type
func (m *MeteoModule) Type() string {
	return m.sensType
}

// Errors get errors count
func (m *MeteoModule) Errors() int {
	return m.err
}

// SetError incremet errors value
func (m *MeteoModule) SetError() {
	m.err++
}

// ResetErrors reset errors value
func (m *MeteoModule) ResetErrors() {
	m.err = 0
}

// SetDelta set new delta value
func (m *MeteoModule) SetDelta(value int) {
	m.delta = value
}

// SetType set new sensor type
func (m *MeteoModule) SetType(value string) {
	m.sensType = value
}

// SetIP set ip address
func (m *MeteoModule) SetIP(value string) {
	m.ip = value
}

// SetChannel set channel number
func (m *MeteoModule) SetChannel(value int) {
	m.channel = value
}

//
// Nested getters and setters
//

// Temp get current temperature value
func (m *MeteoModule) Temp() int {
	return m.data.Temp()
}

// Humidity get current humidity value
func (m *MeteoModule) Humidity() int {
	return m.data.Humidity()
}

// Pressure get current pressure value
func (m *MeteoModule) Pressure() int {
	return m.data.Pressure()
}

// SetErrorValues set error values for sensor
func (m *MeteoModule) SetErrorValues() {
	m.data.SetTemp(SensorErrorValue)
	m.data.SetHumidity(SensorErrorValue)
	m.data.SetPressure(SensorErrorValue)
}

//
// Other functional
//

// Sync data get actual meteo data from remote module
func (m *MeteoModule) SyncData() error {
	meteoData, err := hardware.HdkSyncMeteoData(m.ip, m.channel, m.sensType)
	if err != nil {
		return err
	}

	m.data.SetTemp(meteoData.Temp + m.delta)
	m.data.SetHumidity(meteoData.Humidity)
	m.data.SetPressure(meteoData.Pressure)

	return nil
}
