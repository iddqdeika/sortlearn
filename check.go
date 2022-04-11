package sortlearn

func isSorted(s []int) bool{
	for i, v := range s {
		if i>0 && v < s[i-1]{
			return false
		}
	}
	return true
}

const tryCount = 20

func check(size int, f  func (s []int) []int) bool{
	for i := 0; i < tryCount; i++ {
		if !isSorted(f(generate(size))){
			return false
		}
	}
	return true
}