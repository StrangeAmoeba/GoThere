package tsp

import (
	dt "concurrency-9/dataTypes"
	"math"
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		in, want []dt.GraphEdge
	}{
		{[]dt.GraphEdge{{-1, -1, 6}, {-1, -1, 5}, {-1, -1, 4}, {-1, -1, 3}, {-1, -1, 2}}, []dt.GraphEdge{{-1, -1, 2}, {-1, -1, 3}, {-1, -1, 4}, {-1, -1, 5}, {-1, -1, 6}}},
	}
	flag := false
	for _, c := range cases {
		got := MergeSort(c.in)
		for i := 0; i < len(got)-1; i++ {
			if got[i] != c.want[i] {
				flag = true
			}
		}
	}
	if flag {
		t.Errorf("Failed Sort")
		// t.Log()
	}
}

func TestKruskals(t *testing.T) {
	cases := []struct {
		in   [][]float64
		want []dt.GraphEdge
	}{
		{
			[][]float64{
				[]float64{0, 4, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, 8, math.MaxFloat64},
				[]float64{4, 0, 8, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, 11, math.MaxFloat64},
				[]float64{math.MaxFloat64, 8, 0, 7, math.MaxFloat64, 4, math.MaxFloat64, math.MaxFloat64, 2},
				[]float64{math.MaxFloat64, math.MaxFloat64, 7, 0, 9, 14, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
				[]float64{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, 9, 0, 10, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64},
				[]float64{math.MaxFloat64, math.MaxFloat64, 4, 14, 10, 0, 2, math.MaxFloat64, math.MaxFloat64},
				[]float64{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, 2, 0, 1, 6},
				[]float64{8, 11, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, 1, 0, 7},
				[]float64{math.MaxFloat64, math.MaxFloat64, 2, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, 6, 7, 0},
			},
			[]dt.GraphEdge{
				{0, 1, 4},
				{0, 7, 8},
				{7, 6, 1},
				{6, 5, 2},
				{5, 2, 4},
				{2, 8, 2},
				{3, 2, 7},
				{3, 4, 9},
			},
		},
	}
	flag := false
	for _, c := range cases {
		got := Kruskals(c.in)
		for _, edge := range got {
			if !findEdge(edge, c.want) {
				flag = true
			}
		}
	}
	if flag {
		t.Errorf("%v\n%v", Kruskals(cases[0].in), cases[0].want)
		// t.Log()
	}
}

func findEdge(toFind dt.GraphEdge, slice []dt.GraphEdge) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == toFind || (dt.GraphEdge{slice[i].Dst, slice[i].Src, slice[i].Weight} == toFind) {
			return true
		}
	}
	return false
}
