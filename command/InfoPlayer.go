package command

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

var (
	stats statsv1
	count int
)

type statsv1 struct {
	playertime        int64
	jump              int64
	deaths            int64
	craftitems        int64
	mineblocks        int64
	breakitems        int64
	mobkills          int64
	useitems          int64
	minediamondore    int64
	mineancientdebris int64
}

func InfoPlayer(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	mentions := ctx.Message.Mentions
	var user *discordgo.Member

	if len(mentions) == 1 {
		if mentions[0].ID == "742505851075428423" {
			ctx.Discord.ChannelMessageSend(ctx.Message.ChannelID, "Qui t’a donné l'autorisation de me parler ? xD")
		}

		u, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), mentions[0].ID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		user = u
	}
	if len(mentions) == 0 {
		u, err := ctx.Discord.GuildMember(viper.GetString("GuildID"), ctx.Message.Author.ID)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
		user = u
	}

	t1, _ := user.JoinedAt.Parse()
	_, pseudoMC, t2, _ := mysql.GetWhitelist(user.User.ID)

	requestSQLPlayer(pseudoMC)

	//Compte le nombre de message envoyer par l'utilisateur
	err := framework.DBLyana.QueryRow("SELECT COUNT(uuid) FROM logs WHERE categorie=? AND uuid=?", "msgcount", user.User.ID).Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	messageCount := "\nMessages envoyés : **" + strconv.Itoa(count) + "**"
	if err != nil {
		messageCount = ""
	}

	messagewhitelist := ""
	messagepseudo := " | <:no:742854427207008307> Vous n'êtes pas whitelist"
	statsminecraft := ""
	if len(pseudoMC) > 1 {
		messagewhitelist = "\nVous êtes whitelist depuis : **" + framework.Calculetime(int64(t2), 0) + "**"
		messagepseudo = " | <:CraftingTable:753547645875912736> Pseudo : " + pseudoMC
		statsminecraft = "Minecraft Stats"
	}

	embed := framework.NewEmbed().
		SetTitle("Votre Carte d'identité : "+user.User.Username+messagepseudo).
		SetColor(viper.GetInt("EmbedColor.Informations")).
		AddField("Discord", "Vous êtes arrivé il y à : **"+framework.Calculetime(t1.Unix(), 0)+"**"+messagewhitelist+messageCount, true).
		AddField(statsminecraft, "Temps de jeux : **"+framework.Calculetime(stats.playertime, 2)+"**"+
			"\nNombre de sautes : **"+humanize.Comma(stats.jump)+"**"+
			"\nNombre de morts : **"+humanize.Comma(stats.deaths)+"**"+
			"\nItems craft : **"+humanize.Comma(stats.craftitems)+"**"+
			"\nBloc miné : **"+humanize.Comma(stats.mineblocks)+"**"+
			"\nItems brisés : **"+humanize.Comma(stats.breakitems)+"**"+
			"\nMob tuer : **"+humanize.Comma(stats.mobkills)+"**"+
			"\nItems utilisés : **"+humanize.Comma(stats.useitems)+"**"+
			"\nMinerais de diamant : **"+humanize.Comma(stats.minediamondore)+"**"+
			"\nMinerais d'ancien débris : **"+humanize.Comma(stats.mineancientdebris)+"**", false).MessageEmbed

	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}

func requestSQLPlayer(player string) {
	db := mysql.DbConnMC()
	defer db.Close()

	err := db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='playertime'").Scan(&stats.playertime)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='jump'").Scan(&stats.jump)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='deaths'").Scan(&stats.deaths)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='craftitems'").Scan(&stats.craftitems)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='mineblocks'").Scan(&stats.mineblocks)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='breakitems'").Scan(&stats.breakitems)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='mobkills'").Scan(&stats.mobkills)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='useitems'").Scan(&stats.useitems)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='mine_diamond_ore'").Scan(&stats.minediamondore)
	db.QueryRow("SELECT CONTENT FROM PLAYERDATA WHERE PLAYER='" + player + "' AND VARIABLE='mine_ancient_debris'").Scan(&stats.mineancientdebris)

	return
}
