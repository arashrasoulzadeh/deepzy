package functions

import (
	. "arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/utils"

	"arashrasoulzadeh/deepzy/structs"
	"arashrasoulzadeh/deepzy/types"

	"log"
	"strings"

	"github.com/fatih/color"
)

func Execute(config structs.Config) {
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for i := 0; i < len(config); i++ {

		log.Printf("STEP %s : \n", blue(strings.ToUpper(config[i].Name)))
		if config[i].Become {
			log.Printf("\t%s\t", green(strings.ToUpper("ROOT CHECK")))
			if strings.Compare(utils.GetProcessOwner(), "root") == 0 {
				StepPass()
			} else {
				StepError()
				StepBreak("YOU ARE NOT SUDOER")
			}
			log.Println()
		}
		for s := 0; s < len(config[i].Exec); s++ {
			Log.Printf("%s: %s ", config[i].Exec[s].Type, config[i].Exec[s].Command)
			types.Run(config[i].Exec[s])
		}
	}
	log.Println("")
	log.Println("DONE!")
}
