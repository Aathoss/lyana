package framework

import (
	"github.com/spf13/viper"
	"gitlab.com/lyana/logger"
)

var Cmdliste []string

type (
	Command func(Context)

	CommandStruct struct {
		command Command
		grade   int
		module  string
		help    string
	}

	CmdMap map[string]CommandStruct

	CommandHandler struct {
		cmds CmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

func (handler CommandHandler) GetCmds() CmdMap {
	return handler.cmds
}

func (handler CommandHandler) GetAllCmd(name string) (help string) {
	cmd, _ := handler.cmds[name]
	return cmd.help
}

func (handler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := handler.cmds[name]
	return &cmd.command, found
}

func (handler CommandHandler) Register(name string, command Command, helpmsg string) {
	logger.InfoLogger.Println("Chargement de la commande : " + viper.GetString("PrefixMsg") + name)
	cmdstruct := handler.cmds[name]
	cmdstruct.command = command
	cmdstruct.help = helpmsg
	handler.cmds[name] = cmdstruct

	CountCommand = CountCommand + 1
	Cmdliste = append(Cmdliste, name)
}

func (command CommandStruct) GetHelp() string {
	return command.help
}
