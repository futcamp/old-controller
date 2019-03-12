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
	"fmt"
	"github.com/futcamp/controller/modules/humctrl"
	"github.com/futcamp/controller/net/webserver/handlers/netdata"
	"github.com/futcamp/controller/utils/configs"
	"net/http"

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
	humCtrl *humctrl.HumControl
	dynCfg  *configs.DynamicConfigs
}

// NewHumCtrlHandler make new struct
func NewHumCtrlHandler(hc *humctrl.HumControl, dc *configs.DynamicConfigs) *HumCtrlHandler {
	return &HumCtrlHandler{
		humCtrl: hc,
		dynCfg:  dc,
	}
}

// ProcessHumCtrlAllHandler display actual hum control data for all modules
func (h *HumCtrlHandler) ProcessHumCtrlAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedModule
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range h.humCtrl.AllModules() {
		m := DisplayedModule{
			Name:       mod.Name,
			Sensor:     mod.Sensor,
			Humidity:   mod.Humidity,
			Status:     mod.Data.Status(),
			Threshold:  mod.Data.Threshold(),
			Humidifier: mod.Humidifier,
		}

		mods = append(mods, m)
	}

	netdata.SetRestResponse(data, "humctrl", "Humidity Control", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessHumCtrlSingleHandler display actual hum control data for single module
func (h *HumCtrlHandler) ProcessHumCtrlSingleHandler(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := h.humCtrl.Module(modName)

	m := DisplayedModule{
		Name:       mod.Name,
		Sensor:     mod.Sensor,
		Humidity:   mod.Humidity,
		Status:     mod.Data.Status(),
		Threshold:  mod.Data.Threshold(),
		Humidifier: mod.Humidifier,
	}

	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", m, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessHumCtrlStatus set new hum control status for single module
func (h *HumCtrlHandler) ProcessHumCtrlStatus(modName string, status bool, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse
	var stat string

	// Update status state
	mod := h.humCtrl.Module(modName)
	mod.Data.SetStatus(status)

	// Save configs
	if status {
		stat = "on"
	} else {
		stat = "off"
	}
	h.dynCfg.AddCommand(fmt.Sprintf("humctrl status %s %s", modName, stat))
	h.dynCfg.SaveConfigs()

	// Send response
	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessHumCtrlThreshold set new hum control threshold for single module
func (h *HumCtrlHandler) ProcessHumCtrlThreshold(modName string, plus bool, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update threshold value
	mod := h.humCtrl.Module(modName)
	thresh := mod.Data.Threshold()
	if plus {
		thresh++
	} else {
		thresh--
	}
	mod.Data.SetThreshold(thresh)

	// Save configs
	h.dynCfg.AddCommand(fmt.Sprintf("humctrl threshold %s %d", modName, thresh))
	h.dynCfg.SaveConfigs()

	// Send response
	netdata.SetRestResponse(&data, "humctrl", "Humidity Control", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}
