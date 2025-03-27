# sagam

```
 ▗▄▄▖ ▗▄▖  ▗▄▄▖ ▗▄▖ ▗▖  ▗▖
▐▌   ▐▌ ▐▌▐▌   ▐▌ ▐▌▐▛▚▞▜▌
 ▝▀▚▖▐▛▀▜▌▐▌▝▜▌▐▛▀▜▌▐▌  ▐▌
▗▄▄▞▘▐▌ ▐▌▝▚▄▞▘▐▌ ▐▌▐▌  ▐▌
```

**Sagam** (사감, meaning _teacher who manages and coaches students in school dormitory_ in Korean) is a CLI timer application for [pomodoro technique](https://en.wikipedia.org/wiki/Pomodoro_Technique).

**Sagam** is designed for OS X and uses [`mack`](https://github.com/andybrewer/mack) package to call applescript actions.

## Why?

Recently, I feel like I am losing my ability to keep focus on something.

Rather than struggling with a foggy brain for a long long time, I thought it might be better having small breaks regularly.

It, however, is such a pain setting alarms or timers manually whenever I start focus or rest.

Therefore I built this.

## Features

You set how long you will focus and how long you will rest.

**Sagam** will notify you when to start and when to rest again and again until you quit the program.

## Install

To install the `sagam` command, run

```bash
$ go install github.com/obzva/sagam/cmd/sagam@latest
```

and put the resulting binary in one of your `PATH` directories if `$GOPATH/bin` isn't already in your PATH.

## Usage

### Start

By running

```bash
$ sagam -s [SPRINT_DURATION] -r [REST_DURATION]
```

sagam will set timer for `SPRINT_DURATION` minutes of sprints and `REST_DURATION` of rests again and again.

Both `SPRINT_DURATION` and `REST_DURATION` arguments should be set in minutes.

If you omit flags, default `SPRINT_DURATION` and `REST_DURATION` will be set `25` and `5`.

### Pause and Restart

If you want to pause the timer, simply type `pause` or `p`.

When you're ready, type `start` or `s` to restart the timer.

These commands are case insensitive.

## Reference

[`mack`](https://github.com/andybrewer/mack): Notifying is the most important feature to me, and it is easily done with this package.

[`timer`](https://github.com/ivahaev/timer): To add pause/restart feature, I used this package. I used [my fork version](https://github.com/obzva/timer) because I needed to add [`Restart` method](https://pkg.go.dev/github.com/obzva/timer#Timer.Restart).

## Contribute

It is designed and built for my personal use so it may have a plenty of room to improve.

If you want to add feature or fix something, feel free to contribute or fork.
