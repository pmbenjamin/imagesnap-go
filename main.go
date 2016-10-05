package main

import (
	"github.com/pmbenjamin/imagesnap-go/isg"
	"os"
)

func main() {
	args := os.Args
	env := os.Environ()
	isg.Snap(args, env)
}
