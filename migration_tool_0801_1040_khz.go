// 代码生成时间: 2025-08-01 10:40:25
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// 初始化数据库连接
var db *gorm.DB
var err error

func initDB() {
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }
    fmt.Println("Connected to the database!")
    // 迁移模式
    db.AutoMigrate(&User{}, &Product{})
    fmt.Println("Database migration completed")
}

// 主函数
func main() {
    initDB()
}

// User结构体
type User struct {
    gorm.Model
    Name string
    Email string `gorm:"type:varchar(100);uniqueIndex"`
}

// Product结构体
type Product struct {
    gorm.Model
    Code string  `gorm:"primaryKey"`
    Price uint
}
