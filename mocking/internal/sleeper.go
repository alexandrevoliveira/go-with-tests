package internal

import "time"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleepf   func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.Sleepf(c.Duration)
}
