package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	GuardBegins = iota + 1
	GuardAsleep
	GuardAwake
	GuardStops
)

const timeFormat = "2006-01-02 15:04"

type Log struct {
	time  time.Time
	guard int
	event int
}

// Type GuardMinute is used as the key to a map which counts the times a guard is asleep at a
// certain minute.
type GuardMinute struct {
	guard, sleepingMinute int
}

// Type sortedLogs is a slice type with associated functions to have it be called by sort.Sort().
type sortedLogs []*Log

func (l sortedLogs) Len() int           { return len(l) }
func (l sortedLogs) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l sortedLogs) Less(i, j int) bool { return l[i].time.Before(l[j].time) }

// Function parseLogEntries takes a slice of strings and parses them to Log types
func parseLogEntries(lines []string) []*Log {
	var logs []*Log
	pattern := regexp.MustCompile(`^\[(.+)\] (Guard #(\d+) )?(.+)$`)
	for _, log := range lines {
		matches := pattern.FindAllStringSubmatch(log, 1)
		t, _ := time.Parse(timeFormat, matches[0][1])
		guard, err := strconv.Atoi(matches[0][3])
		if err != nil {
			guard = 0
		}

		// Parse the event in the log entry
		var event int
		switch strings.Fields(matches[0][4])[0] {
		case "begins":
			event = GuardBegins
		case "falls":
			event = GuardAsleep
		case "wakes":
			event = GuardAwake
		}
		logs = append(logs, &Log{t, guard, event})
	}
	sort.Sort(sortedLogs(logs))

	// Populate the Log entries without a guard ID with the guard ID of the log entry before, until
	// a log entry with a guard ID is found.
	var g int
	for _, log := range logs {
		if log.guard != 0 {
			g = log.guard
			continue
		}
		log.guard = g
	}
	return logs
}

func StarOne(input []string) string {
	sleepingGuards := make(map[GuardMinute]int)
	logs := parseLogEntries(input)
	for i, log := range logs {
		if log.event != GuardAsleep {
			continue
		}
		for m := log.time.Minute(); m < logs[i+1].time.Minute(); m++ {
			sleepingGuards[GuardMinute{log.guard, m}] += 1
		}
	}

	// Add up all sleeping minutes of all the guards
	guardSleepingMinutes := make(map[int]int)
	for g := range sleepingGuards {
		guardSleepingMinutes[g.guard] += sleepingGuards[g]
	}

	// Find sleepiest Guard
	var sleepiestGuard, minutes int
	for sleepiestGuard, minutes = range guardSleepingMinutes {
		break
	}
	for g, m := range guardSleepingMinutes {
		if m > minutes {
			sleepiestGuard = g
		}
	}

	// Find minute where sleepiestGuard is most asleep
	var minuteVal, minute int
	for gm, m := range sleepingGuards {
		if gm.guard != sleepiestGuard {
			continue
		}
		if m > minuteVal {
			minuteVal = m
			minute = gm.sleepingMinute
		}
	}

	return strconv.Itoa(sleepiestGuard * minute)
}

func StarTwo(input []string) string {
	sleepingGuards := make(map[GuardMinute]int)
	logs := parseLogEntries(input)
	for i, log := range logs {
		if log.event != 2 {
			continue
		}
		for m := log.time.Minute(); m < logs[i+1].time.Minute(); m++ {
			sleepingGuards[GuardMinute{log.guard, m}] += 1
		}
	}

	// Find guard and minute where a guard is most asleep
	var guard, minuteVal, minute int
	for gm, m := range sleepingGuards {
		if m > minuteVal {
			minuteVal = m
			guard = gm.guard
			minute = gm.sleepingMinute
		}
	}

	return strconv.Itoa(guard * minute)
}

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println("1:", StarOne(input))
	fmt.Println("2:", StarTwo(input))
}
