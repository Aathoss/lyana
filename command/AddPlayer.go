package command

import (
	"strconv"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/spf13/viper"
)

func AddPlayer(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if len(ctx.Args) > 1 {
		mentions := ctx.Message.Mentions
		playermc := ctx.Args[1]

		if len(mentions) == 0 {
			framework.LogsChannel("[:x:] [" + viper.GetString("PrefixMsg") + ctx.Commande + "] Vous n'avez pas mentionné de personne !")
			return
		}

		countuuid, countplayer := mysql.VerifPlayerMC(mentions[0].ID, playermc)
		if countplayer == 1 {
			framework.LogsChannel("[:open_mouth:] [Utilisateur : " + mentions[0].String() + " | Pseudo : " + playermc + "] Ce pseudo existe déjà dans notre base de données.")
			return
		}

		if countuuid == 1 {
			framework.LogsChannel("[:open_mouth:] [Utilisateur : " + mentions[0].String() + " | Pseudo : " + playermc + "] Il n'est pas autorisé d'avoir un double compte...")
			return
		}

		//Ajoute de la personne en base de données + move des grade
		ctx.Discord.ChannelMessageSend(viper.GetString("ChannelID.General"), "<:CraftingTable:753547645875912736> Je viens de craft votre carte d'accès au serveur, nous vous souhaitons la bienvenue parmi nous "+mentions[0].Mention()+".")
		err := mysql.AddWhitelist(mentions[0].ID, playermc)
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		err = ctx.Discord.GuildMemberRoleRemove(viper.GetString("GuildID"), mentions[0].ID, "735281835080286291")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}
		err = ctx.Discord.GuildMemberRoleAdd(viper.GetString("GuildID"), mentions[0].ID, "820404799119818793")
		if err != nil {
			logger.ErrorLogger.Println(err)
		}

		WhitelistPlayerMC(playermc)
	}

	if len(ctx.Args) <= 1 {
		embed := framework.NewEmbed().
			SetTitle("Il semble y avoir une erreur !").
			SetColor(viper.GetInt("EmbedColor.Error")).
			SetDescription("Veuillez respecter ce format : " + viper.GetString("PrefixMsg") + "addplayer <tag_discord> <player_mc>").MessageEmbed

		message, _ := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
		time.Sleep(time.Second * 10)
		ctx.Discord.ChannelMessageDelete(message.ChannelID, message.ID)
	}
}

func WhitelistPlayerMC(player string) {
	for {
		err := framework.ConnectMC[0].Authenticate()
		if err != nil {
			framework.OnlineServer[0] = "offline"
			framework.Connect(0)
		}

		_, err = framework.ConnectMC[0].SendCommand("ewhitelist add " + player)
		if err != nil {
			logger.ErrorLogger.Println("MC-Host : "+viper.GetString("Minecraft."+strconv.Itoa(0)+".Name")+" | Command failed : ", err)
			time.Sleep(time.Second * 5)
			continue
		}
		break
	}
}