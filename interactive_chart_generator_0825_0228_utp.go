// 代码生成时间: 2025-08-25 02:28:01
 * interactive_chart_generator.go
 * This program is an interactive chart generator using the GORM framework in Go.
 * It allows users to input data and create charts dynamically.
 *
 * Author: Your Name
 * Date: YYYY-MM-DD
 */

package main

import (
    "fmt"
    "log"
    "os"
    "github.com/jedib0t/go-pretty/v6/table"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Chart represents the structure of a chart
type Chart struct {
    gorm.Model
    Title string `gorm:"type:varchar(255)"`
    Type  string `gorm:"type:varchar(255)"`
    Data  string `gorm:"type:text"`
}

// DatabaseConfig holds configuration for the database
type DatabaseConfig struct {
    DSN string
}

// dbClient is the global variable for the database connection
var dbClient *gorm.DB

func main() {
    // Setup database connection
    config := DatabaseConfig{DSN: "chart_generator.db"}
    db, err := gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    defer db.Close()
    dbClient = db

    // Migrate the schema
    db.AutoMigrate(&Chart{})

    // Create a new chart
    newChart()
}

// newChart function to create a new chart
func newChart() {
    fmt.Println("Welcome to the Interactive Chart Generator!")
    for {
        fmt.Println("Please enter the chart title: ")
        var title string
        fmt.Scanln(&title)

        fmt.Println("Please enter the chart type (e.g., 'bar', 'pie'): ")
        var chartType string
        fmt.Scanln(&chartType)

        fmt.Println("Please enter chart data (e.g., '10,20,30' for a bar chart): ")
        var data string
        fmt.Scanln(&data)

        // Create a new chart
        chart := Chart{Title: title, Type: chartType, Data: data}
        if err := dbClient.Create(&chart).Error; err != nil {
            log.Println("Failed to create chart: ", err)
            continue
        }

        // Display the chart
        displayChart(chart)

        // Ask if the user wants to create another chart
        fmt.Println("Do you want to create another chart? (yes/no): ")
        var answer string
        fmt.Scanln(&answer)
        if answer != "yes" {
            break
        }
    }
}

// displayChart function to display the chart data
func displayChart(chart Chart) {
    fmt.Println("Chart Title: ", chart.Title)
    fmt.Println("Chart Type: ", chart.Type)
    fmt.Println("Chart Data: ", chart.Data)
    // Here you would add code to generate and display the actual chart,
    // using a charting library or service.
}
