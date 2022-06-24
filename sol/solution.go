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
