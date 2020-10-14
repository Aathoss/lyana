package rcon

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

var (
	c       *Client
	connect bool
)

func RconCommandeList() error {
	openRcon()

	response, err := c.SendCommand("list")
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
		connect = false
		return err
	}

	respFix, err := after(response, "There")
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
		connect = false
		return err
	}

	messageSplit := strings.Fields(respFix)
	framework.OnlinePlayer, err = strconv.Atoi(messageSplit[2])
	if err != nil {
		logger.DebugLogger.Println(respFix)
	}
	framework.MaxOnlinePlayer, err = strconv.Atoi(messageSplit[7])
	if err != nil {
		logger.DebugLogger.Println(respFix)
	}
	framework.ListPlayer = messageSplit[10:]
	framework.OnlineActulise++

	return nil
}

func RconCommandeWhitelistAdd(player string) ([]string, error) {
	openRcon()

	response, err := c.SendCommand("ewhitelist add " + player)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
		connect = false
		return nil, err
	}

	messageSplit := strings.Fields(response)
	return messageSplit, nil
}

func RconCommandeWhitelistRemove(player string) ([]string, error) {
	openRcon()

	response, err := c.SendCommand("ewhitelist remove " + player)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
		connect = false
		return nil, err
	}

	messageSplit := strings.Fields(response)
	return messageSplit, nil
}

func RconCommandeKick(player, raison string) ([]string, error) {
	openRcon()

	response, err := c.SendCommand("kick " + player + " " + raison)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
		connect = false
		return nil, err
	}

	messageSplit := strings.Fields(response)
	return messageSplit, nil
}

func openRcon() {
	for connect == false {
		liaison, err := NewClient(viper.GetString("Minecraft.IP"), viper.GetInt("Minecraft.Port"), viper.GetString("Minecraft.Mdp"))
		if err != nil {
			connect = false
			time.Sleep(10 * time.Second)
			continue
		} else {
			connect = true
			c = liaison
			break
		}
	}
}

func after(value string, a string) (string, error) {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return "", errors.New("after : pos = -1")
	}
	adjustedPos := pos //+ len(a)
	if adjustedPos >= len(value) {
		return "", errors.New("after : adjustedPos >= len(value)")
	}
	return value[adjustedPos:len(value)], nil
}
