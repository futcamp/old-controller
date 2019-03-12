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

import "sync"

type MeteoDBCfg struct {
	IP     string
	Port   int
	User   string
	Passwd string
	Base   string
}

type TimersCfg struct {
	MeteoSensorsDelay int
	MeteoDBDelay      int
	MonitorDelay      int
	DisplayDelay      int
}

type RCliCfg struct {
	UserHash string
}

type Settings struct {
	MeteoDB MeteoDBCfg
	Timers  TimersCfg
	RCli    RCliCfg
}

type DynamicConfigs struct {
	set     Settings
	isSave  bool
	cmds    []string
	mtxSave sync.Mutex
	mtxCmd  sync.Mutex
}

// NewDynamicConfigs make new struct
func NewDynamicConfigs() *DynamicConfigs {
	return &DynamicConfigs{}
}

// AddCommand add new command
func (d *DynamicConfigs) AddCommand(cmd string) {
	d.mtxCmd.Lock()
	d.cmds = append(d.cmds, cmd)
	d.mtxCmd.Unlock()
}

// Commands get new dynamic commands
func (d *DynamicConfigs) Commands() []string {
	var newCmds []string

	d.mtxCmd.Lock()
	newCmds = make([]string, len(d.cmds))
	copy(newCmds, d.cmds)
	d.cmds = make([]string, 0)
	d.mtxCmd.Unlock()

	return newCmds
}

// Settings get settings pointer
func (d *DynamicConfigs) Settings() *Settings {
	return &d.set
}

// SaveConfigs set save configs flag
func (d *DynamicConfigs) SaveConfigs() {
	d.mtxSave.Lock()
	d.isSave = true
	d.mtxSave.Unlock()
}

// SaveConfigs reset save configs flag
func (d *DynamicConfigs) ResetSaveConfigs() {
	d.mtxSave.Lock()
	d.isSave = false
	d.mtxSave.Unlock()
}

// GetSaveConfigs get save configs flag
func (d *DynamicConfigs) GetSaveConfigs() bool {
	var save bool

	d.mtxSave.Lock()
	save = d.isSave
	d.mtxSave.Unlock()

	return save
}
