// Package ui provides utilities for interacting with terminal interfaces.
package ui

import (
	"fmt"
	"time"
)

const (
	frames = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
)

// Spinner is a tool for messaging async information to the terminal.
type Spinner struct {
	stop chan struct{}
}

// Start allocates a channel for the new spinner and starts it spinning in the
// background.  It can be Stop()ed when complete.
func (s *Spinner) Start(format string, args ...interface{}) {
	s.stop = make(chan struct{}, 1)
	msg := fmt.Sprintf(format, args...)

	go func() {
		for {
			for _, r := range frames {
				select {
				case <-s.stop:
					return
				default:
					fmt.Printf("\r%s%s %c%s ", msg, "\x1b[92m", r, "\x1b[39m")
					time.Sleep(time.Millisecond * 100)
				}
			}
		}
	}()

}

// Stop stops the spinner.
func (s *Spinner) Stop() {
	s.stop <- struct{}{}
}

// Pad helper.
func Pad() func() {
	println()
	return func() {
		println()
	}
}
