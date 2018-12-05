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
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/google/logger"
)

type Logger struct {
	LogFile *os.File
	Log  *logger.Logger
}

// NewLogger make new struct
func NewLogger() *Logger {
	return &Logger{}
}

// Init logger initialization
func (l *Logger) Init(path string) {
	var err error
	date := time.Now().Format("2006-01-02")

	// Configuring log module
	l.LogFile, err = os.OpenFile(fmt.Sprintf("%s%s.log", path, date), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Fail to open log file!")
		return
	}

	l.Log = logger.Init("FutcampLogger", true, true, l.LogFile)
}

// LogsList reading logs list from path
func (l *Logger) LogsList(path string) ([]string, error) {
	var logFiles []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			parts := strings.Split(file.Name(), ".")
			if len(parts) < 2 {
				continue
			}

			if parts[1] == "log" {
				logFiles = append(logFiles, file.Name())
			}
		}
	}
	return logFiles, nil
}

// ReadLogByDate reading logs messages by date
func (l *Logger) ReadLogByDate(path string, date string) ([]string, error) {
	var logs []string

	file, _ := os.Open(fmt.Sprintf("%s%s.log", path, date))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	return logs, nil
}

// Free unload logger module
func (l *Logger) Free() {
	if l.Log != nil {
		l.Log.Close()
	}
	if l.LogFile != nil {
		l.LogFile.Close()
	}
}
