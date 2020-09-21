package rcon

import (
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

	messageSplit := strings.Fields(response)
	framework.OnlinePlayer, err = strconv.Atoi(messageSplit[2])
	if err != nil {
		logger.DebugLogger.Println(response)
	}
	framework.MaxOnlinePlayer, err = strconv.Atoi(messageSplit[7])
	if err != nil {
		logger.DebugLogger.Println(response)
	}
	framework.ListPlayer = messageSplit[10:]

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
			logger.ErrorLogger.Println("Open failed", err)
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
