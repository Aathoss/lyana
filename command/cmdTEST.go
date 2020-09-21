package command

import (
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/mysql"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	mysql.VerifInactif()
}
