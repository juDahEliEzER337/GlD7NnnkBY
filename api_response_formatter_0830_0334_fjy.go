// 代码生成时间: 2025-08-30 03:34:01
package main

import (
    "encoding/json"
    "net/http"
    "fmt"
)

// ApiResponseFormatter is a utility to format API responses with a standardized structure.
type ApiResponseFormatter struct {}

// NewApiResponseFormatter creates a new instance of ApiResponseFormatter.
func NewApiResponseFormatter() *ApiResponseFormatter {
    return &ApiResponseFormatter{}
}

// FormatResponse formats the response with a success message and data.
// It returns a JSON response with a status code of 200.
func (formatter *ApiResponseFormatter) FormatResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "status":  "success",
        "data":    data,
    }
    json.NewEncoder(w).Encode(response)
}

// FormatErrorResponse formats the response with an error message and returns a JSON response.
// It returns a JSON response with a status code of 400 or 500 based on the error type.
func (formatter *ApiResponseFormatter) FormatErrorResponse(w http.ResponseWriter, err error, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "status":  "error",
        "message": err.Error(),
    }
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}

// Main function to demonstrate the usage of ApiResponseFormatter.
func main() {
    http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
        formatter := NewApiResponseFormatter()
        formatter.FormatResponse(w, map[string]string{"key": "value"})
    })

    http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
        formatter := NewApiResponseFormatter()
        err := fmt.Errorf("something went wrong")
        formatter.FormatErrorResponse(w, err, http.StatusInternalServerError)
    })

    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", nil)
}
