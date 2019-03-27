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

type MotionData struct {
	activity bool
	curDelay int
	mtxAct   sync.Mutex
	mtxDelay sync.Mutex
}

// Activity get current activity state
func (m *MotionData) Activity() bool {
	var act bool

	m.mtxAct.Lock()
	act = m.activity
	m.mtxAct.Unlock()

	return act
}

// CurDelay get current delay
func (m *MotionData) CurDelay() int {
	var delay int

	m.mtxDelay.Lock()
	delay = m.curDelay
	m.mtxDelay.Unlock()

	return delay
}

// Activity set new activity state
func (m *MotionData) SetActivity(act bool) {
	m.mtxAct.Lock()
	m.activity = act
	m.mtxAct.Unlock()
}

// SetCurDelay set new delay value
func (m *MotionData) SetCurDelay(delay int) {
	m.mtxDelay.Lock()
	m.curDelay = delay
	m.mtxDelay.Unlock()
}
