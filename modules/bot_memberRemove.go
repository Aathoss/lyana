package modules

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func GuildMemberLeave(s *discordgo.Session, leave *discordgo.GuildMemberRemove) {
	//Trafic des membres du discord [leave]
	//leave action
	var msg string

	mysql.RemoveRule(leave.User.ID)

	_, playermc, _, err := mysql.GetWhitelist(leave.User.ID)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	msg = " Viens de partir, il vient d'être retiré retiré de la whitelist pseudo : " + playermc

	_, err = rcon.RconCommandeWhitelistRemove(playermc)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	framework.LogsChannel("[<:downvote:742854427177648190>] " + leave.User.Username + msg)
}
