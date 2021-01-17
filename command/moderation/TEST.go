package moderation

import (
	"gitlab.com/lyana/framework"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	//mysql.VerifInactif()
	//modules.VerifRule(framework.Session)

	//modules.VerifInactif(ctx.Discord)

}
