package modules

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
)

var (
	Version = "0.2.9"
)

func Ready(s *discordgo.Session, Event *discordgo.Event) {
	framework.Session = s

	if viper.GetBool("Dev.test") == false {
		if Event.Type == "READY" {
			s.UpdateStatus(0, viper.GetString("Motd"))
			logger.InfoLogger.Println("Le bot est dispo. [Appuyez sur CTRL+C pour l'arrêter !]")

			framework.LogsChannel("[:tools:] [v:" + Version + "] **Lyana** à correctement démarré accompagné des " + strconv.Itoa(framework.CountCommand+viper.GetInt("ModuleCount")) + " modules.")
		}
	}
}
