package safely

import "log"

// GoDoer is just a function you want to execute inside a go routine.  Closure
// can be used but be careful of concurrency concerns.
type GoDoer func()

// Go spawns the goroutine with added recovery handling
func Go(todo Godoer) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic in safely.Go: %s", err)
			}
		}()
		todo()
	}()
}
