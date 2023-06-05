package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	memo := make(map[string][]int, 0)
	return dfs(expression, memo)
}

func dfs(expression string, memo map[string][]int) []int {
	//返回条件，当表达式只剩数字时，无法继续dfs
	if number, ok := isNumber(expression); ok {
		return []int{number}
	}
	//剪枝法直接返回已知的表达式不同运算优先级可能的计算结果
	if results, found := memo[expression]; found {
		return results
	}
	ans := make([]int, 0)
	//遍历表达式中的运算符，将运算符左右两边的表达式继续按相同逻辑处理，规模缩小，故可认为是一种分治
	for i, _ := range expression {
		char := expression[i]
		if char == '+' || char == '-' || char == '*' {
			leftExp, rightExp := expression[:i], expression[i+1:]
			//运算符左边表达式不同运算优先级下的计算结果
			leftRes := dfs(leftExp, memo)
			//运算符右边表达式不同运算优先级下的计算结果
			rightRes := dfs(rightExp, memo)
			val := 0
			//遍历组合左右表达式的计算结果，得到以该运算符作为最后一个运算符，可能的结果集合
			for _, l := range leftRes {
				for _, r := range rightRes {
					switch char {
					case '+':
						val = l + r
					case '-':
						val = l - r
					case '*':
						val = l * r
					}
					ans = append(ans, val)
				}
			}
			memo[expression] = ans
		}
	}
	//将本层答案返回给上一层
	return ans
}

func isNumber(expression string) (int, bool) {
	val, err := strconv.Atoi(expression)
	if err != nil {
		return -1, false
	}
	return val, true
}

func main() {
	expression := "2*3-4*5"
	ans := diffWaysToCompute(expression)
	fmt.Println(ans)
}
