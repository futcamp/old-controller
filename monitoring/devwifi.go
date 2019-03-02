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

package monitoring

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type WiFiController struct {
	IP string
}

// NewWiFiController make new struct
func NewWiFiController(ip string) *WiFiController {
	return &WiFiController{
		IP: ip,
	}
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
