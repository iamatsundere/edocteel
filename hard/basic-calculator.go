// https://leetcode.com/problems/basic-calculator/
package main

import (
	"fmt"
	"strconv"
)

func pop(stack  *[]string) string {
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret
}

func peek(stack []string) string {
	return stack[len(stack)-1]
}

func sumASimpleEprs(queue []string) int {
	// fmt.Println(queue)
	a := ""
	b := ""
	o := ""

	ax := 0
	bx := 0

	for i := len(queue) - 1; i >= 0; i-- {
		// fmt.Println("a=", a, "b=", b, "o=", o)
		v := queue[i]
		switch v {
		case "-", "+":
			if o != "" {
				ax, _ = strconv.Atoi(a)
				bx, _ = strconv.Atoi(b)
				if o == "-" {
					ax = ax - bx
				} else {
					ax = ax + bx
				}
				a = strconv.Itoa(ax)
				b = ""
			}
			o = v
			break
		default:
			if a != "" && o != "" {
				b = b + v
			} else {
				a = a + v
			}

			break
		}
	}
	if o != "" {
		ax, _ = strconv.Atoi(a)
		bx, _ = strconv.Atoi(b)
		if o == "-" {
			ax = ax - bx
		} else {
			ax = ax + bx
		}
		a = strconv.Itoa(ax)
	}
	ax, _ = strconv.Atoi(a)
	return ax
}

func calculate(s string) int {
	ret := 0
	stack := make([]string, 0)
	queue := make([]string, 0)
	tmp := ""
	n := 0

	for _, c := range s {
		// fmt.Println(stack)
		switch c {
		case ' ':
			break
		case '(', '-', '+':
			stack = append(stack, string(c))
			break
		case ')':
			for len(stack) > 0 {
				tmp = pop(&stack)
				if tmp != "(" {
					queue = append(queue, tmp)
				} else {
					break
				}
			}
			// fmt.Println(queue)
			n = sumASimpleEprs(queue)
			stack = append(stack, strconv.Itoa(n))
			queue = make([]string, 0)
			break
		default:
			stack = append(stack, string(c))
			break
		}
	}

	// fmt.Println(stack)
	// fmt.Println(queue)
	for len(stack) > 0 {
		tmp = pop(&stack)
		if tmp != "(" {
			queue = append(queue, tmp)
		} else {
			break
		}
	}
	ret = sumASimpleEprs(queue)

	return ret
}

func main() {
	fmt.Println(calculate("1 - 11"))
	fmt.Println(calculate("12 + 1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate(" 2-(1 + 2) "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("(1+2)+(1+2)"))
	fmt.Println(calculate("123123123"))
}
