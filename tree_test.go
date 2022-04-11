package sortlearn

import "testing"

func TestTree(t *testing.T){
	if !check(10, treeSort){
		t.Fatalf("treesort is incorrect")
	}
}

func BenchmarkTree1000(b *testing.B){
	for i := 0; i < b.N; i++ {
		check(1000, treeSort)
	}
}

func BenchmarkTree100(b *testing.B){
	for i := 0; i < b.N; i++ {
		check(100, treeSort)
	}
}

func BenchmarkTree10(b *testing.B){
	for i := 0; i < b.N; i++ {
		check(10, treeSort)
	}
}