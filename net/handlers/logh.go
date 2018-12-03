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

	"github.com/futcamp/controller/net/handlers/netdata"
	"github.com/futcamp/controller/utils"
)

type LogHandler struct {
	Log *utils.Logger
}

// NewLogHandler make new struct
func NewLogHandler(log *utils.Logger) *LogHandler {
	return &LogHandler{
		Log: log,
	}
}

// ExistingLogsList get existing log files list
func (l *LogHandler) ProcessExistingLogsList(req *http.Request) ([]byte, error) {
	data := &netdata.RestResponse{}

	logFiles, err := l.Log.LogsList(utils.LogPath)
	if err != nil {
		return nil, err
	}

	netdata.SetRestResponse(data, "logger", "Application Logger", logFiles, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}

// ExistingLogsList get existing log files list
func (l *LogHandler) ProcessLogsByDate(date string, req *http.Request) ([]byte, error) {
	data := &netdata.RestResponse{}

	logs, err := l.Log.ReadLogByDate(utils.LogPath, date)
	if err != nil {
		return nil, err
	}

	netdata.SetRestResponse(data, "logger", "Application Logger", logs, req)

	jData, _ := json.Marshal(data)
	return jData, nil
}