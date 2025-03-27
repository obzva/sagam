package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/obzva/sagam"
)

func main() {
	sFlag := flag.Int("s", 25, "sprint duration in minutes, default is 25")
	rFlag := flag.Int("r", 5, "rest duration in minutes, default is 5")
	repFlag := flag.Int("rep", 1, "number of repetitions of alarm, default is 3")
	flag.Parse()

	sprintDuration := time.Duration(*sFlag) * time.Minute
	restDuration := time.Duration(*rFlag) * time.Minute

	if err := sagam.Run(os.Stdin, os.Stdout, sprintDuration, restDuration, *repFlag); err != nil {
		log.Fatal(err)
	}
}
