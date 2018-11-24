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
	"time"

	"github.com/futcamp/controller/meteo"
	"github.com/futcamp/controller/utils/configs"
)

const (
	AppName = "futcamp"
)

type WebServer struct {
	Cfg   *configs.Configs
	Meteo *meteo.MeteoStation
}

// NewWebServer make new struct
func NewWebServer(cfg *configs.Configs, meteo *meteo.MeteoStation) *WebServer {
	return &WebServer{
		Cfg:   cfg,
		Meteo: meteo,
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
	sensors := w.Meteo.AllSensors()

	SetRestResponse(data, "meteo", "Meteo Station", sensors, req)

	jData, _ := json.Marshal(data)
	resp.Send(string(jData))
}
