package main

import "fmt"

func main() {
	A := []int{1, 2, 4, 5}
	B := []int{2, 3, 5, 6}
	N := 4
	fmt.Println(Solution(A, B, N))
}

func Solution(A []int, B []int, N int) int {
	m := make(map[int][]int)

	for i := 0; i < len(A); i++ {
		m[A[i]] = append(m[A[i]], B[i])
		m[B[i]] = append(m[B[i]], A[i])
	}

	max := 0
	for cur, _ := range m {
		mem := make(map[string]int)
		count := Search(cur, m, 0, mem)
		if max < count {
			max = count
		}
	}
	return max
}

func Search(cur int, m map[int][]int, steps int, mem map[string]int) int {
	if len(m[cur]) == 0 {
		return steps
	}
	for _, child := range m[cur] {
		if _, ok := mem[fmt.Sprintf("%d:%d", child, cur)]; ok {
			continue
		}
		mem[fmt.Sprintf("%d:%d", child, cur)] = 1
		mem[fmt.Sprintf("%d:%d", cur, child)] = 1
		step := Search(child, m, steps+1, mem)
		if steps < step {
			steps = step
		}
	}
	return steps
}
