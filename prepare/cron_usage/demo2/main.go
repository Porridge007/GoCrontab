package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	//需要一个调度协程，谁过期了就执行谁
	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob
	)

	scheduleTable = make(map[string]*CronJob)
	now = time.Now()

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	//启动调度协程序
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				//是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					go func(jobName string) {
						fmt.Println("执行", jobName)
					}(jobName)

					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "下次执行时间：", cronJob.nextTime)
				}
			}
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:

			}
		}
	}()
	time.Sleep(100 * time.Second)
}
