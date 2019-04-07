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
	"github.com/google/logger"
	"net/http"

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/net/webserver/handlers/netdata"
	"github.com/futcamp/controller/utils/configs"

	"github.com/pkg/errors"
)

type DisplayedSecurityModule struct {
	Name   string `json:"name"`
	Opened bool   `json:"opened"`
}

type DisplayedSecurityKey struct {
	User string `json:"user"`
	Key  string `json:"key"`
}

type SecurityHandler struct {
	security *devices.Security
	dynCfg   *configs.DynamicConfigs
}

// NewSecurityHandler make new struct
func NewSecurityHandler(sec *devices.Security, dc *configs.DynamicConfigs) *SecurityHandler {
	return &SecurityHandler{
		security: sec,
		dynCfg:   dc,
	}
}

// ProcessSecurityKeys display actual security keys
func (s *SecurityHandler) ProcessSecurityKeys(req *http.Request) ([]byte, error) {
	var keys []DisplayedSecurityKey
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, key := range *s.security.Keys() {
		k := DisplayedSecurityKey{
			User: key.User,
			Key:  key.Key,
		}

		keys = append(keys, k)
	}

	netdata.SetRestResponse(data, "security", "Security", keys, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessSecurityAllHandler display actual security data for all devices
func (s *SecurityHandler) ProcessSecurityAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedSecurityModule
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range s.security.AllModules() {
		m := DisplayedSecurityModule{
			Name:   mod.Name(),
			Opened: mod.Opened(),
		}

		mods = append(mods, m)
	}

	netdata.SetRestResponse(data, "security", "Security", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessSecuritySingleHandler display actual security data for single mod
func (s *SecurityHandler) ProcessSecuritySingleHandler(modName string, req *http.Request) ([]byte, error) {
	var data netdata.RestResponse

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := s.security.Module(modName)

	m := DisplayedSecurityModule{
		Name:   mod.Name(),
		Opened: mod.Opened(),
	}

	netdata.SetRestResponse(&data, "security", "Security", m, req)

	jData, _ := json.Marshal(&data)
	return jData, nil
}

// ProcessSecurityOpenHandler process open action for sensor
func (s *SecurityHandler) ProcessSecurityOpenHandler(modName string, req *http.Request) error {
	if s.security.Status() {
		s.security.SetAlarm(true)

		mod := s.security.Module(modName)
		mod.SetOpened(true)

		for _, module := range s.security.AllModules() {
			module.SetUpdate(true)
		}

		logger.Infof("Security sensor \"%s\" is opened!", modName)
	}

	return nil
}

// ProcessSecurityStatus set new security status for single mod
func (s *SecurityHandler) ProcessSecurityStatus(status bool, req *http.Request) {
	s.security.SetStatus(status)

	if !status {
		s.security.SetAlarm(false)
	}

	// Update status state
	for _, mod := range s.security.AllModules() {
		mod.SetUpdate(true)
	}
}
