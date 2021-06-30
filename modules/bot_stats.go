package modules

import (
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
)

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de l'obtention du channel,", err)
		return
	}

	if channel.Type == discordgo.ChannelTypeDM {
		logger.InfoLogger.Println("-------------------------------------------------")
		logger.InfoLogger.Println(m.Author.Username)
		logger.InfoLogger.Println(m.Message)
		logger.InfoLogger.Println("-------------------------------------------------")
		return
	}

	count := framework.VerifGrade(m.Member.Roles)
	if count == 1 {
		s.GuildMemberRoleRemove(m.Message.GuildID, m.Author.ID, "757730769023008958")
	}
	mysql.UpdateInactifDiscord(m.Author.ID)

	//Logs le nom de message envoyer par les Utilisateur
	t1 := time.Now()
	insert, err := framework.DBLyana.Prepare("INSERT INTO logs(timestamp, uuid, categorie, content) VALUES(?, ?, ?, ?)")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	defer insert.Close()
	_, err = insert.Exec(t1.Unix(), m.Author.ID, "msgcount", "")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
