// https: //leetcode.com/problems/reorganize-string/
package main

import (
	"fmt"
	"sort"
)

func sortMap(amap map[rune]int) PairList {
	pl := make(PairList, len(amap))
	i := 0
	for k, v := range amap {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func reorganizeString(S string) string {
	ret := ""
	chars := []rune(S)
	pairs := make(map[rune]int)

	for _, char := range chars {
		pairs[char] = pairs[char] + 1
	}

	L := len(S)
	maxTimeOfChar := L / 2
	if L%2 != 0 {
		maxTimeOfChar++
	}

	flag := true
	for _, v := range pairs {
		if v > maxTimeOfChar {
			flag = false
			break
		}
	}

	pairsList := sortMap(pairs)

	if flag {
		intsx := make([]int, L)
		i := 0
		for _, pair := range pairsList {
			c := pair.Key
			v := pair.Value
			for v > 0 {
				if intsx[i] == 0 {
					k := int(c)
					if i == 0 || intsx[i-1] != k {
						intsx[i] = k
						v--
					}
				}
				i++
				if i == L {
					i = 0
				}
			}
		}

		charsx := make([]rune, 0)
		for _, v := range intsx {
			charsx = append(charsx, rune(v))
		}
		ret = string(charsx)
	}

	return ret
}

func main() {
	fmt.Println(reorganizeString("aaabbbb"))
}
