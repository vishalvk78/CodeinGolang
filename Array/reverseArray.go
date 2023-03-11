package main

import "fmt"

//Write a program to reverse an array 

func main() {
    arr := []int{1, 2, 3, 4, 5, 6}
    n := len(arr)
    
    reverseArray(arr, n)
    printArray(arr, n)
}

func reverseArray(arr []int, n int) {
    for i:=0; i<n/2; i++ {
        temp := arr[i]
        arr[i] = arr[n-i-1]
        arr[n-i-1] = temp
    }
}

func printArray(arr []int, n int) {
    for _, val := range arr {
        fmt.Print(val, " ")
    }
}
