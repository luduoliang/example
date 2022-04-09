package common

import "github.com/robfig/cron/v3"

var Cron *cron.Cron

//初始化定时任务
func InitCron() {
	Cron = cron.New(cron.WithSeconds())
	Cron.Start()
	//select {} //阻塞主线程停止
}

//添加定时任务
func AddCronTask(spec string, taskFunc func()) {
	if Cron != nil {
		Cron.AddFunc(spec, func() {
			taskFunc()
		})
	}
}
