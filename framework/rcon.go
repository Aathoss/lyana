package framework

import (
	"errors"
	"strconv"
	"strings"
	"time"

	mcrcon "github.com/Aathoss/lyana/library/package/mc_rcon"
	"github.com/Aathoss/lyana/logger"
	"github.com/spf13/viper"
)

var (
	numInstance int
	notif bool
	pause = time.Second * 10
)

func Connect(num int) {
	for {
		address := viper.GetString("Minecraft."+strconv.Itoa(num)+".IP") + ":" + viper.GetString("Minecraft."+strconv.Itoa(num)+".Port")

		connnect := new(mcrcon.MCConn)
		err := connnect.Open(address, viper.GetString("Minecraft."+strconv.Itoa(num)+".Mdp"))
		if err != nil {
			if notif != true {
				logger.ErrorLogger.Println("MC-Host : "+viper.GetString("Minecraft."+strconv.Itoa(num)+".Name")+" | Open failed", err)
				notif = true
			}
			time.Sleep(pause)
			continue
		}

		err = connnect.Authenticate()
		if err != nil {
			logger.ErrorLogger.Println("MC-Host : "+viper.GetString("Minecraft."+strconv.Itoa(num)+".Name")+" | Auth failed", err)
			time.Sleep(pause)
			continue
		}

		ConnectMC[num] = connnect
		OnlineServer[num] = "online"
		notif = false
		logger.InfoLogger.Println("MC-Host : " + viper.GetString("Minecraft."+strconv.Itoa(num)+".Name") + " | Connexion rcon réussi")
		break
	}
}

func StartRCON(num int) {
	OnlinePlayer[num] = 0
	ListPlayer[num] = ""
	OnlineServer[num] = "offline"
	numInstance = num

	Connect(num)

	for {

		err := ConnectMC[num].Authenticate()
		if err != nil {
			OnlineServer[num] = "offline"
			Connect(num)
		}

		//Count le nombre de joueurs en ligne / liste les pseudo
		resp, err := ConnectMC[num].SendCommand("list")
		if err != nil {
			logger.ErrorLogger.Println("MC-Host : "+viper.GetString("Minecraft."+strconv.Itoa(num)+".Name")+" | Command failed : ", err)
			continue
		}

		if viper.GetString("Minecraft."+strconv.Itoa(num)+".Version") == "1.12" { //There are 1/100 players online:
			respFix, err := after(resp, "There")
			if err != nil {
				logger.ErrorLogger.Println(err)
				time.Sleep(pause)
				continue
			}

			messageSplit := strings.Fields(respFix)
			infoPlayer := strings.Split(messageSplit[2], "/")
			OnlinePlayer[num], err = strconv.Atoi(infoPlayer[0])
			if err != nil {
				logger.ErrorLogger.Println(err)
				time.Sleep(pause)
				continue
			}

			result1 := strings.Join(messageSplit[4:], " ")
			ListPlayer[num] = strings.Replace(result1, "online:", " ", -1)

		}

		if viper.GetString("Minecraft."+strconv.Itoa(num)+".Version") == "1.17" && viper.GetString("Minecraft."+strconv.Itoa(num)+".Version") == "1.18" { //There are 0 of a max of 100 players online:
			respFix, err := after(resp, "There")
			if err != nil {
				logger.ErrorLogger.Println(err)
				time.Sleep(pause)
				continue
			}

			messageSplit := strings.Fields(respFix)
			OnlinePlayer[num], err = strconv.Atoi(messageSplit[2])
			if err != nil {
				logger.ErrorLogger.Println(err)
				time.Sleep(pause)
				continue
			}

			result1 := strings.Join(messageSplit[10:], " ")
			ListPlayer[num] = result1

		}

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
