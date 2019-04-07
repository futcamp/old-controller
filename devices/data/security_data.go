/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2019 Sergey Denisov.
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

package data

import "sync"

type SecurityData struct {
	status   bool
	alarm    bool
	mtxAlarm sync.Mutex
	mtxStat  sync.Mutex
}

//
// Data getters
//

// Status get current status value
func (s *SecurityData) Status() bool {
	var value bool

	s.mtxStat.Lock()
	value = s.status
	s.mtxStat.Unlock()

	return value
}

// Alarm get current alarm value
func (s *SecurityData) Alarm() bool {
	var value bool

	s.mtxAlarm.Lock()
	value = s.alarm
	s.mtxAlarm.Unlock()

	return value
}

// SetAlarm set new alarm value
func (s *SecurityData) SetAlarm(value bool) {
	s.mtxAlarm.Lock()
	s.alarm = value
	s.mtxAlarm.Unlock()
}

// SetStatus set new status value
func (s *SecurityData) SetStatus(value bool) {
	s.mtxStat.Lock()
	s.status = value
	s.mtxStat.Unlock()
}
