package modules

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/unispace/framework"
	"gitlab.com/unispace/logger"
)

var (
	conn *MCConn
)

func OnlinePlayerRcon() {
	if framework.Connection != true {
		conn = OpenConnexionRcon()
	}

	if framework.Connection == true {

		resp, err := conn.SendCommand("list")
		if err != nil {
			logger.InfoLogger.Println("Command failed", err)
			framework.Connection = false
		}

		if framework.Connection == true {
			messageSplit := strings.Fields(resp)
			framework.OnlinePlayer, _ = strconv.Atoi(messageSplit[2])
			framework.MaxOnlinePlayer, _ = strconv.Atoi(messageSplit[7])
			framework.ListPlayer = messageSplit[10:]
		}
	}
}

func WhitelistRcon(player string) []string {
	if framework.Connection != true {
		conn = OpenConnexionRcon()
	}

	resp, err := conn.SendCommand("whitelist add " + player)
	messageSplit := strings.Fields(resp)
	if err != nil {
		logger.InfoLogger.Println("Command failed", err)
		framework.Connection = false
	}

	return messageSplit
}

func UnWhitelistRcon(player string) []string {
	if framework.Connection != true {
		conn = OpenConnexionRcon()
	}

	resp, err := conn.SendCommand("whitelist remove " + player)
	messageSplit := strings.Fields(resp)
	if err != nil {
		logger.InfoLogger.Println("Command failed", err)
		framework.Connection = false
	}

	return messageSplit
}

func KickPlayerRcon(player, raison string) []string {
	if framework.Connection != true {
		conn = OpenConnexionRcon()
	}

	resp, err := conn.SendCommand("kick " + player + " " + raison)
	messageSplit := strings.Fields(resp)
	if err != nil {
		logger.InfoLogger.Println("Command failed", err)
		framework.Connection = false
	}

	return messageSplit
}

func VersionServerRcon() {
	if framework.Connection != true {
		conn = OpenConnexionRcon()
	}

	resp, err := conn.SendCommand("version")
	if err != nil {
		logger.InfoLogger.Println("Command failed", err)
		framework.Connection = false
	}

	if framework.Connection == true {
		messageSplit := strings.Fields(resp)
		framework.VersionMC = strings.Replace(messageSplit[6], "git-Paper-", "", -1)
		framework.BuildMC = strings.Replace(messageSplit[8], ")", "", -1)
	}
}

func OpenConnexionRcon() (conn *MCConn) {
	conn = new(MCConn)
	err := conn.Open(viper.GetString("Minecraft.IP")+":"+viper.GetString("Minecraft.Port"), viper.GetString("Minecraft.Mdp"))
	if err != nil {
		logger.ErrorLogger.Println("Open failed", err)
		framework.Connection = false
		return
	} else {
		framework.Connection = true
	}
	//defer conn.Close()

	if framework.Connection == true {
		err = conn.Authenticate()
		if err != nil {
			logger.ErrorLogger.Println("Auth failed", err)
		}
		LogDiscord("[:tools:] Connexion rcon réussi !")
		framework.Connection = true
	}
	return conn
}
