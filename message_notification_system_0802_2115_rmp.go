// 代码生成时间: 2025-08-02 21:15:15
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "fmt"
    "log"
)

// Message represents a notification message
type Message struct {
    gorm.Model
    Content   string `gorm:"type:varchar(255)"`
    Receiver string `gorm:"type:varchar(255)"`
}

// DBClient is an interface to interact with the database
type DBClient interface {
    CreateMessage(content string, receiver string) error
    FetchMessages(receiver string) ([]Message, error)
}

// messageDBClient is an implementation of DBClient
type messageDBClient struct {
    db *gorm.DB
}

// NewDBClient initializes a new DBClient
func NewDBClient() (*messageDBClient, error) {
    db, err := gorm.Open(sqlite.Open("notification.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&Message{})
    return &messageDBClient{db: db}, nil
}

// CreateMessage adds a new message to the database
func (c *messageDBClient) CreateMessage(content string, receiver string) error {
    message := Message{Content: content, Receiver: receiver}
    result := c.db.Create(&message)
    return result.Error
}

// FetchMessages retrieves messages for a specific receiver
func (c *messageDBClient) FetchMessages(receiver string) ([]Message, error) {
    var messages []Message
    result := c.db.Where(&Message{Receiver: receiver}).Find(&messages)
    return messages, result.Error
}

func main() {
    // Initialize the database client
    dbClient, err := NewDBClient()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Create a new message
    if err := dbClient.CreateMessage("Hello, this is a notification!", "user123"); err != nil {
        log.Printf("Error creating message: %v", err)
    } else {
        fmt.Println("Message created successfully")
    }

    // Fetch all messages for a user
    messages, err := dbClient.FetchMessages("user123")
    if err != nil {
        log.Printf("Error fetching messages: %v", err)
    } else {
        for _, message := range messages {
            fmt.Printf("Message ID: %d, Content: %s, Receiver: %s
", message.ID, message.Content, message.Receiver)
        }
    }
}
