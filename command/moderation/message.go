package moderation

import (
	"net/http"
	"strings"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/bwmarrin/discordgo"
)

var (
	fileName = "lyanamessage.png"
)

func PubliMessage(ctx framework.Context) {
	rm := "!" + ctx.Commande + " " + ctx.Args[0] + " "

	if len(ctx.Message.Attachments) != 0 {

		response, err := http.Get(ctx.Message.Attachments[0].URL)
		if err != nil {
			logger.ErrorLogger.Println(err)
			framework.LogsChannel("Erreur lors de l'envoie du message avec lyana : " + err.Error())
			ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
			return
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			logger.ErrorLogger.Println("Received non 200 response code")
			framework.LogsChannel("Erreur lors de l'envoie du message avec lyana : Received non 200 response code")
			ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
			return
		}

		_, err = ctx.Discord.ChannelMessageSendComplex(ctx.Args[0], &discordgo.MessageSend{
			Content: strings.Replace(ctx.Message.Content, rm, "", -1),
			File: &discordgo.File{
				Name:   fileName,
				Reader: response.Body,
			},
		})
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
	}

	if len(ctx.Message.Attachments) == 0 {
		ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, strings.Replace(ctx.Message.Content, rm, "", -1))
	}

	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
}
