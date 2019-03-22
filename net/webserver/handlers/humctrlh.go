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

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/net/webserver/handlers/netdata"
	"github.com/futcamp/controller/utils/configs"

	"github.com/pkg/errors"
)

type DisplayedModule struct {
	Name       string `json:"name"`
	Sensor     string `json:"sensor"`
	Humidity   int    `json:"temp"`
	Threshold  int    `json:"threshold"`
	Status     bool   `json:"status"`
	Humidifier bool   `json:"humidifier"`
}

type ResultResponse struct {
	Result bool `json:"result"`
}

type HumCtrlHandler struct {
	humCtrl *devices.HumControl
	dynCfg  *configs.DynamicConfigs
}

// NewHumCtrlHandler make new struct
func NewHumCtrlHandler(hc *devices.HumControl, dc *configs.DynamicConfigs) *HumCtrlHandler {
	return &HumCtrlHandler{
		humCtrl: hc,
		dynCfg:  dc,
	}
}

// ProcessHumCtrlAllHandler display actual hum control data for all devices
func (h *HumCtrlHandler) ProcessHumCtrlAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedModule
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range h.humCtrl.AllModules() {
		m := DisplayedModule{
			Name:       mod.Name(),
			Sensor:     mod.Sensor(),
			Humidity:   mod.Humidity(),
			Status:     mod.Status(),
			Threshold:  mod.Threshold(),
			Humidifier: mod.Humidifier(),
		}

		mods = append(mods, m)
	}

	netdata.SetRestResponse(data, "humctrl", "Humidity Control", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessHumCtrlSingleHandler display actual hum control data for single mod
func (h *HumCtrlHandler) ProcessHumCtrlSingleHandler(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := h.humCtrl.Module(modName)

	m := DisplayedModule{
		Name:       mod.Name(),
		Sensor:     mod.Sensor(),
		Humidity:   mod.Humidity(),
		Status:     mod.Status(),
		Threshold:  mod.Threshold(),
		Humidifier: mod.Humidifier(),
	}

	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", m, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessHumCtrlStatus set new hum control status for single mod
func (h *HumCtrlHandler) ProcessHumCtrlStatus(modName string, status bool, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update status state
	mod := h.humCtrl.Module(modName)
	mod.SetStatus(status)
	mod.SetUpdate(true)

	// Send response
	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessHumCtrlSwitchStatus switch current hum control status for single mod
func (h *HumCtrlHandler) ProcessHumCtrlSwitchStatus(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update status state
	mod := h.humCtrl.Module(modName)
	mod.SwitchStatus()
	mod.SetUpdate(true)

	// Send response
	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessHumCtrlSync sync current states with remote module
func (h *HumCtrlHandler) ProcessHumCtrlSync(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update status state
	mod := h.humCtrl.Module(modName)
	mod.SetUpdate(true)

	// Send response
	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessHumCtrlThreshold set new hum control threshold for single mod
func (h *HumCtrlHandler) ProcessHumCtrlThreshold(modName string, plus bool, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update threshold value
	mod := h.humCtrl.Module(modName)
	thresh := (*mod).Threshold()
	if plus {
		thresh++
	} else {
		thresh--
	}
	mod.SetThreshold(thresh)

	// Send response
	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}
