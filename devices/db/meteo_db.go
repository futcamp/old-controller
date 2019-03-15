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

package db

import (
	"fmt"
	"time"

	"github.com/futcamp/controller/utils"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type MeteoDBData struct {
	Sensor   string
	Temp     int
	Humidity int
	Pressure int
	Time     string
	Date     string
}

type DBData struct {
	Temp int
	Hum  int
	Pres int
	Time string
	Date string
}

type MeteoDatabase struct {
	engine *xorm.Engine
	locker *utils.Locker
}

// NewMeteoDatabase make new struct
func NewMeteoDatabase(lck *utils.Locker) *MeteoDatabase {
	return &MeteoDatabase{
		locker: lck,
	}
}

// Load open database
func (m *MeteoDatabase) Connect(ip string, user string, passwd string, db string) error {
	var err error

	m.locker.Lock(utils.MeteoDBName)

	m.engine, err = xorm.NewEngine("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s",
		user, passwd, ip, db))
	if err != nil {
		m.locker.Unlock(utils.MeteoDBName)
		return err
	}

	return nil
}

// AddMeteoData add new record with meteo data
func (m *MeteoDatabase) AddMeteoData(data *MeteoDBData) error {
	date := time.Now().Format("2006-01-02")
	hour := time.Now().Hour()
	nTime := fmt.Sprintf("%d:00", hour)

	isExist, err := m.engine.Table(data.Sensor).Exist(&DBData{
		Time: nTime,
		Date: date,
	})
	if err != nil {
		return err
	}
	if isExist {
		return nil
	}

	mData := DBData{
		Temp: data.Temp,
		Hum:  data.Humidity,
		Pres: data.Pressure,
		Time: nTime,
		Date: date,
	}

	_, err = m.engine.Table(data.Sensor).Insert(&mData)
	if err != nil {
		return err
	}

	return nil
}

// MeteoDataByDate read sensor data by date
func (m *MeteoDatabase) MeteoDataByDate(sensor string, date string) ([]MeteoDBData, error) {
	var data []MeteoDBData

	rows, err := m.engine.Table(sensor).Rows(&DBData{Date: date})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dbData := new(DBData)
	for rows.Next() {
		datum := MeteoDBData{}
		rows.Scan(dbData)

		datum.Temp = dbData.Temp
		datum.Humidity = dbData.Hum
		datum.Pressure = dbData.Pres

		data = append(data, datum)
	}

	return data, nil
}

// MeteoDataClear clear meteo values from sensor table
func (m *MeteoDatabase) MeteoDataClear(sensor string) error {
	_, err := m.engine.Table(sensor).Delete(&DBData{})
	if err != nil {
		return err
	}

	return nil
}

// Unload close database
func (m *MeteoDatabase) Close() {
	m.engine.Close()
	m.locker.Unlock(utils.MeteoDBName)
}
