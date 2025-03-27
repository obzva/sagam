package main

import (
	"log"
	"os"
	"time"

	"github.com/obzva/sagam"
)

func main() {
	// sFlag := flag.Int("s", 25, "sprint duration in minutes, default is 25")
	// rFlag := flag.Int("r", 5, "rest duration in minutes, default is 5")
	// flag.Parse()

	// sprintDuration := time.Duration(*sFlag) * time.Minute
	// restDuration := time.Duration(*rFlag) * time.Minute

	sprintDuration := 10 * time.Second
	restDuration := 5 * time.Second
	if err := sagam.Run(os.Stdin, os.Stdout, sprintDuration, restDuration); err != nil {
		log.Fatal(err)
	}

}
