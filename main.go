package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	bot "github.com/bwmarrin/discordgo"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gitlab.com/unispace/command"
	"gitlab.com/unispace/framework"
	"gitlab.com/unispace/logger"
	"gitlab.com/unispace/modules"
	"gitlab.com/unispace/mysql"
)

// Variable
var (
	CmdHandler *framework.CommandHandler
	Token      = "NzQyNTA1ODUxMDc1NDI4NDIz.XzHGdQ.3e0fddUjcsxT19C1FwtussGA-fk"

	version = "0.2.7"
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
	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	dg, err := bot.New("Bot " + Token)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la session discord,", err)
		return
	}

	dg.AddHandler(func(s *bot.Session, Event *bot.Event) {
		framework.Session = s

		if Event.Type == "READY" {
			s.UpdateStatus(0, viper.GetString("Motd"))
			log.Println("Le bot est dispo.\nAppuyez sur CTRL+C pour l'arrêter !")
			logger.InfoLogger.Println("Le bot est dispo. [Appuyez sur CTRL+C pour l'arrêter !]")

			modules.LogDiscord("[:tools:] [v:" + version + "] **Lyana** à correctement démarré accompagné des " + strconv.Itoa(framework.CountCommand+viper.GetInt("ModuleCount")) + " modules.")
			modules.UpdateOnlinePlayer(framework.Session)
			modules.VersionServerRcon()
			modules.VerifServerMCVersion()
			modules.VerifServerMCBuild()
		}
	})
	dg.AddHandler(commandHandler)
	dg.AddHandler(func(s *discordgo.Session, join *discordgo.GuildMemberAdd) {
		//join action
		s.GuildMemberRoleAdd(viper.GetString("GuildID"), join.User.ID, "742781882852179988")

		embed := modules.NewEmbed().
			SetTitle(join.User.String() + ", je te souhaite la bienvenue parmi nous.").
			SetColor(viper.GetInt("EmbedColor.Bienvenue")).
			SetDescription("Je t'invite à lire notre <#735271074735849564> ainsi que <#735271020575064165>, tu trouveras un maximum d'information pour commencer.\nSi tu à la moindre question, n'hésite pas.\n\nSur ce bon séjour parmi nous. Cordialement Lyana.").MessageEmbed

		s.ChannelMessageSendEmbed(viper.GetString("ChannelID.Trafic"), embed)
		modules.LogDiscord("[<:upvote:742854427454472202>] " + join.User.Username)
	})
	dg.AddHandler(func(s *discordgo.Session, leave *discordgo.GuildMemberRemove) {
		//leave action
		modules.LogDiscord("[<:downvote:742854427177648190>] " + leave.User.Username)
	})
	dg.AddHandler(func(s *discordgo.Session, reac *discordgo.MessageReactionAdd) {
		if reac.UserID == s.State.User.ID {
			return
			log.Println("Bot ajoute un emoji")
		}

		//Acceptation du réglement
		if reac.Emoji.Name == "✅" && reac.ChannelID == viper.GetString("ChannelID.Reglement") && reac.MessageID == viper.GetString("MessageID.Reglement") {
			s.GuildMemberRoleRemove(viper.GetString("GuildID"), reac.UserID, "742781882852179988")
			s.GuildMemberRoleAdd(viper.GetString("GuildID"), reac.UserID, "735281835080286291")
		}
		if reac.Emoji.Name != "✅" && reac.ChannelID == viper.GetString("ChannelID.Reglement") {
			s.MessageReactionRemove(reac.ChannelID, reac.MessageID, reac.Emoji.Name, reac.UserID)
		}
	})

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la connexion,", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	t := time.Tick(1 * time.Minute)
	d := time.Tick(1 * time.Hour)

out:
	for {
		select {
		case <-sc:
			modules.LogDiscord("[:tools:] **Lyana** s'est déconnecté de l'univers !")
			break out
		case <-t:
			// Actualise le nombre de joueurs en ligne toute les 1 minute
			modules.UpdateOnlinePlayer(framework.Session)
			modules.VersionServerRcon()
		case <-d:
			// Actualise le nombre de joueurs en ligne toute les 1 minute
			modules.VerifServerMCVersion()
			modules.VerifServerMCBuild()
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

	messageSplit := strings.Fields(content)
	name := strings.ToLower(messageSplit[0])
	command, found := CmdHandler.Get(name)
	if !found {
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de l'obtention du channel,", err)
		return
	}

	staff := 0
	if channel.Type != discordgo.ChannelTypeDM {
		staff = modules.VerifStaff(m.Member.Roles)
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de l'obtention de la guilde,", err)
		return
	}

	ctx := framework.NewContext(s, guild, channel, user, m, staff, CmdHandler, messageSplit)
	ctx.Args = messageSplit[1:]
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("test", command.Test, "???")

	if viper.GetBool("Dev.test") != true {
		CmdHandler.Register("online", command.OnlinePlayer, "Affiche les joueurs connecté")
		CmdHandler.Register("signal", command.Signalement, "Permet au joueurs whitelist sur le serveur de signaler un autre joueurs commétande une infraction")
		CmdHandler.Register("addplayer", command.AddPlayer, "???")
	}
}
