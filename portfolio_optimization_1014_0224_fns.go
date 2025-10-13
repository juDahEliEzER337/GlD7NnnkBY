// 代码生成时间: 2025-10-14 02:24:25
// portfolio_optimization.go

package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Portfolio represents a collection of investments.
type Portfolio struct {
    gorm.Model
    Name    string   `gorm:"column:name;unique"`
    Stocks []Stock   `gorm:"foreignKey:PortfolioID"`
}

// Stock represents a single investment in the portfolio.
type Stock struct {
    gorm.Model
    PortfolioID uint
    Name        string  `gorm:"column:name"`
    Quantity    uint   `gorm:"column:quantity"`
    Price       float64 `gorm:"column:price"`
}

// Database holds the gorm.DB instance.
type Database struct {
    DB *gorm.DB
}

func main() {
    db := initializeDatabase()
    defer db.DB.Close()

    // Add example data to the database
    addExampleData(db)

    // Fetch and print portfolio data
    portfolios, err := fetchPortfolios(db)
    if err != nil {
        fmt.Printf("Error fetching portfolios: %+v
", err)
        return
    }

    for _, portfolio := range portfolios {
        fmt.Printf("Portfolio: %+v
", portfolio)
        for _, stock := range portfolio.Stocks {
            fmt.Printf("  Stock: %+v
", stock)
        }
    }
}

// initializeDatabase sets up the database connection and migrates the schema.
func initializeDatabase() *Database {
    db, err := gorm.Open(sqlite.Open("invest.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&Portfolio{}, &Stock{})
    return &Database{DB: db}
}

// addExampleData adds example data to the database.
func addExampleData(db *Database) {
    db.DB.Create(&Portfolio{Name: "Example Portfolio"})
    for i := 1; i <= 5; i++ {
        db.DB.Create(&Stock{
            PortfolioID: 1,
            Name:        fmt.Sprintf("Stock %d", i),
            Quantity:    100 * i,
            Price:       10.0 + float64(i),
        })
    }
}

// fetchPortfolios retrieves all portfolios from the database.
func fetchPortfolios(db *Database) ([]Portfolio, error) {
    var portfolios []Portfolio
    if result := db.DB.Preload("Stocks").Find(&portfolios); result.Error != nil {
        return nil, result.Error
    }
    return portfolios, nil
}
