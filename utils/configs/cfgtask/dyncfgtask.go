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

package cfgtask

import (
	"time"

	"github.com/futcamp/controller/utils/configs"
	"github.com/futcamp/controller/utils/startup"

	"github.com/google/logger"
)

const (
	taskDelay = 1
)

// DynConfigsTask dynamic configs task struct
type DynConfigsTask struct {
	reqTimer *time.Timer
	dynCfg   *configs.DynamicConfigs
	startup  *startup.Startup
}

// NewDynConfigsTask make new struct
func NewDynConfigsTask(dc *configs.DynamicConfigs, st *startup.Startup) *DynConfigsTask {
	return &DynConfigsTask{
		dynCfg:  dc,
		startup: st,
	}
}

// TaskHandler process timer loop
func (d *DynConfigsTask) TaskHandler() {
	for {
		<-d.reqTimer.C

		if d.dynCfg.GetSaveConfigs() {
			for _, cmd := range d.dynCfg.Commands() {
				d.startup.AddCmd(cmd)
			}

			logger.Info("DynamicConfigs saving application configs")

			d.startup.SaveAll()
			d.dynCfg.ResetSaveConfigs()
		}

		d.reqTimer.Reset(taskDelay * time.Second)
	}
}

// Start start new timer
func (d *DynConfigsTask) Start() {
	d.reqTimer = time.NewTimer(taskDelay * time.Second)
	d.TaskHandler()
}
