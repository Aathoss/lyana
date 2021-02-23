package modules

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func Stats(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	if viper.GetBool("Dev.PrintMessage") == true {
		log.Println(m.Content)
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
	_, err = insert.Exec(t1.Unix(), m.Author.ID, "msgcount", "")
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
