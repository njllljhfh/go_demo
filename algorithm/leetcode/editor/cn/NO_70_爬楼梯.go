package main

import "fmt"

// leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
    allSteps := []int{1, 2}
    memo := make([]int, n+1)
    for i := 1; i < len(memo); i++ {
        memo[i] = -666
    }
    return dp(n, allSteps, memo)
}

func dp(n int, allSteps, memo []int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 2
    }

    subProblems := make([]int, len(allSteps))

    for i, step := range []int{1, 2} {
        subProblem := n - step
        if subProblem >= 1 {
            if memo[subProblem] != -666 {
                subProblems[i] = memo[subProblem]
            } else {
                subProblems[i] = dp(subProblem, allSteps, memo)
                memo[subProblem] = subProblems[i]
            }
        }
    }

    kind := 0
    for _, v := range subProblems {
        kind += v
    }

    return kind
}

// leetcode submit region end(Prohibit modification and deletion)
func main() {
    fmt.Println(climbStairs(4))
}
