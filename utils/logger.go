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

package utils

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/logger"
)

type Logger struct {
	logFile *os.File
	logger  *logger.Logger
}

// NewLogger make new struct
func NewLogger() *Logger {
	return &Logger{}
}

// Init logger initialization
func (l *Logger) Init(path string) {
	var err error

	// Configuring log module
	l.logFile, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Fail to open log file!")
		return
	}

	var verbose = flag.Bool("verbose", true, "print info level logs to stdout")
	l.logger = logger.Init("FutcampLogger", *verbose, true, l.logFile)
}

// Free unload logger module
func (l *Logger) Free() {
	l.logger.Close()
}
