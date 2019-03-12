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

package humctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CtrlMeteoData controller meteo data
type CtrlHumctrlData struct {
	Status bool `json:"status"`
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

// SyncData get actual data from controller and send cur states
func (w *WiFiController) SyncData(status bool, heater bool) (CtrlHumctrlData, error) {
	var data CtrlHumctrlData

	res, err := http.Get(fmt.Sprintf("http://%s/humctrl?status=%t&humidifier=%t", w.ip, status, heater))
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
