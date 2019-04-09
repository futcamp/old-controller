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

package hardware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	HdkTimeout = 100
)

// HdkHttpSyncData sync data witch remote controller by http request
func HdkHttpSyncData(ip string, request string, response interface{}) error {
	timeout := time.Duration(HdkTimeout * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Get(fmt.Sprintf("http://%s/%s", ip, request))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if response != nil {
		byteBuf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(byteBuf, response)
		if err != nil {
			return err
		}
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
	}

	return nil
}
