// 代码生成时间: 2025-08-17 15:39:11
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// BubbleSort 冒泡排序算法实现
func BubbleSort(arr []int) []int {
    lenArr := len(arr)
    for i := 0; i < lenArr-1; i++ {
        for j := 0; j < lenArr-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j] // 交换元素
            }
        }
    }
    return arr
}

// QuickSort 快速排序算法实现
func QuickSort(arr []int) []int {
    // 如果数组长度小于等于1，直接返回
    if len(arr) <= 1 {
        return arr
    }
    
    // 选择最后一个元素作为基准值
    pivot := arr[len(arr)-1]
    left, right := 0, 0
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] < pivot {
            arr[i], arr[left] = arr[left], arr[i] // 小于基准值的交换到前面
            left++
        }
    }
    arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left] // 将基准值移动到中间
    
    // 递归排序基准值左边和右边的子数组
    QuickSort(arr[:left])
    QuickSort(arr[left+1:])
    return arr
}

// GenerateRandomArray 生成一个指定大小的随机数组
func GenerateRandomArray(size int) []int {
    rand.Seed(time.Now().UnixNano()) // 设置随机种子
    arr := make([]int, size)
    for i := 0; i < size; i++ {
        arr[i] = rand.Intn(100) // 生成0-99的随机数
    }
    return arr
}

func main() {
    size := 20 // 定义数组大小
    fmt.Println("Original array: ", GenerateRandomArray(size))

    // 使用冒泡排序
    sortedArr := BubbleSort(GenerateRandomArray(size))
    fmt.Println("Array after BubbleSort: ", sortedArr)

    // 使用快速排序
    sortedArr = QuickSort(GenerateRandomArray(size))
    fmt.Println("Array after QuickSort: ", sortedArr)
}
