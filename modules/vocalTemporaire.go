package modules

import (
	"regexp"
	"strings"

	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func VocalTemporaire(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {

	/* verifchannel, _ := s.State.GuildChannel(viper.GetString("GuildID"), m.ChannelID)
	if verifchannel.ParentID != viper.GetString("Categorie.TempVoc") {
		return
	} */
	if m.ChannelID == viper.GetString("ChannelID.VocalGeneral") {
		return
	}

	//Crée un nouveau channel vocal quand une personne rejoinds un channel données
	if m.ChannelID == viper.GetString("ChannelID.VocalTempVoc") {
		user, err := s.GuildMember(viper.GetString("GuildID"), m.UserID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		//variable de channel
		channelname := user.User.Username
		channellimite := 10

		usercount, channelid := mysql.Countuser(m.UserID)

		//Evite la duplication de channel inutile
		//Déplace la personne dans sont channel (si le channel existe encore sinon on le recrée)
		if channelid != "" {
			lvlerr := 0

			err = s.GuildMemberMove(viper.GetString("GuildID"), user.User.ID, &channelid)
			if err != nil {
				notmp, _ := regexp.MatchString(`10003`, err.Error())
				if notmp == true {
					err = mysql.UpdateChannelID(user.User.ID, "")
					if err != nil {
						logger.ErrorLogger.Println(err)
						return
					}
				} else {
					lvlerr = 1
					logger.ErrorLogger.Println(err)
					return
				}
			}

			if lvlerr == 0 {
				return
			}
		}

		if usercount == 0 {
			err = mysql.InsertCreation(m.UserID, channelname, channellimite)
			if err != nil {
				logger.ErrorLogger.Println(err)
				return
			}
		} else {
			channelname, channellimite = mysql.ReturnConfigChannel(m.UserID)
		}

		//création du nouveau channel vocal temporaire
		channel, err := s.GuildChannelCreateComplex(viper.GetString("GuildID"), discordgo.GuildChannelCreateData{
			Name:      channelname,
			Type:      2,
			ParentID:  viper.GetString("Categorie.VocalTempVoc"),
			UserLimit: channellimite,
			Position:  7,
		})
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		//Insert le nouveau channel id dans la table
		err = mysql.UpdateChannelID(m.UserID, channel.ID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		//déplacement du membre dans le nouveau channel
		err = s.GuildMemberMove(viper.GetString("GuildID"), user.User.ID, &channel.ID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		return
	}

	//Vérifie les channel vide afin de les supprimé
	if m.ChannelID == "" {
		listchannel := mysql.ReturnChannelIDAll()
		guild, _ := s.State.Guild(m.GuildID)

		if len(listchannel) == 0 {
			return
		}

		for _, channel := range listchannel {
			if len(channel) == 0 {
				continue
			}

			count := 0

			for _, key := range guild.VoiceStates {
				if key.ChannelID == channel {
					count++
				}
			}

			if count == 0 {
				channelrm, err := s.ChannelDelete(channel)
				if err != nil {
					logger.InfoLogger.Println("[" + channel + "] " + err.Error())

					if strings.Count(err.Error(), "10003") >= 1 {
						logger.InfoLogger.Println("[" + channel + "] Channel retirer de la base de données avec succès")
						err = mysql.RemoveChannelID(channel)
						if err != nil {
							logger.DebugLogger.Println(err)
						}
					}
					continue
				}

				err = mysql.RemoveChannelID(channelrm.ID)
				if err != nil {
					logger.DebugLogger.Println(err)
				}
			}
		}
	}
}
