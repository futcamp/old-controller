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

package updater

import (
	"net/http"

	"github.com/inconshreveable/go-update"
)

type Updater struct {
}

// NewUpdater make new struct
func NewUpdater() *Updater {
	return &Updater{}
}

// Update update application
func (u *Updater) Update(url string) error {
	re, err := http.Get(url)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	err = update.Apply(re.Body, update.Options{})
	if err != nil {
		return err
	}

	return nil
}