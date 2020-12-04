package modules

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func TestVocal(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	fmt.Println(m.ChannelID)
	fmt.Println(m.UserID)

	//s.Voic

	if m.ChannelID == "" {
		guild, _ := s.State.Guild(m.GuildID)

		for _, key := range guild.VoiceStates {
			if key.UserID == m.UserID {
				//println(key.UserID, " left channel ", key.ChannelID)
			}
		}
	}
}
