package main

import (
        "fmt"
)

func main() {
        var i = twoSum([]int{2, 7, 11, 15, 23, 44}, 18)
        fmt.Printf("%v\n", i)
        var m = binarySearch([]int{1, 2, 3, 4, 5, 7, 8, 9}, 7)
        fmt.Printf("%d\n", m)
}
