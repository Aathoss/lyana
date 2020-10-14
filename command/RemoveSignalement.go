package command

import (
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func RemoveSignalement(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	mentions := ctx.Message.Mentions
	user, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), mentions[0].ID)
	if err != nil {
		logger.ErrorLogger.Println(err)
		framework.LogsChannel("[:x:] Une erreur c'est produits sur ` " + ctx.Commande + ctx.Args[0] + "`\n" + err.Error())
		return
	}

	pseudomc, msgSanction, msgNotif := mysql.RemoveSanctionLimit(user.User.ID)
	ctx.Discord.ChannelMessageDelete(viper.GetString("ChannelID.Signalement"), msgSanction)
	ctx.Discord.ChannelMessageDelete(viper.GetString("ChannelID.Signalement"), msgNotif)
	rcon.RconCommandeWhitelistAdd(pseudomc)

}
