package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdout = os.Stdout // Redirect stdout to the terminal
	cmd.Stderr = os.Stderr // Redirect stderr to the terminal
	cmd.Stdin = os.Stdin   // Redirect stdin from the terminal
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid:  true, // Create a new session
		Setpgid: true, // Create a new process group
	}
	cmd.Run()

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
