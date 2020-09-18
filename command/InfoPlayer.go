package command

import (
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func InfoPlayer(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	user, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), ctx.Message.Author.ID)
	if err != nil {
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
	}

	t1, _ := user.JoinedAt.Parse()
	_, pseudoMC, t2 := mysql.GetWhitelist(ctx.Message.Author.ID)

	embed := framework.NewEmbed().
		SetTitle("Votre Carte d'identité : "+user.User.Username).
		SetColor(viper.GetInt("EmbedColor.Informations")).
		AddField("Vous êtes arrivé il y à", framework.Calculetime(t1.Unix(), 0), false).
		AddField("Vous êtes whitelist depuis", framework.Calculetime(int64(t2), 0), false).
		AddField("Votre pseudo minecraft", pseudoMC, false).MessageEmbed

	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}
