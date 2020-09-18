package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	bot "github.com/bwmarrin/discordgo"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gitlab.com/lyana/command"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/modules"
	"gitlab.com/lyana/mysql"
)

// Variable
var (
	CmdHandler *framework.CommandHandler
	Token      = "NzQyNTA1ODUxMDc1NDI4NDIz.XzHGdQ.3e0fddUjcsxT19C1FwtussGA-fk"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.InfoLogger.Println("Config file changed:", e.Name)
	})
}

func main() {
	framework.FormatTime(1597530893)

	logger.InfoLogger.Println("\n---------------------------------\nDémarrage du bot en cours")

	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	dg, err := bot.New("Bot " + Token)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la session discord,", err)
		return
	}

	dg.AddHandler(modules.Ready)
	go dg.AddHandler(modules.Stats)
	dg.AddHandler(modules.GuildMemberAdd)
	dg.AddHandler(modules.GuildMemberLeave)
	go dg.AddHandler(modules.ReactionAdd)
	dg.AddHandler(commandHandler)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la connexion,", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	t := time.Tick(1 * time.Minute)

out:
	for {
		select {
		case <-sc:
			framework.LogsChannel("[:tools:] **Lyana** s'est déconnecté de l'univers !")
			break out
		case <-t:
			// Exécute des action
			modules.ExecuteTime()
		}
	}
}

func commandHandler(s *bot.Session, m *bot.MessageCreate) {
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	if viper.GetBool("Dev.PrintMessage") == true {
		log.Println(m.Content)
	}

	if viper.GetBool("Dev.test") != true {
		mysql.NewCountMessage(user.ID)
	}
	framework.CountMsg = framework.CountMsg + 1

	content := m.Content
	if len(content) <= len(viper.GetString("PrefixMsg")) {
		return
	}
	if content[:len(viper.GetString("PrefixMsg"))] != viper.GetString("PrefixMsg") {
		return
	}
	content = content[len(viper.GetString("PrefixMsg")):]
	if len(content) < 1 {
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de l'obtention du channel,", err)
		return
	}

	staff := 0
	if channel.Type != discordgo.ChannelTypeDM {
		staff = framework.VerifStaff(m.Member.Roles)
	}

	checkCmdName := CmdHandler.CheckCmd(content)
	command, found, permission := CmdHandler.Get(checkCmdName, staff)
	if !found {
		return
	}
	if permission == false {
		s.ChannelMessageSendEmbed(m.ChannelID, framework.EmbedPermissionFalse)
		return
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de l'obtention de la guilde,", err)
		return
	}

	ctx := framework.NewContext(s, guild, channel, user, m, CmdHandler, checkCmdName, staff)
	messageSplit := strings.Fields(content)
	if len(strings.Fields(checkCmdName)) == 1 {
		ctx.Args = messageSplit[1:]
	}
	if len(strings.Fields(checkCmdName)) == 2 {
		ctx.Args = messageSplit[2:]
	}
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("test", []string{}, 1, command.Test, "???")
	CmdHandler.Register("test1", []string{}, 1, command.Test1, "???")

	//Commande Modération
	CmdHandler.Register("purge", []string{}, 1, command.Purges, "La commande permet d'effectuer un netoyage d'un channel limite à 2.500 Message")

	//Commande Liée à minecraft
	CmdHandler.Register("fiche", []string{}, 0, command.InfoPlayer, "Permet de voir votre fiche utilisateur")
	CmdHandler.Register("online", []string{}, 0, command.OnlinePlayer, "Affiche les joueurs connecté")
	CmdHandler.Register("signal", []string{}, 0, command.AddSignalement, "Permet au joueurs whitelist sur le serveur de signaler un autre joueurs commétande une infraction")
	CmdHandler.Register("rsignal", []string{}, 1, command.RemoveSignalement, "Permet au staff de retiré un signalement")
	CmdHandler.Register("addplayer", []string{}, 1, command.AddPlayer, "???")

}
