package tsp

import (
	"math"
	"sync"
)

const Inf = math.MaxFloat64

type node struct {
	index int
	dist  float64
	pred  int
	next  *node
}

func Add(head *node, index int, dist float64, pred int) *node {
	newNode := node{index: index, dist: dist, pred: pred}
	if head == nil {
		head = &newNode
	} else {
		temp := head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &newNode
	}
	return head
}

func Delete(head *node, min int) *node {
	temp := head
	prev := head
	if head.index == min {
		head = head.next
		return head
	}
	for temp != nil {
		if temp.index == min {
			prev.next = temp.next
			temp.next = nil
			temp = nil
			break
		}
		prev = temp
		temp = temp.next
	}
	return head
}

// ComputeLocalMin is a Helper function for SingleSourceShortestPath function.
// It returns the minimum of all the local minimas of the given vertices
//
//  Input   : A *node, head of list containing all vertices in cluster.
//  Outputs : A *node containing the Local minimum.
func ComputeLocalMin(mem *node) *node {
	minNode := mem
	temp := mem
	for temp != nil {
		if temp.dist < minNode.dist {
			minNode = temp
		}
		temp = temp.next
	}
	return minNode
}

// ComputeGlobalMin is a Helper function for Dijkstra function.
// It returns the minimum of all the local minimas of the given vertices
//
//  Input   : A []*node slice containing all the Local minimum computed parallely
//	     			by Go routines.
//  Outputs : A *node containing the global minimum.
func ComputeGlobalMin(local []*node) *node {
	minNode := &node{index: -1, dist: Inf}
	for _, v := range local {
		if v != nil {
			if minNode.dist > v.dist {
				minNode = v
			}
		}
	}
	return minNode
}

// ClusterAnalysis function is a Helper function to SingleSourceShortestPath function.
// Input  : A *node, head of the Linked list containing the nodes of Cluster
//	    A single int ,containing source vertex.
//	    A chan *node used to send the Local minimum to the main thread
//          A chan int to send the signal to confirm that go routines exited
//	    A [][]float64 slice, representing the graph
func ClusterAnalysis(members *node, source int,
	ch chan *node, quit chan int, matrix [][]float64) {
	s := source
	var d float64
	d = 0
	for {
		temp := members
		var v float64
		var i int
		for temp != nil {
			v = temp.dist
			i = temp.index
			if matrix[s][i] != 0 && matrix[s][i]+d < v {
				temp.dist = matrix[s][i] + d
				temp.pred = s
			}
			temp = temp.next
		}
		minNode := ComputeLocalMin(members)
		ch <- minNode
		gMin := <-ch
		if minNode != nil {
			if minNode.index == gMin.index {
				members = Delete(members, minNode.index)
			}
		}
		if gMin.index == -1 {
			quit <- 0
			break
		}
		s = gMin.index
		d = gMin.dist
	}
}

// SingleSourceShortestPath is a helper function for Dijkstra function
// The function returns the least weight and corresponding path from a source
// vertex to destination vertex in the graph.
//
//  Input   : A single int, representing the source vertex
//	     A [][]float64 slice ,representing the Adjacency Matrix of the Graph
//  Outputs : A [] float64, containing the minimum distances from source to all vertices
//	     A [][]int slice containing the minimum paths from source to all vertices
func SingleSourceShortestPath(source int, matrix [][]float64) ([]float64, [][]int) {
	partitions := 4
	V := len(matrix)
	ch := make([]chan *node, partitions)
	quit := make(chan int)
	members := make([]*node, partitions)
	cluster := make(map[int]*node)
	cluster[source] = &node{index: source, dist: 0}
	var i int
	for i = 0; i < partitions; i++ {
		ch[i] = make(chan *node)
	}
	for i = 0; i < V; i++ {
		if i != source {
			members[i%partitions] = Add(members[i%partitions], i, Inf, -1)
		}
	}
	for i = 0; i < partitions; i++ {
		go ClusterAnalysis(members[i], source, ch[i], quit, matrix)
	}
	gMin := make([]*node, partitions)
	pred := make(map[int]int)
	pred[source] = -1
	c := 0
	var m *node
	for {
		for i = 0; i < partitions; i++ {
			gMin[i] = <-ch[i]
		}
		m = ComputeGlobalMin(gMin)
		cluster[m.index] = m
		pred[m.index] = m.pred
		for i = 0; i < partitions; i++ {
			ch[i] <- m
		}
		if m.index == -1 {
			break
		}
	}
	for {
		select {
		case <-quit:
			c++
		default:
		}
		if c == partitions {
			break
		}
	}
	var minWeights []float64 = make([]float64, len(matrix))
	var minPath [][]int = make([][]int, len(matrix))
	var pre int
	for i = 0; i < len(matrix); i++ {
		minWeights[i] = cluster[i].dist
		var path []int
		pre = i
		for {
			path = append([]int{pre}, path...)
			pre = pred[pre]
			if pre == -1 {
				break
			}
		}
		minPath[i] = path
	}
	return minWeights, minPath
}

// Dijkstra parses the graph, and returns the least weight of all paths between any
// two vertices in the graph
//  Input   : A [][]float64 slice, representing the Adjacency Matrix of the graph
//  Outputs : A [][]float64 slice, containing the minimum path weights between any two vertices
//	     A [][][]int slice, containing the minimum paths between any two vertices
func Dijkstra(matrix [][]float64) ([][]float64, [][][]int) {
	var minGraph [][]float64 = make([][]float64, len(matrix))
	var minPaths [][][]int = make([][][]int, len(matrix))
	var nodeWg sync.WaitGroup

	nodeWg.Add(len(matrix))
	for src := 0; src < len(matrix); src++ {
		go func(src int) {
			var minpath [][]int
			minGraph[src], minpath = SingleSourceShortestPath(src, matrix)
			minPaths[src] = minpath
			nodeWg.Done()
		}(src)
	}
	nodeWg.Wait()
	return minGraph, minPaths
}
