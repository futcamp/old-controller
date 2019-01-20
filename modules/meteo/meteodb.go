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

package meteo

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
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

type MeteoDatabase struct {
	Database *sql.DB
	FileName string
}

// NewMeteoDatabase make new struct
func NewMeteoDatabase() *MeteoDatabase {
	return &MeteoDatabase{
	}
}

// Load open database
func (m *MeteoDatabase) Connect(ip string, user string, passwd string, db string) error {
	var err error

	m.Database, err = sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s",
		user, passwd, ip, db))
	if err != nil {
		return err
	}

	return nil
}

// AddMeteoData add new record with meteo data
func (m *MeteoDatabase) AddMeteoData(data *MeteoDBData) error {
	date := time.Now().Format("2006-01-02")
	hour := time.Now().Hour()

	rows, err := m.Database.Query(
		fmt.Sprintf("SELECT id FROM %s WHERE time = '%s' AND date = '%s'",
			data.Sensor, fmt.Sprintf("%d:00", hour), date))
	if err != nil {
		return err
	}

	if rows.Next() {
		// Record in table already exists
		rows.Close()
		return nil
	}
	rows.Close()

	_, err = m.Database.Exec(fmt.Sprintf(
		"INSERT INTO %s(temp,humidity,pressure,time,date) VALUES ($1,$2,$3,$4,$5)", data.Sensor),
		data.Temp, data.Humidity, data.Pressure, fmt.Sprintf("%d:00", hour), date)
	if err != nil {
		return err
	}

	return nil
}

// MeteoDataByDate read sensor data by date
func (m *MeteoDatabase) MeteoDataByDate(sensor string, date string) ([]MeteoDBData, error) {
	var data []MeteoDBData

	rows, err := m.Database.Query(
		fmt.Sprintf("SELECT temp,humidity,pressure,time,date FROM %s WHERE date = '%s'",
			sensor, date))
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var mdata MeteoDBData

		err = rows.Scan(
			&mdata.Temp, &mdata.Humidity, &mdata.Pressure, &mdata.Time,
			&mdata.Date)
		if err != nil {
			return data, err
		}

		mdata.Sensor = sensor
		data = append(data, mdata)
	}

	return data, nil
}

// MeteoDataClear clear meteo values from sensor table
func (m *MeteoDatabase) MeteoDataClear(sensor string) error {
	_, err := m.Database.Exec(fmt.Sprintf("TRUNCATE %s", sensor))
	if err != nil {
		return err
	}

	return nil
}

// Unload close database
func (m *MeteoDatabase) Close() {
	m.Database.Close()
}
