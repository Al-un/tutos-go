package beginner

import (
	"fmt"
	"os/exec"
	"runtime"
)

func executeSystemCommand(command string, args ...string) {
	out, err := exec.Command(command, args...).Output()

	if err != nil {
		fmt.Printf("Error for command %v %v :\n", command, args)
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Command %v %v successfully executed\n", command, args)
	output := string(out[:])
	fmt.Println(output)
}

// RunSystemCommands https://tutorialedge.net/golang/executing-system-commands-with-golang/
func RunSystemCommands() {
	if runtime.GOOS == "window" {
		fmt.Println("Can't execute this on a Windows machine")
	} else {
		executeSystemCommand("ls", "-ltr")
		executeSystemCommand("pwd")
		executeSystemCommand("pouet")
	}
}
