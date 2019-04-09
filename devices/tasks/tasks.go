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

type TaskManager interface {
	Start()
	TaskHandler()
}

type DeviceTasks struct {
	cfg        *configs.Configs
	meteoTask  *MeteoTask
	hctrlTask  *HumControlTask
	tctrlTask  *TempControlTask
	lightTask  *LightlTask
	motionTask *MotionTask
	secureTask *SecurityTask
}

// NewDeviceTasks make new struct
func NewDeviceTasks(cfg *configs.Configs, mt *MeteoTask, ht *HumControlTask,
	tt *TempControlTask, lght *LightlTask, mot *MotionTask,
	secTask *SecurityTask) *DeviceTasks {
	return &DeviceTasks{
		cfg:        cfg,
		meteoTask:  mt,
		hctrlTask:  ht,
		tctrlTask:  tt,
		lightTask:  lght,
		motionTask: mot,
		secureTask: secTask,
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
	if d.cfg.Settings().Modules.Tempctrl {
		go d.tctrlTask.Start()
	}
	if d.cfg.Settings().Modules.Light {
		go d.lightTask.Start()
	}
	if d.cfg.Settings().Modules.Motion {
		go d.motionTask.Start()
	}
	if d.cfg.Settings().Modules.Security {
		go d.secureTask.Start()
	}
}
