package modules

import (
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author

	if viper.GetBool("Dev.test") == true {
		return
	}

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		logger.InfoLogger.Println("-------------------------------------------------")
		logger.InfoLogger.Println(m.Author.Username)
		logger.InfoLogger.Println(m.Message.Content)
		logger.InfoLogger.Println("-------------------------------------------------")
		return
	}

	if channel.Type == discordgo.ChannelTypeDM {
		return
	}

	count := framework.VerifGrade(m.Member.Roles)
	if count == 1 {
		s.GuildMemberRoleRemove(m.Message.GuildID, m.Author.ID, "757730769023008958")
	}
	mysql.UpdateInactifDiscord(m.Author.ID)

	//Logs le nombre de message envoyer par les Utilisateur
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
