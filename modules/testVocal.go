package modules

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func TestVocal(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	fmt.Println(m.ChannelID)
	fmt.Println(m.UserID)

	if m.ChannelID == "" { //User disconnected from a voice channel
		guild, _ := s.State.Guild(m.GuildID)

		for _, key := range guild.VoiceStates {
			if key.UserID == m.UserID {
				//This code is never reached as the user was already removed from the VoiceStates array
				println(key.UserID, " left channel ", key.ChannelID)
			}
		}
	}
}
