// 代码生成时间: 2025-08-05 12:58:43
package main

import (
    "fmt"
    "runtime"
    "time"
)

// MemoryUsageAnalyzer is a struct to hold the current memory usage statistics.
type MemoryUsageAnalyzer struct {
    // LastGC is the time when the last garbage collection occurred.
    LastGC time.Time
    // TotalAllocated is the total number of bytes allocated by the program.
    TotalAllocated uint64
    // Frees is the total number of frees by the program.
    Frees uint64
    // LiveObjects is the number of live objects.
    LiveObjects uint64
    // HeapReleased is the amount of heap memory that has been released to the OS.
    HeapReleased uint64
    // HeapInUse is the amount of heap memory in use.
    HeapInUse uint64
    // StackInUse is the amount of stack memory in use.
    StackInUse uint64
    // MSpanInUse is the amount of mspan structures in use.
    MSpanInUse uint64
    // MCacheInUse is the amount of mcache structures in use.
    MCacheInUse uint64
    // BuckHashSys is the amount of memory used by the profiling bucket hash table.
    BuckHashSys uint64
    // GCSys is the amount of memory used by the garbage collector system.
    GCSys uint64
    // OtherSys is the amount of memory used for other system allocations.
    OtherSys uint64
    // NextGC is the target heap size for the next GC.
    NextGC uint64
    // LastGCHeapSize is the size of the heap at the last GC.
    LastGCHeapSize uint64
    // PauseTotal is the total number of nanoseconds spent in GC pause.
    PauseTotal uint64
    // NumGC is the number of GC that have been forced since the program started.
    NumGC uint32
}

// AnalyzeMemoryUsage retrieves and returns the current memory usage statistics.
func AnalyzeMemoryUsage() (MemoryUsageAnalyzer, error) {
    var m MemoryUsageAnalyzer
    var memStats runtime.MemStats

    // Read full memory statistics.
    err := runtime.ReadMemStats(&memStats)
    if err != nil {
        return m, fmt.Errorf("failed to read memory statistics: %w", err)
    }

    // Copy the memory statistics to the MemoryUsageAnalyzer struct.
    m.LastGC = memStats.LastGC
    m.TotalAllocated = memStats.Alloc
    m.Frees = memStats.Frees
    m.LiveObjects = memStats.HeapObjects
    m.HeapReleased = memStats.HeapReleased
    m.HeapInUse = memStats.HeapInuse
    m.StackInUse = memStats.StackInuse
    m.MSpanInUse = memStats.MSpanInuse
    m.MCacheInUse = memStats.MCacheInuse
    m.BuckHashSys = memStats.BuckHashSys
    m.GCSys = memStats.GCSys
    m.OtherSys = memStats.OtherSys
    m.NextGC = memStats.NextGC
    m.LastGCHeapSize = memStats.LastGC
    m.PauseTotal = memStats.PauseTotalNs
    m.NumGC = memStats.NumGC

    return m, nil
}

func main() {
    // Analyze and print memory usage every 5 seconds.
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        analyzer, err := AnalyzeMemoryUsage()
        if err != nil {
            fmt.Printf("Error analyzing memory usage: %v
", err)
            continue
        }

        // Print memory usage statistics.
        fmt.Printf("Memory Usage Analysis:
")
        fmt.Printf("  Last GC: %v
", analyzer.LastGC)
        fmt.Printf("  Total Allocated: %d bytes
", analyzer.TotalAllocated)
        fmt.Printf("  Frees: %d
", analyzer.Frees)
        fmt.Printf("  Live Objects: %d
", analyzer.LiveObjects)
        fmt.Printf("  Heap Released: %d bytes
", analyzer.HeapReleased)
        fmt.Printf("  Heap In Use: %d bytes
", analyzer.HeapInUse)
        fmt.Printf("  Stack In Use: %d bytes
", analyzer.StackInUse)
        fmt.Printf("  MSpan In Use: %d bytes
", analyzer.MSpanInUse)
        fmt.Printf("  MCache In Use: %d bytes
", analyzer.MCacheInUse)
        fmt.Printf("  BuckHashSys: %d bytes
", analyzer.BuckHashSys)
        fmt.Printf("  GCSys: %d bytes
", analyzer.GCSys)
        fmt.Printf("  OtherSys: %d bytes
", analyzer.OtherSys)
        fmt.Printf("  Next GC: %d bytes
", analyzer.NextGC)
        fmt.Printf("  Last GC Heap Size: %d bytes
", analyzer.LastGCHeapSize)
        fmt.Printf("  Pause Total: %d nanoseconds
", analyzer.PauseTotal)
        fmt.Printf("  Number of GC: %d
", analyzer.NumGC)
    }
}