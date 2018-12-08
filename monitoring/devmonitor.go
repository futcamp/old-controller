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
	"sync"

	"github.com/google/logger"
)

type Device struct {
	Name string
	Type string
	IP   string
	Mtx  sync.Mutex
	Stat bool
}

type DeviceMonitor struct {
	Devices    []*Device
	FirstCheck bool
}

// Status get device status
func (d *Device) Status() bool {
	var s bool

	d.Mtx.Lock()
	s = d.Stat
	d.Mtx.Unlock()

	return s
}

// SetStatus set device status
func (d *Device) SetStatus(status bool) {
	d.Mtx.Lock()
	d.Stat = status
	d.Mtx.Unlock()
}

// NewDeviceMonitor make new struct
func NewDeviceMonitor() *DeviceMonitor {
	return &DeviceMonitor{
		FirstCheck: true,
	}
}

// AddDevice add new device
func (d *DeviceMonitor) AddDevice(name string, devType string, ip string) {
	dev := &Device{
		Name: name,
		IP:   ip,
		Type: devType,
		Stat: false,
	}
	d.Devices = append(d.Devices, dev)
}

// AllDevices get devices list
func (d *DeviceMonitor) AllDevices() *[]*Device {
	return &d.Devices
}

// CheckDevices check states of devices
func (d *DeviceMonitor) CheckDevices() {
	for _, device := range d.Devices {
		wdev := NewWiFiController(device.IP)
		status := wdev.DeviceStatus()

		// Status was changed
		if status != device.Stat {
			device.SetStatus(status)
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

	if device.Status() {
		status = "Online"
	} else {
		status = "Offline"
	}

	logger.Infof("Device %s type %s is %s", device.Name,
		device.Type, status)
}
