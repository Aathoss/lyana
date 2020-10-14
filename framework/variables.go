package framework

import bot "github.com/bwmarrin/discordgo"

var (
	CountMsg       int
	SQlRequest     int
	OnlineActulise int
	Session        *bot.Session
)

//Variable framework
var (
	CountCommand int
)

//Minecraft variable
var (
	VersionMC  string
	BuildMC    string
	Connection bool

	OnlinePlayer    int
	MaxOnlinePlayer int
	ListPlayer      []string
	Sanction        [][]string
)
