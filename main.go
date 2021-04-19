package main

import (
	"github.com/AndreiBarbuOz/cormen/cmd"
	"os"
)

func main() {
	cmd.Execute(os.Args[1:])
}
