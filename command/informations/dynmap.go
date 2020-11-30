package informations

import "gitlab.com/lyana/framework"

//DynmapDropURL retourne un message avec les informations de la dynmap
func DynmapDropURL(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	embed := framework.NewEmbed().
		SetTitle("Liens de la dynmap").
		SetColor(0x33E2F0).
		SetDescription("https://map.unispace.fr/" +
			"\n\nSur minecraft vous avez accès à la commande `/marker` afin ajouter des marqueurs :" +
			"\n`/marker home` Permet d'ajouter un marquer exemple : :homes: Home Aathoss" +
			"\n`/marker projets <nom>` Le nom du projet doit être attaché exemple : :flag_white: **Portail-du-Nether**" +
			"\n`/marker shop <nom>` Le nom du shop doit être attaché exemple : **Mr-Gourmand**").MessageEmbed
	//SetDescription("Liens deeeee láàa\n**[Error]**\n‘Please do not contact the deads anymore’").MessageEmbed

	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}
