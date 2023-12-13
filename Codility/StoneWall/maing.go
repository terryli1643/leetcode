package main

import (
	"container/list"
	"fmt"
)

/*
You are going to build a stone wall. The wall should be straight and N meters long, and its thickness should be constant; however, it should have different heights in different places. The height of the wall is specified by an array H of N positive integers. H[I] is the height of the wall from I to I+1 meters to the right of its left end. In particular, H[0] is the height of the wall's left end and H[N−1] is the height of the wall's right end.

The wall should be built of cuboid stone blocks (that is, all sides of such blocks are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

Write a function:

func Solution(H []int) int

that, given an array H of N positive integers specifying the height of the wall, returns the minimum number of blocks needed to build it.

For example, given array H containing N = 9 integers:

  H[0] = 8    H[1] = 8    H[2] = 5
  H[3] = 7    H[4] = 9    H[5] = 8
  H[6] = 7    H[7] = 4    H[8] = 8
the function should return 7. The figure shows one possible arrangement of seven blocks.



Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array H is an integer within the range [1..1,000,000,000].
*/

func main() {
	A := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}
	fmt.Println(Solution(A))
}

func Solution(H []int) int {
	// 8 0 5 2 2 1 0 4 4
	// 4 8

	queue := list.New()

	for _, v := range H {
		if queue.Back() == nil {
			queue.PushBack(v)
			continue
		}

		if v > queue.Back().Value.(int) {
			queue.PushBack(v)
			continue
		}
	loop:
		if queue.Back() != nil {
			e := queue.Remove(queue.Back())
			if e.(int) > v {
				goto loop
			}
		}
		queue.PushBack(v)
	}
	return len(H) - queue.Len()

	// var mem []int

	// for i := range H {
	// 	if len(mem) == 0 && H[i] > 0 {
	// 		mem = append(mem, H[i])
	// 		continue
	// 	}
	// 	diff := H[i]
	// 	if H[i] < H[i-1] {
	// 		mem = append(mem, H[i])
	// 		continue
	// 	}

	// 	for j := i - 1; j >= 0; j-- {
	// 		if diff-H[j] >= 0 {
	// 			diff = diff - H[j]
	// 			continue
	// 		}
	// 		break
	// 	}
	// 	if diff > 0 {
	// 		mem = append(mem, diff)
	// 	}
	// }
	// return len(mem)
}
