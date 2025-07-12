package main

import (
	"fmt"
)

func maxActiveUsers(strArr []string) int {
	subArr := make([][]int, len(strArr))
	for i, v := range strArr {
		_, startByte, _, _, endByte, _ := v[0], v[1], v[2], v[3], v[4], v[5]
		start := startByte - '0'
		end := endByte - '0'
		subArr[i] = []int{int(start), int(end)}
	}

	n := len(subArr)
	maxArr := make([]int, n)

	for i, t := range subArr {
		tStart := t[0]
		tEnd := t[1]
		for _, sub := range subArr {
			start := sub[0]
			end := sub[1]
			if start < tEnd && end >= tStart {
				maxArr[i]++
			}
		}
	}

	var max int
	for _, val := range maxArr {
		if val > max {
			max = val
		}
	}

	return max
}

func main() {
	strArr := []string{"[1, 5]", "[2, 3]", "[4, 8]", "[5, 10]", "[6, 12]"}

	fmt.Println(maxActiveUsers(strArr))
}
