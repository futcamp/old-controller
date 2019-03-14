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

package mod

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type WiFiController struct {
	ip string
}

// NewWiFiController make new struct
func NewWiFiController(ip string) *WiFiController {
	return &WiFiController{
		ip: ip,
	}
}

// SyncData send cur states to controller
func (w *WiFiController) SyncData(status bool, humidifier bool) error {
	res, err := http.Get(fmt.Sprintf("http://%s/humctrl?status=%t&humidifier=%t", w.ip, status, humidifier))
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
