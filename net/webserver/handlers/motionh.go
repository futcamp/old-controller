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

type DisplayedMotionModule struct {
	Name     string `json:"name"`
	Activity bool   `json:"activity"`
	Delay    int    `json:"delay"`
}

type MotionHandler struct {
	motion *devices.Motion
	dynCfg *configs.DynamicConfigs
}

// NewMotionHandler make new struct
func NewMotionHandler(mot *devices.Motion, dc *configs.DynamicConfigs) *MotionHandler {
	return &MotionHandler{
		motion: mot,
		dynCfg: dc,
	}
}

// ProcessMotionAllHandler display actual motion data for all devices
func (m *MotionHandler) ProcessMotionAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedMotionModule
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range m.motion.AllModules() {
		md := DisplayedMotionModule{
			Name:     mod.Name(),
			Activity: mod.Activity(),
			Delay:    mod.CurDelay(),
		}

		mods = append(mods, md)
	}

	netdata.SetRestResponse(data, "motion", "Motion", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessMotionSingleHandler display actual motion data for single mod
func (m *MotionHandler) ProcessMotionSingleHandler(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := m.motion.Module(modName)

	md := DisplayedMotionModule{
		Name:     mod.Name(),
		Activity: mod.Activity(),
		Delay:    mod.CurDelay(),
	}

	netdata.SetRestResponse(&data, "motion", "Motion", md, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessMotionActivity switch motion activity
func (m *MotionHandler) ProcessMotionActivity(modName string, req *http.Request) {
	mod := m.motion.Module(modName)
	mod.SetCurDelay(mod.Delay())
	mod.SetActivity(true)
}
