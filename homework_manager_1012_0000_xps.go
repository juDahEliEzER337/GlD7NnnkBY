// 代码生成时间: 2025-10-12 00:00:18
package main

import (
    "fmt"
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Homework 定义作业实体
type Homework struct {
    gorm.Model
    Title     string  `gorm:"size:255"`
    Content   string  `gorm:"type:text"`
    Deadline  string  `gorm:"size:255"`
    UserID    uint    
    User      User    `gorm:"foreignKey:UserID"`
}

// User 定义用户实体
type User struct {
    gorm.Model
    Name     string  `gorm:"size:255"`
    Email    string  `gorm:"type:varchar(100);uniqueIndex"`
    Homeworks []Homework `gorm:"foreignKey:UserID"`
}

func main() {
    // 连接数据库
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }

    // 迁移模式
    db.AutoMigrate(&User{}, &Homework{})

    // 创建用户
    user := User{Name: "John Doe", Email: "john@example.com"}
    db.Create(&user)

    // 创建作业
    homework := Homework{Title: "Math Homework", Content: "Solve the equation.", Deadline: "2023-12-31", UserID: user.ID}
    db.Create(&homework)

    // 查询所有作业
    var homeworks []Homework
    db.Preload("User").Find(&homeworks)
    fmt.Println("Homeworks:")
    for _, h := range homeworks {
        fmt.Printf("Title: %s, Content: %s, Deadline: %s, User: %s
", h.Title, h.Content, h.Deadline, h.User.Name)
    }
}
