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

type DisplayedLightModule struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type LightHandler struct {
	light  *devices.Light
	dynCfg *configs.DynamicConfigs
}

// NewLightHandler make new struct
func NewLightHandler(lgh *devices.Light, dc *configs.DynamicConfigs) *LightHandler {
	return &LightHandler{
		light:  lgh,
		dynCfg: dc,
	}
}

// ProcessLightAllHandler display actual light data for all devices
func (h *LightHandler) ProcessLightAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedLightModule
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range h.light.AllModules() {
		m := DisplayedLightModule{
			Name:   mod.Name(),
			Status: mod.Status(),
		}

		mods = append(mods, m)
	}

	netdata.SetRestResponse(data, "light", "Light", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessLightSingleHandler display actual light data for single mod
func (h *LightHandler) ProcessLightSingleHandler(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := h.light.Module(modName)

	m := DisplayedLightModule{
		Name:   mod.Name(),
		Status: mod.Status(),
	}

	netdata.SetRestResponse(&data, "light", "Light", m, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessLightStatus set new light status for single mod
func (h *LightHandler) ProcessLightStatus(modName string, status bool, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update status state
	mod := h.light.Module(modName)
	mod.SetStatus(status)
	mod.SetUpdate(true)

	// Send response
	netdata.SetRestResponse(&data, "light", "Light", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessLightSwitchStatus switch current light status for single mod
func (h *LightHandler) ProcessLightSwitchStatus(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update status state
	mod := h.light.Module(modName)
	mod.SwitchStatus()
	mod.SetUpdate(true)

	// Send response
	netdata.SetRestResponse(&data, "light", "Light", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessHumCtrlSync sync current states with remote module
func (h *LightHandler) ProcessLightSync(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse
	var resp ResultResponse

	// Update status state
	mod := h.light.Module(modName)
	mod.SetUpdate(true)

	// Send response
	netdata.SetRestResponse(&data, "light", "Light", resp, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}
