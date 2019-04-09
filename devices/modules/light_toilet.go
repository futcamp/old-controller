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

package modules

type Lamp interface {
	Name() string
	LastState() bool
	SetState(state bool)
}

type ToiletLamp struct {
	name  string
	state bool
}

// NewToiletLamp make new struct
func NewToiletLamp(name string) *ToiletLamp {
	return &ToiletLamp{
		name:  name,
		state: false,
	}
}

// Name get lamp name
func (t *ToiletLamp) Name() string {
	return t.name
}

// LastState get lamp state
func (t *ToiletLamp) LastState() bool {
	return t.state
}

// SetState set new lamp state
func (t *ToiletLamp) SetState(state bool) {
	t.state = state
}
