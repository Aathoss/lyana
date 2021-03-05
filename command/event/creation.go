package event

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func ConstructionEvent(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	t1 := time.Now()

	count, err := mysql.CountIndexEvent()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	framework.EventContructionIndex = count

	titre := strconv.Itoa(framework.EventContructionIndex) + " :pushpin: **Construction de l'évent...**"
	auteur := "nil"
	color := 0xFFFF25
	status := "dev"
	emplacement := "nil"
	eventdate := "nil"
	description := "nil"
	recompense := "nil"
	participant := ""
	updatefooter := t1.Format("2/1 15:04:05")

	message, err := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, &discordgo.MessageEmbed{
		Color:       color,
		Title:       titre,
		Description: "",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Dérnier mise à jour : " + updatefooter,
		},
	})
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	aide, _ := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, &discordgo.MessageEmbed{
		Color: color,
		Title: "Commande de création :",
		Description: "Embed d'évent actualisation toutes les 5 secondes" +
			"\n\n**!event titre** <titre>" +
			"\n**!event gps** <emplacement ou coordonnée de l'évenement>" +
			"\n**!event desc** <description>" +
			"\n**!event date** <années/mois/jours> <heure:minutes>" +
			"\n**!event recompense** <liste des récompense>" +
			"\n**!event auteur** <tags de la personne ou texte>",
	})
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	mysql.CreateEvent(status, titre, auteur, message.ID, ctx.Message.ChannelID, emplacement, eventdate, description, recompense, participant)

	framework.EventConstruction = true
	framework.EventConstructionChannelID = ctx.Message.ChannelID
	framework.EventConstructionMessageID = message.ID
	framework.EventConstructionMessageAide = aide.ID
}
