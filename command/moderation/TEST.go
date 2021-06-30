package moderation

import (
	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/spf13/viper"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	//modules.UpdateOnlinePlayer(10)
	if len(viper.GetString("GlobalMsgSend")) == 0 {
		framework.LogsChannel("[!globalmp] Veuillez définir un message à envoyer !")
		return
	}
	
	dm, err := ctx.Discord.UserChannelCreate("284053267988873226")
	if err != nil {
		logger.DebugLogger.Println(err)
	}
	_, err = ctx.Discord.ChannelMessageSend(dm.ID, viper.GetString("GlobalMsgSend"))
	if err != nil {
		logger.DebugLogger.Println(err)
	}
}

func TestLocal() {
}
