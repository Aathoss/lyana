package modules

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
	"gitlab.com/lyana/rcon"
)

var (
	Minute30 int
)

func ExecuteTime() {
	Minute30++

	UpdateOnlinePlayer(framework.Session)

	if Minute30 >= 30 {
		VerifRule()
		Minute30 = 0
	}
}

func VerifRule() {
	s := &discordgo.Session{}

	liste, err := mysql.VerifRuleTimestamp()
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	if len(liste) >= 1 {
		for _, uid := range liste {
			mysql.RemoveRule(uid)
			s.GuildMemberDeleteWithReason(viper.GetString("GuildID"), uid, "[UniSpace] Vous n'avez pas accepté le règlement du discord sous 3 jours")
			if err != nil {
				logger.ErrorLogger.Println(err)
			}
		}
	}
}

func UpdateOnlinePlayer(session *discordgo.Session) {
	rcon.RconCommandeList()

	test := &discordgo.ChannelEdit{
		Name:     "🪐  Online : " + strconv.Itoa(framework.OnlinePlayer),
		Position: 2,
	}

	_, err := session.ChannelEditComplex(viper.GetString("ChannelID.OnlinePlayer"), test)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
