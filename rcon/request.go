package rcon

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

var (
	c   *Client
	err bool
)

func RconCommandeList() {
	c = openRcon()

	response, err := c.SendCommand("list")
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	if response == "" {
		logger.ErrorLogger.Println("Commande list vide")
		return
	}

	messageSplit := strings.Fields(response)
	framework.OnlinePlayer, _ = strconv.Atoi(messageSplit[2])
	framework.MaxOnlinePlayer, _ = strconv.Atoi(messageSplit[7])
	framework.ListPlayer = messageSplit[10:]
}

func RconCommandeWhitelistAdd(player string) []string {
	c = openRcon()

	response, err := c.SendCommand("whitelist add " + player)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	if response == "" {
		logger.ErrorLogger.Println("Commande list vide")
		return messageSplit
	}

	return messageSplit
}

func RconCommandeWhitelistRemove(player string) []string {
	c = openRcon()

	response, err := c.SendCommand("whitelist remove " + player)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	if response == "" {
		logger.ErrorLogger.Println("Commande list vide")
		return messageSplit
	}

	return messageSplit
}

/* func RconCommandeTest() []string {
	c = openRcon()

	response, err := c.SendCommand("co lookup Aathoss")
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	if response != "" {
		logger.ErrorLogger.Println("Commande list vide")
		return messageSplit
	}
	return messageSplit
} */

func RconCommandeKick(player, raison string) []string {
	c = openRcon()

	response, err := c.SendCommand("kick " + player + " " + raison)
	if err != nil {
		logger.ErrorLogger.Println("Send Command", err)
	}

	messageSplit := strings.Fields(response)
	if response == "" {
		logger.ErrorLogger.Println("Commande list vide")
		return messageSplit
	}

	return messageSplit
}

func openRcon() (c *Client) {
	c, err := NewClient(viper.GetString("Minecraft.IP"), viper.GetInt("Minecraft.Port"), viper.GetString("Minecraft.Mdp"))
	if err != nil {
		logger.ErrorLogger.Println("Open failed", err)
	}
	return c
}
