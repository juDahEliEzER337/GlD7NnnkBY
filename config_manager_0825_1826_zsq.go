// 代码生成时间: 2025-08-25 18:26:38
package main

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// Configuration holds the configuration details
type Configuration struct {
    ID    uint   "gorm:"primaryKey;autoIncrement""
    Key   string
    Value string
}

// ConfigManager is the manager for handling configurations
type ConfigManager struct {
    db *gorm.DB
}

// NewConfigManager creates a new instance of ConfigManager
func NewConfigManager(dbFile string) (*ConfigManager, error) {
    db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil, err
    }
    
    // Migrate the schema
    db.AutoMigrate(&Configuration{})
    return &ConfigManager{db: db}, nil
}

// SaveConfiguration saves the configuration to the database
func (cm *ConfigManager) SaveConfiguration(key string, value string) error {
    var config Configuration
    // Check if the configuration already exists
    result := cm.db.Where(&Configuration{Key: key}).First(&config)
    if result.Error != nil && !result.NotFound() {
        return result.Error
    }
    if result.NotFound() {
        // Create new configuration
        config = Configuration{Key: key, Value: value}
        result = cm.db.Create(&config)
        if result.Error != nil {
            return result.Error
        }
    } else {
        // Update existing configuration
        config.Value = value
        result = cm.db.Save(&config)
        if result.Error != nil {
            return result.Error
        }
    }
    return nil
}

// LoadConfiguration loads the configuration from the database
func (cm *ConfigManager) LoadConfiguration(key string) (string, error) {
    var config Configuration
    // Retrieve the configuration by key
    result := cm.db.Where(&Configuration{Key: key}).First(&config)
    if result.Error != nil {
        return "", result.Error
    }
    return config.Value, nil
}

func main() {
    dbFile := "config.db"
    if _, err := os.Stat(dbFile); os.IsNotExist(err) {
        fmt.Println("Database file does not exist, creating...")
        _, err := NewConfigManager(dbFile)
        if err != nil {
            log.Fatalf("Failed to create config manager: %v", err)
        }
    }
    
    cm, err := NewConfigManager(dbFile)
    if err != nil {
        log.Fatalf("Failed to get config manager: %v", err)
    }
    
    // Save a configuration value
    if err := cm.SaveConfiguration("example_key", "example_value"); err != nil {
        log.Fatalf("Failed to save configuration: %v", err)
    }
    
    // Load a configuration value
    value, err := cm.LoadConfiguration("example_key")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
    fmt.Printf("Loaded configuration value: %s
", value)
}
