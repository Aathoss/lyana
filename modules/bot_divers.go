package modules

import (
	"github.com/bwmarrin/discordgo"
)

func CmdDivers(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author
	if user.ID == s.State.User.ID || user.Bot {
		return
	}

}
