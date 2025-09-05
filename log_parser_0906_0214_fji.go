// 代码生成时间: 2025-09-06 02:14:24
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// LogEntry represents a single log entry with its timestamp, level, and message
type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
}

// ParseLogFile parses a log file and returns a slice of LogEntry objects
func ParseLogFile(filePath string) ([]LogEntry, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    var entries []LogEntry
    scanner := newLineScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parsedEntry, err := parseLine(line)
        if err != nil {
            // Handle parsing error and continue with the next line
            log.Printf("failed to parse line: %s, error: %v", line, err)
            continue
        }
        entries = append(entries, *parsedEntry)
    }
    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("failed to scan file: %w", err)
    }
    return entries, nil
}

// parseLine parses a single line of log and returns a LogEntry
func parseLine(line string) (*LogEntry, error) {
    // Assuming log format: "2023-03-03T12:00:00Z INFO Some log message"
    parts := strings.Fields(line)
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid log line format")
    }
    timestamp, err := time.Parse(time.RFC3339, parts[0]+"T"+parts[1])
    if err != nil {
        return nil, fmt.Errorf("failed to parse timestamp: %w", err)
    }
    level := parts[2]
    message := strings.Join(parts[3:], " ")
    return &LogEntry{Timestamp: timestamp, Level: level, Message: message}, nil
}

// newLineScanner creates a scanner that reads lines from the given reader
func newLineScanner(r io.Reader) *bufio.Scanner {
    return bufio.NewScanner(r)
}

func main() {
    filePath := "example.log"
    entries, err := ParseLogFile(filePath)
    if err != nil {
        fmt.Printf("An error occurred: %v", err)
        return
    }

    for _, entry := range entries {
        fmt.Printf("%s - %s - %s
", entry.Timestamp.Format(time.RFC3339), entry.Level, entry.Message)
    }
}
