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

package configs

type MeteoDBCfg struct {
	IP     string
	Port   int
	User   string
	Passwd string
	Base   string
}

type TimersCfg struct {
	MeteoSensorsDelay int
	MeteoDisplayDelay int
	MeteoDBDelay      int
}

type DynamicConfigs struct {
	MeteoDB MeteoDBCfg
	Timers  TimersCfg
}

// NewDynamicConfigs make new struct
func NewDynamicConfigs() *DynamicConfigs {
	return &DynamicConfigs{}
}
