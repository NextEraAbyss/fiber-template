package config

import (
	"log"

	"github.com/NextEraAbyss/fiber-template/app/schedule"
)

// ScheduleTask 定时任务接口
type ScheduleTask interface {
	Schedule() string // 返回cron表达式
	Task()            // 执行任务
	Start()           // 启动任务
	Stop()            // 停止任务
}

// 所有任务实例
var scheduleTasks []ScheduleTask

// InitTasks 加载所有定时任务
func InitTasks() error {
	// 清空任务列表
	scheduleTasks = []ScheduleTask{}

	// 添加任务
	// 在这里注册您的定时任务
	scheduleTasks = append(scheduleTasks, schedule.NewUpdateStatistics())

	// 这里可以添加更多任务
	// 例如: scheduleTasks = append(scheduleTasks, NewYourTask())

	log.Println("定时任务已加载")

	return nil
}

// BeginTasks 启动所有定时任务
func BeginTasks() {
	for _, task := range scheduleTasks {
		task.Start()
	}
}

// EndTasks 停止所有定时任务
func EndTasks() {
	for _, task := range scheduleTasks {
		task.Stop()
	}
}
