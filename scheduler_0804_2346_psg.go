// 代码生成时间: 2025-08-04 23:46:57
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/robfig/cron/v3"
)

// Scheduler 定时任务调度器结构体
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler 初始化定时任务调度器
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(),
    }
}

// AddJob 添加定时任务
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    _, err := s.cron.AddFunc(spec, cmd)
    if err != nil {
        return err
    }
    return nil
}

// Start 开始执行定时任务
func (s *Scheduler) Start() {
    s.cron.Start()
}

// Stop 停止执行定时任务
func (s *Scheduler) Stop() {
    s.cron.Stop()
}

// ExampleJob 演示任务，打印当前时间
func ExampleJob() {
    fmt.Println("Executing job at", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
    // 创建调度器
    scheduler := NewScheduler()

    // 添加定时任务
    _, err := scheduler.AddJob("* * * * *", ExampleJob)
    if err != nil {
        log.Fatalf("Failed to add job: %v", err)
    }

    // 开始执行定时任务
    scheduler.Start()

    // 模拟长时间运行，实际中可能是无限循环，例如使用select{}
    time.Sleep(10 * time.Minute)

    // 停止执行定时任务
    scheduler.Stop()
}
