// 代码生成时间: 2025-09-22 15:24:32
package main

import (
    "fmt"
    "os"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// TestResult 测试结果的结构体
type TestResult struct {
    ID        uint   
    TestName  string "gorm:column:test_name"
    CreatedAt string "gorm:column:created_at"
    Result    string "gorm:column:result"
}

// ReportGenerator 结构体，包含数据库连接
type ReportGenerator struct {
    db *gorm.DB
}

// NewReportGenerator 创建一个新的ReportGenerator实例
func NewReportGenerator(dbPath string) (*ReportGenerator, error) {
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // 迁移模式
    err = db.AutoMigrate(&TestResult{})
    if err != nil {
        return nil, err
    }

    return &ReportGenerator{db: db}, nil
}

// GenerateReport 生成测试报告
func (r *ReportGenerator) GenerateReport() error {
    var results []TestResult
    // 查询所有测试结果
    if err := r.db.Find(&results).Error; err != nil {
        return err
    }

    // 这里可以添加生成报告的逻辑，例如写入文件或者打印到控制台
    for _, result := range results {
        fmt.Printf("Test Name: %s, Result: %s
", result.TestName, result.Result)
    }
    return nil
}

func main() {
    dbPath := "test.db"
    generator, err := NewReportGenerator(dbPath)
    if err != nil {
        fmt.Printf("Failed to create report generator: %s
", err)
        return
    }
    defer generator.db.Close()

    if err := generator.GenerateReport(); err != nil {
        fmt.Printf("Failed to generate report: %s
", err)
        return
    }
}
