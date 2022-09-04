package config

type Scheduler struct {
	// BEGIN __INCLUDE_EXAMPLE_CRON__
	CronSpecEventExampleUpdater string
	// END __INCLUDE_EXAMPLE_CRON__
	// add more here
}

func (c *Config) initSchedulerConfig(cfg *configIni) {
	appConfig.Scheduler = &Scheduler{
		// BEGIN __INCLUDE_EXAMPLE_CRON__
		CronSpecEventExampleUpdater: cfg.CronSpecEventExampleUpdater,
		// END __INCLUDE_EXAMPLE_CRON__
	}
}
