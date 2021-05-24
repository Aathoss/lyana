package event

import (
	"strconv"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/spf13/viper"
)

func PubliEvent(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	var index int
	
	if len(ctx.Args[0]) != 0 {
		num, _ := strconv.Atoi(ctx.Args[0])
		index = num
	} else {
		index = framework.EventContructionIndex
	}

	if framework.EventConstruction != true && len(ctx.Args[0]) == 0 {
		message, _ := ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "**Il semble ne pas y avoir d'évent en cours de création**")
		time.Sleep(time.Second * 10)
		ctx.Discord.ChannelMessageDelete(message.ChannelID, message.ID)
		return
	}

	mysql.EditStatus(1, index)
	mysql.EditChannelID(viper.GetString("ChannelID.Event"), index)
	tab, err := mysql.GetCreationEvent(index)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	messageid := ConstructionEmbedEvent(1, ctx.Discord, tab)
	ctx.Discord.MessageReactionAdd(viper.GetString("ChannelID.Event"), messageid, ":Yes_Night:742854427987148840")
	mysql.EditMessageID(messageid, index)

	if framework.EventConstruction == true {
		ctx.Discord.ChannelMessageDelete(framework.EventConstructionChannelID, framework.EventConstructionMessageID)
		ctx.Discord.ChannelMessageDelete(framework.EventConstructionChannelID, framework.EventConstructionMessageAide)

		framework.EventConstruction = false
		framework.EventConstructionChannelID = ""
		framework.EventConstructionMessageID = ""
		framework.EventConstructionMessageAide = ""
	}
}
