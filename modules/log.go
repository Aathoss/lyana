package modules

import (
	"log"
	"regexp"

	"github.com/spf13/viper"
	"gitlab.com/unispace/framework"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogDiscord(msg string) {
	if viper.GetBool("Dev.test") != true {
		framework.Session.ChannelMessageSend(viper.GetString("ChannelID.Log"), msg)
	}
}

func ErrorDM(ctx framework.Context, err error) {
	if err != nil {
		notmp, _ := regexp.MatchString(`50007`, err.Error())
		if notmp == true {
			ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, embedMPClose)
		}
	}
}
