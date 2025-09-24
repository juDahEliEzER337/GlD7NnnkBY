// 代码生成时间: 2025-09-24 14:09:05
package main

import (
    "bytes"
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "reflect"
    "syscall"

    "github.com/tealeg/xlsx"
)

// ExcelGenerator represents the excel generator structure
type ExcelGenerator struct {
    db *gorm.DB
}

// NewExcelGenerator creates a new instance of ExcelGenerator
func NewExcelGenerator(db *gorm.DB) *ExcelGenerator {
    return &ExcelGenerator{db: db}
}

// GenerateExcel generates an excel file from a query
func (e *ExcelGenerator) GenerateExcel(query string, outputPath string, data interface{}) error {
    var results []map[string]interface{}
    if err := e.db.Raw(query).Scan(&results).Error; err != nil {
        return fmt.Errorf("failed to execute query: %w", err)
    }

    if len(results) == 0 {
        return nil // No data to process
    }

    // Get the first row to determine the column names
    columns := results[0]
    file := xlsx.NewFile()
    for key := range columns {
        file.AddSheet().SetCellValue(key, 1, 1, key)
    }

    for i, row := range results {
        for key, value := range row {
            cell, _ := xlsx.NewCell()
            cell.Value = fmt.Sprintf("%v", value)
            file.SetCellValue(key, i+2, 1, cell)
        }
    }

    // Write the output to a file
    f, err := syscall.Create(outputPath)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer f.Close()
    writer, err := xlsx.NewWriter(f)
    if err != nil {
        return fmt.Errorf("failed to create writer: %w", err)
    }
    writer.WriteFile(file)
    if err := writer.Close(); err != nil {
        return fmt.Errorf("failed to close writer: %w", err)
    }

    return nil
}

func main() {
    // Initialize a new SQLite database
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&YourModel{})

    // Create a new ExcelGenerator
    excelGen := NewExcelGenerator(db)

    // Define the file path and query
    outputPath := "output.xlsx"
    query := "SELECT * FROM your_table"
    // Define the data type to fetch
    var data []YourModel

    // Generate the Excel file
    if err := excelGen.GenerateExcel(query, outputPath, &data); err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Excel file generated successfully")
    }
}

// YourModel represents your data model
type YourModel struct {
    ID       uint   "gorm:\"primaryKey\" json:\"id\""
    Name     string 
    // Add other fields as per your schema
}

// Ensure YourModel adheres to the gorm.Model interface
func (YourModel) TableName() string {
    return "your_table"
}
