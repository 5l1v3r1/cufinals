package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/unixpickle/cufinals"
)

func main() {
	var schedule []cufinals.Entry
	if len(os.Args) == 3 {
		contents, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not read file:", err)
			os.Exit(1)
		}
		if err := json.Unmarshal(contents, &schedule); err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse file:", err)
			os.Exit(1)
		}
	} else if len(os.Args) == 2 {
		var err error
		schedule, err = cufinals.FetchSchedule(cufinals.DefaultScheduleURL)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not fetch schedule:", err)
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Usage: roomstats <short room name> [schedule.json]")
		os.Exit(1)
	}

	roomName := os.Args[1]
	relevant := EntryList{}
	for _, entry := range schedule {
		if entry.Room.ShortName == roomName {
			relevant = append(relevant, entry)
		}
	}
	sort.Sort(relevant)

	for _, entry := range relevant {
		fmt.Println(entry)
	}
}
