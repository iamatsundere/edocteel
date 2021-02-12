// https://leetcode.com/problems/destination-city/
package main

func destCity(paths [][]string) string {
	ret := ""
	pairs := make(map[string]bool)
	for _, path := range paths {
		pairs[path[0]] = true
	}
	for _, path := range paths {
		if !pairs[path[1]] {
			ret = path[1]
			break
		}
	}
	return ret
}

func main() {
}
