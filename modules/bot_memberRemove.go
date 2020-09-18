package modules

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/mysql"
)

func GuildMemberLeave(s *discordgo.Session, leave *discordgo.GuildMemberRemove) {
	//Trafic des membres du discord [leave]
	//leave action

	framework.LogsChannel("[<:downvote:742854427177648190>] " + leave.User.Username)
	mysql.RemoveRule(leave.User.ID)
}
