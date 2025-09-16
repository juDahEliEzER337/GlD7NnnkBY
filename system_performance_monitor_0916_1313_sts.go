// 代码生成时间: 2025-09-16 13:13:30
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/load"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/net"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// SystemStats 结构体用于存储系统性能监控数据
type SystemStats struct {
    ID        uint      `gorm:"primaryKey"`
    Timestamp time.Time `gorm:"type:datetime"`
    CPUUsage  float64   
    MemUsage  uint64    
    DiskUsage uint64    
    NetUsage  uint64    
}

// SystemPerformanceMonitor 系统性能监控工具
type SystemPerformanceMonitor struct {
    db *gorm.DB
}

// NewSystemPerformanceMonitor 初始化系统性能监控工具
func NewSystemPerformanceMonitor(dbPath string) (*SystemPerformanceMonitor, error) {
    var db *gorm.DB
    var err error
    // 连接数据库
    db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // 自动迁移数据库模式
    db.AutoMigrate(&SystemStats{})

    return &SystemPerformanceMonitor{db: db}, nil
}

// Monitor 监控系统性能并存储到数据库
func (spm *SystemPerformanceMonitor) Monitor() error {
    // 获取CPU使用率
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return err
    }
    var cpuUsage float64
    if len(cpuPercent) > 0 {
        cpuUsage = cpuPercent[0]
    }

    // 获取内存使用量
    memStat, err := mem.VirtualMemory()
    if err != nil {
        return err
    }
    var memUsage uint64 = memStat.Used

    // 获取磁盘使用量
    diskStat, err := disk.Usage="/"
    if err != nil {
        return err
    }
    var diskUsage uint64 = diskStat.Used

    // 获取网络使用量
    netStat, err := net.IOCounters()
    if err != nil {
        return err
    }
    var netUsage uint64 = netStat.BytesSent + netStat.BytesRecv

    // 存储监控数据
    stats := SystemStats{
        Timestamp: time.Now(),
        CPUUsage:  cpuUsage,
        MemUsage:  memUsage,
        DiskUsage: diskUsage,
        NetUsage:  netUsage,
    }

    if err := spm.db.Create(&stats).Error; err != nil {
        return err
    }

    return nil
}

func main() {
    dbPath := "system_stats.db"
    monitor, err := NewSystemPerformanceMonitor(dbPath)
    if err != nil {
        log.Fatal("Failed to initialize system performance monitor: ", err)
    }

    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        if err := monitor.Monitor(); err != nil {
            fmt.Println("Error monitoring system performance: ", err)
        }
    }
}
