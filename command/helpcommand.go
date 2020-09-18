package command

import (
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
)

func HelpCommand(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	dm, _ := ctx.Discord.UserChannelCreate(ctx.Message.Author.ID)

	var cmd string
	var cmdhelp string

	cmds := framework.Cmdliste

	for _, cmdStruct := range cmds {
		cmd = cmd + viper.GetString("PrefixMsg") + cmdStruct + "\n"
		cmdhelp = cmdhelp + ctx.CmdHandler.GetAllCmd(cmdStruct) + "\n"
	}

	embedHelp := framework.NewEmbed().
		SetTitle("Liste des commande :").
		SetColor(0x6E318E).
		AddField("Commande", cmd, true).
		AddField("Info", cmdhelp, true).MessageEmbed

	_, err := ctx.Discord.ChannelMessageSendEmbed(dm.ID, embedHelp)
	framework.ErrorDM(ctx, err)
}
