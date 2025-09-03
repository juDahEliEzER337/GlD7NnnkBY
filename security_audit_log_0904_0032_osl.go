// 代码生成时间: 2025-09-04 00:32:05
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// AuditLog represents a security audit log entry
type AuditLog struct {
	ID        uint   	`gorm:"primaryKey"`
	Timestamp string 	`gorm:"type:datetime"`
	Action    string 	`gorm:"type:varchar(255)"`
	UserID    uint   	`gorm:"index"`
	UserName  string 	`gorm:"type:varchar(255)"`
	Details   string 	`gorm:"type:text"`
}

// InitDatabase initializes the database connection
func InitDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("audit_log.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// CreateAuditLogEntry creates a new audit log entry
func CreateAuditLogEntry(db *gorm.DB, action, userName, details string, userID uint) error {
	auditLog := AuditLog{
		Timestamp: fmt.Sprintf("%v", GormCurrentTime()),
		Action:    action,
		UserID:    userID,
		UserName:  userName,
		Details:   details,
	}
	result := db.Create(&auditLog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GormCurrentTime returns the current time formatted in a way suitable for GORM's datetime field
func GormCurrentTime() string {
	return fmt.Sprintf("%v", time.Now())
}

func main() {
	// Initialize the database connection
	db := InitDatabase()
	
	// Migrate the schema
	db.AutoMigrate(&AuditLog{})
	
	// Example usage: create an audit log entry
	err := CreateAuditLogEntry(db, "USER_LOGIN", "example_user", "User logged in successfully", 1)
	if err != nil {
		fmt.Println("Error creating audit log entry: ", err)
	} else {
		fmt.Println("Audit log entry created successfully.")
	}
}
