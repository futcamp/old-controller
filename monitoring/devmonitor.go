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

	"github.com/futcamp/controller/notifier"

	"github.com/google/logger"
)

type DeviceMonitor struct {
	name       string
	notify     *notifier.Notifier
	devices    []MonitoringDevice
	firstCheck bool
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
func (d *DeviceMonitor) AddDevice(device MonitoringDevice) {
	d.devices = append(d.devices, device)
}

// DeleteDevice delete device from storage
func (d *DeviceMonitor) DeleteDevice(name string) {
	var newDevs []MonitoringDevice

	for _, dev := range d.devices {
		if dev.Name() != name {
			newDevs = append(newDevs, dev)
		}
	}

	d.devices = newDevs
}

// AllDevices get devices list
func (d *DeviceMonitor) AllDevices() *[]MonitoringDevice {
	return &d.devices
}

// CheckDevices check states of devices
func (d *DeviceMonitor) CheckDevices() {
	for _, device := range d.devices {
		wdev := NewWiFiController(device.IP())
		status := wdev.DeviceStatus()

		// Status was changed
		if status != device.Status() {
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
func (d *DeviceMonitor) SendNotify(device MonitoringDevice) {
	var status string
	var message string

	if device.Status() {
		status = "online"
	} else {
		status = "offline"
	}

	message = fmt.Sprintf("device \"%s\" module \"%s\" is \"%s\"", device.Name,
		device.Type, status)

	logger.Infof("Monitor %s", message)
	d.notify.SendNotify("Monitor:", message)
}
