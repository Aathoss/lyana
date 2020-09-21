package modules

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

func GuildMemberLeave(s *discordgo.Session, leave *discordgo.GuildMemberRemove) {
	//Trafic des membres du discord [leave]
	//leave action

	mysql.RemoveRule(leave.User.ID)
	_, playermc, _, err := mysql.GetWhitelist(leave.User.ID)
	if err != nil {
		rcon.RconCommandeWhitelistRemove(playermc)
	}

	framework.LogsChannel("[<:downvote:742854427177648190>] " + leave.User.Username + "Viens de partir, il vient d'être retiré retiré de la whitelist pseudo : " + playermc)
}
