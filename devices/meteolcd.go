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

package devices

import (
	"github.com/futcamp/controller/devices/modules"
)

type MeteoDisplay struct {
	displays map[string]modules.Indicator
	sensors  []string
}

// NewMeteoDisplay make new struct
func NewMeteoDisplay() *MeteoDisplay {
	lcd := make(map[string]modules.Indicator)
	return &MeteoDisplay{
		displays: lcd,
	}
}

// AddDisplay add new display
func (m *MeteoDisplay) AddDisplay(name string, display modules.Indicator) {
	m.displays[name] = display
}

// DeleteDisplay add new display
func (m *MeteoDisplay) DeleteDisplay(name string) {
	delete(m.displays, name)
}

// Displays get all displays list
func (m *MeteoDisplay) Displays() []modules.Indicator {
	var displays []modules.Indicator

	for _, lcd := range m.displays {
		displays = append(displays, lcd)
	}

	return displays
}

// Display get single display struct
func (m *MeteoDisplay) Display(name string) modules.Indicator {
	return m.displays[name]
}
