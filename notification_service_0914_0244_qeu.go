// 代码生成时间: 2025-09-14 02:44:31
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Message struct represents a message to be sent in the notification system
type Message struct {
    gorm.Model
    Content string `gorm:"type:text"`
    Status  string
}

// NotificationService struct encapsulates the functionality for the notification system
type NotificationService struct {
    db *gorm.DB
}

// NewNotificationService initializes a new NotificationService with a database connection
func NewNotificationService(db *gorm.DB) *NotificationService {
    return &NotificationService{db: db}
}

// AddMessage adds a new message to the system
func (s *NotificationService) AddMessage(content string) (*Message, error) {
    message := &Message{Content: content, Status: "pending"}
    if err := s.db.Create(message).Error; err != nil {
        return nil, err
    }
    return message, nil
}

// GetMessages retrieves all messages from the system
func (s *NotificationService) GetMessages() ([]Message, error) {
    var messages []Message
    if err := s.db.Find(&messages).Error; err != nil {
        return nil, err
    }
    return messages, nil
}

// MarkMessageAsSent updates the status of a message to 'sent'
func (s *NotificationService) MarkMessageAsSent(id uint) error {
    if err := s.db.Model(&Message{}).Where("id = ?", id).Update(
        "Status", "sent").Error; err != nil {
        return err
    }
    return nil
}

func main() {
    // Initialize the database connection
    db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database"))
    }
    defer db.Close()
    autoMigrate(db)

    // Create a new notification service
    svc := NewNotificationService(db)

    // Add a message to the system
    message, err := svc.AddMessage("Hello, this is a test message!")
    if err != nil {
        fmt.Println("Error adding message: ", err)
        return
    }
    fmt.Printf("Added message with ID: %d
", message.ID)

    // Mark the message as sent
    if err := svc.MarkMessageAsSent(message.ID); err != nil {
        fmt.Println("Error marking message as sent: ", err)
        return
    }
    fmt.Println("Message marked as sent")

    // Retrieve all messages
    messages, err := svc.GetMessages()
    if err != nil {
        fmt.Println("Error getting messages: ", err)
        return
    }
    fmt.Println("Messages: ")
    for _, msg := range messages {
        fmt.Printf("ID: %d, Content: %s, Status: %s
", msg.ID, msg.Content, msg.Status)
    }
}

// autoMigrate is a helper function to create the necessary database tables
func autoMigrate(db *gorm.DB) {
    if err := db.AutoMigrate(&Message{}); err != nil {
        panic("Failed to auto migrate: " + err.Error())
    }
}
