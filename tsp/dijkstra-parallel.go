package tsp

import (
	"math"
	"sync"
)

const Inf = math.MaxFloat64

type node struct{
	index int
	dist float64
	pred int
	next *node
}

func Add(head *node,index int,dist float64,pred int) (*node){
	newNode := node{index:index,dist:dist,pred:pred}
	if head == nil{
		head = &newNode
	}else{
		temp := head
		for temp.next != nil{
			temp = temp.next
		}
		temp.next = &newNode
	}
	return head
}

func Delete(head *node,min int) (*node){
	temp := head
	prev := head
	if head.index == min{
		head = head.next
		return head
	}
	for temp != nil{
		if temp.index == min{
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
// Input   : A *node, head of list containing all vertices in cluster.
// Outputs : A *node containing the Local minimum.
func ComputeLocalMin(mem *node) (*node){
	minNode := mem
	temp := mem
	for temp != nil{
		if temp.dist < minNode.dist{minNode = temp}
		temp = temp.next
	}
	return minNode
}

// ComputeGlobalMin is a Helper function for Dijkstra function.
// It returns the minimum of all the local minimas of the given vertices
// 
// Input   : A []*node slice containing all the Local minimum computed parallely
//			 by Go routines.
// Outputs : A *node containing the global minimum.
func ComputeGlobalMin(local []*node) (*node){
	minNode := &node{index : -1,dist:Inf}
	for _,v := range local{
		if v != nil{
			if minNode.dist > v.dist{
				minNode = v
			}
		}
	}
	return minNode
}

// ClusterAnalysis function is a Helper function to SingleSourceShortestPath function.
// Input  : A *node, head of the Linked list containing the nodes of Cluster
//			A single int ,containing source vertex.
//			A single int ,containing dest vertex
//			A chan *node used to send the Local minimum to the main thread
//			A chan int to send the signal to confirm that go routines exited
//			A [][]float64 slice, representing the graph
func ClusterAnalysis(members *node,source int,
	dest int,ch chan *node,quit chan int,matrix [][] float64){
	//TODO : can use priority queues instead of linked list to store vertices
	s := source
	var d float64
	d = 0
	for {
		temp := members
		var v float64
		var i int
		for temp != nil{
			v = temp.dist
			i = temp.index
			if matrix[s][i] != 0 && matrix[s][i] + d < v{
				temp.dist = matrix[s][i] + d
				temp.pred = s
			}
			temp = temp.next
		}
		minNode := ComputeLocalMin(members)
		ch <- minNode
		globalMin := <- ch
		if minNode != nil{
			if minNode.index == globalMin.index{
				members = Delete(members,minNode.index)
			}
		}
		if globalMin.index == dest{
			quit <- 0
			break
		}
		s = globalMin.index
		d = globalMin.dist
	}	
}

// SingleSourceShortestPath is a helper function for Dijkstra function.
// The function returns the least weight and corresponding path from a source
// vertex to destination vertex in the graph.
//
// Input   : A single int, representing the source vertex
//		     A single int, representing the destination vertex
//		     A [][]float64 slice ,representing the Adjacency Matrix of the Graph
// Outputs : A single float64, containing the minimum distance from source to vertex
//			 A []int slice containing the minimum path from source to destination
func SingleSourceShortestPath(source int,dest int,matrix [][]float64) (float64, []int){
	if(source == dest){
		var p[] int
		p = append(p,source)
		return 0,p	
	}
	partitions := 4
	V := len(matrix)
	ch := make([]chan *node,partitions)
	quit := make(chan int)
	members := make([]*node,partitions)
	var i int
	for i = 0 ;i < partitions ; i++{
		ch[i] = make(chan *node)
	}
	for i = 0 ; i < V ;i++{
		if i != source{
			members[i%partitions] = Add(members[i%partitions],i,Inf,-1)
		}
	}
	for i = 0 ; i< partitions ; i++{
		go ClusterAnalysis(members[i],source,dest,ch[i],quit,matrix)
	}
	globalMin := make([]*node,partitions)
	pred := make(map[int]int)
	pred[source] = -1
	c := 0
	var m *node 
	for {
		for i = 0 ; i< partitions ;i++{
			globalMin[i] = <-ch[i]
		}
		m = ComputeGlobalMin(globalMin)
		pred[m.index] = m.pred
		for i = 0 ; i< partitions ;i++{ 
			ch[i] <- m
		}
		if m.index == dest{
			break
		}
	}
	for{
		select{
		case <- quit:
			c++
		default:
		}
		if c == partitions{break}
	}
	var path[]int
	pre := dest
	for{
		path = append([]int{pre},path...)
		pre = pred[pre]
		if pre == -1{break}
	}
	return m.dist,path
}

// Dijkstra parses the graph, and returns the least weight of all paths between any
// two vertices in the graph
// Input   : A [][]float64 slice, representing the Adjacency Matrix of the graph
// Outputs : A [][]float64 slice, containing the minimum path weights between any two vertices
//			 A [][][]int slice, containing the minimum paths between any two vertices
func Dijkstra(matrix [][]float64) ([][]float64,[][][]int){
	var minGraph [][]float64 = make([][]float64, len(matrix))
	var minPaths [][][]int = make([][][]int,len(matrix))
	var nodeWg sync.WaitGroup

	nodeWg.Add(len(matrix)*len(matrix))
	for src := 0 ; src < len(matrix) ;src++{ 
		minGraph[src] = make([]float64,len(matrix))
		minPaths[src] = make([][]int,len(matrix))
		for dest := 0 ;dest < len(matrix) ;dest++{
			go func(src int,dest int) {
				var minpath[]int 
				minGraph[src][dest],minpath = SingleSourceShortestPath(src,dest,matrix)
				minPaths[src][dest] = minpath
				nodeWg.Done()
			}(src,dest)
		}
	}
	nodeWg.Wait()
	return minGraph,minPaths
}
