package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Aathoss/lyana/command"
	"github.com/Aathoss/lyana/command/event"
	"github.com/Aathoss/lyana/command/informations"
	"github.com/Aathoss/lyana/command/moderation"
	"github.com/Aathoss/lyana/command/stats"
	"github.com/Aathoss/lyana/command/vocaltemporaire"
	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/modules"
	"github.com/bwmarrin/discordgo"
	bot "github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// Variable
var (
	CmdHandler *framework.CommandHandler
)

func main() {
	//framework.LoadConfiguration()

	CmdHandler = framework.NewCommandHandler()
	registerCommands()

	dg, err := bot.New("Bot " + viper.GetString("ID"))
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la session discord,", err)
		return
	}

	dg.AddHandler(modules.Ready)
	dg.AddHandler(modules.Stats)
	dg.AddHandler(modules.VocalTemporaire)
	dg.AddHandler(modules.GuildMemberAdd)
	dg.AddHandler(modules.GuildMemberLeave)
	dg.AddHandler(modules.ReactionAdd)
	dg.AddHandler(modules.ReactionRemove)
	dg.AddHandler(modules.LevelingMessages)
	dg.AddHandler(commandHandler)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
	err = dg.Open()
	if err != nil {
		logger.ErrorLogger.Println("Erreur lors de la connexion,", err)
		return
	}

	go modules.MinecraftCheckShop(10)
	go modules.VerifCandid(10)
	go modules.UpdateOnlinePlayer(5)
	go event.UpdateEvent(5)
	go func() {
		for {
			consoleReader := bufio.NewReader(os.Stdin)

			input, _ := consoleReader.ReadString('\n')

			input = strings.ToLower(input)

			if strings.HasPrefix(input, "bye") {
				framework.DBLyana.Close()
				if viper.GetBool("Mysql.Minecraft.online") == true {
					framework.DBMinecraft.Close()
				}

				framework.LogsChannel("[:tools:] **Lyana** s'est déconnecté de l'univers !")

				fmt.Println("\nUptime : " + framework.Calculetime(stats.StartTime.Unix(), 0) +
					"\nMessage total : " + strconv.Itoa(framework.CountMsg) +
					"\nRoutine : " + strconv.Itoa(runtime.NumGoroutine()) +
					"\n\nAllez bonne route ++ \n")
				os.Exit(100)
			}
		}
	}()

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
	framework.Session = s
	user := m.Author

	if user.ID == s.State.User.ID || user.Bot {
		return
	}

	if viper.GetBool("Dev.PrintMessage") == true {
		log.Println(m.Content)
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
	CmdHandler.Register("test", []string{}, 1, moderation.Test, "???")

	//Commande Modération
	CmdHandler.Register("stats", []string{}, 1, stats.Statistique, "Returne les statistique du bot")
	CmdHandler.Register("purge", []string{}, 1, moderation.Purges, "La commande permet d'effectuer un netoyage d'un channel limite à 2.500 Message")
	CmdHandler.Register("grade", []string{}, 0, moderation.Grade, "Affiche la conversion des grade")
	CmdHandler.Register("help", []string{}, 0, moderation.HelpCommand, "Affiche la liste des commande")
	CmdHandler.Register("lyana", []string{}, 1, moderation.PubliMessage, "Publi un message avec Lyana")
	CmdHandler.Register("globalmp", []string{}, 1, moderation.MessageGlobalMp, "Envoie un message global à tout les membres")

	//Commande Liée à minecraft
	CmdHandler.Register("fiche", []string{"profils", "profil"}, 0, command.InfoPlayer, "Permet de voir votre fiche utilisateur/player")
	CmdHandler.Register("online", []string{}, 0, command.OnlinePlayer, "Affiche les joueurs connecté")
	CmdHandler.Register("addplayer", []string{}, 1, command.AddPlayer, "???")

	//Commande d'informations
	CmdHandler.Register("map", []string{}, 0, informations.DynmapDropURL, "Affiche le liens de la dynmap")
	CmdHandler.Register("globalstats", []string{}, 1, informations.StatsUnispaceV1, "???")

	//Commande vocal VocalTemporaire
	CmdHandler.Register("vtitre", []string{}, 0, vocaltemporaire.VocalTempEditTitre, "Modifie le titre de votre channel vocal temporaire")
	CmdHandler.Register("vlimite", []string{}, 0, vocaltemporaire.VocalTempEditLimit, "Modifie le nombre de memebre dans votre channel temporaire")

	//Commande event
	CmdHandler.Register("event cree", []string{}, 1, event.ConstructionEvent, "Démarre la création d'un évent")
	CmdHandler.Register("event titre", []string{}, 1, event.EditTitre, "Modifie le titre durant la création")
	CmdHandler.Register("event gps", []string{}, 1, event.EditEmplacement, "Modifie la localisation durant la création")
	CmdHandler.Register("event desc", []string{}, 1, event.EditDescription, "Modifie la description lors de la création")
	CmdHandler.Register("event date", []string{}, 1, event.EditDate, "Modifie la date lors de la création")
	CmdHandler.Register("event recompense", []string{}, 1, event.EditRecompense, "Modifie la liste de récompense lors de la création")
	CmdHandler.Register("event auteur", []string{}, 1, event.EditAuteur, "Modifie l'auteur durant la création")
	CmdHandler.Register("event publi", []string{}, 1, event.PubliEvent, "Publi la création de l'évent pour tout le monde")
	CmdHandler.Register("event termine", []string{}, 1, event.EventTermine, "Publi la création de l'évent pour tout le monde")
}
