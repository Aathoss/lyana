package stats

import (
	"runtime"
	"strconv"
	"time"

	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
)

var (
	StartTime      = time.Now()
	DebugUserMysql = 0
)

func Statistique(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	inactif, err := mysql.CompteInactif()
	if err != nil {
		inactif = -1
	}
	tNow := time.Now()

	membersgrade, err := ctx.Discord.State.Guild(ctx.Guild.ID)
	if err != nil {
		logger.DebugLogger.Println(err)
		return
	}

	//Débug mysql count le nombre d'utilisateur actif
	err = framework.DBLyana.QueryRow("SELECT COUNT(*) FROM information_schema.PROCESSLIST WHERE Host LIKE '172.18.0.10%'").Scan(&DebugUserMysql)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	embed := framework.NewEmbed().
		SetTitle("Statistique bot :").
		SetColor(0x5f27cd).
		AddField("Informations Discord", "\nMembre : **"+strconv.Itoa(membersgrade.MemberCount-3)+"**"+
			"\nInactif : **"+strconv.Itoa(inactif)+"**/**"+strconv.Itoa(membersgrade.MemberCount-3)+"**"+
			"\nMessage total : **"+strconv.Itoa(framework.CountMsg)+"**", false).
		AddField("Informations Bot",
			"V. Go : **"+runtime.Version()+"**"+
				"\nV. Discord : **"+discordgo.VERSION+"**"+
				"\nRoutine : **"+strconv.Itoa(runtime.NumGoroutine())+"**"+
				"\nUptime : **"+framework.Calculetime(StartTime.Unix(), 0)+"**"+
				"\nMémoire utiliser : **"+humanize.Bytes(stats.Alloc)+"** / **"+humanize.Bytes(stats.Sys)+"**"+
				"\n\n[Débug] Nombre de pool : **"+strconv.Itoa(DebugUserMysql)+"**", false).
		SetFooter(ctx.Message.Author.Username + " | Date : " + tNow.Format("2/1 15:04:05")).MessageEmbed
	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	return
}
