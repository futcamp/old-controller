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

import "github.com/futcamp/controller/devices/hardware"

type Indicator interface {
	Name() string
	IP() string
	Sensors() *[]string
	SetIP(ip string)
	AddSensor(sensor string)
	SyncData(sensor string, temp int, hum int, pres int) error
}

type DisplayModule struct {
	name    string
	ip      string
	sensors []string
}

// NewDisplayModule make new struct
func NewDisplayModule(name string) *DisplayModule {
	return &DisplayModule{
		name: name,
	}
}

//
// Simple data getters and setters
//

// Name get current display name
func (d *DisplayModule) Name() string {
	return d.name
}

// IP get current display ip address
func (d *DisplayModule) IP() string {
	return d.ip
}

// Name get current display sensors list
func (d *DisplayModule) Sensors() *[]string {
	return &d.sensors
}

// SetIP set new display ip address
func (d *DisplayModule) SetIP(ip string) {
	d.ip = ip
}

// AddSensor add new displayed meteo sensor
func (d *DisplayModule) AddSensor(sensor string) {
	d.sensors = append(d.sensors, sensor)
}

//
// Other functional
//

// Sync data get actual meteo data from remote module
func (m *DisplayModule) SyncData(sensor string, temp int, hum int, pres int) error {
	_, err := hardware.HdkSyncDisplayData(m.ip, sensor, temp, hum, pres)
	if err != nil {
		return err
	}

	return nil
}
