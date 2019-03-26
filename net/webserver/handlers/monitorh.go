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

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/net/webserver/handlers/netdata"
	"github.com/futcamp/controller/utils"
)

type DisplayedDevice struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Online bool   `json:"online"`
}

type MonitorHandler struct {
	Log     *utils.Logger
	Monitor *monitoring.DeviceMonitor
}

// NewMonitorHandler make new struct
func NewMonitorHandler(log *utils.Logger,
	mon *monitoring.DeviceMonitor) *MonitorHandler {
	return &MonitorHandler{
		Log:     log,
		Monitor: mon,
	}
}

// ProcessMonitoring process monitoring of devices
func (m *MonitorHandler) ProcessMonitoring(req *http.Request) ([]byte, error) {
	var devices []DisplayedDevice
	data := &netdata.RestResponse{}

	for _, dev := range *m.Monitor.AllDevices() {
		device := DisplayedDevice{
			Name:dev.Name(),
			Type:dev.Type(),
		}
		device.Online = dev.Status()
		devices = append(devices, device)
	}
	netdata.SetRestResponse(data, "monitoring", "devices status monitoring", devices, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}
