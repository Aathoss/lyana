package moderation

import (
	"fmt"
	"strconv"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/spf13/viper"
)

func MessageGlobalMp(ctx framework.Context) {

	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	membersgrade, err := ctx.Discord.State.Guild(ctx.Guild.ID)
	if err != nil {
		logger.DebugLogger.Println(err)
		return
	}

	if len(viper.GetString("GlobalMsgSend")) == 0 {
		framework.LogsChannel("[!globalmp] Veuillez définir un message à envoyer !")
		return
	}

	countMembersTemp := 100
	countMembers := membersgrade.MemberCount
	idMembers := ""
	count := 0
	countMP := 0

	for {

		if countMembers >= countMembersTemp {
			countMembers = countMembers - countMembersTemp

		} else {
			countMembersTemp = countMembers
			countMembers = 0
		}
		fmt.Println("  ")

		membre, err := ctx.Discord.GuildMembers(viper.GetString("GuildID"), idMembers, countMembersTemp)
		if err != nil {
			logger.DebugLogger.Println(err)
			continue
		}

		for _, key := range membre {
			count++
			logger.InfoLogger.Println(strconv.Itoa(count) + " | " + key.User.ID + " | " + key.User.Username)
			if count <= viper.GetInt("GlobalMsgSendReprise") {
				continue
			}

			idMembers = key.User.ID

			_, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), key.User.ID)
			if err != nil {
				countMP++
				logger.ErrorLogger.Println("-> Mp Close")
				framework.LogsChannel(":mailbox_with_mail: **Erreur mp :x: ** : N°**" + strconv.Itoa(count) + "** | " + key.User.ID + " | " + key.User.Username)
				continue
			} else {
				dm, err := ctx.Discord.UserChannelCreate(key.User.ID)
				if err != nil {
					logger.DebugLogger.Println(err)
					continue
				}
				_, err = ctx.Discord.ChannelMessageSend(dm.ID, viper.GetString("GlobalMsgSend"))
				if err != nil {
					logger.DebugLogger.Println(err)
					countMP++
					logger.ErrorLogger.Println("-> Mp Close")
					framework.LogsChannel(":mailbox_with_mail: **Erreur mp :x: ** : N°**" + strconv.Itoa(count) + "** | " + key.User.ID + " | " + key.User.Username)
					continue
				}
			}
			logger.InfoLogger.Println("-> Mp Envoyer")
			logger.InfoLogger.Println(" ")
			framework.LogsChannel(":mailbox_with_mail: **Message envoyé** : N°**" + strconv.Itoa(count) + "** | " + key.User.ID + " | " + key.User.Username)

		}

		if countMembers == 0 {
			fmt.Println(countMP)
			break
		}
		continue
	}
}
