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

import "github.com/futcamp/controller/devices/data"

type MotionController interface {
	Name() string
	IP() string
	SetIP(ip string)
	SetDelay(delay int)
	Delay() int
	Activity() bool
	CurDelay() int
	SetActivity(act bool)
	SetCurDelay(delay int)
	AddLamp(lamp string)
	Lamps() *[]string
	AlreadyOn() bool
	SetAlreadyOn(on bool)
}

type MotionModule struct {
	name      string
	ip        string
	delay     int
	lamps     []string
	alreadyOn bool
	data      data.MotionData
}

// NewMotionModule make new struct
func NewMotionModule(name string) *MotionModule {
	return &MotionModule{
		name: name,
	}
}

//
// Simple data getters and setters
//

// Name get current module name
func (m *MotionModule) Name() string {
	return m.name
}

// Delay get current delay value
func (m *MotionModule) Delay() int {
	return m.delay
}

// IP get current module ip address
func (m *MotionModule) IP() string {
	return m.ip
}

// Lamps get current lamps names
func (m *MotionModule) Lamps() *[]string {
	return &m.lamps
}

// AlreadyOn get current already on state
func (m *MotionModule) AlreadyOn() bool {
	return m.alreadyOn
}

// SetIP set new module ip address
func (m *MotionModule) SetIP(ip string) {
	m.ip = ip
}

// SetDelay set new module delay
func (m *MotionModule) SetDelay(delay int) {
	m.delay = delay
}

// AddLamp add new lamp
func (m *MotionModule) AddLamp(lamp string) {
	m.lamps = append(m.lamps, lamp)
}

// SetChannel set new module channel
func (m *MotionModule) SetAlreadyOn(on bool) {
	m.alreadyOn = on
}

//
// Nested getters and setters
//

// Activity get current activity state
func (m *MotionModule) Activity() bool {
	return m.data.Activity()
}

// CurDelay get current delay value
func (m *MotionModule) CurDelay() int {
	return m.data.CurDelay()
}

// SetActivity set new activity state
func (m *MotionModule) SetActivity(act bool) {
	m.data.SetActivity(act)
}

// SetCurDelay set new delay value
func (m *MotionModule) SetCurDelay(delay int) {
	m.data.SetCurDelay(delay)
}
