package framework

import (
	"errors"
	"strconv"
	"strings"
	"time"

	mcrcon "github.com/Kelwing/mc-rcon"
	"github.com/spf13/viper"
	"gitlab.com/lyana/logger"
)

var (
	conn  *mcrcon.MCConn
	notif bool
	pause = time.Second * 10
)

func connect(nom string, num int) {
	for {
		address := viper.GetString("Minecraft."+nom+".IP") + ":" + viper.GetString("Minecraft."+nom+".Port")

		connnect := new(mcrcon.MCConn)
		err := connnect.Open(address, viper.GetString("Minecraft."+nom+".Mdp"))
		if err != nil {
			if notif != true {
				logger.ErrorLogger.Println("MC-Host : "+nom+" | Open failed", err)
				notif = true
			}
			time.Sleep(pause)
			continue
		}

		err = connnect.Authenticate()
		if err != nil {
			logger.ErrorLogger.Println("MC-Host : "+nom+" | Auth failed", err)
			time.Sleep(pause)
			continue
		}

		conn = connnect
		OnlineServer[num] = "online"
		notif = false
		logger.InfoLogger.Println("MC-Host : " + nom + " | Connexion rcon réussi")
		break
	}
}

func StartRCON(nom string, num int) {
	OnlinePlayer[num] = 0
	ListPlayer[num] = ""
	OnlineServer[num] = "offline"

	connect(nom, num)

	for {
		err := conn.Authenticate()
		if err != nil {
			OnlineServer[num] = "offline"
			connect(nom, num)
		}

		//Count le nombre de joueurs en ligne / liste les pseudo
		resp, err := conn.SendCommand("list")
		if err != nil {
			logger.ErrorLogger.Println("MC-Host : "+nom+" | Command failed : ", err)
			continue
		}
		//fmt.Println("Session : " + nom + " | ID : " + strconv.Itoa(num) + " | Retour : " + resp)

		respFix, err := after(resp, "There")
		if err != nil {
			continue
		}

		messageSplit := strings.Fields(respFix)
		OnlinePlayer[num], err = strconv.Atoi(messageSplit[2])
		if err != nil {
			logger.DebugLogger.Println(resp)
			continue
		}

		result1 := strings.Join(messageSplit[10:], " ")
		ListPlayer[num] = result1
		time.Sleep(pause)
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
