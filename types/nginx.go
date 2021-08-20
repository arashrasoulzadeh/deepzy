package types

import (
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/utils"

	"log"
	"strings"

	"github.com/fatih/color"
)

func runNginx(execstruct structs.ExecStruct) {
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("\t%s", green(strings.ToUpper(execstruct.Name)))

	if utils.StringInSlice(execstruct.Command, []string{"restart", "reload", "stop", "start"}) {
		logger.StepVerboseExec(execstruct)
		utils.RunCustomBashCommand("/", execstruct.PassOnError, "systemctl "+execstruct.Command+" nginx")
	} else {
		logger.StepVerboseError(execstruct)
	}

}
