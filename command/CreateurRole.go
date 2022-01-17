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
}
