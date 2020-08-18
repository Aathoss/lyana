package command

import (
	"strconv"
	"strings"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/modules"
	"gitlab.com/lyana/rcon"
)

func OnlinePlayer(ctx framework.Context) {
	var listeplayers string
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	rcon.RconCommandeList()

	for i := 0; i < len(framework.ListPlayer); i++ {
		listeplayers = listeplayers + framework.ListPlayer[i]
	}

	listeplayers = strings.Replace(listeplayers, ",", " - ", -1)

	embedHelp := modules.NewEmbed().
		SetTitle("Joueurs en ligne " + strconv.Itoa(framework.OnlinePlayer) + " / " + strconv.Itoa(framework.MaxOnlinePlayer)).
		SetColor(0x725F7C).
		SetDescription(listeplayers).MessageEmbed

	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embedHelp)
}
