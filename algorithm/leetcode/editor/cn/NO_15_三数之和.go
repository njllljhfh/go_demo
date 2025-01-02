package main

import (
    "fmt"
    "sort"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 官方解法
// func threeSum(nums []int) [][]int {
//     n := len(nums)
//     sort.Ints(nums)
//     ans := make([][]int, 0)
//
//     // 枚举 a
//     for first := 0; first < n; first++ {
//         // 需要和上一次枚举的数不相同
//         if first > 0 && nums[first] == nums[first-1] {
//             continue
//         }
//         // c 对应的指针初始指向数组的最右端
//         third := n - 1
//         target := -1 * nums[first]
//         // 枚举 b
//         for second := first + 1; second < n; second++ {
//             // 需要和上一次枚举的数不相同
//             if second > first+1 && nums[second] == nums[second-1] {
//                 continue
//             }
//             // 需要保证 b 的指针在 c 的指针的左侧
//             for second < third && nums[second]+nums[third] > target {
//                 third--
//             }
//             // 如果指针重合，随着 b 后续的增加
//             // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
//             if second == third {
//                 break
//             }
//             if nums[second]+nums[third] == target {
//                 ans = append(ans, []int{nums[first], nums[second], nums[third]})
//             }
//         }
//     }
//     return ans
// }
// leetcode submit region end(Prohibit modification and deletion)

// 其他道友的解法
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    ans := make([][]int, 0)
    if nums[n-1] < 0 {
        return ans
    }

    for i := 0; i < n-2; i++ {
        if nums[i] > 0 {
            return ans
        }
        a := nums[i]
        // 双指针找值
        l, r := i+1, n-1
        for l < r {
            b := nums[l]
            c := nums[r]
            if a+b+c < 0 {
                l++
            } else if a+b+c > 0 {
                r--
            } else {
                ans = append(ans, []int{a, b, c})
                for r > 0 && nums[r] == c {
                    r--
                }
                for l < n && nums[l] == b {
                    l++
                }
            }
        }
        // 跳过 nums[i] 右边与 nums[i] 重复的数
        for i < n-2 && nums[i+1] == a {
            i++
        }
    }
    return ans
}

func main() {
    fmt.Printf("%v \n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
    // fmt.Printf("%v \n", threeSum([]int{1, 2, -2, -1}))
    // fmt.Printf("%v \n", threeSum([]int{-4, -1, -1, 0, 1, 2}))
}
