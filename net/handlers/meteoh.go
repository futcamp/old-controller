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
	"github.com/futcamp/controller/net/handlers/nettools"
	"github.com/futcamp/controller/utils/configs"
	"github.com/pkg/errors"
)

type MeteoHandler struct {
	MeteoCfg *configs.MeteoConfigs
	Meteo    *meteo.MeteoStation
}

// NewMeteoHandler make new struct
func NewMeteoHandler(mCfg *configs.MeteoConfigs, meteo *meteo.MeteoStation) *MeteoHandler {
	return &MeteoHandler{
		MeteoCfg: mCfg,
		Meteo:    meteo,
	}
}

// ProcessMeteoAllHandler display actual meteo data for all sensors
func (m *MeteoHandler) ProcessMeteoAllHandler(req *http.Request) ([]byte, error) {
	data := &nettools.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	sensors := m.Meteo.AllSensors()
	nettools.SetRestResponse(data, "meteo", "Meteo Station", sensors, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessMeteoDBHandler display meteo data for concrete sensor by date
func (m *MeteoHandler) ProcessMeteoDBHandler(sensor string, date string, req *http.Request) ([]byte, error) {
	data := &nettools.RestResponse{}

	if req.Method != http.MethodGet {
		return nil, errors.New("Bad request method")
	}

	db := meteo.NewMeteoDB(m.MeteoCfg.Settings().Database.Path)
	err := db.Load()
	if err != nil {
		return nil, err
	}

	sensors, err := db.MeteoDataByDate(sensor, date)
	db.Unload()
	if err != nil {
		return nil, err
	}
	nettools.SetRestResponse(data, "meteo", "Meteo Station", sensors, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ProcessMeteoDBClearHandler delete sensor's values from table
func (m *MeteoHandler) ProcessMeteoDBClearHandler(sensor string, req *http.Request) error {
	if req.Method != http.MethodPut {
		return errors.New("Bad request method")
	}

	db := meteo.NewMeteoDB(m.MeteoCfg.Settings().Database.Path)
	err := db.Load()
	if err != nil {
		return err
	}

	err = db.MeteoDataClear(sensor)
	db.Unload()
	if err != nil {
		return err
	}

	return nil
}
