package event

import (
	"strconv"

	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/mysql"
)

func EventTermine(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	num, _ := strconv.Atoi(ctx.Args[0])
	mysql.EditStatus(3, num)
}
