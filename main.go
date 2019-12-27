package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error no command provided")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		var args []string
		if len(os.Args) < 3 {
			fmt.Println("Error not enough args")
			os.Exit(1)
		} else if len(os.Args) > 3 {
			args = os.Args[3:]
		}
		run(os.Args[2], args)
	}
}

func run(name string, args []string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	isolate(cmd, syscall.CLONE_NEWUSER)

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: unable to run", err)
		os.Exit(1)
	}
}

func isolate(cmd *exec.Cmd, flags uintptr) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: flags,
	}
}
