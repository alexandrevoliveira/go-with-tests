package main

import (
	"os"
	"time"

	c "mocking/internal"
)

func main() {
	sleeper := &c.ConfigurableSleeper{Duration: 1 * time.Second, Sleepf: time.Sleep}
	c.Countdown(os.Stdout, sleeper)
}
