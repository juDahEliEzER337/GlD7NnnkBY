// 代码生成时间: 2025-10-09 00:00:30
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Member represents a member of the membership system.
type Member struct {
    gorm.Model
    Points   int
    Username string
}

// PointService provides operations for managing member points.
type PointService struct {
    db *gorm.DB
}

// NewPointService initializes a new PointService with a database connection.
func NewPointService(db *gorm.DB) *PointService {
    return &PointService{db: db}
}

// AddPoints adds points to a member.
func (s *PointService) AddPoints(memberID uint, points int) error {
    member := Member{}
    if err := s.db.First(&member, memberID).Error; err != nil {
        return err
    }
    member.Points += points
    if err := s.db.Save(&member).Error; err != nil {
        return err
    }
    return nil
}

// SubtractPoints subtracts points from a member.
func (s *PointService) SubtractPoints(memberID uint, points int) error {
    member := Member{}
    if err := s.db.First(&member, memberID).Error; err != nil {
        return err
    }
    if member.Points < points {
        return fmt.Errorf("not enough points")
    }
    member.Points -= points
    if err := s.db.Save(&member).Error; err != nil {
        return err
    }
    return nil
}

// AutoMigrate ensures that the Member table is created in the database.
func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(&Member{}).Error
}

func main() {
    // Initialize a new database connection using SQLite.
    db, err := gorm.Open(sqlite.Open("membership.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Create the Member table if it does not exist.
    if err := AutoMigrate(db); err != nil {
        panic("failed to migrate database")
    }

    // Initialize the PointService with the database connection.
    pointService := NewPointService(db)

    // Example usage: Adding and subtracting points from a member.
    // Add points to a member with ID 1.
    if err := pointService.AddPoints(1, 100); err != nil {
        fmt.Println("Error adding points: ", err)
    }

    // Subtract points from a member with ID 1.
    if err := pointService.SubtractPoints(1, 50); err != nil {
        fmt.Println("Error subtracting points: ", err)
    }
}
