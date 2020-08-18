package command

import (
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/modules"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func Signalement(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	count := mysql.SelectCount("membre", "tag_discord", ctx.Message.Author.ID)
	if count == 1 {
		if len(ctx.MessageSplit) >= 2 {
			player_mc := ctx.MessageSplit[1]
			raison := ctx.MessageSplit[2:]
			var raisonString string

			for i := 0; i <= len(raison)-1; i++ {
				raisonString = raisonString + "|" + raison[i]
			}
			raisonString = strings.Replace(raisonString, "|", " ", -1)

			countPlayer := mysql.VerifPlayerMC(player_mc)
			if countPlayer == 1 {
				rcon.RconCommandeWhitelistRemove(player_mc)
				rcon.RconCommandeKick(player_mc, "Vous êtes actuellement suspecté d'avoir enfreint les règles ! Nous vous invitons à vous rendre sur le discord dans le channel #signalement-de-joueur")

				embed := modules.NewEmbed().
					SetTitle("Signialement  de joueurs !").
					SetColor(viper.GetInt("EmbedColor.Signalement")).
					AddField("Informateur", ctx.Message.Author.Username, true).
					AddField("Suspect", player_mc, true).
					AddField("Raison", raisonString, false).MessageEmbed

				ctx.Discord.ChannelMessageSendEmbed(viper.GetString("ChannelID.Signalement"), embed)
				ctx.Discord.ChannelMessageSend(viper.GetString("ChannelID.Signalement"), "> Le staff va lancer une vérification des que possible <@&743945966813708369>")

			}

			if countPlayer == 0 {
				embed := modules.NewEmbed().
					SetTitle("D'après ma sonde, aucun joueur whitelist n'existe avec ce pseudo !").
					SetColor(viper.GetInt("EmbedColor.Error")).MessageEmbed
				ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
			}

		}

		if len(ctx.MessageSplit) < 2 {
			embedHelp := modules.NewEmbed().
				SetTitle("Il semble y avoir une erreur !").
				SetColor(viper.GetInt("EmbedColor.Error")).
				SetDescription("Veuillez respecter ce format : " + viper.GetString("PrefixMsg") + "signal <player> <raison>").MessageEmbed

			ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embedHelp)
		}
	}

	if count == 0 {
		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, modules.EmbedPermissionFalse)
	}
}
