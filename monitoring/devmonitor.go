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

package monitoring

import (
	"fmt"
	"sync"

	"github.com/futcamp/controller/notifier"

	"github.com/google/logger"
)

// Structure with device fields
type Device struct {
	Name string
	Type string
	IP   string
	Mtx  sync.Mutex
	Stat bool
}

type DeviceMonitor struct {
	name string
	notify     *notifier.Notifier
	devices    []*Device
	firstCheck bool
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
func NewDeviceMonitor(ntf *notifier.Notifier) *DeviceMonitor {
	return &DeviceMonitor{
		firstCheck: true,
		notify:     ntf,
	}
}

// SetName set monitor name
func (d *DeviceMonitor) SetName(name string) {
	d.name = name
}

// AddDevice add new device
func (d *DeviceMonitor) AddDevice(name string, devType string, ip string) {
	dev := &Device{
		Name: name,
		IP:   ip,
		Type: devType,
		Stat: false,
	}
	d.devices = append(d.devices, dev)
}

// AllDevices get devices list
func (d *DeviceMonitor) AllDevices() *[]*Device {
	return &d.devices
}

// CheckDevices check states of devices
func (d *DeviceMonitor) CheckDevices() {
	for _, device := range d.devices {
		wdev := NewWiFiController(device.IP)
		status := wdev.DeviceStatus()

		// Status was changed
		if status != device.Stat {
			device.SetStatus(status)
			d.SendNotify(device)
		} else if d.firstCheck {
			d.SendNotify(device)
		}
	}

	if d.firstCheck {
		d.firstCheck = false
	}
}

// SendNotify send notify with status
func (d *DeviceMonitor) SendNotify(device *Device) {
	var status string
	var message string

	if device.Status() {
		status = "Online"
	} else {
		status = "Offline"
	}

	message = fmt.Sprintf("Device %s type %s is %s", device.Name,
		device.Type, status)

	logger.Infof(message)
	d.notify.SendNotify("Monitor", message)
}
