// 代码生成时间: 2025-09-17 20:09:08
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User represents the model for our users
type User struct {
    gorm.Model
    Name string
    Email string `gorm:"type:varchar(100);uniqueIndex"`
    Age uint
}

// Database connection settings
const dbPath = "./test.db"

// SetupDatabase initializes the database connection and migrates the schema
func SetupDatabase() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&User{})
    return db, nil
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, userData *User) error {
    // Save the user data
    if err := db.Create(userData).Error; err != nil {
        return err
    }
    return nil
}

// GetUser retrieves a user from the database by ID
func GetUser(db *gorm.DB, id uint) (User, error) {
    var user User
    // Find the user
    if result := db.First(&user, id).Error; result != nil {
        return user, result
    }
    return user, nil
}

func main() {
    db, err := SetupDatabase()
    if err != nil {
        fmt.Println("Error setting up database: \%+v", err)
        return
    }
    defer db.Migrator.Close()

    // Example user creation
    user := User{Name: "John Doe", Email: "john.doe@example.com", Age: 30}
    if err := CreateUser(db, &user); err != nil {
        fmt.Println("Error creating user: \%+v", err)
        return
    }
    fmt.Println("User created successfully")

    // Example user retrieval
    retrievedUser, err := GetUser(db, user.ID)
    if err != nil {
        fmt.Println("Error retrieving user: \%+v", err)
        return
    }
    fmt.Printf("Retrieved User: %+v\
", retrievedUser)
}
