package modules

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

var (
	Minute30 int
)

func ExecuteTime() {
	Minute30++

	mysql.UpdateInactifPlayer()

	if Minute30 >= 30 {
		VerifRule(framework.Session)
		VerifInactif(framework.Session)

		Minute30 = 0
	}
}

//VerifCandid [GoRoutine]
func VerifCandid(secondeboucle time.Duration) {
	logger.InfoLogger.Println("----- [GoRoutine] Démarrage de la boucle VerifCandid")

	session := framework.Session

	type jsonsheet struct {
		SpreadsheetID string `json:"spreadsheetId"`
		ValueRanges   []struct {
			Range          string     `json:"range"`
			MajorDimension string     `json:"majorDimension"`
			Values         [][]string `json:"values"`
		} `json:"valueRanges"`
	}

	clesAPI, iddoc := "AIzaSyCRv40zgbaSwTXN270cdWTAWfEwbNvSIFI", "1FPseonUXhNOxciTwtNx00HQIF5V0-RvJRek9Lu53CZc"
	url1 := "https://sheets.googleapis.com/v4/spreadsheets/" + iddoc + "/values:batchGet?majorDimension=COLUMNS&ranges=A2%3AA1000&access_token=" + clesAPI + "&key=" + clesAPI
	url2 := "https://sheets.googleapis.com/v4/spreadsheets/" + iddoc + "/values:batchGet?majorDimension=COLUMNS&ranges=C2%3AC1000&access_token=" + clesAPI + "&key=" + clesAPI

	veriftemp := 0
	enattente := 0

	for {
		body1, err := framework.RequestAPI("GET", url1)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		var sheet jsonsheet
		err = json.Unmarshal(body1, &sheet)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		if len(sheet.ValueRanges) == 0 {
			continue
		}
		count1 := len(sheet.ValueRanges[0].Values[0])

		//----------------------

		body2, err := framework.RequestAPI("GET", url2)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		err = json.Unmarshal(body2, &sheet)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		count2 := len(sheet.ValueRanges[0].Values[0])
		enattente = count2 - count1

		if enattente > 0 && veriftemp != enattente {
			tNow := time.Now()
			embedNotif := framework.NewEmbed().
				SetTitle(":open_mouth: " + strconv.Itoa(enattente) + " Candidature en attente.").
				SetColor(0x43C2EB).
				SetURL("https://unispace.page.link/mVFa").
				SetDescription("Date : " + tNow.Format("2/1 15:04:05")).MessageEmbed

			_, err = session.ChannelMessageSendEmbed("735273466051297433", embedNotif)
			if err != nil {
				logger.ErrorLogger.Println(err)
			}
			veriftemp = enattente
		}
		if veriftemp > enattente {
			veriftemp = 0
		}
		time.Sleep(time.Second * secondeboucle)
	}
	logger.InfoLogger.Println("----- [GoRoutine] Arrêt de la boucle VerifCandid")
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
				logger.ErrorLogger.Println("[uuid: " + string(inf[1]) + " ]" + err.Error())
				continue
			}

			semaine, err := strconv.Atoi(inf[3])
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}
			semaine = semaine + 1

			if semaine == 1 {
				err = session.GuildMemberRoleAdd(viper.GetString("GuildID"), user.User.ID, "757730769023008958")
				if err != nil {
					logger.ErrorLogger.Println(err)
					continue
				}
				framework.LogsChannel("[:zzz:] " + user.User.String() + " inactif depuis " + strconv.Itoa(semaine) + " semaines")
			} else {
				if semaine%4 == 0 {
					framework.LogsChannel("[:zzz:] " + user.User.String() + " inactif depuis " + strconv.Itoa(semaine) + " semaines")
				}

			}

			mysql.UpdateMembresInactif(inf[1])
			time.Sleep(1 * time.Second)
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

func UpdateOnlinePlayer(secondeboucle time.Duration) {
	onlineplayer := 0

	for {
		time.Sleep(time.Second * secondeboucle)
		var count int
		session := framework.Session

		for _, val := range framework.OnlinePlayer {
			count = count + val
		}

		if onlineplayer != count {
			editchannel := &discordgo.ChannelEdit{
				Name:     "🪐  Online : " + strconv.Itoa(count),
				Position: 2,
			}

			_, err := session.ChannelEditComplex(viper.GetString("ChannelID.OnlinePlayer"), editchannel)
			if err != nil {
				logger.ErrorLogger.Println(err)
				continue
			}
			onlineplayer = count
		}

	}
}
