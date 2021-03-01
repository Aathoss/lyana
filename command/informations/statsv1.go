package informations

import (
	"github.com/dustin/go-humanize"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/mysql"
)

var (
	stats statsv1
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

//StatsUnispaceV1 retourn les satistique globale du serveur unispace v1
func StatsUnispaceV1(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)
	requestSQL()

	embed := framework.NewEmbed().
		SetTitle("Statistique global de Unispace V1").
		SetColor(0x3A0071).
		SetDescription("Temps de jeux : **" + framework.Calculetime(stats.playertime, 2) + "**" +
			"\nNombre de sautes : **" + humanize.Comma(stats.jump) + "**" +
			"\nNombre de morts : **" + humanize.Comma(stats.deaths) + "**" +
			"\nItems craft : **" + humanize.Comma(stats.craftitems) + "**" +
			"\nBloc miné : **" + humanize.Comma(stats.mineblocks) + "**" +
			"\nItems brisés : **" + humanize.Comma(stats.breakitems) + "**" +
			"\nMob tuer : **" + humanize.Comma(stats.mobkills) + "**" +
			"\nItems utilisés : **" + humanize.Comma(stats.useitems) + "**" +
			"\nMinerais de diamant : **" + humanize.Comma(stats.minediamondore) + "**" +
			"\nMinerais d'ancien débris : **" + humanize.Comma(stats.mineancientdebris) + "**").MessageEmbed
	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}

func requestSQL() {
	db := mysql.DbConnMC()

	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='playertime'").Scan(&stats.playertime)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='jump'").Scan(&stats.jump)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='deaths'").Scan(&stats.deaths)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='craftitems'").Scan(&stats.craftitems)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='mineblocks'").Scan(&stats.mineblocks)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='breakitems'").Scan(&stats.breakitems)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='mobkills'").Scan(&stats.mobkills)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='useitems'").Scan(&stats.useitems)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='mine_diamond_ore'").Scan(&stats.minediamondore)
	db.QueryRow("SELECT SUM(CONTENT) FROM PLAYERDATA WHERE VARIABLE='mine_ancient_debris'").Scan(&stats.mineancientdebris)

	return
}
