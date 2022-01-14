package command

import (
	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/spf13/viper"
)

func CreateurRole(ctx framework.Context) {
	var roleStatus bool
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	user, err := ctx.Discord.GuildMember(ctx.Message.GuildID, ctx.Message.Author.ID)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	
	for _, roleID := range user.Roles {
		if roleID == "931472614302556170" {
			roleStatus = true
		}
	}

	if roleStatus {
		ctx.Discord.GuildMemberRoleRemove(viper.GetString("GuildID"), ctx.Message.Author.ID, "931472614302556170")
	} else {
		ctx.Discord.GuildMemberRoleAdd(viper.GetString("GuildID"), ctx.Message.Author.ID, "931472614302556170")
	}
	
	/* err = ctx.Discord.GuildMemberRoleRemove(viper.GetString("GuildID"), mentions[0].ID, "735281835080286291")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	err = ctx.Discord.GuildMemberRoleAdd(viper.GetString("GuildID"), mentions[0].ID, "820404799119818793")
	if err != nil {
		logger.ErrorLogger.Println(err)
	} */

	/* if len(ctx.Args) <= 1 {
		embed := framework.NewEmbed().
			SetTitle("Il semble y avoir une erreur !").
			SetColor(viper.GetInt("EmbedColor.Error")).
			SetDescription("Veuillez respecter ce format : " + viper.GetString("PrefixMsg") + "addplayer <tag_discord> <player_mc>").MessageEmbed

		message, _ := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
		time.Sleep(time.Second * 10)
		ctx.Discord.ChannelMessageDelete(message.ChannelID, message.ID)
	} */
}
