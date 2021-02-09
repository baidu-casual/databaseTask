package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

//Exec executes terminal commands
func Exec() {
	out, err := exec.Command("bash", "command.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
