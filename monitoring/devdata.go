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

package monitoring

import "sync"

type DeviceData struct {
	mtxStat sync.Mutex
	status  bool
}

// Status get current device status
func (d *DeviceData) Status() bool {
	var stat bool

	d.mtxStat.Lock()
	stat = d.status
	d.mtxStat.Unlock()

	return stat
}

// SetStatus set new device status
func (d *DeviceData) SetStatus(status bool) {
	d.mtxStat.Lock()
	d.status = status
	d.mtxStat.Unlock()
}
