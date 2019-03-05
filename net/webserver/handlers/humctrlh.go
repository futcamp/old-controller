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

	"github.com/futcamp/controller/modules/humctrl"
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

type HumCtrlHandler struct {
	humCtrl *humctrl.HumControl
	dynCfg  *configs.DynamicConfigs
}

// NewHumCtrlHandler make new struct
func NewHumCtrlHandler(hc *humctrl.HumControl) *HumCtrlHandler {
	return &HumCtrlHandler{
		humCtrl: hc,
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
		mData := mod.ModuleData()
		sData := mod.ServerData()

		m := DisplayedModule{
			Name:       mod.Name,
			Sensor:     mod.Sensor,
			Humidity:   sData.Hum,
			Status:     sData.Status,
			Threshold:  sData.Threshold,
			Humidifier: mData.Humidifier,
		}

		mods = append(mods, m)
	}

	netdata.SetRestResponse(data, "humctrl", "Humidity Control", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessHumCtrlSingleHandler display actual hum control data for single module
func (h *HumCtrlHandler) ProcessHumCtrlSingleHandler(modName string, req *http.Request) ([]byte, error) {
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := h.humCtrl.Module(modName)
	mData := mod.ModuleData()
	sData := mod.ServerData()

	m := DisplayedModule{
		Name:       mod.Name,
		Sensor:     mod.Sensor,
		Humidity:   sData.Hum,
		Status:     sData.Status,
		Threshold:  sData.Threshold,
		Humidifier: mData.Humidifier,
	}

	netdata.SetRestResponse(data, "humctrl", "Humidity Control", m, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}
