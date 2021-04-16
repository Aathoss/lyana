package event

import (
	"strconv"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/mysql"
)

func EventTermine(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	num, _ := strconv.Atoi(ctx.Args[0])
	mysql.EditStatus(3, num)
}
