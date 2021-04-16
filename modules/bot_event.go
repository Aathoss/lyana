package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
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
