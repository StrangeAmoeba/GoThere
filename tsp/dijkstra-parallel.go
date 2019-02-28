package tsp

import (
	"math"
)

const inf = math.MaxFloat64

type node struct{
	index int
	dist float64
	pred int
	next *node
}

//linked list functions
func add(head *node,index int,dist float64,pred int) (*node){
	new_node := node{index:index,dist:dist,pred:pred}
	if head == nil{
		head = &new_node
	}else{
		temp := head
		for temp.next != nil{
			temp = temp.next
		}
		temp.next = &new_node
	}
	return head
}

func delete(head *node,min int) (*node){
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

func compute_min(mem *node) (*node){
	min_node := mem
	temp := mem
	for temp != nil{
		if temp.dist < min_node.dist{min_node = temp}
		temp = temp.next
	}
	return min_node
}

//function to find the global minimum
func find_min(local []*node) (*node){
	min_node := &node{dist:inf}
	for _,v := range local{
		if v != nil{
			if min_node.dist > v.dist{
				min_node = v
			}
		}
	}
	return min_node
}

//computing the local minimum in the cluster of vertices
func compute(members *node,source int,
	dest int,ch chan *node,quit chan int,adj_matrix [][] float64){
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
			if adj_matrix[s-1][i-1] != 0 && adj_matrix[s-1][i-1] + d < v{
				temp.dist = adj_matrix[s-1][i-1] + d
				temp.pred = s
			}
			temp = temp.next
		}
		min_node := compute_min(members)
		ch <- min_node
		g_min := <- ch
		cluster = add(cluster,g_min.index,g_min.dist,g_min.pred)
		if min_node != nil{
			if min_node.index == g_min.index{
				members = delete(members,min_node.index)
			}
		}
		if g_min.index == dest || g_min.index == 0{
			quit <- 0
			break
		}
		s = g_min.index
		d = g_min.dist
	}	
}

func singleSourceShortestPath(source int,dest int,adj_matrix [][]float64) (map[int]int,float64){
	num_partitions = 4
	var mem *node
	ch := make([]chan *node,num_partitions)
	quit := make(chan int)
	members := make([]*node,num_partitions)
	for i = 0 ;i < num_partitions ; i++{
		ch[i] = make(chan *node)
	}
	for i = 1 ; i <= V ;i++{
		if i != source{
			members[i%num_partitions] = add(members[i%num_partitions],i,inf,-1)
			mem = add(mem,i,inf,-1)
		}
	}
	for i = 0 ; i< num_partitions ; i++{
		go compute(members[i],source,dest,ch[i],adj_matrix,quit,preds[i])
	}
	g_min := make([]*node,num_partitions)
	pred = make(map[int]int)
	c := 0
	var m *node
	// computing global minimum 
	for {
		for i = 0 ; i< num_partitions ;i++{
			g_min[i] = <-ch[i]
		}
		m = find_min(g_min)
		pred[m.index] = m.pred
		for i = 0 ; i< num_partitions ;i++{ // broadcasting it to all the nodes
			ch[i] <- m
		}
		if m.index == dest || m.index == 0{
			break
		}
	}
	for{
		select{
		case <- quit:
			c++
		default:
		}
		if c == 3{break}
	}
	return pred,m.dist //returning minimum path and minimum distance for a given source vertex pair
}
