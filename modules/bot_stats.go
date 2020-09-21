package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/command"
	"gitlab.com/lyana/mysql"
)

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	mysql.UpdateInactifDiscord(m.Author.ID)

	if viper.GetBool("Dev.test") != true {
		command.SetCountMsg()
	}
}
