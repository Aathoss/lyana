package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func InfoPlayer(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	mentions := ctx.Message.Mentions
	var user *discordgo.Member

	if len(mentions) == 1 {
		u, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), mentions[0].ID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		user = u
	}
	if len(mentions) == 0 {
		u, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), ctx.Message.Author.ID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		user = u
	}

	t1, _ := user.JoinedAt.Parse()
	_, pseudoMC, t2, _ := mysql.GetWhitelist(user.User.ID)

	embed := framework.NewEmbed().
		SetTitle("Votre Carte d'identité : "+user.User.Username+" <:CraftingTable:753547645875912736> Pseudo : "+pseudoMC).
		SetColor(viper.GetInt("EmbedColor.Informations")).
		AddField("Vous êtes arrivé il y à", framework.Calculetime(t1.Unix(), 0), true).
		AddField("Vous êtes whitelist depuis", framework.Calculetime(int64(t2), 0), true).MessageEmbed

	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}
