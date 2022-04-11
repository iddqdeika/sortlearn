package sortlearn

import "math/rand"


func generate(l int) []int{
	res := make([]int, l)
	for i := 0; i < l; i++ {
		res[i] = rand.Intn(100)
	}
	return res
}
