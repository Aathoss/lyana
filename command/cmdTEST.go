package command

import (
	"github.com/spf13/viper"
	"gitlab.com/unispace/framework"
	"gitlab.com/unispace/modules"
)

func Test(ctx framework.Context) {
	ctx.Discord.ChannelMessageDelete(ctx.TextChannel.ID, ctx.Message.ID)

	if ctx.Staff >= 1 {
		embed := modules.NewEmbed().
			SetTitle(ctx.User.String() + ", je te souhaite la bienvenue parmi nous.").
			SetColor(viper.GetInt("EmbedColor.EmbedColor")).
			SetDescription("Je t'invite à lire notre <#735271074735849564> ainsi que <#735271020575064165>, tu trouveras un maximum d'information pour commencer.\nSi tu à la moindre question, n'hésite pas.\n\nSur ce bon séjour parmi nous. Cordialement Lyana.").MessageEmbed

		ctx.Discord.ChannelMessageSendEmbed(ctx.Message.ChannelID, embed)
	}
}
