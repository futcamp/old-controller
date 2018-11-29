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
	"strings"
	"time"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/net/handlers"
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
	MeteoHdl *handlers.MeteoHandler
}

// NewWebServer make new struct
func NewWebServer(cfg *configs.Configs, meteo *meteo.MeteoStation,
	mCfg *configs.MeteoConfigs, mh *handlers.MeteoHandler) *WebServer {
	return &WebServer{
		Cfg:      cfg,
		MeteoCfg: mCfg,
		Meteo:    meteo,
		MeteoHdl: mh,
	}
}

// Start run web server
func (w *WebServer) Start(ip string, port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", w.IndexHandler)
	if w.Cfg.Settings().Modules.Meteo {
		mux.HandleFunc(fmt.Sprintf("/api/%s/meteo/", configs.ApiVersion), w.MeteoHandler)
	}

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
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	if len(args) == 6 {
		// Clear all sensor meteo data
		if args[5] == "clear" {
			err := w.MeteoHdl.ProcessMeteoDBClearHandler(args[4], req)
			if err != nil {
				logger.Error(err.Error())
				resp.SendFail(err.Error())
				return
			}
			resp.SendOk()
			return
		}

		// Get sensor's meteo data by date
		data, err := w.MeteoHdl.ProcessMeteoDBHandler(args[4], args[5], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	// Get all meteo sensors data
	data, err := w.MeteoHdl.ProcessMeteoAllHandler(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(data))
	return
}
