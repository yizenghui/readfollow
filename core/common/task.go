package common

//----------
// 二维码扫描任务
//----------

import "errors"

type (
	// QrCodeTask 二维码任务调试
	QrCodeTask struct {
		ID     int32
		Ticket string
		Value  string
	}
)

var (
	//tasks 二维码任务列表
	tasks       = map[int32]*QrCodeTask{}
	seq   int32 = 100001
)

//CreateTask 创建签名任务
func CreateTask() int32 {
	t := &QrCodeTask{
		ID: seq,
	}
	tasks[t.ID] = t
	seq++
	return t.ID
}

// GetTask 获取任务
func GetTask(id int32) (task *QrCodeTask, err error) {
	if task, ok := tasks[id]; ok {
		return task, nil
	}
	return tasks[id], errors.New("不存在的任务")
}

// SetValue 更新任务 (获取到值了)
func SetValue(id int32, value string) error {
	task, err := GetTask(id)
	if err == nil {
		task.Value = value
	}
	return err
}

// SetTicket 设置二维码 Ticket
func SetTicket(id int32, ticket string) error {
	task, err := GetTask(id)
	if err == nil {
		task.Ticket = ticket
	}
	return err
}

// CheckTicket 检查二维码 Ticket
func CheckTicket(id int32, ticket string) bool {
	task, err := GetTask(id)
	if err == nil {
		if task.Ticket == ticket {
			return true
		}
	}
	return false
}

// GetTaskCompleteValue 检查二维码 Ticket
func GetTaskCompleteValue(id int32) (string, error) {
	task, err := GetTask(id)
	if err == nil {
		if task.Value != "" {
			return task.Value, nil
		}
		err = errors.New("任务未完成")
	}
	return "", err
}

// DeleteTask 删除任务
func DeleteTask(id int32) {
	delete(tasks, id)
}

// DeleteAllTask 删除所有任务
func DeleteAllTask() {
	for k := range tasks {
		delete(tasks, k)
	}
}
