package main

import "fmt"

var (
	UNCOLORED, RED, GREEN = 0, 1, 2
	color                 []int
	valid                 bool
)

func isBipartite(graph [][]int) bool {
	n := len(graph)
	valid = true
	color = make([]int, n)
	for i := 0; i < n && valid; i++ {
		if color[i] == UNCOLORED {
			dfs(i, RED, graph)
		}
	}
	return valid
}

func dfs(node, c int, graph [][]int) {
	color[node] = c
	cNext := RED
	if c == RED {
		cNext = GREEN
	}
	for _, next := range graph[node] {
		if color[next] == UNCOLORED {
			dfs(next, cNext, graph)
			if !valid {
				return
			}
		} else if color[next] != cNext {
			valid = false
			return
		}
	}
}

func main() {
	graph := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	ans := isBipartite(graph)
	fmt.Println(ans)
}
