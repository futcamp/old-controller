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

package airctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CtrlAirData controller air sync data
type CtrlAirData struct {
	Name   string `json:"name"`
	Synced bool   `json:"synced"`
}

type WiFiController struct {
	ip string
}

// NewWiFiController make new struct
func NewWiFiController(ip string) *WiFiController {
	return &WiFiController{
		ip: ip,
	}
}

// SyncThermoData get actual air data from controller
func (w *WiFiController) SyncAirData() (CtrlAirData, error) {
	var data CtrlAirData

	res, err := http.Get(fmt.Sprintf("http://%s/sync", w.ip))
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	byteBuf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(byteBuf, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// SwitchRelay switch humidity control device relay
func (w *WiFiController) SwitchRelay(state bool) error {
	res, err := http.Get(fmt.Sprintf("http://%s/relay?state=%t", w.ip, state))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}
