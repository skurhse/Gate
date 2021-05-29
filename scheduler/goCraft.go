package scheduler

import (
	"fmt"
	"github.com/gobuffalo/envy"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"os/signal"
)

var redisPool = &redis.Pool{
	MaxActive: 50,
	MaxIdle:   50,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", fmt.Sprintf("%s:%s", envy.Get("REDIS_ADDRESS", "localhost"), envy.Get("REDIS_PORT", "6379")))
	},
}

var enqueue = work.NewEnqueuer("jh_bot_pool", redisPool)

type Context struct {
	customerID int64
}

func InitWorker() {
	pool := work.NewWorkerPool(Context{}, 50, "gate_pool", redisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)
	pool.Middleware((*Context).FindCustomer)

	//// Map the name of jobs to handler functions
	pool.JobWithOptions(OHLC, work.JobOptions{Priority: 1, MaxFails: 2}, (*Context).ohlc)
	//pool.JobWithOptions(GameRequest, work.JobOptions{Priority: 1, MaxFails: 2}, (*Context).sendGameRequest)
	//pool.JobWithOptions(Broadcast, work.JobOptions{Priority: 2, MaxFails: 2}, (*Context).Broadcast)
	//
	//// Customize options:
	//pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	log.Print("Starting job: ", job.Name)
	return next()
}

func (c *Context) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
	if _, ok := job.Args["customer_id"]; ok {
		c.customerID = job.ArgInt64("customer_id")
		if err := job.ArgError(); err != nil {
			return err
		}
	}
	return next()
}
