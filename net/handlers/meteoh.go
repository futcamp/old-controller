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
	"fmt"
	"net/http"
	"strings"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/utils/configs"
	"github.com/google/logger"
)

// ProcessMeteoHandler process meteo handler
func ProcessMeteoHandler(m *meteo.MeteoStation, mCfg *configs.MeteoConfigs,
	writer *http.ResponseWriter, req *http.Request)  {
	data := &RestResponse{}
	resp := NewResponse(writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	// Get sensors data by date
	if len(args) == 6 && req.Method == http.MethodGet {
		db := meteo.NewMeteoDB(mCfg.Settings().Database.Path)
		err := db.Load()
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}

		sensors, err := db.MeteoDataByDate(args[4], args[5])
		db.Unload()
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		SetRestResponse(data, "meteo", "Meteo Station", sensors, req)
		fmt.Print(sensors)

		jData, _ := json.Marshal(data)
		resp.Send(string(jData))
		return
	}

	// Get actual meteo data from all sensors
	if req.Method != http.MethodGet {
		logger.Error("Bad request method")
		resp.SendFail("Bad request method")
		return
	}

	sensors := m.AllSensors()
	SetRestResponse(data, "meteo", "Meteo Station", sensors, req)

	jData, _ := json.Marshal(data)
	resp.Send(string(jData))
}
