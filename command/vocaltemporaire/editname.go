package vocaltemporaire

import (
	"strings"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
)

func VocalTempEditTitre(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	usercount, returnchannel := mysql.Countuser(ctx.User.ID)
	channelname := ctx.User.Username
	channellimite := 10
	characterlimit := ""

	if usercount == 0 {
		mysql.InsertCreation(ctx.User.ID, channelname, channellimite)
	} else {
		channelname, channellimite = mysql.ReturnConfigChannel(ctx.User.ID)
	}

	if len(ctx.Args) > 100 {
		channelname = strings.Join(ctx.Args[:100], " ")
		characterlimit = "\n:warning: Vous avez dépassé la limite de 100 caractères."
	} else {
		channelname = strings.Join(ctx.Args, " ")
	}

	//Mise à jour du channel vocal (si il existe)
	if returnchannel != "" {
		ctx.Discord.ChannelEditComplex(returnchannel, &discordgo.ChannelEdit{
			Name:      channelname,
			UserLimit: channellimite,
			Position:  7,
		})
	}

	err := mysql.UpdateChannelName(ctx.User.ID, channelname)
	if err != nil {
		logger.ErrorLogger.Println(err)
		ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, ":unamused: ... Je viens de rencontrer une erreur, si le problème persiste contacter le staff")
	} else {
		ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "<:yes_tick:742854426867531867> Modification du nom du channel vocal : "+channelname+characterlimit)
	}
}
