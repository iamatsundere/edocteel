// https://leetcode.com/problems/find-and-replace-in-string/
package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Index  int
	Source string
	Target string
}

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	ret := ""
	flag := 0

	items := make([]Item, 0)
	for i, index := range indexes {
		items = append(items, Item{index, sources[i], targets[i]})
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Index < items[j].Index
	})

	// fmt.Println(S)
	// fmt.Println(items)

	for _, item := range items {
		ret = ret + S[flag:item.Index]
		source := item.Source
		l := len(source)
		s := S[item.Index : item.Index+l]

		if s == source {
			ret += item.Target
		} else {
			ret += s
		}

		flag = item.Index + l
	}

	if flag < len(S) {
		ret += S[flag:len(S)]
	}

	return ret
}

func main() {
	fmt.Println(findReplaceString("vmokgggqzp", []int{3, 5, 1}, []string{"kg", "ggq", "mo"}, []string{"s", "so", "bfr"}))
}
