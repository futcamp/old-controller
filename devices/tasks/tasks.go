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

package tasks

import "github.com/futcamp/controller/utils/configs"

type DeviceTasks struct {
	cfg       *configs.Configs
	meteoTask *MeteoTask
	hctrlTask *HumControlTask
}

// NewDeviceTasks make new struct
func NewDeviceTasks(cfg *configs.Configs, mt *MeteoTask, ht *HumControlTask) *DeviceTasks {
	return &DeviceTasks{
		cfg:       cfg,
		meteoTask: mt,
		hctrlTask: ht,
	}
}

// RunTasks run all devices tasks
func (d *DeviceTasks) RunTasks() {
	if d.cfg.Settings().Modules.Meteo {
		go d.meteoTask.Start()
	}
	if d.cfg.Settings().Modules.Humctrl {
		go d.hctrlTask.Start()
	}
}
