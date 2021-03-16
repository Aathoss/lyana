package moderation

import (
	"gitlab.com/lyana/framework"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	//modules.UpdateOnlinePlayer(10)
}

func TestLocal() {
}
