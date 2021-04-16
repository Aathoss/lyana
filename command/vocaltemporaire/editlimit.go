package vocaltemporaire

import (
	"strconv"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
)

func VocalTempEditLimit(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	usercount, returnchannel := mysql.Countuser(ctx.User.ID)
	channelname := ctx.User.Username
	channellimite := 10

	if usercount == 0 {
		mysql.InsertCreation(ctx.User.ID, channelname, channellimite)
	} else {
		channelname, channellimite = mysql.ReturnConfigChannel(ctx.User.ID)
	}

	if num, err := strconv.Atoi(ctx.Args[0]); err == nil {
		channellimite = num

		if channellimite < 1 || channellimite > 100 {
			ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, ":unamused: ... Merci de définir un nombre entre 1 & 99")
			return
		}

		//Mise à jour du channel vocal (si il existe)
		if returnchannel != "" {
			ctx.Discord.ChannelEditComplex(returnchannel, &discordgo.ChannelEdit{
				Name:      channelname,
				UserLimit: channellimite,
				Position:  7,
			})
		}

		err := mysql.UpdateChannelUserLimit(ctx.User.ID, channellimite)
		if err != nil {
			logger.ErrorLogger.Println(err)
			ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, ":unamused: ... Je viens de rencontrer une erreur, si le problème persiste contacter le staff")
		} else {
			ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "<:yes_tick:742854426867531867> Modification du nombre de membre dans le channel : "+strconv.Itoa(channellimite))
		}
	} else {
		ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, ":unamused: ... Merci de définir un nombre entre 1 & 99")
	}
}
