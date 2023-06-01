package main

import "fmt"

func solveNQueens() [][]string {
    var res [][]string
    var board [8][8]bool
    dfs(board, 0, &res)
    return res
}

func dfs(board [8][8]bool, row int, res *[][]string) {
    if row == 8 {
        temp := make([]string, 8)
        for i := 0; i < 8; i++ {
            var s string
            for j := 0; j < 8; j++ {
                if board[i][j] {
                    s += "Q"
                } else {
                    s += "."
                }
            }
            temp[i] = s
        }
        *res = append(*res, temp)
        return
    }

    for col := 0; col < 8; col++ {
        if isValid(board, row, col) {
            board[row][col] = true
            dfs(board, row+1, res)
            board[row][col] = false
        }
    }
}

func isValid(board [8][8]bool, row, col int) bool {
    // 检查列是否有冲突
    for i := 0; i < row; i++ {
        if board[i][col] {
            return false
        }
    }

    // 检查左上到右下对角线是否有冲突
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] {
            return false
        }
    }

    // 检查右上到左下对角线是否有冲突
    for i, j := row-1, col+1; i >= 0 && j < 8; i, j = i-1, j+1 {
        if board[i][j] {
            return false
        }
    }

    return true
}

func main() {
    res := solveNQueens()
    for _, r := range res {
        for _, s := range r {
            fmt.Println(s)
        }
        fmt.Println()
    }
}