package main

import (
	"fmt"
	"os"

	"github.com/rogwilco/quartermaster/internal/app/ui"
)

func main() {
	if err := ui.Run(); err != nil {
		fmt.Printf("Error running application: %v\n", err)
		os.Exit(1)
	}
}
