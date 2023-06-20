package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if nil != err {
		fmt.Fprintf(os.Stderr, "could not initialize SDL: %v", err)
		os.Exit(2)
	}

	defer sdl.Quit()
}
