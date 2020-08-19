package rcon

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

var (
	c       *Client
	connect bool
)

func RconCommandeList() {
	openRcon()

	response, err := c.SendCommand("list")
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	framework.OnlinePlayer, _ = strconv.Atoi(messageSplit[2])
	framework.MaxOnlinePlayer, _ = strconv.Atoi(messageSplit[7])
	framework.ListPlayer = messageSplit[10:]
}

func RconCommandeWhitelistAdd(player string) []string {
	openRcon()

	response, err := c.SendCommand("ewhitelist add " + player)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	return messageSplit
}

func RconCommandeWhitelistRemove(player string) []string {
	openRcon()

	response, err := c.SendCommand("ewhitelist remove " + player)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	return messageSplit
}

func RconCommandeKick(player, raison string) []string {
	openRcon()

	response, err := c.SendCommand("kick " + player + " " + raison)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	return messageSplit
}

func openRcon() {
	for connect == false {
		liaison, err := NewClient(viper.GetString("Minecraft.IP"), viper.GetInt("Minecraft.Port"), viper.GetString("Minecraft.Mdp"))
		if err != nil {
			logger.ErrorLogger.Println("Open failed", err)
			connect = false
			continue
		} else {
			connect = true
			c = liaison
			break
		}
	}
}
