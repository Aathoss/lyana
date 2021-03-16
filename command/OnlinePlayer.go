package command

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
)

func OnlinePlayer(ctx framework.Context) {
	if ctx.TextChannel.ParentID != viper.GetString("Categorie.Information") {
		var onlineliste string
		ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

		fmt.Println(framework.PositionServeur)
		for num, val := range framework.PositionServeur {
			listeplayers := strings.Replace(framework.ListPlayer[num], ",", " <:tirer:804348409024741408> ", -1)

			if framework.OnlineServer[num] == "online" {
				onlineliste = onlineliste + ":green_circle: **" + strings.Title(val) + "**\n**" + strconv.Itoa(framework.OnlinePlayer[num]) + "** Joueurs : " + listeplayers
			} else {
				onlineliste = onlineliste + ":red_circle: **" + strings.Title(val) + "**"
			}

			onlineliste = onlineliste + "\n\n"
		}

		tNow := time.Now()
		embedHelp := framework.NewEmbed().
			//SetTitle("Joueurs en ligne " + strconv.Itoa(framework.OnlinePlayer) + " / " + strconv.Itoa(framework.MaxOnlinePlayer)).
			SetColor(0x725F7C).
			SetFooter(ctx.Message.Author.Username + " | Date : " + tNow.Format("2/1 15:04:05")).
			SetDescription(onlineliste).MessageEmbed

		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embedHelp)
	}
}
