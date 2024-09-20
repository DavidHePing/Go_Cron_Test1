package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// Seconds field, required
	job := cron.New(cron.WithSeconds())
	job.AddFunc("*/1 * * * * *", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})

	job.Start()

	time.Sleep(1 * time.Hour)
}
