package main

import "fmt"

const NAME = "322_零钱兑换"

// leetcode submit region begin(Prohibit modification and deletion)
// coinChange 自底向上使用 dp table 来消除重叠子问题
// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。
func coinChange(coins []int, amount int) int {
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

// leetcode submit region end(Prohibit modification and deletion)

func main() {
    fmt.Printf("%s \n", NAME)
    fmt.Println("------------------------------------------")
    coins := []int{1, 2, 5}
    amount := 11
    coinNum := coinChange(coins, amount)
    fmt.Printf("最少的硬币个数: %d\n", coinNum)
}
