// 代码生成时间: 2025-09-24 00:42:33
@author: Your Name
@date: 2023-11-24
@version: 1.0
*/

package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// User represents the user entity with necessary fields for authentication
type User struct {
    gorm.Model
    Username string `gorm:"column:username;uniqueIndex"`
    Password string `gorm:"column:password"`
}

// AuthService handles user authentication
type AuthService struct {
    db *gorm.DB
}

// NewAuthService creates an instance of AuthService
func NewAuthService() *AuthService {
    // Initialize and connect to SQLite database
    db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    return &AuthService{db: db}
}

// Authenticate checks if the provided username and password are correct
func (a *AuthService) Authenticate(username, password string) error {
    var user User
    // Fetch user by username
    result := a.db.Where(&User{Username: username}).First(&user)
    if result.Error != nil {
        return fmt.Errorf("authentication failed: %w", result.Error)
    }

    // Hash the provided password and compare with the stored hash
    passwordHash := sha256.Sum256([]byte(password))
    if hex.EncodeToString(passwordHash[:]) != user.Password {
        return fmt.Errorf("authentication failed: incorrect password")
    }

    return nil
}

func main() {
    authService := NewAuthService()
    err := authService.Authenticate("exampleUser", "examplePassword")
    if err != nil {
        fmt.Println("Authentication failed: ", err)
    } else {
        fmt.Println("Authentication successful")
    }
}
