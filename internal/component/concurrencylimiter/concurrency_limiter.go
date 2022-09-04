package concurrencylimiter

const (
	DefaultConcurrentLimit = 10
)

type limiter struct {
	limit    int
	routines chan int
}

func NewConcurrencyLimiter(limit int) ConcurrencyLimiter {
	if limit <= 0 {
		limit = DefaultConcurrentLimit
	}
	l := &limiter{
		limit:    limit,
		routines: make(chan int, limit),
	}
	for i := 0; i < l.limit; i++ {
		l.routines <- i
	}
	return l
}

// Run execute the passed function when the limiter has vacant goroutine to execute it.
// The limit of goroutine this limiter is able to spawn can be set upon initialization.
// Calling Run(job) does not block the next process if there is a vacant goroutine.
// If there is no goroutine available, it will block the next process until a goroutine can take the job.
func (l *limiter) Run(job func()) int {
	routine := <-l.routines
	go func() {
		defer func() {
			l.routines <- routine
		}()
		job()
	}()
	return routine
}

// Wait will block until all goroutine within the limiter is done executing all the jobs.
func (l *limiter) Wait() {
	for i := 0; i < l.limit; i++ {
		<-l.routines
	}
}
