package utils

import (
	. "arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"fmt"
	"github.com/briandowns/spinner"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func GetProcessOwner() string {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		StepBreak(err)
	}
	user := ""
	if runtime.GOOS == "windows" {
		user = strings.TrimRight(string(stdout), "\r\n")
	} else {
		user = strings.TrimRight(string(stdout), "\n")
	}

	return user
}
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func RenderCommand(cmd string, args []structs.ExecArgs) string {

	for i := 0; i < len(args); i++ {
		cmd = strings.ReplaceAll(cmd, "{"+args[i].Key+"}", args[i].Value)
	}
	return cmd
}
func RunCustomBashCommand(path string, pass_on_error bool, command string, s *spinner.Spinner) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()

	s.Stop()
	fmt.Printf("\r")
	if err != nil {
		if !pass_on_error {
			StepError("pass on error is false")
			Log.Printf("%s => %s", fmt.Sprint(err), output)
			StepBreak(err)
		} else {
			StepError("pass on error is on, passing")
		}
	} else {
		StepPass()
	}
}
func RunBashCommand(execstep structs.ExecStruct) {
	RunCustomBashCommand(execstep.Path, execstep.PassOnError, RenderCommand(execstep.Command, execstep.Args), nil)
}
