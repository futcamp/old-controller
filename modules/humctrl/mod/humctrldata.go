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

package mod

import "sync"

type HumCtrlData struct {
	threshold  int
	humidity   int
	status     bool
	humidifier bool
	mtxThresh  sync.Mutex
	mtxHum     sync.Mutex
	mtxStat    sync.Mutex
	mtxHumfier sync.Mutex
}


//
// Data getters
//


// Threshold get current threshold value
func (h *HumCtrlData) Threshold() int {
	var value int

	h.mtxThresh.Lock()
	value = h.threshold
	h.mtxThresh.Unlock()

	return value
}

// Humidity get current humidity value
func (h *HumCtrlData) Humidity() int {
	var value int

	h.mtxHum.Lock()
	value = h.humidity
	h.mtxHum.Unlock()

	return value
}

// Status get current status value
func (h *HumCtrlData) Status() bool {
	var value bool

	h.mtxStat.Lock()
	value = h.status
	h.mtxStat.Unlock()

	return value
}

// Humidifier get current humidifier value
func (h *HumCtrlData) Humidifier() bool {
	var value bool

	h.mtxHumfier.Lock()
	value = h.humidifier
	h.mtxHumfier.Unlock()

	return value
}


//
// Data setters
//


// SetThreshold set new threshold value
func (h *HumCtrlData) SetThreshold(value int) {
	h.mtxThresh.Lock()
	h.threshold = value
	h.mtxThresh.Unlock()
}

// SetHumidity set new humidity value
func (h *HumCtrlData) SetHumidity(value int) {
	h.mtxHum.Lock()
	h.humidity = value
	h.mtxHum.Unlock()
}

// SetStatus set new status value
func (h *HumCtrlData) SetStatus(value bool) {
	h.mtxStat.Lock()
	h.status = value
	h.mtxStat.Unlock()
}

// SetHumidifier set new humidifier value
func (h *HumCtrlData) SetHumidifier(value bool) {
	h.mtxHumfier.Lock()
	h.humidifier = value
	h.mtxHumfier.Unlock()
}