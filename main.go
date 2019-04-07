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

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/devices/db"
	"github.com/futcamp/controller/devices/tasks"
	"github.com/futcamp/controller/monitoring"
	"github.com/futcamp/controller/net/rcli"
	"github.com/futcamp/controller/net/webserver"
	"github.com/futcamp/controller/net/webserver/handlers"
	"github.com/futcamp/controller/notifier"
	"github.com/futcamp/controller/updater"
	"github.com/futcamp/controller/utils"
	"github.com/futcamp/controller/utils/configs"
	"github.com/futcamp/controller/utils/configs/cfgtask"
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
	container.Provide(startup.NewStartup)
	container.Provide(tasks.NewDeviceTasks)

	container.Provide(devices.NewMeteoStation)
	container.Provide(tasks.NewMeteoTask)
	container.Provide(db.NewMeteoDatabase)
	container.Provide(devices.NewMeteoDisplay)

	container.Provide(devices.NewHumidityControl)
	container.Provide(tasks.NewHumControlTask)
	container.Provide(handlers.NewHumCtrlHandler)

	container.Provide(devices.NewTemperatureControl)
	container.Provide(tasks.NewTempControlTask)
	container.Provide(handlers.NewTempCtrlHandler)

	container.Provide(devices.NewLight)
	container.Provide(tasks.NewLightlTask)
	container.Provide(handlers.NewLightHandler)

	container.Provide(devices.NewMotion)
	container.Provide(tasks.NewMotionTask)
	container.Provide(handlers.NewMotionHandler)

	container.Provide(devices.NewSecurity)
	container.Provide(handlers.NewSecurityHandler)

	container.Provide(handlers.NewMeteoHandler)
	container.Provide(monitoring.NewDeviceMonitor)
	container.Provide(monitoring.NewMonitorTask)
	container.Provide(handlers.NewMonitorHandler)

	container.Provide(configs.NewConfigs)
	container.Provide(configs.NewDynamicConfigs)
	container.Provide(cfgtask.NewDynConfigsTask)
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
