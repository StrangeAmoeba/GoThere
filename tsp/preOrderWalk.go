package tsp

import (
	dt "concurrency-9/dataTypes"
	// "fmt"
)

var done []int
var tr []int

func preOrderWalk(edges []dt.GraphEdge) []int {
	newTr := make([]int, 0)
	newDone := make([]int, 0)
	tr = newTr
	done = newDone
	var edgeSym []dt.GraphEdge
	for i := 0; i < len(edges); i++ {
		edgeSym = append(edgeSym, dt.GraphEdge{edges[i].Dst, edges[i].Src, edges[i].Weight})
		edgeSym = append(edgeSym, edges[i])
	}
	tr := preOrderNode(edgeSym, 0)
	return tr
}

func findInSlice(toFind int, slice []int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == toFind {
			return true
		}
	}
	return false
}

func preOrderNode(edges []dt.GraphEdge, node int) []int {
	tr = append(tr, node)
	done = append(done, node)
	for i := 0; i < len(edges); i++ {
		if edges[i].Src == node && !findInSlice(edges[i].Dst, done) {
			preOrderNode(edges, edges[i].Dst)
		}
	}
	trD := make([]int, len(tr))
	copy(trD, tr)
	// tr, done = nil, nil
	return trD

}
