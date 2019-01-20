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

package utils

import (
	"time"
)

const (
	taskDelay = 1
)

// logTask logger task struct
type LogTask struct {
	Log             *Logger
	ReqTimer        *time.Timer
	DisplaysCounter int
	SensorsCounter  int
	DatabaseCounter int
	LastDate        string
}

// NewLogTask make new struct
func NewLogTask(log *Logger) *LogTask {
	return &LogTask{
		Log:      log,
		LastDate: "",
	}
}

// TaskHandler process timer loop
func (l *LogTask) TaskHandler() {
	l.LastDate = time.Now().Format("2006-01-02")

	for {
		<-l.ReqTimer.C

		curDate := time.Now().Format("2006-01-02")
		if curDate != l.LastDate {
			l.LastDate = curDate
			l.Log.Free()
			l.Log.Init(LogPath)
		}

		l.ReqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (l *LogTask) Start() {
	l.ReqTimer = time.NewTimer(taskDelay * time.Second)
	l.TaskHandler()
}
