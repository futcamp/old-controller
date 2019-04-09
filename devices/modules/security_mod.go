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

package modules

import (
	"sync"

	"github.com/futcamp/controller/devices/hardware"
)

type SecurityController interface {
	Update() bool
	Name() string
	IP() string
	Error() bool
	Opened() bool
	SetError(err bool)
	SetIP(ip string)
	SetUpdate(state bool)
	SetOpened(state bool)
	SyncAlarm(alarm bool) error
}

type SecurityModule struct {
	name      string
	ip        string
	error     bool
	update    bool
	channel   int
	opened    bool
	mtxOpened sync.Mutex
	mtxUpdate sync.Mutex
}

// NewSecurityModule make new struct
func NewSecurityModule(name string) *SecurityModule {
	return &SecurityModule{
		name:   name,
		update: true,
	}
}

//
// Simple data getters and setters
//

// Update get current update state
func (s *SecurityModule) Update() bool {
	var value bool

	s.mtxUpdate.Lock()
	value = s.update
	s.mtxUpdate.Unlock()

	return value
}

// Name get mod name
func (s *SecurityModule) Name() string {
	return s.name
}

// IP get current mod ip
func (s *SecurityModule) IP() string {
	return s.ip
}

// Opened get current open status
func (s *SecurityModule) Opened() bool {
	var state bool

	s.mtxOpened.Lock()
	state = s.opened
	s.mtxOpened.Unlock()

	return state
}

// Error get current error state
func (s *SecurityModule) Error() bool {
	return s.error
}

// SetOpenned set new open status
func (s *SecurityModule) SetOpened(state bool) {
	s.mtxOpened.Lock()
	s.opened = state
	s.mtxOpened.Unlock()
}

// SetError set new error state
func (s *SecurityModule) SetError(err bool) {
	s.error = err
}

// SetIP set new ip address
func (s *SecurityModule) SetIP(ip string) {
	s.ip = ip
}

// SetUpdate set new update state
func (s *SecurityModule) SetUpdate(state bool) {
	s.mtxUpdate.Lock()
	s.update = state
	s.mtxUpdate.Unlock()
}

//
// Other functional
//

// SyncData sync data with module
func (s *SecurityModule) SyncAlarm(alarm bool) error {
	_, err := hardware.HdkSendSecurityAlarm(s.IP(), alarm)
	if err != nil {
		return err
	}

	return nil
}
