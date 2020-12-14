package command

import (
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func AddPlayer(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if len(ctx.Args) > 1 {
		mentions := ctx.Message.Mentions
		playermc := ctx.Args[1]

		if len(mentions) == 0 {
			framework.LogsChannel("[:x:] [" + viper.GetString("PrefixMsg") + ctx.Commande + "] Vous n'avez pas mentionné de personne !")
			return
		}

		countuuid, countplayer := mysql.VerifPlayerMC(mentions[0].ID, playermc)
		if countplayer == 1 {
			framework.LogsChannel("[:open_mouth:] [Utilisateur : " + mentions[0].String() + " | Pseudo : " + playermc + "] Ce pseudo existe déjà dans notre base de données.")
			return
		}

		if countuuid == 1 {
			framework.LogsChannel("[:open_mouth:] [Utilisateur : " + mentions[0].String() + " | Pseudo : " + playermc + "] Il n'est pas autorisé d'avoir un double compte...")
			return
		}

		resp, err := rcon.RconCommandeWhitelistAdd(playermc)
		if err != nil {
			logger.ErrorLogger.Println(err)
			framework.LogsChannel("[:x:] [" + viper.GetString("PrefixMsg") + ctx.Commande + " " + mentions[0].ID + " " + playermc + "] Une erreur c'est produits sur le whitelist rcon")
			return
		}

		if resp[4] == "§7Whitelisted" {
			ctx.Discord.ChannelMessageSend(viper.GetString("ChannelID.General"), "<:CraftingTable:753547645875912736> Je viens de craft votre carte d'accès au serveur, nous vous souhaitons la bienvenue parmi nous "+mentions[0].Mention()+".")
			err = mysql.AddWhitelist(mentions[0].ID, playermc)
			if err != nil {
				logger.ErrorLogger.Println(err)
			}
			return
		}
	}

	if len(ctx.Args) <= 1 {
		embed := framework.NewEmbed().
			SetTitle("Il semble y avoir une erreur !").
			SetColor(viper.GetInt("EmbedColor.Error")).
			SetDescription("Veuillez respecter ce format : " + viper.GetString("PrefixMsg") + "addplayer <tag_discord> <player_mc>").MessageEmbed

		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	}
}
