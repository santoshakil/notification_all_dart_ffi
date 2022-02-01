package main

import (
	"C"

	"github.com/gen2brain/beeep"
)

func main() {
	// Beep()
	Notify()
	// Alert()
}

// export Notify
func Beep() {
	beeep.Beep(1.0, 1)
}

//export Notify
func Notify() {
	beeep.Notify("Title", "MessageBody", "assets/information.png")
}

//export Alert
func Alert() {
	beeep.Alert("Title", "MessageBody", "assets/warning.png")
}
