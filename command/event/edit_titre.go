package event

import (
	"strconv"
	"strings"
	"time"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func EditTitre(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if framework.EventConstruction != true {
		return
	}

	if len(ctx.Args[0]) == 0 {
		return
	}

	titre := "n°" + strconv.Itoa(framework.EventContructionIndex) + " :pushpin: " + strings.Join(ctx.Args[0:], " ")

	err := mysql.EditTitre(titre, framework.EventContructionIndex)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	message, _ := ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "**Nouveau titre : **"+titre)
	time.Sleep(time.Second * 10)
	ctx.Discord.ChannelMessageDelete(message.ChannelID, message.ID)
}
