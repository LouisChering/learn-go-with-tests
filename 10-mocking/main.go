package main

import (
	"os"

	. "github.com/louischering/learn-go-with-tests/10-mocking/mocking"
)

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
