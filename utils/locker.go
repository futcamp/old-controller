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

package utils

import (
	"sync"

	"github.com/pkg/errors"
)

type Locker struct {
	Locks map[string]*sync.Mutex
}

// NewLocker make new struct
func NewLocker() *Locker {
	l := &Locker{}

	l.Locks = make(map[string]*sync.Mutex)

	return l
}

// AddLock add new lock
func (l *Locker) AddLock(name string) {
	l.Locks[name] = &sync.Mutex{}
}

// Lock lock mutex
func (l *Locker) Lock(name string) error {
	lck := l.Locks[name]

	if lck == nil {
		return errors.New("Lock not found")
	}
	lck.Lock()

	return nil
}

// Unlock unlock mutex
func (l *Locker) Unlock(name string) error {
	lck := l.Locks[name]

	if lck == nil {
		return errors.New("Lock not found")
	}
	lck.Unlock()

	return nil
}