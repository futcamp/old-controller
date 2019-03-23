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

package hardware

import (
	"fmt"
)

type HdkMeteoData struct {
	Temp     int `json:"temp"`
	Humidity int `json:"humidity"`
	Pressure int `json:"pressure"`
}

type HdkModResponse struct {
	Module string `json:"module"`
	Device string `json:"device"`
}

// HdkSyncMeteoData get actual meteo data from controller
func HdkSyncMeteoData(ip string, channel int, sensType string) (*HdkMeteoData, error) {
	data := &HdkMeteoData{}

	err := HdkHttpSyncData(ip, fmt.Sprintf("meteo?chan=%d&type=%s", channel, sensType), data)

	return data, err
}

// HdkSyncDisplayData send actual meteo data to display controller
func HdkSyncDisplayData(ip string, sensor string, temp int, hum int, pres int) (*HdkModResponse, error) {
	resp := &HdkModResponse{}

	err := HdkHttpSyncData(ip, fmt.Sprintf("display?sensor=%s&temp=%d&hum=%d&pres=%d", sensor, temp, hum, pres), resp)

	return resp, err
}

// HdkSyncDisplayData send cur humctrl states to controller
func HdkSyncHumCtrlData(ip string, status bool, humidifier bool) (*HdkModResponse, error) {
	resp := &HdkModResponse{}

	err := HdkHttpSyncData(ip, fmt.Sprintf("humctrl?status=%t&humidifier=%t", status, humidifier), resp)

	return resp, err
}

// HdkSyncTempCtrlData send cur tempctrl states to controller
func HdkSyncTempCtrlData(ip string, status bool, heater bool) (*HdkModResponse, error) {
	resp := &HdkModResponse{}

	err := HdkHttpSyncData(ip, fmt.Sprintf("tempctrl?status=%t&heater=%t", status, heater), resp)

	return resp, err
}

// HdkSyncLightData send cur light states to controller
func HdkSyncLightData(ip string, channel int, status bool) (*HdkModResponse, error) {
	resp := &HdkModResponse{}

	err := HdkHttpSyncData(ip, fmt.Sprintf("light?chan=%d&status=%t", channel, status), resp)

	return resp, err
}