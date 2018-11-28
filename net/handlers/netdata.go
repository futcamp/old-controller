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

package handlers

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/futcamp/controller/utils/configs"
)

type RestResponse struct {
	Module      string      `json:"module"`
	Description string      `json:"description"`
	Api         string      `json:"api"`
	Time        string      `json:"time"`
	Host        string      `json:"host"`
	Protocol    string      `json:"protocol"`
	Method      string      `json:"method"`
	Sensors     interface{} `json:"sensors"`
	Hash        string      `json:"hash"`
	HashType    string      `json:"hash_type"`
}

// SetRestResponse set values for response
func SetRestResponse(restData *RestResponse, module string, desc string,
	sensors interface{}, req *http.Request) {
	// Add info
	restData.Module = module
	restData.Description = desc
	restData.Api = configs.ApiVersion
	restData.Time = time.Now().Format("2006.01.02 15:04:06")
	restData.Hash = ""
	restData.HashType = "SHA512"
	restData.Host = req.Host
	restData.Protocol = req.Proto
	restData.Method = req.Method
	restData.Sensors = sensors

	// Calc response hash
	byteData, _ := json.Marshal(restData)
	h := sha512.New()
	h.Write(byteData)
	hOut := h.Sum(nil)
	for i := 0; i < len(hOut); i++ {
		restData.Hash += fmt.Sprintf("%x", hOut[i])
	}
}
