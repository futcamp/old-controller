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

package data

import "sync"

type MeteoData struct {
	temperature int
	humidity    int
	pressure    int
	mtxTemp     sync.Mutex
	mtxHum      sync.Mutex
	mtxPres     sync.Mutex
}

// Temp get current temperature value
func (m *MeteoData) Temp() int {
	var value int

	m.mtxTemp.Lock()
	value = m.temperature
	m.mtxTemp.Unlock()

	return value
}

// Humidity get current humidity value
func (m *MeteoData) Humidity() int {
	var value int

	m.mtxHum.Lock()
	value = m.humidity
	m.mtxHum.Unlock()

	return value
}

// Pressure get current pressure value
func (m *MeteoData) Pressure() int {
	var value int

	m.mtxPres.Lock()
	value = m.pressure
	m.mtxPres.Unlock()

	return value
}

// SetTemp set new temperature value
func (m *MeteoData) SetTemp(value int) {
	m.mtxTemp.Lock()
	m.temperature = value
	m.mtxTemp.Unlock()
}

// SetHumidity set new humidity value
func (m *MeteoData) SetHumidity(value int) {
	m.mtxHum.Lock()
	m.humidity = value
	m.mtxHum.Unlock()
}

// SetPressure set new pressure value
func (m *MeteoData) SetPressure(value int) {
	m.mtxPres.Lock()
	m.pressure = value
	m.mtxPres.Unlock()
}