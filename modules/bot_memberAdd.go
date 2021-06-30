package modules

import (
	"github.com/Aathoss/lyana/framework"
	"github.com/Aathoss/lyana/logger"
	"github.com/Aathoss/lyana/mysql"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func GuildMemberAdd(s *discordgo.Session, join *discordgo.GuildMemberAdd) {
	//ajoute d'un grade lors de l'arrivé
	s.GuildMemberRoleAdd(viper.GetString("GuildID"), join.User.ID, "742781882852179988")

	
	/* invites, err := s.GuildInvites(viper.GetString("GuildID"))
	if err != nil {
		logger.ErrorLogger.Println(err)
	}


	for _, invite := range invites {
		println(invite.TargetUser.ID)
		println(invite.Uses)

		 if invite.TargetUser.ID == join.User.ID {
			framework.LogsChannel("[<:upvote:742854427454472202>] :mag_right:  " + join.User.Mention() + " Invité par " + invite.Inviter.Mention())
		}
	}

	if "693473406108041218" == join.User.ID {
		return
	} */

	//affichage d'un message de bienvenue
	embed := framework.NewEmbed().
		SetTitle("<:EnderPearl:753547886637350922> Whoo une téléportation viens d'avoir lieu !").
		SetColor(viper.GetInt("EmbedColor.Bienvenue")).
		SetDescription("Bienvenue parmi nous " + join.User.String() + "\nJe t'invite à lire notre #📚règlement ainsi que #📖présentation\nSur ce bon séjour parmi nous. Cordialement").MessageEmbed

	_, err := s.ChannelMessageSendEmbed(viper.GetString("ChannelID.Trafic"), embed)
	framework.LogsChannel("[<:upvote:742854427454472202>] " + join.User.Mention())

	//ajoute en bdd la personne concernant le règlement
	t1, err := join.JoinedAt.Parse()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}
	mysql.AddRule(join.User.ID, t1.Unix())
}
