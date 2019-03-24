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
func (l *LightHandler) ProcessLightAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedLightModule
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range l.light.AllModules() {
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
func (l *LightHandler) ProcessLightSingleHandler(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := l.light.Module(modName)

	m := DisplayedLightModule{
		Name:   mod.Name(),
		Status: mod.Status(),
	}

	netdata.SetRestResponse(&data, "light", "Light", m, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessLightStatus set new light status for single mod
func (l *LightHandler) ProcessLightStatus(modName string, status bool, req *http.Request) error {
	// Update status state
	mod := l.light.Module(modName)
	mod.SetStatus(status)
	mod.SetUpdate(true)

	return nil
}

// ProcessLightSwitchStatus switch current light status for single mod
func (l *LightHandler) ProcessLightSwitchStatus(modName string, req *http.Request) error {
	// Update status state
	mod := l.light.Module(modName)
	mod.SwitchStatus()
	mod.SetUpdate(true)

	return nil
}

// ProcessHumCtrlSync sync current states with remote module
func (l *LightHandler) ProcessLightSync(modName string, req *http.Request) error {
	// Update status state
	mod := l.light.Module(modName)
	mod.SetUpdate(true)

	return nil
}
