// 代码生成时间: 2025-10-01 02:38:25
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "fmt"
)

// Notification represents the structure for a notification entity
type Notification struct {
    gorm.Model
    Title   string `gorm:"type:varchar(255);"`
    Message string `gorm:"type:text;"`
}

// NotificationService defines the methods for interact with notifications
type NotificationService struct {
    db *gorm.DB
}

// NewNotificationService creates a new instance of NotificationService
func NewNotificationService(db *gorm.DB) *NotificationService {
    return &NotificationService{db: db}
}

// CreateNotification adds a new notification to the database
func (service *NotificationService) CreateNotification(title, message string) error {
    notification := Notification{Title: title, Message: message}
    if err := service.db.Create(&notification).Error; err != nil {
        return err
    }
    return nil
}

// GetAllNotifications retrieves all notifications from the database
func (service *NotificationService) GetAllNotifications() ([]Notification, error) {
    var notifications []Notification
    if err := service.db.Find(&notifications).Error; err != nil {
        return nil, err
    }
    return notifications, nil
}

// main function to demonstrate the usage of NotificationService
func main() {
    // Initialize SQLite DB
    db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Notification{})

    // Create a NotificationService instance
    notificationService := NewNotificationService(db)

    // Create some notifications
    if err := notificationService.CreateNotification("This is a title", "This is a message"); err != nil {
        log.Fatalf("error creating notification: %v", err)
    }

    // Get all notifications and print them
    allNotifications, err := notificationService.GetAllNotifications()
    if err != nil {
        log.Fatalf("error getting all notifications: %v", err)
    }
    for _, notification := range allNotifications {
        fmt.Printf("Title: %s, Message: %s
", notification.Title, notification.Message)
    }
}