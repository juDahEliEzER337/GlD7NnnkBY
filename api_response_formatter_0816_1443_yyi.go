// 代码生成时间: 2025-08-16 14:43:44
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ApiResponseFormatter is a struct that holds the data for the API response
type ApiResponseFormatter struct {
    Data      interface{} `json:"data"`
    Message   string    `json:"message"`
    ErrorCode string    `json:"error_code"`
}

// NewApiResponseFormatter creates a new ApiResponseFormatter instance
func NewApiResponseFormatter(data interface{}, message string, errorCode string) ApiResponseFormatter {
    return ApiResponseFormatter{
        Data:      data,
        Message:   message,
        ErrorCode: errorCode,
    }
}

// Response is a function that formats the API response with the given status code and ApiResponseFormatter
func Response(c *gin.Context, response ApiResponseFormatter, statusCode int) {
    c.IndentedJSON(statusCode, response)
}

// ErrorResponse is a helper function to generate an error response
func ErrorResponse(c *gin.Context, errorCode, message string) {
    Response(c, NewApiResponseFormatter(nil, message, errorCode), http.StatusInternalServerError)
}

// SuccessResponse is a helper function to generate a success response
func SuccessResponse(c *gin.Context, data interface{}, message string) {
    Response(c, NewApiResponseFormatter(data, message, ""), http.StatusOK)
}

// main function to run the server
func main() {
    r := gin.Default()

    // Define a sample API endpoint
    r.GET("/api/hello", func(c *gin.Context) {
        SuccessResponse(c, gin.H{"hello": "world"}, "Hello World")
    })

    // Start the server
    r.Run()
}
