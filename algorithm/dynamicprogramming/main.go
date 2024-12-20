package main

import "fmt"

/*
动态规划解题套路框架
https://labuladong.online/algo/essential-technique/dynamic-programming-framework/
*/

func fib1(N int) int {
    // 备忘录全初始化为 0
    memo := make([]int, N+1)
    // 进行带备忘录的递归
    return dp1(memo, N)
}

// dp1 带着备忘录进行递归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func dp1(memo []int, n int) int {
    // base case
    if n == 0 || n == 1 {
        return n
    }
    // 已经计算过，不用再计算了
    if memo[n] != 0 {
        return memo[n]
    }
    memo[n] = dp1(memo, n-1) + dp1(memo, n-2)
    return memo[n]
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// fib2 自底向上的 dp 数组的迭代（递推）解法
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func fib2(N int) int {
    if N == 0 {
        return 0
    }

    // dp table
    dp := make([]int, N+1)

    // base case
    dp[0], dp[1] = 0, 1

    // 状态转移
    for i := 2; i <= N; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[N]
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// 细心的读者会发现，根据斐波那契数列的状态转移方程，当前状态 n 只和之前的 n-1, n-2 两个状态有关，
// 其实并不需要那么长的一个 DP table 来存储所有的状态，只要想办法存储之前的两个状态就行了。
// 所以，可以进一步优化，把空间复杂度降为 O(1)。这也就是我们最常见的计算斐波那契数的算法：
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func fib3(n int) int {
    if n == 0 || n == 1 {
        // base case
        return n
    }
    // 分别代表 dp[i - 1] 和 dp[i - 2]
    dp_i_1, dp_i_2 := 1, 0
    for i := 2; i <= n; i++ {
        // dp[i] = dp[i - 1] + dp[i - 2];
        dp_i := dp_i_1 + dp_i_2
        // 滚动更新
        dp_i_2 = dp_i_1
        dp_i_1 = dp_i
    }
    return dp_i_1
}

/*
这一般是动态规划问题的最后一步优化，如果我们发现每次状态转移只需要 DP table 中的一部分，
那么可以尝试缩小 DP table 的大小，只记录必要的数据，从而降低空间复杂度。

上述例子就相当于把 DP table 的大小从 n 缩小到 2，即把空间复杂度下降了一个量级。我会在后文
对动态规划发动降维打击 进一步讲解这个压缩空间复杂度的技巧，一般来说用来把一个二维的 DP table 压缩成一维，即把空间复杂度从 O(n^2) 压缩到 O(n)。

有人会问，动态规划的另一个重要特性「最优子结构」，怎么没有涉及？下面会涉及。
斐波那契数列的例子严格来说不算动态规划，因为没有涉及求最值，以上旨在说明重叠子问题的消除方法，演示得到最优解法逐步求精的过程。
下面，看第二个例子，凑零钱问题。
*/

func main() {
    // fmt.Println(fib1(5))
    // fmt.Println(fib2(5))
    fmt.Println(fib3(5))
}
