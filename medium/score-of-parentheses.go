// https://leetcode.com/problems/score-of-parentheses/
package main

import (
	"fmt"
	"strconv"
)

func pop(stack *[]string) string {
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret
}

func peek(stack []string) string {
	return stack[len(stack)-1]
}

func sumArray(arr []string) int {
	tmp := 0
	for _, v := range arr {
		vx, _ := strconv.Atoi(v)
		tmp += vx
	}
	return tmp
}

func fillStrings(a, b string) (string, string) {
	for len(a) < len(b) {
		a = "0" + a
	}
	for len(a) > len(b) {
		b = "0" + b
	}
	return a, b
}

func addNumberAsString(a, b string) string {
	ret := ""
	a, b = fillStrings(a, b)
	remember := 0
	for i := len(a) - 1; i >= 0; i-- {
		tmp := int(a[i]+b[i]-'0'-'0') + remember
		remember = tmp / 10
		tmp = tmp % 10
		ret = string(tmp+'0') + ret
	}

	if remember > 0 {
		ret = string(remember+'0') + ret
	}

	return ret
}

func scoreOfParentheses(S string) int {
	stack := make([]string, 0)

	for _, c := range S {
		// fmt.Println(string(c), stack)
		if c == '(' {
			stack = append(stack, string(c))
		} else {
			tmp := pop(&stack)
			// fmt.Println(len(stack))

			if tmp == "(" {
				stack = append(stack, "1")
			} else {
				tmp = addNumberAsString(tmp, tmp)
				pop(&stack)
				stack = append(stack, tmp)
			}

			tmp = "0"
			for len(stack) > 1 && peek(stack) != "(" {
				tmp = addNumberAsString(tmp, pop(&stack))
			}
			if tmp != "0" {
				stack = append(stack, tmp)
			}
		}
	}

	// fmt.Println(stack)
	return sumArray(stack)
}

func main() {
	fmt.Println(scoreOfParentheses("()"))
	fmt.Println(scoreOfParentheses("(())"))
	fmt.Println(scoreOfParentheses("(()())"))
	fmt.Println(scoreOfParentheses("()()()"))
	fmt.Println(scoreOfParentheses("(()(()))"))
	// fmt.Println(addNumberAsString("5", "5"))
	fmt.Println(scoreOfParentheses("(()(()()))"))
}
