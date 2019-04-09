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

package webserver

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/futcamp/controller/devices"
	"github.com/futcamp/controller/net/webserver/handlers"
	"github.com/futcamp/controller/utils/configs"

	"github.com/google/logger"
)

const (
	AppName = "futcamp"
)

type WebServer struct {
	cfg         *configs.Configs
	meteo       *devices.MeteoStation
	meteoHdl    *handlers.MeteoHandler
	logHdl      *handlers.LogHandler
	monitorHdl  *handlers.MonitorHandler
	humCtrlHdl  *handlers.HumCtrlHandler
	tempCtrlHdl *handlers.TempCtrlHandler
	lightHdl    *handlers.LightHandler
	motionHdl   *handlers.MotionHandler
	securityHdl *handlers.SecurityHandler
}

// NewWebServer make new struct
func NewWebServer(cfg *configs.Configs, meteo *devices.MeteoStation, mh *handlers.MeteoHandler,
	lh *handlers.LogHandler, mnh *handlers.MonitorHandler,
	hch *handlers.HumCtrlHandler, tch *handlers.TempCtrlHandler,
	lgh *handlers.LightHandler, moth *handlers.MotionHandler, sech *handlers.SecurityHandler) *WebServer {
	return &WebServer{
		cfg:         cfg,
		meteo:       meteo,
		meteoHdl:    mh,
		logHdl:      lh,
		monitorHdl:  mnh,
		humCtrlHdl:  hch,
		tempCtrlHdl: tch,
		lightHdl:    lgh,
		motionHdl:   moth,
		securityHdl:sech,
	}
}

// Start run web server
func (w *WebServer) Start(ip string, port int) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", w.IndexHandler)
	mux.HandleFunc(fmt.Sprintf("/api/%s/log/", configs.ApiVersion), w.LogHandler)
	mux.HandleFunc(fmt.Sprintf("/api/%s/monitoring/", configs.ApiVersion), w.MonitorHandler)
	if w.cfg.Settings().Modules.Meteo {
		mux.HandleFunc(fmt.Sprintf("/api/%s/meteo/", configs.ApiVersion), w.MeteoHandler)
	}
	if w.cfg.Settings().Modules.Humctrl {
		mux.HandleFunc(fmt.Sprintf("/api/%s/humctrl/", configs.ApiVersion), w.HumCtrlHandler)
	}
	if w.cfg.Settings().Modules.Tempctrl {
		mux.HandleFunc(fmt.Sprintf("/api/%s/tempctrl/", configs.ApiVersion), w.TempCtrlHandler)
	}
	if w.cfg.Settings().Modules.Light {
		mux.HandleFunc(fmt.Sprintf("/api/%s/light/", configs.ApiVersion), w.LightHandler)
	}
	if w.cfg.Settings().Modules.Light {
		mux.HandleFunc(fmt.Sprintf("/api/%s/motion/", configs.ApiVersion), w.MotionHandler)
	}
	if w.cfg.Settings().Modules.Security {
		mux.HandleFunc(fmt.Sprintf("/api/%s/security/", configs.ApiVersion), w.SecurityHandler)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", ip, port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	return server.ListenAndServe()
}

// IndexHandler index handler with app information
func (w *WebServer) IndexHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, AppName)
	resp.SendOk()
}

// LogHandler logger handler
func (w *WebServer) LogHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	if len(args) == 5 && args[4] != "" {
		logs, err := w.logHdl.ProcessLogsByDate(args[4], req)
		if err != nil {
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(logs))
		return
	}

	logs, err := w.logHdl.ProcessExistingLogsList(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(logs))
}

// MeteoHandler meteo station requests handler
func (w *WebServer) MeteoHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	// Get single sensor data
	if len(args) == 5 && args[4] != "" {
		data, err := w.meteoHdl.ProcessMeteoSingleHandler(args[4], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	if len(args) == 6 && args[5] != "" {
		// Clear all sensor meteo data
		if args[5] == "clear" {
			err := w.meteoHdl.ProcessMeteoDBClearHandler(args[4], req)
			if err != nil {
				logger.Error(err.Error())
				resp.SendFail(err.Error())
				return
			}
			resp.SendOk()
			return
		}

		// Get sensor's meteo data by date
		data, err := w.meteoHdl.ProcessMeteoDBHandler(args[4], args[5], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	// Get all meteo sensors data
	data, err := w.meteoHdl.ProcessMeteoAllHandler(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(data))
	return
}

// MonitorHandler devices monitoring handler
func (w *WebServer) MonitorHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)

	devices, err := w.monitorHdl.ProcessMonitoring(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}

	resp.Send(string(devices))
}

// HumCtrlHandler hum control requests handler
func (w *WebServer) HumCtrlHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	// Change status
	if len(args) == 7 && args[5] == "status" && args[6] != "" {
		if req.Method == "PUT" {
			switch args[6] {
			case "on":
				err := w.humCtrlHdl.ProcessHumCtrlStatus(args[4], true, req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			case "off":
				err := w.humCtrlHdl.ProcessHumCtrlStatus(args[4], false, req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			case "switch":
				err := w.humCtrlHdl.ProcessHumCtrlSwitchStatus(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			default:
				resp.SendFail("Bad status request")
				return
			}
		} else {
			resp.SendFail("Bad status request")
			return
		}
	}

	// Change threshold
	if len(args) == 7 && args[5] == "threshold" && args[6] != "" {
		if req.Method == "PUT" {
			if args[5] == "threshold" {
				switch args[6] {
				case "plus":
					err := w.humCtrlHdl.ProcessHumCtrlThreshold(args[4], true, req)
					if err != nil {
						logger.Error(err.Error())
						resp.SendFail(err.Error())
						return
					}
					resp.SendOk()
					return

				case "minus":
					err := w.humCtrlHdl.ProcessHumCtrlThreshold(args[4], false, req)
					if err != nil {
						logger.Error(err.Error())
						resp.SendFail(err.Error())
						return
					}
					resp.SendOk()
					break

				default:
					resp.SendFail("Bad threshold request")
					break
				}
				return
			} else {
				resp.SendFail("Bad threshold request")
				return
			}
		} else {
			resp.SendFail("Bad threshold request method")
			return
		}
	}

	// Sync states with remote module
	if len(args) == 6 && args[5] != "" {
		if req.Method == "PUT" {
			if args[5] == "sync" {
				err := w.humCtrlHdl.ProcessHumCtrlSync(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return
			} else {
				resp.SendFail("Bad sync request")
				return
			}
		} else {
			resp.SendFail("Bad sync request method")
			return
		}
	}

	// Get single mod data
	if len(args) == 5 && args[4] != "" {
		data, err := w.humCtrlHdl.ProcessHumCtrlSingleHandler(args[4], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	// Get all humctrl devices data
	data, err := w.humCtrlHdl.ProcessHumCtrlAllHandler(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(data))
	return
}

// TempCtrlHandler temp control requests handler
func (w *WebServer) TempCtrlHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	// Change status
	if len(args) == 7 && args[5] == "status" && args[6] != "" {
		if req.Method == "PUT" {
			switch args[6] {
			case "on":
				err := w.tempCtrlHdl.ProcessTempCtrlStatus(args[4], true, req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			case "off":
				err := w.tempCtrlHdl.ProcessTempCtrlStatus(args[4], false, req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			case "switch":
				err := w.tempCtrlHdl.ProcessTempCtrlSwitchStatus(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			default:
				resp.SendFail("Bad status request")
				return
			}
		} else {
			resp.SendFail("Bad status request")
			return
		}
	}

	// Change threshold
	if len(args) == 7 && args[5] == "threshold" && args[6] != "" {
		if req.Method == "PUT" {
			if args[5] == "threshold" {
				switch args[6] {
				case "plus":
					err := w.tempCtrlHdl.ProcessTempCtrlThreshold(args[4], true, req)
					if err != nil {
						logger.Error(err.Error())
						resp.SendFail(err.Error())
						return
					}
					resp.SendOk()
					return

				case "minus":
					err := w.tempCtrlHdl.ProcessTempCtrlThreshold(args[4], false, req)
					if err != nil {
						logger.Error(err.Error())
						resp.SendFail(err.Error())
						return
					}
					resp.SendOk()
					break

				default:
					resp.SendFail("Bad threshold request")
					break
				}
				return
			} else {
				resp.SendFail("Bad threshold request")
				return
			}
		} else {
			resp.SendFail("Bad threshold request method")
			return
		}
	}

	// Sync states with remote module
	if len(args) == 6 && args[5] != "" {
		if req.Method == "PUT" {
			if args[5] == "sync" {
				err := w.tempCtrlHdl.ProcessTempCtrlSync(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return
			} else {
				resp.SendFail("Bad sync request")
				return
			}
		} else {
			resp.SendFail("Bad sync request method")
			return
		}
	}

	// Get single mod data
	if len(args) == 5 && args[4] != "" {
		data, err := w.tempCtrlHdl.ProcessTempCtrlSingleHandler(args[4], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	// Get all tempctrl devices data
	data, err := w.tempCtrlHdl.ProcessTempCtrlAllHandler(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(data))
	return
}

// LightHandler light requests handler
func (w *WebServer) LightHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	// Change status
	if len(args) == 7 && args[5] == "status" && args[6] != "" {
		if req.Method == "PUT" {
			switch args[6] {
			case "on":
				err := w.lightHdl.ProcessLightStatus(args[4], true, req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			case "off":
				err := w.lightHdl.ProcessLightStatus(args[4], false, req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			case "switch":
				if args[4] == "toilet-set" {
					w.lightHdl.ProcessToiletSwitchStatus(req)
					resp.SendOk()
					return
				}

				err := w.lightHdl.ProcessLightSwitchStatus(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			default:
				resp.SendFail("Bad status request")
				return
			}
		} else {
			resp.SendFail("Bad status request")
			return
		}
	}

	// Sync states with remote module
	if len(args) == 6 && args[5] != "" {
		if req.Method == "PUT" {
			if args[5] == "sync" {
				err := w.lightHdl.ProcessLightSync(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return
			} else {
				resp.SendFail("Bad sync request")
				return
			}
		} else {
			resp.SendFail("Bad sync request method")
			return
		}
	}

	// Get single mod data
	if len(args) == 5 && args[4] != "" {
		data, err := w.lightHdl.ProcessLightSingleHandler(args[4], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	// Get all light devices data
	data, err := w.lightHdl.ProcessLightAllHandler(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(data))
	return
}

// SecurityHandler security requests handler
func (w *WebServer) SecurityHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	// Change status
	if len(args) == 6 && args[4] == "status" && args[6] != "" {
		if req.Method == "PUT" {
			switch args[5] {
			case "on":
				w.securityHdl.ProcessSecurityStatus(true, req)
				resp.SendOk()
				return

			case "off":
				w.securityHdl.ProcessSecurityStatus(false, req)
				resp.SendOk()
				return

			default:
				resp.SendFail("Bad status request")
				return
			}
		} else {
			resp.SendFail("Bad status request")
			return
		}
	}

	// Open request
	if len(args) == 6 && args[5] != "" {
		if req.Method == "PUT" {
			switch (args[5]) {
			case "open":
				err := w.securityHdl.ProcessSecurityOpenHandler(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return

			default:
				resp.SendFail("Bad sync request")
				return
			}
		} else {
			resp.SendFail("Bad sync request method")
			return
		}
	}

	// Open request
	if len(args) == 6 && args[5] != "" {
		if req.Method == "PUT" {
			if args[5] == "open" {
				err := w.securityHdl.ProcessSecurityOpenHandler(args[4], req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.SendOk()
				return
			} else {
				resp.SendFail("Bad open request")
				return
			}
		} else {
			resp.SendFail("Bad open request method")
			return
		}
	}

	// Keys request
	if len(args) == 5 && args[4] != "" {
		if req.Method == "GET" {
			if args[4] == "keys" {
				data, err := w.securityHdl.ProcessSecurityKeys(req)
				if err != nil {
					logger.Error(err.Error())
					resp.SendFail(err.Error())
					return
				}
				resp.Send(string(data))
			}
		} else {
			resp.SendFail("Bad keys request method")
			return
		}
	}

	// Get single mod data
	if len(args) == 5 && args[4] != "" {
		data, err := w.securityHdl.ProcessSecuritySingleHandler(args[4], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	// Get all light devices data
	data, err := w.securityHdl.ProcessSecurityAllHandler(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(data))
	return
}

// MotionHandler light requests handler
func (w *WebServer) MotionHandler(writer http.ResponseWriter, req *http.Request) {
	resp := NewResponse(&writer, configs.AppName)
	args := strings.Split(req.RequestURI, "/")

	// Activity states with remote module
	if len(args) == 6 && args[5] != "" {
		if req.Method == "PUT" {
			if args[5] == "activity" {
				w.motionHdl.ProcessMotionActivity(args[4], req)
				resp.SendOk()
				return
			} else {
				resp.SendFail("Bad sync request")
				return
			}
		} else {
			resp.SendFail("Bad sync request method")
			return
		}
	}

	// Get single mod data
	if len(args) == 5 && args[4] != "" {
		data, err := w.motionHdl.ProcessMotionSingleHandler(args[4], req)
		if err != nil {
			logger.Error(err.Error())
			resp.SendFail(err.Error())
			return
		}
		resp.Send(string(data))
		return
	}

	// Get all light devices data
	data, err := w.motionHdl.ProcessMotionAllHandler(req)
	if err != nil {
		logger.Error(err.Error())
		resp.SendFail(err.Error())
		return
	}
	resp.Send(string(data))
	return
}
