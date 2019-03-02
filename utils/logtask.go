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
	"fmt"
	"os"
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
			l.Log.Free()

			// Remove if last log file is empty
			fileName := fmt.Sprint("%s/%s.log", LogPath, l.LastDate)
			file, err := os.Open(fileName)
			if err == nil {
				stat, err := file.Stat()
				if err == nil {
					if stat.Size() == 0 {
						os.Remove(fileName)
					}
				}
			}

			l.LastDate = curDate
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
