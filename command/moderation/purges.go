package moderation

import (
	"strconv"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/spf13/viper"
)

func Purges(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if len(ctx.Args) < 1 {
		embed := framework.NewEmbed().
			SetTitle("Syntaxe : `" + viper.GetString("PrefixMsg") + ctx.Commande + " <nombre_de_message>`").MessageEmbed
		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
		return
	}

	limiteMsg, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	if limiteMsg > 2500 {
		embed := framework.NewEmbed().
			SetTitle("Vous ne pouvez pas effectuer une purge supérieure à 2.500.").MessageEmbed
		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
		return
	}

	t1, err := ctx.Message.Timestamp.Parse()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	var limit = limiteMsg
	var horsdate int
	var beforeID string
	messages := []string{}

	for limit > 0 {
		var l int

		if limit > 100 {
			l = 100
		} else {
			l = limit
		}

		msgs, err := ctx.Discord.ChannelMessages(ctx.Message.ChannelID, l, beforeID, "", "")
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}

		if len(msgs) == 0 {
			horsdate += limit
			limit = 0
			break
		}

		beforeID = msgs[len(msgs)-1].ID

		var count int

		for _, msg := range msgs {
			t2, err := msg.Timestamp.Parse()
			if err != nil {
				logger.ErrorLogger.Println(err)
				return
			}
			diff := t1.Sub(t2).Seconds()

			count++
			if int64(diff) < (60*60*24*14)-(30) {
				messages = append(messages, msg.ID)
			} else {
				horsdate++
			}
		}
		limit -= count
	}

	err = ctx.Discord.ChannelMessagesBulkDelete(ctx.Message.ChannelID, messages)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	framework.LogsChannel("[<:ClearMessages:753547497330180156>] **" + strconv.Itoa(limiteMsg) + "** Messages à traiter **| " + strconv.Itoa(limiteMsg-horsdate) + "** Supprimés **| " + strconv.Itoa(horsdate) + "** Échec")
}
