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

package main

import (
	"fmt"
	"github.com/futcamp/controller/updater"

	"github.com/futcamp/controller/modules/meteo"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/net/rcli"
	"github.com/futcamp/controller/net/webserver"
	"github.com/futcamp/controller/net/webserver/handlers"
	"github.com/futcamp/controller/notifier"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"
	"github.com/futcamp/controller/utils/startup"
	"github.com/futcamp/controller/utils/startup/io"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	container.Provide(utils.NewLogger)
	container.Provide(utils.NewLogTask)
	container.Provide(utils.NewLocker)
	container.Provide(handlers.NewLogHandler)
	container.Provide(notifier.NewNotifier)
	container.Provide(updater.NewUpdater)

	container.Provide(meteo.NewMeteoStation)
	container.Provide(meteo.NewMeteoTask)
	container.Provide(meteo.NewMeteoDatabase)
	container.Provide(meteo.NewMeteoDisplays)
	container.Provide(startup.NewStartup)

	container.Provide(handlers.NewMeteoHandler)
	container.Provide(monitoring.NewDeviceMonitor)
	container.Provide(monitoring.NewMonitorTask)
	container.Provide(handlers.NewMonitorHandler)

	container.Provide(configs.NewConfigs)
	container.Provide(configs.NewDynamicConfigs)
	container.Provide(io.NewStartupIO)
	container.Provide(io.NewStartupMods)

	container.Provide(webserver.NewWebServer)
	container.Provide(rcli.NewRCliServer)
	container.Provide(NewApplication)

	err := container.Invoke(func(app *Application) {
		app.Start()
		app.Free()
	})

	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
	}
}
