// 代码生成时间: 2025-08-20 05:10:40
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// TestReport 代表测试报告的数据结构
type TestReport struct {
    gorm.Model
    Title       string  `gorm:"column:title;type:varchar(255)"`
    Description string  `gorm:"column:description;type:text"`
    Status      string  `gorm:"column:status;type:varchar(50)"`
    Result      string  `gorm:"column:result;type:text"`
    Version     string  `gorm:"column:version;type:varchar(50)"`
    CreatedAt   string  `gorm:"column:created_at;type:varchar(50)"`
    UpdatedAt   string  `gorm:"column:updated_at;type:varchar(50)"`
}

// DBConfig 数据库配置
type DBConfig struct {
    DBName string
}

// NewDB 连接数据库并返回*gorm.DB对象
func NewDB(config DBConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", config.DBName)
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    // 迁移模式，确保数据库结构是最新的
    db.AutoMigrate(&TestReport{})
    return db, nil
}

// GenerateTestReport 生成测试报告
func GenerateTestReport(db *gorm.DB, title, description, status, result, version string) (*TestReport, error) {
    // 创建一个新的测试报告
    report := TestReport{
        Title:       title,
        Description: description,
        Status:      status,
        Result:      result,
        Version:     version,
        CreatedAt:   fmt.Sprintf("%v", now()),
        UpdatedAt:   fmt.Sprintf("%v", now()),
    }
    // 保存测试报告到数据库
    if err := db.Create(&report).Error; err != nil {
        return nil, err
    }
    return &report, nil
}

func now() int64 {
    return time.Now().Unix()
}

func main() {
    // 数据库配置
    config := DBConfig{DBName: "test.db"}
    // 连接数据库
    db, err := NewDB(config)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Migrator.Close()
    
    // 生成测试报告
    report, err := GenerateTestReport(db, "Example Test", "This is an example test.", "Passed", "Test Result", "1.0.0")
    if err != nil {
        log.Printf("Failed to generate test report: %v", err)
        return
    }
    
    fmt.Printf("Test report created: %+v
", report)
}
