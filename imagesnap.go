package isg

import (
	"os/exec"
	"syscall"
)

// Snap takes a picture via the imagesnap binary
func Snap(args, env []string) {
	binary, lookErr := exec.LookPath("imagesnap")

	if lookErr != nil {
		panic(lookErr)
	}

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}

}
