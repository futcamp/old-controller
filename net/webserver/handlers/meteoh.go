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

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/devices/db"
	"github.com/futcamp/controller/net/webserver/handlers/netdata"
	"github.com/futcamp/controller/utils/configs"

	"github.com/pkg/errors"
)

type DisplayedSensor struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Temp     int    `json:"temp"`
	Humidity int    `json:"humidity"`
	Pressure int    `json:"pressure"`
}

type MeteoHandler struct {
	meteo   *devices.MeteoStation
	meteoDB *db.MeteoDatabase
	dynCfg  *configs.DynamicConfigs
}

// NewMeteoHandler make new struct
func NewMeteoHandler(meteo *devices.MeteoStation, mdb *db.MeteoDatabase) *MeteoHandler {
	return &MeteoHandler{
		meteo:   meteo,
		meteoDB: mdb,
	}
}

// ProcessMeteoAllHandler display actual meteo data for all sensors
func (m *MeteoHandler) ProcessMeteoAllHandler(req *http.Request) ([]byte, error) {
	var mods []DisplayedSensor
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, mod := range m.meteo.AllModules() {
		s := DisplayedSensor{
			Name:     mod.Name(),
			Type:     mod.Type(),
			Temp:     mod.Temp(),
			Humidity: mod.Humidity(),
			Pressure: mod.Pressure(),
		}

		mods = append(mods, s)
	}

	netdata.SetRestResponse(data, "meteo", "Meteo Station", mods, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessMeteoDBHandler display meteo data for concrete sensor by date
func (m *MeteoHandler) ProcessMeteoDBHandler(sensor string, date string, req *http.Request) ([]byte, error) {
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	db := m.dynCfg.Settings().MeteoDB
	err := m.meteoDB.Connect(db.IP, db.User, db.Passwd, db.Base)
	if err != nil {
		return nil, err
	}
	defer m.meteoDB.Close()

	sensors, err := m.meteoDB.MeteoDataByDate(sensor, date)
	if err != nil {
		return nil, err
	}
	netdata.SetRestResponse(data, "meteo", "Meteo Station", sensors, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessMeteoDBClearHandler delete sensor's values from table
func (m *MeteoHandler) ProcessMeteoDBClearHandler(sensor string, req *http.Request) error {
	if req.Method != http.MethodPut {
		return errors.New("Bad request method")
	}

	db := m.dynCfg.Settings().MeteoDB
	err := m.meteoDB.Connect(db.IP, db.User, db.Passwd, db.Base)
	if err != nil {
		return err
	}
	defer m.meteoDB.Close()

	err = m.meteoDB.MeteoDataClear(sensor)
	if err != nil {
		return err
	}

	return nil
}

// ProcessMeteoSingleHandler display actual meteo data for single sensors
func (m *MeteoHandler) ProcessMeteoSingleHandler(sensorName string, req *http.Request) ([]byte, error) {
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	mod := m.meteo.Module(sensorName)

	s := DisplayedSensor{
		Name:     mod.Name(),
		Type:     mod.Type(),
		Temp:     mod.Temp(),
		Humidity: mod.Humidity(),
		Pressure: mod.Pressure(),
	}

	netdata.SetRestResponse(data, "meteo", "Meteo Station", s, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}
