package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/command"
)

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	if viper.GetBool("Dev.test") != true {
		command.SetCountMsg()
	}
}
