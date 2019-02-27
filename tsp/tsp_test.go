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

func TestKruskals(t *testing.T) {

}
