package framework

import (
	bot "github.com/bwmarrin/discordgo"
	"gitlab.com/lyana/logger"
)

type Context struct {
	Discord      *bot.Session
	Guild        *bot.Guild
	TextChannel  *bot.Channel
	User         *bot.User
	Message      *bot.MessageCreate
	Staff        int
	Args         []string
	MessageSplit []string

	// dependency injection?
	CmdHandler *CommandHandler
}

func NewContext(discord *bot.Session, guild *bot.Guild, textChannel *bot.Channel,
	user *bot.User, message *bot.MessageCreate, staff int, cmdHandler *CommandHandler, messageSplit []string) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.Staff = staff
	ctx.CmdHandler = cmdHandler
	ctx.MessageSplit = messageSplit
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
