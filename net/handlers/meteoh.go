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

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/net/handlers/netdata"
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
	MeteoCfg *configs.MeteoConfigs
	Meteo    *meteo.MeteoStation
	MeteoDB  *meteo.MeteoDatabase
}

// NewMeteoHandler make new struct
func NewMeteoHandler(mCfg *configs.MeteoConfigs, meteo *meteo.MeteoStation,
	mdb *meteo.MeteoDatabase) *MeteoHandler {
	return &MeteoHandler{
		MeteoCfg: mCfg,
		Meteo:    meteo,
		MeteoDB:  mdb,
	}
}

// ProcessMeteoAllHandler display actual meteo data for all sensors
func (m *MeteoHandler) ProcessMeteoAllHandler(req *http.Request) ([]byte, error) {
	var sensors []DisplayedSensor
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	for _, sensor := range m.Meteo.AllSensors() {
		mdata := sensor.MeteoData()

		s := DisplayedSensor{
			Name:     sensor.Name,
			Temp:     mdata.Temp,
			Humidity: mdata.Humidity,
			Pressure: mdata.Pressure,
		}

		sensors = append(sensors, s)
	}

	netdata.SetRestResponse(data, "meteo", "Meteo Station", sensors, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessMeteoDBHandler display meteo data for concrete sensor by date
func (m *MeteoHandler) ProcessMeteoDBHandler(sensor string, date string, req *http.Request) ([]byte, error) {
	data := &netdata.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	err := m.MeteoDB.Load()
	if err != nil {
		return nil, err
	}

	sensors, err := m.MeteoDB.MeteoDataByDate(sensor, date)
	m.MeteoDB.Unload()
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

	err := m.MeteoDB.Load()
	if err != nil {
		return err
	}

	err = m.MeteoDB.MeteoDataClear(sensor)
	m.MeteoDB.Unload()
	if err != nil {
		return err
	}

	return nil
}
