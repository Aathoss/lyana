package stats

import (
	"runtime"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

var (
	Version   = "0.3.0"
	startTime = time.Now()
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
				"\nUptime : **"+framework.Calculetime(startTime.Unix(), 0)+"**"+
				"\nMémoire utiliser : **"+humanize.Bytes(stats.Alloc)+"** / **"+humanize.Bytes(stats.Sys)+"**"+
				"\nNombre de requête SQL : **"+strconv.Itoa(framework.SQlRequest)+"**"+
				"\nNombre d'actualisation channel online : **"+strconv.Itoa(framework.OnlineActulise)+"**", false).
		SetFooter(ctx.Message.Author.Username + " | Date : " + tNow.Format("2/1 15:04:05")).MessageEmbed
	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	return
}
