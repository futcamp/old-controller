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

type DisplayedTempModule struct {
	Name        string `json:"name"`
	Sensor      string `json:"sensor"`
	Temperature int    `json:"temp"`
	Threshold   int    `json:"threshold"`
	Status      bool   `json:"status"`
	Heater      bool   `json:"heater"`
}

type TempCtrlHandler struct {
	tempCtrl *devices.TempControl
	dynCfg   *configs.DynamicConfigs
}

// NewTempCtrlHandler make new struct
func NewTempCtrlHandler(tc *devices.TempControl, dc *configs.DynamicConfigs) *TempCtrlHandler {
	return &TempCtrlHandler{
		tempCtrl: tc,
		dynCfg:   dc,
	}
}

// ProcessTempCtrlAllHandler display actual temp control data for all devices
func (t *TempCtrlHandler) ProcessTempCtrlAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedTempModule
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range t.tempCtrl.AllModules() {
		m := DisplayedTempModule{
			Name:        mod.Name(),
			Sensor:      mod.Sensor(),
			Temperature: mod.Temperature(),
			Status:      mod.Status(),
			Threshold:   mod.Threshold(),
			Heater:      mod.Heater(),
		}

		mods = append(mods, m)
	}

	netdata.SetRestResponse(data, "tempctrl", "Temperature Control", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessTempCtrlSingleHandler display actual temp control data for single mod
func (t *TempCtrlHandler) ProcessTempCtrlSingleHandler(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := t.tempCtrl.Module(modName)

	m := DisplayedTempModule{
		Name:        mod.Name(),
		Sensor:      mod.Sensor(),
		Temperature: mod.Temperature(),
		Status:      mod.Status(),
		Threshold:   mod.Threshold(),
		Heater:      mod.Heater(),
	}

	netdata.SetRestResponse(&data, "tempctrl", "Temperature Control", m, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessTempCtrlStatus set new temp control status for single mod
func (t *TempCtrlHandler) ProcessTempCtrlStatus(modName string, status bool, req *http.Request) error {
	// Update status state
	mod := t.tempCtrl.Module(modName)
	mod.SetStatus(status)
	mod.SetUpdate(true)

	return nil
}

// ProcessTempCtrlSwitchStatus switch current temp control status for single mod
func (t *TempCtrlHandler) ProcessTempCtrlSwitchStatus(modName string, req *http.Request) error {
	// Update status state
	mod := t.tempCtrl.Module(modName)
	mod.SwitchStatus()
	mod.SetUpdate(true)

	return nil
}

// ProcessTempCtrlSync sync current states with remote module
func (t *TempCtrlHandler) ProcessTempCtrlSync(modName string, req *http.Request) error {
	// Update status state
	mod := t.tempCtrl.Module(modName)
	mod.SetUpdate(true)

	return nil
}

// ProcessTempCtrlThreshold set new temp control threshold for single mod
func (t *TempCtrlHandler) ProcessTempCtrlThreshold(modName string, plus bool, req *http.Request) error {
	// Update threshold value
	mod := t.tempCtrl.Module(modName)
	thresh := mod.Threshold()
	if plus {
		thresh++
	} else {
		thresh--
	}
	mod.SetThreshold(thresh)

	return nil
}
