package sortlearn

func bubbleSort(s []int) []int{
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if s[j] < s[j-1]{
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	return s
}
