package framework

import (
	bot "github.com/bwmarrin/discordgo"
	mcrcon "gitlab.com/lyana/library/package/mc_rcon"
)

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

	OnlinePlayer []int
	ListPlayer   []string
	OnlineServer []string
	Sanction     [][]string
	ConnectMC    []*mcrcon.MCConn

	//Gestion de la création d'évent
	EventConstruction            bool
	EventConstructionChannelID   string
	EventConstructionMessageID   string
	EventConstructionMessageAide string
	EventContructionIndex        int
)
