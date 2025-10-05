// 代码生成时间: 2025-10-05 21:22:59
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Patient 代表病人信息
type Patient struct {
    gorm.Model
    Name    string
    Age     int
    Height  float64
    Weight  float64
    Records []Record
}

// Record 代表病人的健康记录
type Record struct {
    gorm.Model
    PatientID uint
    Patient   Patient
    HR       int    // 心率
    BloodO2  float64 // 血氧
    Glucose  float64 // 血糖
    Temperature float64 // 体温
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("health_monitor.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    defer db.Close()

    // 自动迁移模式，确保表结构是最新的
    db.AutoMigrate(&Patient{}, &Record{})

    // 创建新的病人记录
    patient := Patient{Name: "John Doe", Age: 30, Height: 175.5, Weight: 75.0}
    if err := db.Create(&patient).Error; err != nil {
        log.Fatalf("failed to create patient: %v", err)
    }

    // 创建新的健康记录
    record := Record{
        PatientID: patient.ID,
        HR:       70,
        BloodO2:  98.5,
        Glucose:  5.0,
        Temperature: 37.0,
    }
    if err := db.Create(&record).Error; err != nil {
        log.Fatalf("failed to create record: %v", err)
    }

    // 查询病人信息及其健康记录
    var patientWithRecords Patient
    if err := db.Preload("Records").First(&patientWithRecords, patient.ID).Error; err != nil {
        log.Fatalf("failed to find patient: %v", err)
    }

    fmt.Printf("Patient: %+v
", patientWithRecords)
}
