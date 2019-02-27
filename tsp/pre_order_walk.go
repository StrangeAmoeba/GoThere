package tsp

import (
	dt "concurrency-9/data_types"
	// "fmt"
)

func pre_order_walk(edges []dt.Graph_edge) []int {
	var edge_sym []dt.Graph_edge
	for i := 0; i < len(edges); i++ {
		edge_sym = append(edge_sym, dt.Graph_edge{edges[i].Dst, edges[i].Src, edges[i].Weight})
		edge_sym = append(edge_sym, edges[i])
	}
	tr := pre_order_node(edge_sym, 0)
	return tr
}

func find_in_slice(to_find int, slice []int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == to_find {
			return true
		}
	}
	return false
}

func pre_order_node(edges []dt.Graph_edge, node int) []int {
	var done []int
	var tr []int
	tr = append(tr, node)
	done = append(done, node)
	for i := 0; i < len(edges); i++ {
		if edges[i].Src == node && !find_in_slice(edges[i].Dst, done) {
			pre_order_node(edges, edges[i].Dst)
		}
	}
	return tr
}
