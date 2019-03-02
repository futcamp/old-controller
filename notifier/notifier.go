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

package notifier

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"strconv"
)

type Notifier struct {
	name   string
	apiKey string
	chatID []string
}

// NewNotifier make new struct
func NewNotifier() *Notifier {
	return &Notifier{}
}

// SetName set new name for notifier
func (n *Notifier) SetName(name string) {
	n.name = name
}

// SetApiKey set new telegram bot API key
func (n *Notifier) SetApiKey(key string) {
	n.apiKey = key
}

// AddChatID add new telegram chat id
func (n *Notifier) AddChatID(id string) {
	n.chatID = append(n.chatID, id)
}

// SendNotify send notify to chat
func (n *Notifier) SendNotify(module string, message string) error {
	bot, err := tgbotapi.NewBotAPI(n.apiKey)
	if err != nil {
		return err
	}
	bot.Debug = false

	for _, chatID := range n.chatID {
		id, err := strconv.Atoi(chatID)
		if err != nil {
			return err
		}
		msg := tgbotapi.NewMessage(int64(id), fmt.Sprintf("<b>%s</b> %s", module, message))
		msg.ParseMode = "HTML"
		bot.Send(msg)
	}

	return nil
}
