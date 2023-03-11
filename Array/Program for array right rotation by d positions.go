package main

import "fmt"

//Program for array left rotation by d positions.

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7}
    
    n := len(arr)
    
    d :=3
    
    rotateArray(arr, n, d)
    printArray(arr)
}

func rotateArray(arr []int, n, d int) {
    for i:=0; i<d; i++ {
        temp := arr[n-1]
        for j:=n-2; j>=0; j-- {
            arr[j+1]=arr[j]
        }
        arr[0]=temp
    }
}


func printArray(arr []int){
    for _, val := range arr {
        fmt.Print(val, " ")
    }
}
