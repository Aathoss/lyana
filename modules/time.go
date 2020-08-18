package modules

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/rcon"
)

func UpdateOnlinePlayer(session *discordgo.Session) {
	rcon.RconCommandeList()
	_, err := session.ChannelEdit(viper.GetString("ChannelID.OnlinePlayer"), "🪐  Online : "+strconv.Itoa(framework.OnlinePlayer))
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
