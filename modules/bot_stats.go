package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/mysql"
)

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	count := framework.VerifGrade(m.Member.Roles)
	if count == 1 {
		s.GuildMemberRoleRemove(m.Message.GuildID, m.Author.ID, "757730769023008958")
	}
	mysql.UpdateInactifDiscord(m.Author.ID)

	if viper.GetBool("Dev.test") != true {
		framework.CountMsg++
	}
}
