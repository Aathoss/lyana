package modules

import (
	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
)

func GuildMemberLeave(s *discordgo.Session, leave *discordgo.GuildMemberRemove) {
	//Trafic des membres du discord [leave]
	//leave action
	var msg string

	mysql.RemoveRule(leave.User.ID)

	_, playermc, _, err := mysql.GetWhitelist(leave.User.ID)
	if err != nil {
		msg = " Viens de partir, d'après mes informations il n'était pas whitelist."
		framework.LogsChannel("[<:downvote:742854427177648190>] " + leave.User.Username + msg)

		logger.ErrorLogger.Println(err)
		return
	}
	msg = " Viens de partir, il vient d'être retiré retiré de la whitelist pseudo : " + playermc

	/* _, err = rcon.RconCommandeWhitelistRemove(playermc)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	} */
	mysql.DeleteUserWhitelist(leave.User.ID)
	framework.LogsChannel("[<:downvote:742854427177648190>] " + leave.User.Username + msg)
}
