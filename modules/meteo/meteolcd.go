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

package meteo

type MeteoDisplay struct {
	Name    string
	IP      string
	sensors []string
}

// NewMeteoDisplay make new struct
func NewMeteoDisplay(name string, ip string) *MeteoDisplay {
	return &MeteoDisplay{
		Name: name,
		IP:   ip,
	}
}

// AddDisplayingSensor add new displaying new sensor
func (m *MeteoDisplay) AddDisplayingSensor(name string) {
	m.sensors = append(m.sensors, name)
}

// Sensors get displaying meteo sensors
func (m *MeteoDisplay) Sensors() *[]string {
	return &m.sensors
}

// DisplayMeteo send display command to controller
func (m *MeteoDisplay) DisplayMeteo(sensor string, temp int, hum int, pres int) error {
	ctrlSensor := &CtrlMeteoData{
		Temp:     temp,
		Humidity: hum,
		Pressure: pres,
	}

	// Send data to controller
	ctrl := NewWiFiControllerDisplay(m.IP)
	return ctrl.DisplayMeteoData(sensor, ctrlSensor)
}

type MeteoDisplays struct {
	displays map[string]*MeteoDisplay
	sensors  []string
}

// NewMeteoDisplays make new struct
func NewMeteoDisplays() *MeteoDisplays {
	lcd := make(map[string]*MeteoDisplay)
	return &MeteoDisplays{
		displays: lcd,
	}
}

// AddMeteoDisplay add new display
func (m *MeteoDisplays) AddMeteoDisplay(name string, display *MeteoDisplay) {
	m.displays[name] = display
}

// Displays get all displays list
func (m *MeteoDisplays) Displays() []*MeteoDisplay {
	var displays []*MeteoDisplay

	for _, lcd := range m.displays {
		displays = append(displays, lcd)
	}

	return displays
}

// Display get single display struct
func (m *MeteoDisplays) Display(name string) *MeteoDisplay {
	return m.displays[name]
}
