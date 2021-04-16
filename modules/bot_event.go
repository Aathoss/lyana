package modules

import (
	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	ready bool
)

func Ready(s *discordgo.Session, Event *discordgo.Event) {
	framework.Session = s

	if Event.Type == "READY" && ready == false {
		ready = true
		s.UpdateGameStatus(0, viper.GetString("Motd"))
		logger.InfoLogger.Println("Le bot est dispo. [Appuyez sur CTRL+C pour l'arrêter !]")

		framework.LogsChannel("[:tools:] [v:" + framework.Version + "] **Lyana** à correctement démarré")
	}
}
