// 代码生成时间: 2025-10-03 17:13:46
// slow_query_analyzer.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql" // Import your database driver
)

// SlowQuery represents a slow query with its details
type SlowQuery struct {
    ID        uint      `gorm:"primary_key"`
    Query     string    `gorm:"type:text"`
    Duration  time.Duration
    CreatedAt time.Time
}

// AnalyzerConfig is the configuration for the slow query analyzer
type AnalyzerConfig struct {
    DSN          string  // Data Source Name
    Threshold    float64 // Threshold for slow queries in seconds
}

// SlowQueryAnalyzer is the main struct for analyzing slow queries
type SlowQueryAnalyzer struct {
    db         *sql.DB
    config     AnalyzerConfig
    slowQueries []SlowQuery
}

// NewSlowQueryAnalyzer creates a new instance of SlowQueryAnalyzer
func NewSlowQueryAnalyzer(config AnalyzerConfig) (*SlowQueryAnalyzer, error) {
    var db *sql.DB
    var err error

    db, err = sql.Open("mysql", config.DSN)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    return &SlowQueryAnalyzer{db: db, config: config}, nil
}

// Analyze queries the database for slow queries based on the threshold and returns them
func (a *SlowQueryAnalyzer) Analyze() ([]SlowQuery, error) {
    var queries []SlowQuery
    err := a.db.QueryRow("SHOW PROFILES WHERE Query_time > ?", a.config.Threshold).Scan(&queries)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve slow queries: %w", err)
    }
    a.slowQueries = queries
    return queries, nil
}

func main() {
    config := AnalyzerConfig{
        DSN:          "user:password@tcp(127.0.0.1:3306)/dbname",
        Threshold:    0.5, // 500ms
    }

    analyzer, err := NewSlowQueryAnalyzer(config)
    if err != nil {
        log.Fatal(err)
    }

    slowQueries, err := analyzer.Analyze()
    if err != nil {
        log.Fatal(err)
    }

    for _, sq := range slowQueries {
        fmt.Printf("Slow Query: %s
", sq.Query)
        fmt.Printf("Duration: %f seconds
", sq.Duration.Seconds())
        fmt.Printf("Timestamp: %s

", sq.CreatedAt)
    }
}