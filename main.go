package main

import (
	"flag"
	"fmt"
	"log"
	"weigh/sizecalc"
)

func main() {
	// Define flags
	unit := flag.String("unit", "GB", "Size unit (KB, MB, GB)")
	verbose := flag.Bool("verbose", false, "Enable verbose error messages")

	// Parse CLI arguments
	flag.Parse()
	files := flag.Args()

	// Check if files were provided
	if len(files) == 0 {
		log.Fatal("Error: Please provide at least one file or folder.")
	}

	// Iterate over each file/folder and calculate the size
	for _, file := range files {
		size, err := sizecalc.CalculateSize(file)
		if err != nil {
			if *verbose {
				log.Printf("Error processing '%s': %v\n", file, err)
			} else {
				log.Printf("Error: '%s' not found or inaccessible.\n", file)
			}
			continue
		}
		// Convert size based on the unit flag
		humanReadableSize := sizecalc.FormatSize(size, *unit)
		fmt.Printf("%s: %s\n", file, humanReadableSize)
	}
}
