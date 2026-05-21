package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/patlehmann1/space_coast_fl_launches/api"
	"github.com/patlehmann1/space_coast_fl_launches/display"
)

const version = "1.0.0"

func main() {
	count := flag.Int("count", 5, "number of launches to show (1–10)")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("fl-launches v%s\n", version)
		return
	}

	if *count < 1 {
		*count = 1
	}
	if *count > 10 {
		*count = 10
	}

	launches, err := api.FetchUpcoming(*count)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	display.Render(launches)
}
