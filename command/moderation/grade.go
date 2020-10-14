package moderation

import "gitlab.com/lyana/framework"

func Grade(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.Message.ChannelID, ctx.Message.ID)

	embed := framework.NewEmbed().
		SetTitle("Conversion des grade").
		SetColor(0x33E2F0).
		AddField("Staff", "_Fondateur -->_ **Chef d'acquisition**"+
			"\n_Administrateur -->_ **Chef de Communication**"+
			"\n_Responsable -->_ **Ingénieur en Chef**"+
			"\n_Modérateur -->_ **Antivortex**"+
			"\n_Ambassadrice -->_ **Gestion d'amarrage**", false).
		AddField("Bot", "_Bot Lyana -->_ **Satellite**"+
			"\n_Bot Music -->_ **Nanosatellite**", false).
		AddField("Membre", "_VIP -->_ **Martien**"+
			"\n_Joueurs -->_ **Astronaute**"+
			"\n_Transfert -->_ **Débris Spatial**"+
			"\n_Inactif -->_ **Biostase**", false).MessageEmbed

	ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
}
