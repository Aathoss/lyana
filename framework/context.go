package framework

import (
	"github.com/Aathoss/lyana/logger"
	bot "github.com/bwmarrin/discordgo"
)

type Context struct {
	Discord     *bot.Session
	Guild       *bot.Guild
	TextChannel *bot.Channel
	User        *bot.User
	Message     *bot.MessageCreate
	Commande    string
	Args        []string
	NiveauStaff int

	// dependency injection?
	CmdHandler *CommandHandler
}

func NewContext(discord *bot.Session, guild *bot.Guild, textChannel *bot.Channel,
	user *bot.User, message *bot.MessageCreate, cmdHandler *CommandHandler, cmdName string, staff int) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.CmdHandler = cmdHandler
	ctx.Commande = cmdName
	ctx.NiveauStaff = staff
	return ctx
}

func (ctx Context) Reply(content string) *bot.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, content)
	if err != nil {
		logger.ErrorLogger.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}
