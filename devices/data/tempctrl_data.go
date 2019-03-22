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

type TempCtrlData struct {
	threshold   int
	temperature int
	status      bool
	heater      bool
	mtxThresh   sync.Mutex
	mtxTemp     sync.Mutex
	mtxStat     sync.Mutex
	mtxHeater   sync.Mutex
}

//
// Data getters
//

// Threshold get current threshold value
func (t *TempCtrlData) Threshold() int {
	var value int

	t.mtxThresh.Lock()
	value = t.threshold
	t.mtxThresh.Unlock()

	return value
}

// Temperature get current temperature value
func (t *TempCtrlData) Temperature() int {
	var value int

	t.mtxTemp.Lock()
	value = t.temperature
	t.mtxTemp.Unlock()

	return value
}

// Status get current status value
func (t *TempCtrlData) Status() bool {
	var value bool

	t.mtxStat.Lock()
	value = t.status
	t.mtxStat.Unlock()

	return value
}

// Heater get current heater value
func (t *TempCtrlData) Heater() bool {
	var value bool

	t.mtxHeater.Lock()
	value = t.heater
	t.mtxHeater.Unlock()

	return value
}

//
// Data setters
//

// SetThreshold set new threshold value
func (t *TempCtrlData) SetThreshold(value int) {
	t.mtxThresh.Lock()
	t.threshold = value
	t.mtxThresh.Unlock()
}

// SetTemperature set new temperature value
func (t *TempCtrlData) SetTemperature(value int) {
	t.mtxTemp.Lock()
	t.temperature = value
	t.mtxTemp.Unlock()
}

// SetStatus set new status value
func (t *TempCtrlData) SetStatus(value bool) {
	t.mtxStat.Lock()
	t.status = value
	t.mtxStat.Unlock()
}

// SetHeater set new heater value
func (t *TempCtrlData) SetHeater(value bool) {
	t.mtxHeater.Lock()
	t.heater = value
	t.mtxHeater.Unlock()
}
