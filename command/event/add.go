package event

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
)

func CreationEvent(ctx framework.Context) {
	t1 := time.Now()

	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	ctx.Discord.ChannelMessageSendEmbed(viper.GetString("ChannelID.Event"), &discordgo.MessageEmbed{
		Color: 0xFFFF25,
		Title: ":pushpin: **Construction de l'embed d'event titre**",
		Description: `
		**Emplacement :satellite_orbital: :** serveur vanilla
		**Début à :clock: :** 16h le 25 Janvier

		**Description :clipboard: :**
		L'évent se déroule en 2 parti :
		La premier, sera une capture de drapeau .......
		La deuxième ..... on sans tape :')
		J'ai besoin de contenu donc chute
		J'avais pas d'inspi :confused:
		`,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Dérnier mise à jour : " + t1.Format("2/1 15:04:05"),
		},
	})
}
