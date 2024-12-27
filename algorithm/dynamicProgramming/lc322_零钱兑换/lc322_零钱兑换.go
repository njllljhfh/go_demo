package main

import (
    "fmt"
    "math"
)

/*
322. 零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

难度：中等

https://leetcode.cn/problems/coin-change/description/
*/

// // 伪码框架
// func coinChange(coins []int, amount int) int {
//     // 题目要求的最终结果是 dp(amount)
//     return dp(coins, amount)
// }
//
// // 定义：要凑出金额 n，至少要 dp(coins, n) 个硬币
// func dp(coins []int, n int) int {
//     // 初始化为最大值
//     res := math.MaxInt
//
//     // 做选择，选择需要硬币最少的那个结果
//     for _, coin := range coins {
//         res = min(res, 1+subProblem)
//     }
//     return res
// }

// coinChange 自定向下递归
// 时间复杂度：O(k^n)
func coinChange(coins []int, amount int) int {
    // 题目要求的最终结果是 dp(amount)
    return dp(coins, amount)
}

// 定义：要凑出金额 n，至少要 dp(coins, n) 个硬币
func dp(coins []int, amount int) int {
    // base case
    if amount == 0 {
        return 0
    }
    if amount < 0 {
        return -1
    }

    res := math.MaxInt
    for _, coin := range coins {
        // 计算子问题的结果
        subProblem := dp(coins, amount-coin)
        // 子问题无解则跳过
        if subProblem == -1 {
            continue
        }
        // 在子问题中选择最优解，然后加一
        res = minNum(res, subProblem+1)
    }

    if res == math.MaxInt {
        return -1
    }
    return res
}

func minNum(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func main() {
    coins := []int{1, 2, 5}
    amount := 11
    coinNum := coinChange(coins, amount)
    fmt.Printf("最少的硬币个数: %d\n", coinNum)
}
