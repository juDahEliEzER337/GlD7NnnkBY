// 代码生成时间: 2025-10-03 01:30:22
This tool allows for sending HTTP requests to test different API endpoints.
*/

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

// ApiResponse defines the structure for API response data
type ApiResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    interface{} `json:"data"`
}

func main() {
    // Read API test data from a file
    testDataFile := "test_data.json"
    if len(os.Args) > 1 {
        testDataFile = os.Args[1]
    }
    testDataReader, err := os.Open(testDataFile)
    if err != nil {
        fmt.Println("Error opening test data file:", err)
        return
    }
    defer testDataReader.Close()

    // Decode test data from JSON
    var testData map[string]interface{}
    jsonData, err := ioutil.ReadAll(testDataReader)
    if err != nil {
        fmt.Println("Error reading test data:", err)
        return
    }
    err = json.Unmarshal(jsonData, &testData)
    if err != nil {
        fmt.Println("Error decoding test data:", err)
        return
    }

    // Iterate through test data and send API requests
    for endpoint, data := range testData {
        fmt.Printf("Testing endpoint: %s
", endpoint)
        var requestBody bytes.Buffer
        json.NewEncoder(&requestBody).Encode(data)

        response, err := http.Post(endpoint, "application/json", &requestBody)
        if err != nil {
            fmt.Printf("Error sending request to %s: %s
", endpoint, err)
            continue
        }
        defer response.Body.Close()

        // Read response data
        responseBody, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("Error reading response from %s: %s
", endpoint, err)
            continue
        }

        // Print API response
        fmt.Printf("Response from %s:
%s
", endpoint, responseBody)
    }
}
