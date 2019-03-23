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

type LightData struct {
	status  bool
	mtxStat sync.Mutex
}

//
// Data getters
//

// Status get current status value
func (l *LightData) Status() bool {
	var value bool

	l.mtxStat.Lock()
	value = l.status
	l.mtxStat.Unlock()

	return value
}

//
// Data setters
//

// SetStatus set new status value
func (l *LightData) SetStatus(value bool) {
	l.mtxStat.Lock()
	l.status = value
	l.mtxStat.Unlock()
}
