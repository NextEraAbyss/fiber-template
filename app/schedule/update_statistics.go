package schedule

import (
	"log"

	"github.com/robfig/cron/v3"
)

// UpdateStatistics 更新统计数据的定时任务（示例）
// 实现ScheduleTask接口
type UpdateStatistics struct {
	cron *cron.Cron
}

// Schedule 返回定时任务的执行时间，使用cron表达式
func (t *UpdateStatistics) Schedule() string {
	// 每30秒执行一次（示例）
	return "*/30. * * * * *"
}

// Task 定时任务的执行逻辑
func (t *UpdateStatistics) Task() {
	log.Println("执行示例定时任务")

	// 这里添加您的业务逻辑
	log.Println("示例任务执行完成")
}

// Start 启动定时任务
func (t *UpdateStatistics) Start() {
	t.cron = cron.New(cron.WithSeconds())
	_, err := t.cron.AddFunc(t.Schedule(), t.Task)
	if err != nil {
		log.Printf("启动定时任务失败: %v", err)
		return
	}
	t.cron.Start()
	log.Printf("定时任务已启动: [%s]", t.Schedule())
}

// Stop 停止定时任务
func (t *UpdateStatistics) Stop() {
	if t.cron != nil {
		t.cron.Stop()
		log.Println("定时任务已停止")
	}
}

// NewUpdateStatistics 创建并返回一个新的UpdateStatistics实例
func NewUpdateStatistics() *UpdateStatistics {
	return &UpdateStatistics{}
}
