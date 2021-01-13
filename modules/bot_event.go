package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/command/stats"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/rcon"
)

var (
	ready bool
)

func Ready(s *discordgo.Session, Event *discordgo.Event) {
	framework.Session = s

	if Event.Type == "READY" && ready == false {
		ready = true
		s.UpdateStatus(0, viper.GetString("Motd"))
		logger.InfoLogger.Println("Le bot est dispo. [Appuyez sur CTRL+C pour l'arrêter !]")

		framework.LogsChannel("[:tools:] [v:" + stats.Version + "] **Lyana** à correctement démarré")
		rcon.Openrcon()
	}
}
