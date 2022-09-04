package handler

import (
	"log"

	"github.com/muhammad-fakhri/archetype-be/internal/config"
	"github.com/muhammad-fakhri/archetype-be/internal/usecase"
	"github.com/robfig/cron/v3"
)

type CronTask int

const (
	EmptyTask CronTask = iota
	// BEGIN __INCLUDE_EXAMPLE_CRON__
	UpdateEventExampleTask
	// END __INCLUDE_EXAMPLE_CRON__
)

type CronParam struct {
	Handler  CronHandler
	Schedule cron.Schedule
}

type CronHandler func()

func Init(usecase usecase.Usecase, config *config.Scheduler, task CronTask) (param CronParam) {
	switch task {
	// BEGIN __INCLUDE_EXAMPLE_CRON__
	case UpdateEventExampleTask:
		param = CronParam{
			Handler:  UpdateEventExample(usecase.UpdateEventExample),
			Schedule: parseSchedule(config.CronSpecEventExampleUpdater),
		}
		// END __INCLUDE_EXAMPLE_CRON__
	default:
		log.Fatalf("failed init cron. err:invalid task %d", task)
	}

	return
}

func parseSchedule(config string) (schedule cron.Schedule) {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	schedule, err := parser.Parse(config)
	if err != nil {
		log.Fatalf("invalid cron specs, err:%v", err)
	}

	return schedule
}
