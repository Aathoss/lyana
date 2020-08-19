package command

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/modules"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func AddPlayer(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if ctx.Staff >= 1 {
		if len(ctx.MessageSplit) >= 2 {
			tag_discord := ctx.MessageSplit[1]
			player_mc := ctx.MessageSplit[2]

			userID := strings.Replace(ctx.MessageSplit[1], "<@!", "", -1)
			userID = strings.Replace(userID, ">", "", -1)
			user, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), userID)
			if err != nil {
				logger.ErrorLogger.Println(err)
				modules.LogDiscord("[:x:] Une erreur c'est produits sur `!addplayer <" + tag_discord + "> <" + player_mc + ">`\n" + err.Error())
				return
			}

			resp := rcon.RconCommandeWhitelistAdd(player_mc)
			fmt.Println(resp)
			if resp[4] == "§7Whitelisted" {
				ctx.Discord.ChannelMessageSend(viper.GetString("ChannelID.General"), "> Votre candidature à était accepté, je viens de procéder à la whitelist de votre pseudo mc. \n> Nous vous souhaitons un agréable séjour "+tag_discord+".")
				mysql.AddWhitelist(user.User.ID, player_mc)
			}
		}

		if len(ctx.MessageSplit) < 2 {
			embed := modules.NewEmbed().
				SetTitle("Il semble y avoir une erreur !").
				SetColor(viper.GetInt("EmbedColor.Error")).
				SetDescription("Veuillez respecter ce format : " + viper.GetString("PrefixMsg") + "addplayer <tag_discord> <player_mc>").MessageEmbed

			ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
		}
	}
	if ctx.Staff < 1 {
		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, modules.EmbedPermissionFalse)
	}
}
