package job

import "github.com/GanEasy/wxrankapi/repository"

//TaskSpider struct  采集任务
type TaskSpider struct {
	URL string
}

//RunTask 实现 Task.RunTask()
func (t *TaskSpider) RunTask() {
	repository.Post(t.URL)
}
