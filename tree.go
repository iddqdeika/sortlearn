package sortlearn

func treeSort(s []int) []int{
	n := fill(s)
	res := make([]int, 0, len(s))
	n.collect(&res)
	return res
}

type node struct {
	v int
	lessOrEq *node
	greater *node
}

func (n *node) put(i int){
	if i <= n.v{
		if n.lessOrEq == nil{
			n.lessOrEq = &node{v: i}
			return
		}
		n.lessOrEq.put(i)
		return
	}
	if n.greater == nil{
		n.greater = &node{v: i}
		return
	}
	n.greater.put(i)
	return
}

func fill(s []int) *node{
	r := &node{v: s[0]}
	for _, v := range s[1:] {
		r.put(v)
	}
	return r
}

func (n *node) collect(res *[]int) {
	if n.lessOrEq != nil{
		n.lessOrEq.collect(res)
	}
	*res = append(*res, n.v)
	if n.greater != nil{
		n.greater.collect(res)
	}
}
