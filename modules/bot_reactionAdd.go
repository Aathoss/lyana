package modules

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/mysql"
)

func ReactionAdd(s *discordgo.Session, reac *discordgo.MessageReactionAdd) {
	if reac.UserID == s.State.User.ID {
		log.Println("Bot ajoute un emoji")
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
}
