// 代码生成时间: 2025-09-08 10:39:56
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

// CalculateSHA256 takes a string and returns its SHA256 hash.
func CalculateSHA256(input string) (string, error) {
    // Create a new SHA256 hash generator.
    hash := sha256.New()
    
    // Write the input to the hash generator.
    _, err := hash.Write([]byte(input))
    if err != nil {
        return "", err
    }
    
    // Compute the hash.
    hashedBytes := hash.Sum(nil)
    
    // Convert the hash to a hexadecimal string.
    return hex.EncodeToString(hashedBytes), nil
}

func main() {
    // Example usage of CalculateSHA256 function.
    input := "Hello, Gophers!"
    hash, err := CalculateSHA256(input)
    if err != nil {
        fmt.Printf("Error calculating hash: %s
", err)
    } else {
        fmt.Printf("The SHA256 hash of '%s' is: %s
", input, hash)
    }
}
