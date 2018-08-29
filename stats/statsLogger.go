package stats

import (
	"log"
	"os"
)

// statsLogger is used for recording/saving/printing alerts
var statsLogger *log.Logger

// SetupLogger initializes stats logger to print to console
func SetupLogger() {

	// Create logger
	statsLogger = log.New(os.Stdout, "", 0)
}
