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

package net

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/futcamp/controller/meteo"
	"github.com/futcamp/controller/utils/configs"
	"github.com/google/logger"
)

const (
	AppName = "futcamp"
)

type WebServer struct {
	Cfg      *configs.Configs
	MeteoCfg *configs.MeteoConfigs
	Meteo    *meteo.MeteoStation
}

// NewWebServer make new struct
func NewWebServer(cfg *configs.Configs, meteo *meteo.MeteoStation, mCfg *configs.MeteoConfigs) *WebServer {
	return &WebServer{
		Cfg:      cfg,
		MeteoCfg: mCfg,
		Meteo:    meteo,
	}
}

// Start run web server
func (w *WebServer) Start(ip string, port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", w.IndexHandler)
	mux.HandleFunc(fmt.Sprintf("/api/%s/meteo/", configs.ApiVersion), w.MeteoHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", ip, port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	return server.ListenAndServe()
}

// IndexHandler index handler with app information
func (w *WebServer) IndexHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, AppName)
	resp.SendOk()
}

// MeteoHandler meteo station requests handler
func (w *WebServer) MeteoHandler(writer http.ResponseWriter, req *http.Request) {
	data := &RestResponse{}
	resp := NewResponse(&writer, AppName)
	args := strings.Split(req.RequestURI, "/")

	// Get sensors data by date
	if len(args) == 6 && req.Method == http.MethodGet {
		db := meteo.NewMeteoDB(w.MeteoCfg.Settings().Database.Path)
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

	sensors := w.Meteo.AllSensors()
	SetRestResponse(data, "meteo", "Meteo Station", sensors, req)

	jData, _ := json.Marshal(data)
	resp.Send(string(jData))
}
