package main

import (
	"os"

	"github.com/anks333/parking_lot/modes"
)

func main() {
	if len(os.Args) > 1 {
		modes.NewFileMode(os.Args[1]).Start()
		return
	}
	modes.NewInterActiveMode().Start()
}
