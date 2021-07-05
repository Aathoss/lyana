package framework

import (
	"regexp"

	"github.com/spf13/viper"
)

func LogsChannel(msg string) {
	if viper.GetBool("Dev.test") != true {
		Session.ChannelMessageSend(viper.GetString("ChannelID.Log"), msg)
	}
}

func LogsRolePolicier(msg string) {
	Session.ChannelMessageSend("785611617026834432", msg)
}

func ErrorDM(ctx Context, err error) {
	if err != nil {
		notmp, _ := regexp.MatchString(`50007`, err.Error())
		if notmp == true {
			ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, EmbedMPClose)
		}
	}
}
