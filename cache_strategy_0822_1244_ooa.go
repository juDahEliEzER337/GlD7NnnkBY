// 代码生成时间: 2025-08-22 12:44:55
package main

import (
    "fmt"
    "time"
    "github.com/go-redis/redis/v8" // 使用第三方库实现缓存
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// CacheStrategy 缓存策略接口
type CacheStrategy interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}, duration time.Duration) error
}

// RedisCache 实现了 CacheStrategy 接口，使用 Redis 作为缓存存储
type RedisCache struct {
    Client *redis.Client
}

// NewRedisCache 初始化 Redis 缓存
func NewRedisCache() *RedisCache {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis 地址
        Password: "",               // Redis 密码，没有则为空字符串
        DB:       0,                // 默认数据库
    })
    return &RedisCache{Client: rdb}
}

// Get 从 Redis 中获取缓存值
func (rc *RedisCache) Get(key string) (interface{}, error) {
    result, err := rc.Client.Get(ctx, key).Result()
    if err != nil && err != redis.Nil {
        return nil, err
    }
    return result, nil
}

// Set 在 Redis 中设置缓存值
func (rc *RedisCache) Set(key string, value interface{}, duration time.Duration) error {
    err := rc.Client.Set(ctx, key, value, duration).Err()
    if err != nil {
        return err
    }
    return nil
}

// DatabaseModel 数据库模型
type DatabaseModel struct {
    gorm.Model
    Value string
}

// DatabaseCache 封装数据库和缓存操作
type DatabaseCache struct {
    DB     *gorm.DB
    Cache  RedisCache
}

// NewDatabaseCache 初始化数据库和缓存
func NewDatabaseCache() (*DatabaseCache, error) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    _ = db.AutoMigrate(&DatabaseModel{})
    cache := NewRedisCache()
    return &DatabaseCache{DB: db, Cache: *cache}, nil
}

// GetModelWithCache 先从缓存中获取，如果没有则从数据库获取，最后设置缓存
func (dc *DatabaseCache) GetModelWithCache(key string) (*DatabaseModel, error) {
    var model DatabaseModel
    cacheResult, err := dc.Cache.Get(key)
    if err != nil {
        return nil, err
    }
    if cacheResult != nil {
        fmt.Println("Cache hit")
        err = dc.DB.Where("value = ?