package tsp

import (
	dt "concurrency-9/data_types"
	"testing"
)

func TestSort(t *testing.T) {
	cases := []struct {
		in, want []dt.Graph_edge
	}{
		{[]dt.Graph_edge{{-1, -1, 6}, {-1, -1, 5}, {-1, -1, 4}, {-1, -1, 3}, {-1, -1, 2}}, []dt.Graph_edge{{-1, -1, 2}, {-1, -1, 3}, {-1, -1, 4}, {-1, -1, 5}, {-1, -1, 6}}},
	}
	flag := false
	for _, c := range cases {
		got := Merge_Sort(c.in)
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

// func TestKruskals(t *testing.T) {
// 	cases := []struct {
// 		in [][]float64, want []dt.Graph_edge
// 	}{
// 		{
// 			[][]float64{
// 				[]float64{0,4,zz,zz,zz,zz,zz,8,zz},
// 				[]float64{zz,0,8,zz,zz,zz,zz,11,zz},
// 				[]float64{zz,zz,0,7,zz,4,zz,zz,zz},
// 				[]float64{zz,zz,zz,0,9,14,zz,zz,zz},
// 				[]float64{zz,zz,zz,zz,0,zz,zz,zz,zz},
// 				[]float64{zz,zz,zz,zz,10,0,zz,zz,zz},
// 				[]float64{zz,zz,zz,zz,zz,2,0,zz,zz},
// 				[]float64{zz,zz,zz,zz,zz,zz,1,0,7},
// 				[]float64{zz,zz,2,zz,zz,zz,6,zz,0}
// 		}
// 			[]dt.Graph_edge{{-1, -1, 2}, {-1, -1, 3}, {-1, -1, 4}, {-1, -1, 5}, {-1, -1, 6}}
// 		},
// 	}
// 	flag := false
// 	for _, c := range cases {
// 		got := Merge_Sort(c.in)
// 		for i := 0; i < len(got)-1; i++ {
// 			if got[i] != c.want[i] {
// 				flag = true
// 			}
// 		}
// 	}
// 	if flag {
// 		t.Errorf("Failed Sort")
// 		// t.Log()
// 	}
// }
