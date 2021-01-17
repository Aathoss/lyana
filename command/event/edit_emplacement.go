package event

import (
	"strings"
	"time"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func EditEmplacement(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if framework.EventConstruction != true {
		return
	}

	if len(ctx.Args[0]) == 0 {
		return
	}

	emplacement := strings.Join(ctx.Args[0:], " ")

	err := mysql.EditEmplacement(emplacement, framework.EventContructionIndex)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	message, _ := ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "**Emplacement de l'évent : **"+emplacement)
	time.Sleep(time.Second * 10)
	ctx.Discord.ChannelMessageDelete(message.ChannelID, message.ID)
}
