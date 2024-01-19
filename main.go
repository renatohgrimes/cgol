package main

import (
	"cgol/cgol"
	"time"
)

func main() {
	cm := cgol.NewManager(15, 50)
	for {
		cm.ComputeNextGeneration()
		cgol.Draw(cm)
		time.Sleep(500 * time.Millisecond)
	}
}
