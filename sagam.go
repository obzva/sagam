package sagam

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/andybrewer/mack"
	"github.com/obzva/timer"
)

const (
	openingText string = ` ▗▄▄▖ ▗▄▖  ▗▄▄▖ ▗▄▖ ▗▖  ▗▖
▐▌   ▐▌ ▐▌▐▌   ▐▌ ▐▌▐▛▚▞▜▌
 ▝▀▚▖▐▛▀▜▌▐▌▝▜▌▐▛▀▜▌▐▌  ▐▌
▗▄▄▞▘▐▌ ▐▌▝▚▄▞▘▐▌ ▐▌▐▌  ▐▌

type 'pause' or 'p' to pause the timer
type 'start' or 's' to restart the timer
the commands are case insensitive

`

	statusSprint = "sprint"
	statusRest   = "rest"
)

func printOpening(w io.Writer) {
	fmt.Fprint(w, openingText)
}

func sprint(timer *timer.Timer) string {
	<-timer.C
	return statusRest
}

func rest(timer *timer.Timer) string {
	<-timer.C
	return statusSprint
}

func sayAndNotify(reps int, s string) error {
	for range reps {
		if err := mack.Say(s); err != nil {
			return err
		}
		if err := mack.Notify(s, "Sagam"); err != nil {
			return err
		}
	}
	return nil
}

func formatMinute(d time.Duration) string {
	res := fmt.Sprintf("%d minute", d/time.Minute)
	if d > 1*time.Minute {
		res += "s"
	}
	return res
}

func Run(r io.Reader, w io.Writer, sprintDuration, restDuration time.Duration, reps int) error {
	if err := sayAndNotify(reps, "Let's get started!"); err != nil {
		return err
	}
	printOpening(w)

	// initialize timer
	t := timer.NewTimer(1 * time.Nanosecond)

	status := statusSprint

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			command := strings.ToLower(scanner.Text())
			switch command {
			case "pause", "p":
				t.Pause()
				fmt.Fprintf(w, "[%s] timer paused: remaining %v\n", time.Now().Format("15:04:05"), t.Duration.Truncate(time.Second))
			case "start", "s":
				t.Start()
				fmt.Fprintf(w, "[%s] timer started\n", time.Now().Format("15:04:05"))
			default:
				fmt.Fprintf(w, "[%s] unknown command: %s\n", time.Now().Format("15:04:05"), command)
			}
		}
	}()

	sprintCount := 0

	for {
		switch status {
		case statusSprint:
			if err := sayAndNotify(reps, fmt.Sprintf("Sprint for %s", formatMinute(sprintDuration))); err != nil {
				return err
			}
			sprintCount++
			fmt.Fprintf(w, "[%s] SPRINT %d\n", time.Now().Format("15:04:05"), sprintCount)
			t.Restart(sprintDuration)
			next := sprint(t)
			status = next
		case statusRest:
			if err := sayAndNotify(reps, fmt.Sprintf("Rest for %s", formatMinute(restDuration))); err != nil {
				return err
			}
			fmt.Fprintf(w, "[%s] REST\n", time.Now().Format("15:04:05"))
			t.Restart(restDuration)
			next := rest(t)
			status = next
		}
	}
}
