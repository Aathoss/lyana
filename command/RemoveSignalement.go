package command

import (
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func RemoveSignalement(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	userID := strings.Replace(ctx.Args[0], "<@!", "", -1)
	userID = strings.Replace(userID, ">", "", -1)
	user, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), userID)
	if err != nil {
		logger.ErrorLogger.Println(err)
		framework.LogsChannel("[:x:] Une erreur c'est produits sur ` " + ctx.Commande + ctx.Args[0] + "`\n" + err.Error())
		return
	}

	msgSanction, msgNotif := mysql.RemoveSanctionLimit(user.User.ID)
	ctx.Discord.ChannelMessageDelete(viper.GetString("ChannelID.Signalement"), msgSanction)
	ctx.Discord.ChannelMessageDelete(viper.GetString("ChannelID.Signalement"), msgNotif)

}
