package main

import (
	"time"

	"github.com/spencercdixon/go/ui"
)

func main() {
	s := &ui.Spinner{}
	s.Start("Doing something important")
	time.Sleep(5 * time.Second)
	s.Stop()
}
