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

package main

import (
	"fmt"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/net"
	"github.com/futcamp/controller/net/handlers"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	container.Provide(utils.NewLogger)
	container.Provide(utils.NewLogTask)
	container.Provide(utils.NewLocker)
	container.Provide(handlers.NewLogHandler)
	container.Provide(configs.NewConfigs)
	container.Provide(configs.NewMeteoConfigs)
	container.Provide(meteo.NewMeteoStation)
	container.Provide(meteo.NewMeteoTask)
	container.Provide(handlers.NewMeteoHandler)
	container.Provide(meteo.NewMeteoDatabase)
	container.Provide(net.NewWebServer)
	container.Provide(NewApplication)

	err := container.Invoke(func(app *Application) {
		app.Start()
		app.Free()
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
