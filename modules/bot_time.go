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
	go mysql.UpdateInactifPlayer()

	if Minute30 >= 30 {
		VerifRule(framework.Session)
		VerifInactif(framework.Session)

		Minute30 = 0
	}
}

//VerifInactif Vérifie la liste des inactif discord/mc
func VerifInactif(session *discordgo.Session) {
	liste, err := mysql.VerifInactif()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	if len(liste) != 0 {
		for _, inf := range liste {
			user, err := session.GuildMember(viper.GetString("GuildID"), inf[1])
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}

			semaine, err := strconv.Atoi(inf[3])
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}

			err = session.GuildMemberRoleAdd(viper.GetString("GuildID"), user.User.ID, "757730769023008958")
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}

			framework.LogsChannel("[:zzz:] " + user.User.String() + " inactif depuis " + strconv.Itoa(semaine+1) + " semaines")
			mysql.UpdateMembresInactif(inf[1])
		}
	}
}

func VerifRule(session *discordgo.Session) {
	liste, err := mysql.VerifRuleTimestamp()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	if len(liste) >= 1 {
		for _, uuid := range liste {
			user, err := session.GuildMember(viper.GetString("GuildID"), uuid)
			if err != nil {
				logger.ErrorLogger.Println(err)
				return
			}

			err = mysql.RemoveRule(uuid)
			if err != nil {
				logger.ErrorLogger.Println(err)
				return
			}

			err = session.GuildMemberDeleteWithReason(viper.GetString("GuildID"), uuid, "[UniSpace] Vous n'avez pas accepté le règlement du discord sous 3 jours")
			if err != nil {
				logger.ErrorLogger.Println(err)
				return
			}

			framework.LogsChannel("[<:downvote:742854427177648190>] " + user.User.String() + " n'as pas validé les règles sous 3 jours")
		}
	}
}

func UpdateOnlinePlayer(session *discordgo.Session) {
	err := rcon.RconCommandeList()
	if err != nil {
		return
	}

	editchannel := &discordgo.ChannelEdit{
		Name:     "🪐  Online : " + strconv.Itoa(framework.OnlinePlayer),
		Position: 2,
	}

	_, err = session.ChannelEditComplex(viper.GetString("ChannelID.OnlinePlayer"), editchannel)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
}
