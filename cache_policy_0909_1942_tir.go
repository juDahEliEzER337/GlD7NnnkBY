// 代码生成时间: 2025-09-09 19:42:05
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/go-redis/redis/v8"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// CacheItem represents an item stored in the cache
type CacheItem struct {
    Key       string    `gorm:"primaryKey"`
    Value     string    
    CreatedAt time.Time 
}

// CacheManager provides an interface to manage cache
type CacheManager struct {
    db     *gorm.DB
    redis  *redis.Client
    TTL     time.Duration // Time to live for cache items
}

// NewCacheManager initializes a new CacheManager
func NewCacheManager(db *gorm.DB, redis *redis.Client, ttl time.Duration) *CacheManager {
    return &CacheManager{db: db, redis: redis, TTL: ttl}
}

// Set stores a value in the cache
func (cm *CacheManager) Set(key string, value string) error {
    // Store in Redis
    if err := cm.redis.Set(key, value, cm.TTL).Err(); err != nil {
        return err
    }

    // Store in Database (GORM)
    item := CacheItem{Key: key, Value: value, CreatedAt: time.Now()}
    result := cm.db.Create(&item)
    return result.Error
}

// Get retrieves a value from the cache
func (cm *CacheManager) Get(key string) (string, error) {
    // Try to get from Redis first
    if value, err := cm.redis.Get(key).Result(); err == nil {
        return value, nil
    }

    // If not found in Redis, try to get from Database (GORM)
    var item CacheItem
    if result := cm.db.First(&item, key); result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return "", nil // Return nil if not found in database
        }
        return "", result.Error
    }

    // If found in database, store it back in Redis for future retrievals
    if err := cm.redis.Set(item.Key, item.Value, cm.TTL).Err(); err != nil {
        log.Printf("Failed to set value in Redis: %v", err)
    }

    return item.Value, nil
}

func main() {
    // Initialize database connection
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Migrate schema
    db.AutoMigrate(&CacheItem{})

    // Initialize Redis connection
    redisClient := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // Check connection
    pong, err := redisClient.Ping().Result()
    if err != nil {
        log.Fatal("Failed to connect to Redis: ", err)
    } else {
        fmt.Printf("Connected to Redis: %s
", pong)
    }

    // Create CacheManager with 10 minutes TTL
    cacheManager := NewCacheManager(db, redisClient, 10*time.Minute)

    // Set a cache item
    if err := cacheManager.Set("test_key", "Hello, World!"); err != nil {
        log.Fatal("Failed to set cache item: ", err)
    }

    // Get a cache item
    value, err := cacheManager.Get("test_key")
    if err != nil {
        log.Fatal("Failed to get cache item: ", err)
    }

    fmt.Printf("Cache value: %s
", value)
}