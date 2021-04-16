package command

import (
	"strconv"
	"strings"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/spf13/viper"
)

func OnlinePlayer(ctx framework.Context) {
	if ctx.TextChannel.ParentID != viper.GetString("Categorie.Information") {
		var (
			onlineliste string
			domaine     string
		)

		ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

		countMAP := len(viper.GetStringMapString("Minecraft"))
		for num := 0; num < countMAP; num++ {
			listeplayers := strings.Replace(framework.ListPlayer[num], ",", " <:tirer:804348409024741408> ", -1)

			if len(viper.GetString("Minecraft."+strconv.Itoa(num)+".Domaine")) != 0 {
				domaine = "\n:white_small_square: **IP :** " + viper.GetString("Minecraft."+strconv.Itoa(num)+".Domaine")
			}

			if framework.OnlineServer[num] == "online" {
				onlineliste = onlineliste + ":green_circle: **" + viper.GetString("Minecraft."+strconv.Itoa(num)+".Name") + domaine + "**\n:white_small_square: **" + strconv.Itoa(framework.OnlinePlayer[num]) + "** Joueurs : " + listeplayers
			} else {
				onlineliste = onlineliste + ":red_circle: **" + viper.GetString("Minecraft."+strconv.Itoa(num)+".Name") + "**"
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
