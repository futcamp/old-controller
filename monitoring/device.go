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

package monitoring

type MonitoringDevice interface {
	Name() string
	Type() string
	IP() string
	Status() bool
	SetStatus(status bool)
}

type Device struct {
	name    string
	devType string
	ip      string
	data    DeviceData
}

// NewDevice make new struct
func NewDevice(name string, devType string, ip string) *Device {
	return &Device{
		name:    name,
		devType: devType,
		ip:      ip,
	}
}

//
// Simple data getters and setters
//

// Name get device name
func (d *Device) Name() string {
	return d.name
}

// Type get device type
func (d *Device) Type() string {
	return d.devType
}

// IP get device ip address
func (d *Device) IP() string {
	return d.ip
}

// Status get current device status
func (d *Device) Status() bool {
	return d.data.Status()
}

//
// Nested getters and setters
//

// Status set new device status
func (d *Device) SetStatus(status bool) {
	d.data.SetStatus(status)
}
