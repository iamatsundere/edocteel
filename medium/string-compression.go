// https: //leetcode.com/problems/string-compression/
package main

import "fmt"

func compress(chars []byte) int {
	ret := 0
	count := 1
	s := ""

	for i := range chars {
		if i != 0 {
			if chars[i] != chars[i-1] {
				if count != 1 {
					ret += (len(fmt.Sprintf("%d", count)) + 1)
					s += string(chars[i-1]) + fmt.Sprintf("%d", count)
				} else {
					ret++
					s += string(chars[i-1])
				}
				count = 1
			} else {
				count++
			}
		}
		// fmt.Println(string(chars[i]), l, count)
	}

	if count != 1 {
		ret += (len(fmt.Sprintf("%d", count)) + 1)
		s += string(chars[len(chars)-1]) + fmt.Sprintf("%d", count)
	} else {
		ret++
		s += string(chars[len(chars)-1])
	}

	copy(chars, []byte(s))
	// fmt.Println(chars, s)
	return ret
}

func main() {
	fmt.Println(compress([]byte("aabbcc")))
	fmt.Println(compress([]byte("abbbbbbbbbbbb")))
	fmt.Println(compress([]byte("Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of \"de Finibus Bonorum et Malorum\" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, \"Lorem ipsum dolor sit amet..\", comes from a line in section 1.10.32.")))
}
