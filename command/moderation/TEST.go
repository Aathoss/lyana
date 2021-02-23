package moderation

import (
	"gitlab.com/lyana/framework"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

}

func TestLocal() {
}
