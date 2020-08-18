package command

import (
	"fmt"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/rcon"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	if ctx.Staff >= 1 {
		msg := rcon.RconCommandeWhitelistAdd("Aathoss")
		fmt.Println(msg)
	}
}
