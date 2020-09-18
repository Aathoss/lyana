package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"gitlab.com/lyana/framework"
	"gitlab.com/lyana/logger"
	"gitlab.com/lyana/mysql"
)

func GuildMemberAdd(s *discordgo.Session, join *discordgo.GuildMemberAdd) {
	//ajoute d'un grade lors de l'arrivé
	s.GuildMemberRoleAdd(viper.GetString("GuildID"), join.User.ID, "742781882852179988")

	//affichage d'un message de bienvenue
	embed := framework.NewEmbed().
		SetTitle("<:EnderPearl:753547886637350922> Whoo une téléportation viens d'avoir lieu !").
		SetColor(viper.GetInt("EmbedColor.Bienvenue")).
		SetDescription("Bienvenue parmi nous " + join.User.String() + "\nJe t'invite à lire notre #📚règlement ainsi que #📖présentation\nSur ce bon séjour parmi nous. Cordialement").MessageEmbed

	_, err := s.ChannelMessageSendEmbed(viper.GetString("ChannelID.Trafic"), embed)
	framework.LogsChannel("[<:upvote:742854427454472202>] " + join.User.Username)

	//ajoute en bdd la personne concernant le règlement
	t1, err := join.JoinedAt.Parse()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	mysql.AddRule(join.User.ID, t1.Unix())
}
