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

package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FailResponse struct with fail response values
type FailResponse struct {
	Service string `json:"service"`
	Result  bool   `json:"result"`
	Reason  string `json:"reason"`
}

// OkResponse struct with ok response values
type OkResponse struct {
	Service string `json:"service"`
	Result  bool   `json:"result"`
}

// Response is used for simple answ to client
type Response struct {
	writer  *http.ResponseWriter
	service string
}

// NewResponse make new struct
func NewResponse(writer *http.ResponseWriter, service string) *Response {
	return &Response{
		writer:  writer,
		service: service,
	}
}

// SendFail simple sends fail result to client
func (r *Response) SendFail(reason string) {
	var failResp FailResponse

	failResp.Service = r.service
	failResp.Result = false
	failResp.Reason = reason

	byteStr, _ := json.Marshal(&failResp)
	(*r.writer).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*r.writer).WriteHeader(http.StatusForbidden)
	fmt.Fprintf(*r.writer, string(byteStr))
}

// SendOk simple sending successful result to client
func (r *Response) SendOk() {
	var okResp OkResponse

	okResp.Service = r.service
	okResp.Result = true

	byteStr, _ := json.Marshal(&okResp)
	(*r.writer).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*r.writer).WriteHeader(http.StatusOK)
	fmt.Fprintf(*r.writer, string(byteStr))
}

// Send abstraction of response sending
func (r *Response) Send(response string) {
	(*r.writer).Header().Set("Content-Type", "application/json; charset=utf-8")
	(*r.writer).WriteHeader(http.StatusOK)
	fmt.Fprintf(*r.writer, response)
}
