package console

import (
	"log"
	"time"

	"go-cron-example/console/commands"

	"github.com/robfig/cron"
)

func schedule(c *cron.Cron) {
	c.AddFunc("* * * * * *", commands.Test)
}

func Default() {
	log.Println("Starting...")

	c := cron.New()
	c.Start()

	schedule(c)

	timer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second * 10)
		}
	}
}
