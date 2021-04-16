package modules

import (
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func ReactionAdd(s *discordgo.Session, reac *discordgo.MessageReactionAdd) {
	if reac.UserID == s.State.User.ID {
		return
	}

	//Acceptation du réglement
	if reac.Emoji.Name == "✅" && reac.ChannelID == viper.GetString("ChannelID.Reglement") && reac.MessageID == viper.GetString("MessageID.Reglement") {
		s.GuildMemberRoleRemove(viper.GetString("GuildID"), reac.UserID, "742781882852179988")
		s.GuildMemberRoleAdd(viper.GetString("GuildID"), reac.UserID, "735281835080286291")
		mysql.RemoveRule(reac.UserID)
	}
	if reac.Emoji.Name != "✅" && reac.ChannelID == viper.GetString("ChannelID.Reglement") {
		s.MessageReactionRemove(reac.ChannelID, reac.MessageID, reac.Emoji.Name, reac.UserID)
	}

	//Ajoute une réaction pour participé à l'événement
	if reac.Emoji.Name == "Yes_Night" && reac.ChannelID == viper.GetString("ChannelID.Event") {
		mysql.ReactionParticipants(0, reac.MessageID, reac.UserID)
	} else if reac.Emoji.Name != "Yes_Night" && reac.ChannelID == viper.GetString("ChannelID.Event") {
		s.MessageReactionRemove(reac.ChannelID, reac.MessageID, reac.Emoji.Name, reac.UserID)
	}
}

func ReactionRemove(s *discordgo.Session, reac *discordgo.MessageReactionRemove) {
	if reac.UserID == s.State.User.ID {
		return
	}

	//remove une réaction pour participé à l'événement
	if reac.Emoji.Name == "Yes_Night" && reac.ChannelID == viper.GetString("ChannelID.Event") {
		mysql.ReactionParticipants(1, reac.MessageID, reac.UserID)
	}

}
