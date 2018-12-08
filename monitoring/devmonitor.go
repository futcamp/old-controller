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

package monitoring

import (
	"github.com/google/logger"
)

type Device struct {
	Name   string
	Type   string
	IP     string
	Status bool
}

type DeviceMonitor struct {
	Devices    []*Device
	FirstCheck bool
}

// NewDeviceMonitor make new struct
func NewDeviceMonitor() *DeviceMonitor {
	return &DeviceMonitor{
		FirstCheck: true,
	}
}

// AddDevice add device
func (d *DeviceMonitor) AddDevice(name string, devType string, ip string) {
	dev := &Device{
		Name:   name,
		IP:     ip,
		Type:   devType,
		Status: false,
	}
	d.Devices = append(d.Devices, dev)
}

// CheckDevices check states of devices
func (d *DeviceMonitor) CheckDevices() {
	for _, device := range d.Devices {
		wdev := NewWiFiController(device.IP)
		status := wdev.DeviceStatus()

		// Status was changed
		if status != device.Status {
			device.Status = status
			d.SendNotify(device)
		} else if d.FirstCheck {
			d.SendNotify(device)
		}
	}
	
	if d.FirstCheck {
		d.FirstCheck = false
	}
}

// SendNotify send notify with status
func (d *DeviceMonitor) SendNotify(device *Device) {
	var status string

	if device.Status {
		status = "Online"
	} else {
		status = "Offline"
	}

	logger.Infof("Device %s type %s is %s", device.Name,
		device.Type, status)
}
