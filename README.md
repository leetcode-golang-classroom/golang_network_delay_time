# golang_network_delay_time

You are given a network of `n` nodes, labeled from `1` to `n`. You are also given `times`, a list of travel times as directed edges `times[i] = (ui, vi, wi)`, where `ui` is the source node, `vi` is the target node, and `wi` is the time it takes for a signal to travel from source to target.

We will send a signal from a given node `k`. Return *the **minimum** time it takes for all the* `n` *nodes to receive the signal*. If it is impossible for all the `n` nodes to receive the signal, return `-1`.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2019/05/23/931_example_1.png](https://assets.leetcode.com/uploads/2019/05/23/931_example_1.png)

```
Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2

```

**Example 2:**

```
Input: times = [[1,2,1]], n = 2, k = 1
Output: 1

```

**Example 3:**

```
Input: times = [[1,2,1]], n = 2, k = 2
Output: -1

```

**Constraints:**

- `1 <= k <= n <= 100`
- `1 <= times.length <= 6000`
- `times[i].length == 3`
- `1 <= ui, vi <= n`
- `ui != vi`
- `0 <= wi <= 100`
- All the pairs `(ui, vi)` are **unique**. (i.e., no multiple edges.)

## 解析

題目給定一個矩陣 times, 一個整數 n, 還有一個整數 k

其中每個 entry,  times[i] = [$u_i, v_i, w_i]$ 代表 從在一條從 $u_i$ 到 $v_i$ 的路徑 且花費是 $w_i$

n 代表具有 label 1 到 n 個 vertex

k 代表從 label k 的 vertex 出發

要求寫一個演算法來找出從 k 發送封包:

如果可以傳送封包到所有 vertex的話, 封包傳達到所有 vertex 所需花費的最少時間

如果不能傳送到所有 vertex 則回傳 -1

這題的困難點在於如何找到最小花費路徑

首先需要透過 times 矩陣來建立 adjacencyList 

然後 透過 **[Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)**

因為封包可以同時運發

所以直覺的作法是對 adjacencyList 做 BFS

透過 minHeap 找出當下鄰近的 vertex 中篩選出要找最小花費的邊

並且把累計 weight 的最大值紀錄下來

為了避免重複走，要有一個 HashSet visit 來紀錄走訪過的 vertex

假設走完所有 adjacencyList 的 vertex 

假設 visit 的長度 是 n 代表可以走完所有 vertex

否則就是沒有辦法走完


![](https://i.imgur.com/Hm2HNTX.png)

## 程式碼
```go
package sol

import "container/heap"

type AdjacentNode struct {
	Weight, Node int
}
type AdjacentMinHeap []AdjacentNode

func (h *AdjacentMinHeap) Len() int {
	return len(*h)
}
func (h *AdjacentMinHeap) Less(i, j int) bool {
	return (*h)[i].Weight < (*h)[j].Weight
}
func (h *AdjacentMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *AdjacentMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *AdjacentMinHeap) Push(value interface{}) {
	*h = append(*h, value.(AdjacentNode))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func networkDelayTime(times [][]int, n int, k int) int {
	// create adjacencyList
	adjacencyMap := make(map[int]AdjacentMinHeap)
	for _, t := range times {
		source := t[0]
		target := t[1]
		weight := t[2]
		adjacencyMap[source] = append(adjacencyMap[source], AdjacentNode{Weight: weight, Node: target})
	}
	time := 0
	visit := make(map[int]struct{})
	// start from k
	priorityQueue := &AdjacentMinHeap{AdjacentNode{Weight: 0, Node: k}}
	heap.Init(priorityQueue)
	// Dijkstra's algorithm
	for priorityQueue.Len() != 0 {
		node := heap.Pop(priorityQueue).(AdjacentNode)
		if _, ok := visit[node.Node]; ok {
			continue
		}
		visit[node.Node] = struct{}{}
		time = max(time, node.Weight)
		adjList := adjacencyMap[node.Node]
		for _, adjNode := range adjList {
			if _, ok := visit[adjNode.Node]; !ok {
				heap.Push(priorityQueue, AdjacentNode{Weight: node.Weight + adjNode.Weight, Node: adjNode.Node})
			}
		}
	}
	if len(visit) == n {
		return time
	}
	return -1
}

```
## 困難點

1. 理解 **[Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) 如何尋找最小花費路徑**
2. 理解 MinHeap

## Solve Point

- [x]  透過給定的 times 來建立 adjacencyList
- [x]  透過 **[Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm) 來找尋花費 最少的路徑**
- [x]  透過 MinHeap 與 BFS 來實作 **[Dijkstra's algorithm](https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)**