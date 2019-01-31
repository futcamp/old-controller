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

package startup

import (
	"strings"

	"github.com/futcamp/controller/utils/startup/io"
)

const (
	FirstParam = 3
)

type Startup struct {
	mods     *io.StartupMods
	fileName string
}

// NewStartup make new struct
func NewStartup(m *io.StartupMods) *Startup {
	return &Startup{
		mods: m,
	}
}

// Load load startup-configs file
func (s *Startup) Load(fileName string) error {
	s.fileName = fileName
	return s.mods.LoadFromFile(fileName)
}

// ExecCmd exec new command
func (s *Startup) ExecCmd(cmd string) error {
	var params []string

	parts := strings.Split(cmd, " ")

	if parts[0] == "no" {
		return s.mods.DeleteModCommand(s.fileName, parts[0], parts[1], parts[2])
	} else {
		if len(parts) > FirstParam {
			for i := FirstParam; i < len(parts); i++ {
				params = append(params, parts[i])
			}
		}
		return s.mods.ExecModCommand(s.fileName, parts[0], parts[1], parts[2], params)
	}
	return nil
}

// SaveAll save all commands to startup-configs file
func (s *Startup) SaveAll() error {
	return s.mods.SaveCommands(s.fileName)
}
