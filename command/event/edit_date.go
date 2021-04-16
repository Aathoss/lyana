package event

import (
	"strconv"
	"strings"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
)

func EditDate(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if framework.EventConstruction != true {
		return
	}

	if len(ctx.Args[0]) == 0 {
		return
	}

	format := "2006/01/02 15:04"
	date, err := time.Parse(format, strings.Join(ctx.Args[0:2], " "))
	if err != nil {
		return
	}

	err = mysql.EditDate(strconv.Itoa(int(date.Unix())), framework.EventContructionIndex)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	message, _ := ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "**Date défini :** \n"+date.Format("le 02/01/2006 à 15:04:05")+"\nDans : "+framework.Calculetime(date.Unix(), 1))
	time.Sleep(time.Second * 10)
	ctx.Discord.ChannelMessageDelete(message.ChannelID, message.ID)
}
