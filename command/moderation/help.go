package moderation

import (
	"regexp"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/spf13/viper"
)

func HelpCommand(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	//dm, _ := ctx.Discord.UserChannelCreate(ctx.Message.Author.ID)
	var cmdalias string
	tab := [][]string{}
	cmds := framework.Cmdliste

	for _, cmdStruct := range cmds {
		var info []string
		_, alias, grade, help := ctx.CmdHandler.GetAllCmd(cmdStruct)

		if grade <= ctx.NiveauStaff {
			if alias > 0 {
				cmdalias = cmdalias + (viper.GetString("PrefixMsg") + cmdStruct) + " "
			}
			if alias == 0 {
				info = append(info, (viper.GetString("PrefixMsg") + cmdStruct), cmdalias, help)
				tab = append(tab, info)
				cmdalias = ""
			}
		}
	}

	embedHelp := framework.NewEmbed().
		SetTitle("Liste des commande :").
		SetColor(0x6E318E).
		MessageHelp(tab).MessageEmbed

	_, err := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embedHelp)
	if err != nil {
		notmp, _ := regexp.MatchString(`50007`, err.Error())
		if notmp == true {
			ctx.Discord.ChannelMessageSendEmbed(ctx.TextChannel.ID, framework.EmbedMPClose)
		}
		logger.ErrorLogger.Println(err)
		return
	}
}
