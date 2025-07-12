package main

import (
	"fmt"
	"strconv"
)

func shortestPath(args []string) string {
	// adjMap := map[rune][]rune{
	// 	'A': {'B'},
	// 	'B': {'A', 'C', 'D'},
	// 	'C': {'B', 'D'},
	// 	'D': {'B', 'C'},
	// }

	n, _ := strconv.Atoi(args[0])

	adjMap := make(map[string][]string, n)
	for i := 1; i <= n; i++ {
		adjMap[args[i]] = make([]string, 0)
	}

	for i := n + 1; i <= n*2; i++ {
		from, _, to := string(args[i][0]), args[i][1], string(args[i][2])
		adjMap[from] = append(adjMap[from], to)
		adjMap[to] = append(adjMap[to], from)
	}

	visited := make(map[string]bool, n)
	pathMap := make(map[string]string, n)

	from, to := args[1], args[n]
	pathMap[from] = from
	traverse(adjMap, pathMap, visited, from, to)

	return pathMap[to]
}

func traverse(adjMap map[string][]string, pathMap map[string]string, visited map[string]bool, from, to string) {
	visited[from] = true

	for _, v := range adjMap[from] {
		if visited[v] {
			continue
		}

		if path, ok := pathMap[v]; ok {
			if len(path) > len(pathMap[from])+3 {
				pathMap[v] = fmt.Sprintf("%s->%s", pathMap[from], v)
			}
		} else {
			pathMap[v] = fmt.Sprintf("%s->%s", pathMap[from], v)
		}

		if v == to {
			return
		}

		traverse(adjMap, pathMap, visited, v, to)
	}
}

func main() {
	args := []string{"5", "A", "B", "C", "D", "E", "A-B", "B-D", "C-D", "D-E", "E-B"}

	fmt.Println(shortestPath(args))
}
