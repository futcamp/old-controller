/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2018 Sergey Denisov.
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

package meteo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CtrlMeteoData controller meteo data
type CtrlMeteoData struct {
	Temp     int `json:"temp"`
	Humidity int `json:"humidity"`
	Pressure int `json:"pressure"`
}

type WiFiController struct {
	SensorType string
	IP         string
	Channel    int
}

// NewWiFiController make new struct
func NewWiFiController(sType string, ip string, ch int) *WiFiController {
	return &WiFiController{
		SensorType: sType,
		IP:         ip,
		Channel:    ch,
	}
}

// NewWiFiControllerDisplay make new struct for display
func NewWiFiControllerDisplay(ip string) *WiFiController {
	return &WiFiController{
		IP: ip,
	}
}

// SyncMeteoData get actual meteo data from controller
func (w *WiFiController) SyncMeteoData() (CtrlMeteoData, error) {
	var data CtrlMeteoData

	res, err := http.Get(fmt.Sprintf("http://%s/meteo?chan=%d&type=%s",
		w.IP, w.Channel, w.SensorType))
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

// DisplayMeteoData get actual meteo data from controller
func (w *WiFiController) DisplayMeteoData(sensor string, data *CtrlMeteoData) error {
	request := fmt.Sprintf("http://%s/display?sensor=%s&temp=%d&hum=%d&pres=%d",
		w.IP, sensor, data.Temp, data.Humidity, data.Pressure)

	res, err := http.Get(request)
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

// DeviceStatus device online status
func (w *WiFiController) DeviceStatus() bool {
	request := fmt.Sprintf("http://%s/", w.IP)

	res, err := http.Get(request)
	if err != nil {
		return false
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}

	return true
}
