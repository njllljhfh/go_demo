package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
var emptyFlag = -666

func climbStairs(n int) int {
    // 支持任意无重复的自然数组
    allSteps := []int{1, 2}
    memo := make([]int, n+1)
    for i := 0; i < len(memo); i++ {
        memo[i] = emptyFlag
    }
    return dp(n, allSteps, memo)
}

func dp(n int, allSteps, memo []int) int {
    // 例如，对于[1,2]的情况，n=2 的子问题是 2-1=1 和 2-2=0,
    if n == 0 {
        return 1
    }

    if memo[n] != emptyFlag {
        return memo[n]
    }

    // 子问题
    subProblems := make([]int, len(allSteps))

    for i, step := range allSteps {
        subProblem := n - step
        if subProblem >= 0 {
            if memo[subProblem] != emptyFlag {
                // 从缓存中获取子问题
                subProblems[i] = memo[subProblem]
            } else {
                // 缓存子问题
                subProblems[i] = dp(subProblem, allSteps, memo)
                memo[subProblem] = subProblems[i]
            }
        }
    }

    kind := 0
    for _, v := range subProblems {
        kind += v
    }
    memo[n] = kind

    return kind
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
    fmt.Println(climbStairs(4))
}
