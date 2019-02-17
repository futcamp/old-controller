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

package rcli

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"

	pb "github.com/futcamp/controller/net/rcli/rcliproto"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/startup"

	"github.com/google/logger"
	"google.golang.org/grpc"
)

type RCliServer struct {
	startup  *startup.Startup
	userHash string
}

// NewRCliServer make new struct
func NewRCliServer(stp *startup.Startup) *RCliServer {
	return &RCliServer{
		startup: stp,
	}
}

// SetHash set new user hash
func (r *RCliServer) SetHash(hash string) {
	r.userHash = hash
}

// LoginCheck process login check request
func (r *RCliServer) LoginCheck(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp := &pb.LoginResponse{}

	if in.LoginHash == r.userHash {
		resp.Result = true
	} else {
		resp.Result = false
	}

	return resp, nil
}

// SendCmd process send command request
func (r *RCliServer) SendCmd(ctx context.Context, in *pb.CmdRequest) (*pb.CmdResponse, error) {
	resp := &pb.CmdResponse{}

	if in.LoginHash != r.userHash {
		resp.Result = false
		return resp, nil
	}

	logger.Infof("Remote CLI command applying...")

	if in.Command == "reload" {
		logger.Infof("Rebooting the system...")
		exec.Command("reboot")
		resp.Result = true
		return resp, nil
	}

	if in.Command == "write startup" {
		err := r.startup.SaveAll()
		if err != nil {
			resp.Result = false
			return resp, err
		}
		resp.Result = true
		return resp, nil
	}

	err := r.startup.ExecCmd(in.Command)
	if err != nil {
		resp.Result = false
		return resp, err
	}

	err = r.startup.SaveAll()
	if err != nil {
		resp.Result = false
		return resp, err
	}

	resp.Result = true
	return resp, nil
}

// ComandsList send startup-configs
func (r *RCliServer) ComandsList(ctx context.Context, in *pb.CmdListRequest) (*pb.CmdListResponse, error) {
	resp := &pb.CmdListResponse{}

	if in.LoginHash != r.userHash {
		resp.Result = false
		return resp, nil
	}

	file, err := os.Open(utils.StartupCfgPath)
	if err != nil {
		resp.Result = false
		return resp, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resp.Cmds = append(resp.Cmds, scanner.Text())
	}

	resp.Result = true
	return resp, nil
}

// UpdateStartup update startup-configs
func (r *RCliServer) UpdateStartup(ctx context.Context, in *pb.UpdateStartupRequest) (*pb.UpdateStartupResponse, error) {
	resp := &pb.UpdateStartupResponse{}

	if in.LoginHash != r.userHash {
		resp.Result = false
		return resp, nil
	}

	file, err := os.OpenFile(utils.StartupCfgPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		resp.Result = false
		return resp, err
	}
	defer file.Close()

	for _, cmd := range in.Cmds {
		file.WriteString(cmd)
		file.WriteString("\n")
	}

	resp.Result = true
	return resp, nil
}

// SendLinuxCmd process send linux command
func (r *RCliServer) SendLinuxCmd(ctx context.Context, in *pb.LinuxCmdRequest) (*pb.LinuxCmdResponse, error) {
	var cmd *exec.Cmd

	resp := &pb.LinuxCmdResponse{}

	if in.LoginHash != r.userHash {
		resp.Result = false
		return resp, nil
	}

	if in.Args == "" {
		cmd = exec.Command(in.Command)
	} else if in.SubArgs == "" {
		cmd = exec.Command(in.Command, in.Args)
	} else {
		cmd = exec.Command(in.Command, in.Args, in.SubArgs)
	}

	output, err := cmd.Output()
	if err != nil {
		resp.Result = false
		return resp, err
	}

	resp.Output = string(output)

	resp.Result = true
	return resp, nil
}

// Start start gRPC remote CLI server
func (r *RCliServer) Start(addr string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterRemoteCliServer(s, r)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
