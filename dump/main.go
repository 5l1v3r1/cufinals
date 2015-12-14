package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/unixpickle/cufinals"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: dump <output.json>")
		os.Exit(1)
	}

	results, err := cufinals.FetchSchedule(cufinals.DefaultScheduleURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error fetching schedule:", err)
		os.Exit(1)
	}

	data, _ := json.Marshal(results)
	if err := ioutil.WriteFile(os.Args[1], data, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write file:", err)
		os.Exit(1)
	}
}
