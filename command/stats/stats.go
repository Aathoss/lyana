package stats

import (
	"runtime"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"gitlab.com/lyana/framework"
)

var (
	startTime = time.Now()
)

func Statistique(ctx framework.Context) {

	//dm, _ := ctx.Discord.UserChannelCreate(ctx.User.ID)
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	if len(ctx.Args) == 0 {
		stats := runtime.MemStats{}
		runtime.ReadMemStats(&stats)

		embed := framework.NewEmbed().
			SetTitle("Statistique bot :").
			SetColor(0x5f27cd).
			//SetDescription("Commande : "+viper.GetString("PrefixMsg")+ctx.Commande+" <options : **staff** / **ticket**>").
			AddField("Informations basic",
				"go version : **"+runtime.Version()+"**"+
					"\ndiscordgo version : **"+discordgo.VERSION+"**"+
					"\nuptime : **"+framework.Calculetime(startTime.Unix(), 0)+"**"+
					"\nMémoire utiliser : **"+humanize.Bytes(stats.Alloc)+"** / **"+humanize.Bytes(stats.Sys)+"**"+
					"\nNombre de GoRoutine : **"+strconv.Itoa(runtime.NumGoroutine())+"**"+
					"\nMessage total : **"+strconv.Itoa(framework.CountMsg)+"**"+
					"\nNombre de requête SQL : **"+strconv.Itoa(framework.SQlRequest)+"**"+
					"\nNombre d'actualisation channel online : **"+strconv.Itoa(framework.OnlineActulise)+"**", true).MessageEmbed
		_, err := ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
		framework.ErrorDM(ctx, err)
		return
	}

	if ctx.Args[0] == "staff" {

	}

	if ctx.Args[0] == "ticket" {

	}

}
