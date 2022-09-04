package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "net/http/pprof"

	"github.com/muhammad-fakhri/archetype-be/cmd/cron"
	cronhandler "github.com/muhammad-fakhri/archetype-be/cmd/cron/handler"
	"github.com/muhammad-fakhri/archetype-be/cmd/subscriber"
	subhandler "github.com/muhammad-fakhri/archetype-be/cmd/subscriber/handler"
	web "github.com/muhammad-fakhri/archetype-be/cmd/webservice"
	"github.com/muhammad-fakhri/archetype-be/internal/config"
)

// running mode
const (
	// BEGIN __INCLUDE_EXAMPLE_CRON__
	cronEventUpdaterExampleMode = "cron-event-example-updater"
	// END __INCLUDE_EXAMPLE_CRON__
	// BEGIN __INCLUDE_EXAMPLE__
	subscriberUpdateConfigMode = "subscriber-update-config"
	// END __INCLUDE_EXAMPLE__
	webServiceMode = "web-service"
)

var (
	mode string
)

func init() {
	flag.StringVar(&mode, "mode", webServiceMode, "service run mode")
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	config.Init()

	var stopFunc func()
	switch mode {
	// BEGIN __INCLUDE_EXAMPLE_CRON__
	case cronEventUpdaterExampleMode:
		stopFunc = cron.Start(cronhandler.UpdateEventExampleTask)
	// END __INCLUDE_EXAMPLE_CRON__
	// BEGIN __INCLUDE_EXAMPLE__
	case subscriberUpdateConfigMode:
		stopFunc = subscriber.Start(subhandler.UpdateSystemConfigWithDeadLetterTask)
	// END __INCLUDE_EXAMPLE__
	default:
		stopFunc = web.Start()
	}

	// will wait until terminate signal or interrupt happened
	for {
		<-c
		log.Println("terminate service")
		stopFunc()
		os.Exit(0)
	}
}
