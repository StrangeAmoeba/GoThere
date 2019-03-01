package tsp

import (
	dt "concurrency-9/dataTypes"
	// "fmt"
)

var done []int
var tr []int

// preOrderWalk converts the mst into a suitable form to
// get its pre order walkm by calling preOrderNode
//
//	Input: edges, the list of edges in the mst i.e. []dataTypes.GraphEdge
//	Output: tr, the preorder walk of the mst i.e. []int
func preOrderWalk(edges []dt.GraphEdge, firstDst int) []int {
	newTr := make([]int, 0)
	newDone := make([]int, 0)
	tr = newTr
	done = newDone
	var edgeSym []dt.GraphEdge
	for i := 0; i < len(edges); i++ {
		edgeSym = append(edgeSym, dt.GraphEdge{edges[i].Dst, edges[i].Src, edges[i].Weight})
		edgeSym = append(edgeSym, edges[i])
	}
	tr := preOrderNode(edgeSym, firstDst)
	return tr
}

// findInSlice finds if a given element belongs to a given slice
//
//  Input: toFind, the element we want to find i.e. int
//		  slice, the slice we want to search in i.e. []int
//  Output: bool, representing existence of the element in the slice.
func findInSlice(toFind int, slice []int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == toFind {
			return true
		}
	}
	return false
}

// preOrderNode is a reecursive function which gets the preorder walk of
// the given node.
//
// Input: edges, the edges in the mst i.e. []dataTypes.GraphEdge
// 	   node, the node we want to compute preoderwalk of i.e. int
// Output: trD, the preorderwalk rooted at the node i.e. []int
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
