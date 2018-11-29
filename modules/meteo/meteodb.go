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

package meteo

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type MeteoDBData struct {
	Sensor   string
	Temp     int
	Humidity int
	Pressure int
	Altitude int
	Time     string
	Date     string
}

type MeteoDB struct {
	FileName string
	Database *sql.DB
	Mtx      sync.Mutex
}

// NewMeteoDB make new struct
func NewMeteoDB(fname string) *MeteoDB {
	return &MeteoDB{
		FileName: fname,
	}
}

// Load open database
func (m *MeteoDB) Load() error {
	var err error

	m.Database, err = sql.Open("sqlite3", m.FileName)
	if err != nil {
		return err
	}

	return nil
}

// AddMeteoData add new record with meteo data
func (m *MeteoDB) AddMeteoData(data *MeteoDBData) error {
	date := time.Now().Format("2006-01-02")
	hour := time.Now().Hour()

	rows, err := m.Database.Query(
		fmt.Sprintf("SELECT id FROM %s WHERE time=\"%s\" AND date=\"%s\"",
			data.Sensor, fmt.Sprintf("%d:00", hour), date))
	if err != nil {
		return err
	}
	if rows.Next() {
		// Record in table already exists
		return nil
	}

	stmt, err := m.Database.Prepare(fmt.Sprintf(
		"INSERT INTO %s(temp,humidity,pressure,time,date) VALUES (%d,%d,%d,\"%s\",\"%s\")",
			data.Sensor, data.Temp, data.Humidity, data.Pressure,
			fmt.Sprintf("%d:00", hour), date))
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

// MeteoDataByDate read sensor data by date
func (m *MeteoDB) MeteoDataByDate(sensor string, date string) ([]MeteoDBData, error) {
	var data []MeteoDBData

	rows, err := m.Database.Query(
		fmt.Sprintf("SELECT temp,humidity,pressure,time,date FROM %s WHERE date=\"%s\"",
			sensor, date))
	if err != nil {
		return data, err
	}

	for rows.Next() {
		var datum MeteoDBData

		err = rows.Scan(
			&datum.Temp, &datum.Humidity, &datum.Pressure, &datum.Time,
			&datum.Date)
		if err != nil {
			return data, err
		}

		datum.Sensor = sensor
		data = append(data, datum)
	}

	return data, nil
}

// MeteoDataClear clear meteo values from sensor table
func (m *MeteoDB) MeteoDataClear(sensor string) error {
	stmt, err := m.Database.Prepare(fmt.Sprintf(
		"DELETE FROM %s", sensor))
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

// Unload close database
func (m *MeteoDB) Unload() {
	if m.Database != nil {
		m.Database.Close()
	}
}
