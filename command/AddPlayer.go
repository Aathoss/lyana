package command

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func AddPlayer(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if len(ctx.Args) >= 1 {
		tagdiscord := ctx.Args[0]
		playermc := ctx.Args[1]

		userID := strings.Replace(ctx.Args[0], "<@!", "", -1)
		userID = strings.Replace(userID, ">", "", -1)
		user, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), userID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			framework.LogsChannel("[:x:] Une erreur c'est produits sur `" + ctx.Commande + " " + tagdiscord + " " + playermc + "`\n" + err.Error() + "`\n" + "uuid discord : " + user.User.ID)
			return
		}

		resp, err := rcon.RconCommandeWhitelistAdd(playermc)
		if err != nil {
			logger.ErrorLogger.Println(err)
			framework.LogsChannel("[:x:] Une erreur c'est produits sur `" + ctx.Commande + " " + tagdiscord + " " + playermc + "`\n" + err.Error() + "`\n" + "uuid discord : " + user.User.ID)
			return
		}
		fmt.Println(resp)
		if resp[4] == "§7Whitelisted" {
			ctx.Discord.ChannelMessageSend(viper.GetString("ChannelID.General"), "<:CraftingTable:753547645875912736> Je viens de craft votre carte d'accès au serveur, nous vous souhaitons la bienvenue parmi nous "+tagdiscord+".")
			mysql.AddWhitelist(user.User.ID, playermc)
			return
		}
	}

	if len(ctx.Args) < 1 {
		embed := framework.NewEmbed().
			SetTitle("Il semble y avoir une erreur !").
			SetColor(viper.GetInt("EmbedColor.Error")).
			SetDescription("Veuillez respecter ce format : " + viper.GetString("PrefixMsg") + "addplayer <tag_discord> <player_mc>").MessageEmbed

		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	}
}
