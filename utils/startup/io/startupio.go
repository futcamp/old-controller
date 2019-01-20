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

package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StartupCmd struct {
	Module  string
	Command string
	Device  string
	Args    []string
}

type StartupIO struct {
	cmds []*StartupCmd
}

// NewStartupIO make new struct
func NewStartupIO() *StartupIO {
	return &StartupIO{}
}

// LoadCommands load commands from file
func (s *StartupIO) LoadCommands(fileName string, cmdHandler func(*StartupCmd)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text()[0] == '!' {
			continue
		}

		parts := strings.Split(scanner.Text(), " ")

		newCmd := &StartupCmd{
			Module:  parts[0],
			Command: parts[1],
			Device:  parts[2],
		}

		if len(parts) > 3 {
			for i := 3; i < len(parts); i++ {
				newCmd.Args = append(newCmd.Args, parts[i])
			}
		}

		cmdHandler(newCmd)
		s.cmds = append(s.cmds, newCmd)
	}

	return nil
}

// SaveCommands save all commands to file
func (s *StartupIO) SaveCommands(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	lastModule := ""
	lastDevice := ""
	for _, curCmd := range s.cmds {
		if lastModule != curCmd.Module {
			file.WriteString("!\n")
			lastModule = curCmd.Module
		}
		if lastDevice != curCmd.Device {
			file.WriteString("!\n")
			lastDevice = curCmd.Device
		}

		outStr := fmt.Sprintf("%s %s %s", curCmd.Module, curCmd.Command, curCmd.Device)
		for _, arg := range curCmd.Args {
			outStr += " "
			outStr += arg
		}
		outStr += "\n"

		file.WriteString(outStr)
	}
	return nil
}

// AddCommand add new command
func (s *StartupIO) AddCommand(mod string, cmd string, dev string, args []string) {
	var newCmds []*StartupCmd

	for _, curCmd := range s.cmds {
		if curCmd.Module == mod {
			if curCmd.Device == dev {
				if curCmd.Command == cmd {
					curCmd.Args = args
					return
				}
			}
		}
	}

	c := &StartupCmd{
		Module:  mod,
		Command: cmd,
		Device:  dev,
		Args:    args,
	}

	found := false
	for _, curCmd := range s.cmds {
		if curCmd.Module == mod && curCmd.Device == dev {
			found = true
			if curCmd.Command[0] == 'a' && curCmd.Command[1] == 'd' && curCmd.Command[2] == 'd' {
				newCmds = append(newCmds, curCmd)
				newCmds = append(newCmds, c)
			} else {
				newCmds = append(newCmds, c)
				newCmds = append(newCmds, curCmd)
			}
		} else {
			newCmds = append(newCmds, curCmd)
		}
	}

	if !found {
		s.cmds = append(s.cmds, c)
	} else {
		s.cmds = newCmds
	}
}

// DeleteCommand delete command
func (s *StartupIO) DeleteCommand(mod string, cmd string, dev string) {
	var newCmds []*StartupCmd

	// All delete
	if cmd[0] == 'a' && cmd[1] == 'd' && cmd[2] == 'd' {
		for _, curCmd := range s.cmds {
			if !(curCmd.Module == mod && curCmd.Device == dev) {
				newCmds = append(newCmds, curCmd)
			}
		}
		s.cmds = make([]*StartupCmd, len(newCmds))
		copy(s.cmds, newCmds)
		return
	}

	// Single delete
	for _, curCmd := range s.cmds {
		if !(curCmd.Module == mod && curCmd.Device == dev && curCmd.Command == cmd) {
			newCmds = append(newCmds, curCmd)
		}
	}

	s.cmds = make([]*StartupCmd, len(newCmds))
	copy(s.cmds, newCmds)
}
