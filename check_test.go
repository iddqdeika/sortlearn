package sortlearn

import "testing"

var (
	correctSets = [][]int{
		{1,2,3,4,5},
		{1,5,7,8},
		{-999,0,1111},
	}
	incorrectSets = [][]int{
		{1,3,2},
		{-999,0,999,1},
	}
)

func TestCheck(t *testing.T) {
	for _, set := range correctSets {
		if !isSorted(set){
			t.Errorf("must be check result true for %v", set)
		}
	}
	for _, set := range incorrectSets {
		if isSorted(set){
			t.Errorf("must be check result true for %v", set)
		}
	}
}
