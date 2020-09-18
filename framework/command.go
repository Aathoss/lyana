package framework

import "regexp"

var (
	Cmdliste      []string
	commandNumero int
)

type (
	Command func(Context)

	CommandStruct struct {
		command  Command
		grade    int
		help     string
		cmdalias int
		cmdnum   int
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

func (handler CommandHandler) CheckCmd(content string) (name string) {
	for _, cmd := range Cmdliste {
		matched, _ := regexp.MatchString(cmd, content)
		/* fmt.Print(err)
		fmt.Print(" | ")
		fmt.Print(matched)
		fmt.Print(" | ")
		fmt.Println(cmd) */
		if matched == true {
			name = cmd
		}
	}
	return name
}

func (handler CommandHandler) Get(name string, gradelvl int) (*Command, bool, bool) {
	cmd, found := handler.cmds[name]
	if gradelvl >= cmd.grade {
		return &cmd.command, found, true
	}
	return &cmd.command, found, false
}

func (handler CommandHandler) Register(name string, alias []string, gradelvl int, command Command, helpmsg string) {
	//logger.DebugLogger.Println("Chargement de la commande : " + viper.GetString("PrefixMsg") + name)
	niveaucmd := 0
	alias = append(alias, name)
	//fmt.Println(len(name))

	for _, aliasCmd := range alias {
		cmdstruct := handler.cmds[aliasCmd]
		cmdstruct.command = command
		cmdstruct.grade = gradelvl
		cmdstruct.help = helpmsg
		cmdstruct.cmdalias = niveaucmd
		cmdstruct.cmdnum = commandNumero
		handler.cmds[aliasCmd] = cmdstruct

		niveaucmd++
		Cmdliste = append(Cmdliste, aliasCmd)

		//fmt.Print(aliasCmd + " | ")
		//fmt.Println(cmdstruct)
	}
	commandNumero++
}

func (command CommandStruct) GetHelp() string {
	return command.help
}
