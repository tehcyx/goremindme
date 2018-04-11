package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/0xAX/notificator"
)

var notify *notificator.Notificator

// Options struct with all CLI args
type Options struct {
	period  *string
	elapse  *string
	message *string
	task    *string
}

var opts Options
var timeRegex = regexp.MustCompile(`^([0-9]{0,3})(s|m|h|d)$`)

func parseArgs(args []string) Options {
	opts.period = flag.String("p", "", "Period for the reminder to run, e.g. 10m, 5m, 1h, 1d (Required)")
	opts.elapse = flag.String("e", "", "Time between consecutive reminders, e.g. 2m, 5h, 1d")
	opts.message = flag.String("m", "", "Message to be displayed as reminder, e.g. 'Stay focused'")
	opts.task = flag.String("t", "", "Task name, e.g. 'Task notifier' (Required)")

	flag.Parse()

	if *opts.period == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *opts.task == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !timeRegex.MatchString(*opts.period) {
		fmt.Println("Error parsing period")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *opts.elapse != "" {
		if !timeRegex.MatchString(*opts.elapse) {
			fmt.Println("Error parsing elapse")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	return opts
}

func main() {
	opts := parseArgs(os.Args)

	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "goremindme",
	})

	res := timeRegex.FindStringSubmatch(*opts.period)
	timeVal, _ := strconv.Atoi(res[1])
	modifierVal := getModifier(res[2])

	notificationExpireTimerChan := time.NewTimer(time.Duration(timeVal) * modifierVal).C

	if *opts.elapse != "" {
		res = timeRegex.FindStringSubmatch(*opts.elapse)
		if len(res) > 0 {
			timeVal, _ = strconv.Atoi(res[1])
		} else {
			timeVal = 0
		}
		modifierVal = getModifier(res[2])
	}
	elapseTickerChan := time.NewTicker(time.Duration(timeVal) * modifierVal).C

	for {
		select {
		case <-notificationExpireTimerChan:
			notify.Push(*opts.task, *opts.message, "/home/user/icon.png", notificator.UR_CRITICAL)
			return
		case <-elapseTickerChan:
			notify.Push(*opts.task, *opts.message, "/home/user/icon.png", notificator.UR_CRITICAL)
		}
	}
}

func getModifier(val string) time.Duration {
	switch val {
	case "s":
		return time.Second
	case "m":
		return time.Minute
	case "h":
		return time.Hour
	case "d":
		return time.Hour * 24
	}
	return time.Second
}
