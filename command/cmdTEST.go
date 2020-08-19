package command

import (
	"gitlab.com/lyana/framework"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	if ctx.Staff >= 1 {
		//rcon.RconCommandeTest(ctx.MessageSplit[1])
	}
}
