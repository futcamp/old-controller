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

package tasks

import (
	"time"

	"github.com/futcamp/controller/devices"
)

const (
	secureTaskDelay = 10
)

type SecurityTask struct {
	security *devices.Security
	reqTimer *time.Timer
}

// NewSecurityTask make new struct
func NewSecurityTask(sec *devices.Security) *SecurityTask {
	return &SecurityTask{
		security: sec,
	}
}

// SecurityTask process timer loop
func (s *SecurityTask) TaskHandler() {
	for {
		<-s.reqTimer.C

		if s.security.Alarm() {
			s.security.SendAlarm(true)
		} else {
			s.security.SendAlarm(false)
		}

		s.reqTimer.Reset(secureTaskDelay * time.Second)
	}
}

// Start start new timer
func (s *SecurityTask) Start() {
	s.reqTimer = time.NewTimer(secureTaskDelay * time.Second)
	s.TaskHandler()
}
