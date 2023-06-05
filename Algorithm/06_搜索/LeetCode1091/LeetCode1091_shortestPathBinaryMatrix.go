package main

import "fmt"

// Array Step1: 定义一个队列，并实现压入、弹出等方法
type Array struct {
	nodes []interface{}
}

type Next struct {
	x int
	y int
}

// 寻求最优解一般使用宽度优先遍历，最先遍历完的即为最优解
func shortestPathBinaryMatrix(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
	if grid[0][0] != 0 || grid[M-1][N-1] != 0 || M == 0 {
		return -1
	}
	//Step2: 定义一个路径集合记录已经遍历过的路径
	path := make([][]bool, M)
	for r, _ := range path {
		path[r] = make([]bool, N)
	}
	q := NewArray()
	//Step3: 将第一个元素压入队列 & 将第一个元素路径加入path
	q.Push([]int{0, 0})
	path[0][0] = true

	ans := 0
	//Step4: 循环遍历直到队列为空
	for !q.IsEmpty() {
		//Step5: 求当前层宽度 & 在当前层中遍历
		size := q.Len()
		for i := 0; i < size; i++ {
			//Step6: 依次弹出当前层元素 & 判断是否触发返回条件
			curNode := q.Pop()
			curX := curNode.([]int)[0]
			curY := curNode.([]int)[1]
			if curX == M-1 && curY == N-1 {
				return ans + 1
			}
			next := []Next{
				{curX + 1, curY}, {curX + 1, curY + 1}, {curX, curY + 1}, {curX - 1, curY + 1},
				{curX - 1, curY}, {curX - 1, curY - 1}, {curX, curY - 1}, {curX + 1, curY - 1},
			}
			//Step7: 将当前元素所连接的下一个[满足条件]的元素依次压入队列 & 记录在path中
			for _, n := range next {
				if 0 <= n.x && n.x < M && 0 <= n.y && n.y < N && path[n.x][n.y] == false && grid[n.x][n.y] == 0 {
					path[n.x][n.y] = true
					q.Push([]int{n.x, n.y})
				}
			}
		}
		//Step8 ans++
		ans++
	}
	return -1
}

func NewArray() *Array {
	return &Array{}
}

func (arr *Array) Pop() interface{} {
	first := arr.nodes[0]
	arr.nodes = arr.nodes[1:]
	return first
}

func (arr *Array) Push(node interface{}) {
	arr.nodes = append(arr.nodes, node)
}

func (arr *Array) Len() int {
	return len(arr.nodes)
}

func (arr *Array) IsEmpty() bool {
	return len(arr.nodes) == 0
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	ans := shortestPathBinaryMatrix(grid)
	fmt.Println(ans)
}
