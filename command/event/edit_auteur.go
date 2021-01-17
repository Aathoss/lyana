package event

import (
	"strings"
	"time"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func EditAuteur(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if framework.EventConstruction != true {
		return
	}

	if len(ctx.Args[0]) == 0 {
		return
	}

	auteur := strings.Join(ctx.Args[0:], " ")

	err := mysql.EditAuteur(auteur, framework.EventContructionIndex)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	message, _ := ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "**Nouvelle Auteur : **"+auteur)
	time.Sleep(time.Second * 10)
	ctx.Discord.ChannelMessageDelete(message.ChannelID, message.ID)
}
