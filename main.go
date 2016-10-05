package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	imageSnap()
}

func imageSnap() {
	binary, lookErr := exec.LookPath("imagesnap")
	if lookErr != nil {
		panic(lookErr)
	}

	args := os.Args
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}
}
