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
	"context"
	"fmt"
	"net"

	"github.com/futcamp/controller/utils/startup"
	pb "github.com/futcamp/controller/net/rcli/rcliproto"

	"google.golang.org/grpc"
)

type RCliServer struct {
	startup *startup.Startup
	userHash string
}

// NewRCliServer make new struct
func NewRCliServer() *RCliServer {
	return &RCliServer{
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

	err := r.startup.ExecCmd(in.Command)
	if err != nil {
		resp.Result = false
		return resp, err
	}

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