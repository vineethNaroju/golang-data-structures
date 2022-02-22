package main

import (
	"fmt"
	"math/rand"
)

type Ufds struct {
	n      int
	parent []int
}

func NewUfds(N int) *Ufds {
	uf := &Ufds{
		n: N,
	}

	for i := 0; i < N; i++ {
		uf.parent = append(uf.parent, i)
	}

	return uf
}

func (uf *Ufds) GetParent(idx int) int {
	temp := idx

	stk := []int{}

	for uf.parent[temp] != temp {
		stk = append(stk, temp)
		temp = uf.parent[temp]
	}

	for _, val := range stk {
		uf.parent[val] = temp
	}

	return temp
}

func (uf *Ufds) Merge(a, b int) {
	if a > b {
		a, b = b, a
	}

	pa, pb := uf.GetParent(a), uf.GetParent(b)

	uf.parent[pb] = pa
}

func (uf *Ufds) Size() int {
	mp := make(map[int]int)

	for _, val := range uf.parent {
		if _, ok := mp[val]; ok {
			mp[val]++
		} else {
			mp[val] = 1
		}
	}

	return len(mp)
}

func main() {
	N := 100

	rand := func() int {
		return rand.Intn(N)
	}

	ufds := NewUfds(N)

	prev, curr := 0, 0

	for i := 0; i < N*N; i++ {
		curr = rand()

		ufds.Merge(prev, curr)

		fmt.Println(ufds.Size())

		prev = curr
	}

}
