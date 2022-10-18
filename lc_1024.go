package gobar

import (
	"strconv"
	"strings"
)

func LC1024(nums []int, ops []byte) string {
	res := lc1024(nums, ops)
	if len(res) == 0 {
		return ""
	}
	return strings.Join(lc1024(nums, ops)[0], " ; ")
}

func lc1024(nums []int, ops []byte) [][]string {
	var ans [][]string
	var list [3]string
	var dfs func([]int, []byte, [3]string, int)
	dfs = func(nums []int, ops []byte, list [3]string, p int) {
		for i := 0; i < len(nums)-1 && len(ans) == 0; i++ {
			for j := i + 1; j < len(nums) && len(ans) == 0; j++ {
				for k, op := range ops {
					a, b := nums[i], nums[j]
					if op == '/' && b == 0 {
						continue
					}
					res, str := calc(a, b, op)
					list[p] = str

					if p == 2 {
						if res == 1024 {
							ans = append(ans, []string{list[0], list[1], list[2]})
						}
						continue
					}

					nums[j] = res
					nums = append(nums[:i], nums[i+1:]...)
					ops = append(ops[:k], ops[k+1:]...)

					dfs(nums, ops, list, p+1)

					ops = append(ops[:k], append([]byte{op}, ops[k:]...)...)
					nums = append(nums[:i], append([]int{a}, nums[i:]...)...)
					nums[j] = b
				}
			}
		}
	}
	dfs(nums, ops, list, 0)
	return ans
}

func calc(a, b int, op byte) (int, string) {
	ans := 0
	switch op {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	//**
	case 'p':
		ans = pow(a, b)
	case '%':
		ans = a % b
	case '/':
		ans = a / b
	case '|':
		ans = a | b
	case '&':
		ans = a & b
	case '^':
		ans = a ^ b
	case '<':
		ans = a << b
	case '>':
		ans = a >> b
	}
	s1 := strconv.Itoa(a)
	s2 := strconv.Itoa(b)
	return ans, s1 + string([]byte{op}) + s2
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		half := pow(x, n>>1)
		return half * half
	} else {
		return pow(x, n-1) * x
	}
}
