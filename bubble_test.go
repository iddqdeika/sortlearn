package sortlearn

import "testing"

func TestBubble(t *testing.T){
	if !check(10, bubbleSort){
		t.Fatalf("bubblesort is incorrect")
	}
}

func BenchmarkBubble1000(b *testing.B){
	for i := 0; i < b.N; i++ {
		check(1000, bubbleSort)
	}
}

func BenchmarkBubble100(b *testing.B){
	for i := 0; i < b.N; i++ {
		check(100, bubbleSort)
	}
}

func BenchmarkBubble10(b *testing.B){
	for i := 0; i < b.N; i++ {
		check(10, bubbleSort)
	}
}