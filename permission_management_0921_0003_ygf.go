// 代码生成时间: 2025-09-21 00:03:04
package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

// User represents a user with permissions
type User struct {
    gorm.Model
    Username string
    Permissions []Permission `gorm:"many2many:user_permissions;"`
}

// Permission represents a permission for a user
type Permission struct {
    gorm.Model
    Name string
}

// UserPermission is the join table for User and Permission many-to-many relationship
type UserPermission struct {
    User uint
    Permission uint
}

// PermissionManager handles all operations related to user permissions
type PermissionManager struct {
    db *gorm.DB
}

// NewPermissionManager creates a new PermissionManager with a database connection
func NewPermissionManager() *PermissionManager {
    db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }
    db.AutoMigrate(&User{}, &Permission{}, &UserPermission{})
    return &PermissionManager{db: db}
}

// AddUser adds a new user with the given username
func (pm *PermissionManager) AddUser(username string) error {
    user := User{Username: username}
    result := pm.db.Create(&user)
    return result.Error
}

// AddPermission adds a new permission with the given name
func (pm *PermissionManager) AddPermission(name string) error {
    permission := Permission{Name: name}
    result := pm.db.Create(&permission)
    return result.Error
}

// AssignPermission assigns a permission to a user
func (pm *PermissionManager) AssignPermission(userID uint, permissionID uint) error {
    result := pm.db.Model(&User{}).Where(&User{ID: userID}).Assign("Permissions", []Permission{Permission{ID: permissionID}})
    return result.Error
}

// RevokePermission revokes a permission from a user
func (pm *PermissionManager) RevokePermission(userID uint, permissionID uint) error {
    result := pm.db.Model(&User{}).Where(&User{ID: userID}).Updates(map[string]interface{}{"Permissions": []Permission{Permission{ID: permissionID}}{ID: 0}})
    return result.Error
}

func main() {
    pm := NewPermissionManager()
    
    // Add new user
    if err := pm.AddUser("johndoe"); err != nil {
        log.Println("Error adding user: ", err)
    }
    
    // Add new permission
    if err := pm.AddPermission("read"); err != nil {
        log.Println("Error adding permission: ", err)
    }
    
    // Assign permission to user
    if err := pm.AssignPermission(1, 1); err != nil {
        log.Println("Error assigning permission: ", err)
    }
    
    // Revoke permission from user
    if err := pm.RevokePermission(1, 1); err != nil {
        log.Println("Error revoking permission: ", err)
    }
    
    fmt.Println("User permissions management system is running...")
}