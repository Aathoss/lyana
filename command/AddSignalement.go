package command

import (
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func AddSignalement(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	count := mysql.SelectCount("sanction", "uid", ctx.Message.Author.ID)
	if count == 0 {

		count = mysql.SelectCount("membre", "tag_discord", ctx.Message.Author.ID)
		if count == 1 {

			if len(ctx.Args) >= 1 {

				playermc := ctx.Args[0]
				if playermc == "Aathoss" || playermc == "Paulth04" || playermc == "Bajoux" || playermc == "Laihela_" {
					embed := framework.NewEmbed().
						SetTitle("Vous n'êtes pas autorisé à signaler un membre du staff").
						SetColor(viper.GetInt("EmbedColor.Error")).MessageEmbed

					ctx.Discord.ChannelMessageSendEmbed(viper.GetString("ChannelID.Signalement"), embed)
					return
				}

				raison := ctx.Args[1:]
				var raisonString string

				for i := 0; i <= len(raison)-1; i++ {
					raisonString = raisonString + "|" + raison[i]
				}
				raisonString = strings.Replace(raisonString, "|", " ", -1)

				_, countPlayer := mysql.VerifPlayerMC(ctx.Message.Author.ID, playermc)
				if countPlayer == 1 {

					rcon.RconCommandeWhitelistRemove(playermc)
					rcon.RconCommandeKick(playermc, "Vous êtes actuellement suspecté d'avoir enfreint les règles ! Nous vous invitons à vous rendre sur le discord dans le channel #signalement-de-joueur")

					embed := framework.NewEmbed().
						SetTitle("Signialement  de joueurs !").
						SetColor(viper.GetInt("EmbedColor.Signalement")).
						AddField("Informateur", ctx.Message.Author.Username, true).
						AddField("Suspect", playermc, true).
						AddField("Raison", raisonString, false).MessageEmbed

					info, _ := ctx.Discord.ChannelMessageSendEmbed(viper.GetString("ChannelID.Signalement"), embed)
					infoNotif, _ := ctx.Discord.ChannelMessageSend(viper.GetString("ChannelID.Signalement"), "> Le staff va lancer une vérification des que possible <@&743945966813708369>")
					mysql.AddSanctionLimit(ctx.Message.Author.ID, playermc, info.ID, infoNotif.ID)
					return

				}

				embed := framework.NewEmbed().
					SetTitle("D'après ma sonde, aucun joueur whitelist n'existe avec ce pseudo !").
					SetColor(viper.GetInt("EmbedColor.Error")).MessageEmbed
				ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
				return

			}

			embedHelp := framework.NewEmbed().
				SetTitle("Il semble y avoir une erreur !").
				SetColor(viper.GetInt("EmbedColor.Error")).
				SetDescription("Veuillez respecter ce format : " + viper.GetString("PrefixMsg") + "signal <player> <raison>").MessageEmbed

			ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embedHelp)
			return
		}

		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, framework.EmbedPermissionFalse)
		return
	}
	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, framework.EmbedLimite)
	return
}
