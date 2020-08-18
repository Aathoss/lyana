package command

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"gitlab.com/lyana/framework"
)

// credit to github.com/iopred/bruxism for this command

var (
	startTime = time.Now()
	countMsg  int
)

func SetCountMsg() {
	countMsg = countMsg + 1
}

func getDurationString(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%02d:%02d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}

func InfoCommand(ctx framework.Context) {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)

	fmt.Println("go version : ", runtime.Version())
	fmt.Println("discordgo version : ", discordgo.VERSION)
	fmt.Println("uptime : ", getDurationString(time.Now().Sub(startTime)))
	fmt.Sprintf("Mémoire utiliser : %s / %s", humanize.Bytes(stats.Alloc),
		humanize.Bytes(stats.Sys))
	fmt.Println("Nombre de GoRoutine : ", strconv.Itoa(runtime.NumGoroutine()))
	fmt.Println("\nMessage total : ", countMsg)
}
