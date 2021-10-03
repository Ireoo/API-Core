package debug

import (
	"log"

	"github.com/gookit/color"
)

var debug = true

func SetDebug(d bool) {
	debug = d
}

func Trace(message interface{}) {
	if debug {
		log.Println(color.FgGreen.Sprintf(color.Bold.Sprintf(" ✔ ")) + color.FgGreen.Render(message))
	}
}

func Error(message interface{}) {
	if debug {
		log.Println(color.FgRed.Sprintf(color.Bold.Sprintf(" ✘ ")) + color.FgRed.Render(message))
	}
}

func Info(message interface{}) {
	if debug {
		log.Println(color.Bold.Sprintf(" ➽ ") + message.(string))
	}
}

func Finish(message interface{}) {
	if debug {
		log.Println(color.FgCyan.Sprintf(color.Bold.Sprintf(" 💯 ")) + color.FgCyan.Render(message))
	}
}
