package framework

import bot "github.com/bwmarrin/discordgo"

var (
	CountMsg       int
	OnlineActulise int
	Session        *bot.Session

	//Variable framework
	CountCommand int

	//Minecraft variable
	VersionMC  string
	BuildMC    string
	Connection bool

	OnlinePlayer    int
	MaxOnlinePlayer int
	ListPlayer      []string
	Sanction        [][]string

	//Gestion de la création d'évent
	EventConstruction            bool
	EventConstructionChannelID   string
	EventConstructionMessageID   string
	EventConstructionMessageAide string
	EventContructionIndex        int
)
