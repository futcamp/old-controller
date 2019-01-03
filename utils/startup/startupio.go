/*******************************************************************/
/*
/* Future Camp Project
/*
/* Copyright (C) 2018 Sergey Denisov.
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
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StartupCmd struct {
	Module  string
	Command string
	Device  string
	Type    string
	Value   interface{}
}

type StartupIO struct {
	cmds []*StartupCmd
}

// NewStartupIO make new struct
func NewStartupIO() *StartupIO {
	return &StartupIO{}
}

// AddCommand add new command to list of commands
func (s *StartupIO) AddCommand(module string, cmd string, dev string, valType string, value interface{}) {
	var newCmds []*StartupCmd
	var found bool

	newCmd := &StartupCmd{
		Module:  module,
		Command: cmd,
		Device:  dev,
		Type:    valType,
		Value:   value,
	}

	found = false
	for _, cmd := range s.cmds {
		if cmd.Module == newCmd.Module && cmd.Command == newCmd.Command && cmd.Device == newCmd.Device {
			cmd.Value = newCmd.Value
			newCmds = append(newCmds, cmd)
			found = true
		} else {
			newCmds = append(newCmds, cmd)
		}
	}

	if !found {
		newCmds = append(newCmds, newCmd)
	}

	copy(s.cmds, newCmds)
}

// DeleteCommand delete command from list of commands
func (s *StartupIO) DeleteCommand(module string, cmd string, dev string) {
	var newCmds []*StartupCmd

	for _, c := range s.cmds {
		if c.Module != module || c.Command != cmd || c.Device != dev {
			newCmds = append(newCmds, c)
		}
	}

	if len(newCmds) == 0 {
		s.cmds = make([]*StartupCmd, 0)
	} else {
		copy(s.cmds, newCmds)
	}
}

// SaveCommands save all commands to startup-configs file
func (s *StartupIO) SaveCommands(fileName string) error {
	var strCmd string

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, cmd := range s.cmds {
		if cmd.Device == "" {
			cmd.Device = "nil"
		}

		switch cmd.Type {
		case "string":
			strCmd = fmt.Sprintf("%s %s %s %s %s", cmd.Module, cmd.Command, cmd.Device, cmd.Type, cmd.Value.(string))
			break

		case "bool":
			strCmd = fmt.Sprintf("%s %s %s %s %t", cmd.Module, cmd.Command, cmd.Device, cmd.Type, cmd.Value.(bool))
			break

		case "int":
			strCmd = fmt.Sprintf("%s %s %s %s %d", cmd.Module, cmd.Command, cmd.Device, cmd.Type, cmd.Value.(int))
			break
		}

		_, err = file.WriteString(strCmd)
		if err != nil {
			return err
		}
	}

	return nil
}

// LoadCommands load all commands from startup-configs file
func (s *StartupIO) LoadCommands(fileName string, cmdHandler func(string, string, string, interface{})) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		newCmd := &StartupCmd{
			Module:  parts[0],
			Command: parts[1],
			Device:  parts[2],
			Type:    parts[3],
		}

		switch newCmd.Type {
		case "string":
			newCmd.Value = parts[4]
			cmdHandler(newCmd.Module, newCmd.Command, newCmd.Device, newCmd.Value)
			break

		case "bool":
			newCmd.Value, err = strconv.ParseBool(parts[4])
			if err != nil {
				return err
			}
			cmdHandler(newCmd.Module, newCmd.Command, newCmd.Device, newCmd.Value)
			break

		case "int":
			newCmd.Value, err = strconv.Atoi(parts[4])
			if err != nil {
				return err
			}
			cmdHandler(newCmd.Module, newCmd.Command, newCmd.Device, newCmd.Value)
			break
		}

		s.cmds = append(s.cmds, newCmd)
	}

	return nil
}
