package main

import "fmt"

func main() {
    //a := [][]int{
    //    {1, 2, 3, 4},
    //    {5, 6, 7, 8},
    //    {9, 10, 11, 12},
    //    {13, 14, 15, 16},
    //}

    a := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    rotate(a)

    fmt.Println(a)
}

func test(matrix [][]int) {
    n := len(matrix)
    for i := 0; i < len(matrix)/2; i++ {
        for j := i; j < n-1-i; j++ {
            //matrix[i][j], matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i] = matrix[n-1-j][i], matrix[n-1-j][n-1-j], matrix[j][n-1-i], matrix[i][j]

            matrix[i][j] = matrix[j][i]

        }
    }
}

func rotate(matrix [][]int) {
    n := len(matrix)
    for i := 0; i < len(matrix)/2; i++ {
        for j := i; j < n-1-i; j++ {
            matrix[i][j], matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i] = matrix[n-1-j][i], matrix[n-1-j][n-1-j], matrix[j][n-1-i], matrix[i][j]
        }
    }
}
