package concurrencylimiter

// ConcurrencyLimiter represents the internal concurrency limiter interface.
type ConcurrencyLimiter interface {
	// Run execute the passed function when the limiter has vacant goroutine to execute it.
	// The limit of goroutine this limiter is able to spawn can be set upon initialization.
	// Calling Run(job) does not block the next process if there is a vacant goroutine.
	// If there is no goroutine available, it will block the next process until a goroutine can take the job.
	Run(job func()) int
	// Wait will block until all goroutine within the limiter is done executing all the jobs.
	Wait()
}
