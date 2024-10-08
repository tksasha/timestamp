package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	//nolint:forbidigo
	fmt.Printf(
		"%04d%02d%02d%02d%02d%02d\n",
		today.Year(),
		today.Month(),
		today.Day(),
		today.Hour(),
		today.Minute(),
		today.Second(),
	)
}
