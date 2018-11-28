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
	"fmt"
	"net/http"
	"time"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/net/handlers"
	"github.com/futcamp/controller/net/handlers/nettools"
	"github.com/futcamp/controller/utils/configs"
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
	resp := nettools.NewResponse(&writer, AppName)
	resp.SendOk()
}

// MeteoHandler meteo station requests handler
func (w *WebServer) MeteoHandler(writer http.ResponseWriter, req *http.Request) {
	handlers.ProcessMeteoHandler(w.Meteo, w.MeteoCfg, &writer, req)
}
