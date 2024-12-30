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

// coinChange 自顶向下递归
// 时间复杂度：O(k^n)
func coinChange1(coins []int, amount int) int {
    // 题目要求的最终结果是 dp(amount)
    return dp1(coins, amount)
}

// 定义：要凑出金额 n，至少要 dp(coins, n) 个硬币
func dp1(coins []int, amount int) int {
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
        subProblem := dp1(coins, amount-coin)
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

// ------------------------------------------------------------------

// coinChange2 带备忘录的递归
// 类似之前斐波那契数列的例子，只需要稍加修改，就可以通过备忘录消除子问题：
func coinChange2(coins []int, amount int) int {
    memo := make([]int, amount+1)
    for i := range memo {
        memo[i] = -666
    }
    // 备忘录初始化为一个不会被取到的特殊值，代表还未被计算
    return dp2(coins, amount, memo)
}

func dp2(coins []int, amount int, memo []int) int {
    if amount == 0 {
        return 0
    }
    if amount < 0 {
        return -1
    }
    // 查备忘录，防止重复计算
    if memo[amount] != -666 {
        return memo[amount]
    }

    res := math.MaxInt32
    for _, coin := range coins {
        // 计算子问题的结果
        subProblem := dp2(coins, amount-coin, memo)
        // 子问题无解则跳过
        if subProblem == -1 {
            continue
        }
        // 在子问题中选择最优解，然后加一
        res = min(res, subProblem+1)
    }
    // 把计算结果存入备忘录
    if res == math.MaxInt32 {
        memo[amount] = -1
    } else {
        memo[amount] = res
    }
    return memo[amount]
}

// ------------------------------------------------------------------

// coinChange3 自底向上使用 dp table 来消除重叠子问题
// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。
func coinChange3(coins []int, amount int) int {
    // 数组大小为 amount + 1，初始值也为 amount + 1
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1
    }

    // base case
    dp[0] = 0
    // 外层 for 循环在遍历所有状态的所有取值
    for i := 0; i < len(dp); i++ {
        // 内层 for 循环在求所有选择的最小值
        for _, coin := range coins {
            // 子问题无解，跳过
            if i-coin < 0 {
                continue
            }
            dp[i] = min(dp[i], 1+dp[i-coin])
        }
    }
    if dp[amount] == amount+1 {
        return -1
    }
    return dp[amount]
}

func main() {
    coins := []int{1, 2, 5}
    amount := 11
    // coinNum := coinChange1(coins, amount)
    // coinNum := coinChange2(coins, amount)
    coinNum := coinChange3(coins, amount)
    fmt.Printf("最少的硬币个数: %d\n", coinNum)
}
