package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

const DefaultBackupDataJobRunSpec = "@every 3s"

func main() {
	cron := cron.New(cron.WithSeconds())
	cron.AddJob(DefaultBackupDataJobRunSpec, &ReminderEmail{})
	cron.Start()
	time.Sleep(10 * time.Second)

}

type ReminderEmail struct {
}

func (re *ReminderEmail) Run() {
	fmt.Println("Every 3 sec send reminder emails")

}
