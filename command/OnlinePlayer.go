package command

import (
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/rcon"
)

func OnlinePlayer(ctx framework.Context) {
	if ctx.TextChannel.ParentID != viper.GetString("Categorie.Information") {
		var listeplayers string
		ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

		rcon.RconCommandeList()

		for i := 0; i < len(framework.ListPlayer); i++ {
			listeplayers = listeplayers + framework.ListPlayer[i]
		}

		listeplayers = strings.Replace(listeplayers, ",", " - ", -1)

		tNow := time.Now()
		embedHelp := framework.NewEmbed().
			SetTitle("Joueurs en ligne " + strconv.Itoa(framework.OnlinePlayer) + " / " + strconv.Itoa(framework.MaxOnlinePlayer)).
			SetColor(0x725F7C).
			SetFooter(ctx.Message.Author.Username + " | Date : " + tNow.Format("2/1 15:04:05")).
			SetDescription(listeplayers).MessageEmbed

		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embedHelp)
	}
}
