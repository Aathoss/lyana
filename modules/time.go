package modules

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/unispace/framework"
	"gitlab.com/unispace/logger"
)

func UpdateOnlinePlayer(session *discordgo.Session) {
	OnlinePlayerRcon()
	_, err := session.ChannelEdit(viper.GetString("ChannelID.OnlinePlayer"), "🪐  Online : "+strconv.Itoa(framework.OnlinePlayer))
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}
