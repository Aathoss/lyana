package event

import (
	"strconv"
	"strings"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

//0 = Update de l'embed
//1 = Création de l'embed
func ConstructionEmbedEvent(situation int, session *discordgo.Session, tab []string) (embedmessageid string) {

	color := 0
	prepare := ""

	//Vérifie et modifie le status du message
	if tab[1] == "dev" {
		color = 0x2A84A1
		prepare = "**:gear: :heavy_minus_sign: :gear: :heavy_minus_sign: [ En Construction ] :heavy_minus_sign: :gear: :heavy_minus_sign: :gear:**\n"
	} else if tab[1] == "en cours" {
		color = 0x0CBA01
		prepare = "**:clapper: :heavy_minus_sign: :clapper: :heavy_minus_sign: [ Évent en cours ] :heavy_minus_sign: :clapper: :heavy_minus_sign: :clapper:**\n"

	} else if tab[1] == "prepterminer" {
		color = 0xC12600
		prepare = "**:x: :heavy_minus_sign: :x: :heavy_minus_sign: [ Évent Terminé ] :heavy_minus_sign: :x: :heavy_minus_sign: :x:**\n"

	} else if tab[1] == "à venir" {
		color = 0xFFFF25
		prepare = "**:hourglass_flowing_sand: :heavy_minus_sign: :hourglass_flowing_sand: :heavy_minus_sign: [ Évent à venir ] :heavy_minus_sign: :hourglass_flowing_sand: :heavy_minus_sign: :hourglass_flowing_sand:**\n"
	}

	//vérifie et construit la parti organisateur de l'évent
	if tab[10] != "nil" {
		user, _ := session.GuildMember(viper.GetString("GuildID"), tab[10])
		if user == nil {
			prepare = prepare + "\n**Organisateur :speech_balloon:** :heavy_minus_sign: " + tab[10]

		} else {
			prepare = prepare + "\n**Organisateur :speech_balloon:** :heavy_minus_sign: " + user.Mention()
		}
	}

	//vérifie et construit la parti emplacement
	if tab[5] != "nil" {
		prepare = prepare + "\n**Emplacement :satellite_orbital:** :heavy_minus_sign: " + tab[5]
	}

	//vérifie et construit la parti programmation / décompte de l'évent
	if tab[7] != "nil" {
		dateunix, _ := strconv.Atoi(tab[7])
		date := time.Unix(int64(dateunix), 0)
		prepare = prepare + "\n**Programmé le :clock:** :heavy_minus_sign: " + date.Format("02/01/2006 à 15:04:05")
		if tab[1] == "à venir" || tab[1] == "dev" {
			prepare = prepare + "\n:heavy_minus_sign: :heavy_minus_sign: :heavy_minus_sign::heavy_minus_sign::heavy_minus_sign: :heavy_minus_sign: décompte --> " + framework.Calculetime(date.Unix(), 1)
		}
	}

	//vérifie et construit la parti description de l'évent
	if tab[6] != "nil" {
		prepare = prepare + "\n\n**Description :clipboard:**\n" + strings.Replace(tab[6], `\n`, "\n", -1)
	}

	//vérifie et construit la parti récompense de l'évent
	if tab[8] != "nil" {
		prepare = prepare + "\n\n**Récompense :gift:**\n" + strings.Replace(tab[8], `\n`, "\n", -1)
	}

	//vérifie et construit la parti participant à l'évent
	if tab[9] != "" {
		countprepare := 1975 - len(prepare)
		personnes := strings.Split(tab[9], ",")
		mentions := ""

		for _, p := range personnes {
			user, err := session.GuildMember(viper.GetString("GuildID"), p)
			if err != nil {
				continue
			}

			if len(mentions) == 0 {
				mentions = mentions + user.User.Mention()
			} else {
				mentions = mentions + " | " + user.User.Mention()
			}
		}

		if len(mentions) >= countprepare {
			prepare = prepare + "\n\n**Participant " + strconv.Itoa(len(personnes)) + " :tickets:**\n" + mentions[0:countprepare] + " ..."
		} else {
			prepare = prepare + "\n\n**Participant " + strconv.Itoa(len(personnes)) + " :tickets:**\n" + mentions
		}
	}

	t1 := time.Now()
	updatefooter := t1.Format("1/2 à 15:04:05")

	//update de l'embed
	if situation == 0 {
		message, err := session.ChannelMessageEditEmbed(tab[3], tab[2], &discordgo.MessageEmbed{
			Title:       tab[4],
			Color:       color,
			Description: prepare,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Dérnier mise à jour : " + updatefooter,
			},
		})
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		return message.ID
	}

	if situation == 1 {
		message, err := session.ChannelMessageSendEmbed(tab[3], &discordgo.MessageEmbed{
			Title:       tab[4],
			Color:       color,
			Description: prepare,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Dérnier mise à jour : " + updatefooter,
			},
		})
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		return message.ID
	}

	return ""
}
