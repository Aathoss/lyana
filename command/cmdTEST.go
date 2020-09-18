package command

import (
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	if viper.GetBool("Dev.test") == true {
		mysql.VerifRuleTimestamp()
	}
}

func Test1(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	if viper.GetBool("Dev.test") == true {
		t1, err := ctx.Message.Timestamp.Parse()
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		mysql.AddRule(ctx.Message.Author.ID, t1.Unix())
	}
}
