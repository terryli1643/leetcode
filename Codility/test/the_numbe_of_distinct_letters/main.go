package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
We are given two strings P and Q, each consisting of N lowercase English letters. For each position in the strings, we have to choose one letter from either P or Q, in order to construct a new string S, such that the number of distinct letters in S is minimal. Our task is to find the number of distinct letters in the resulting string S.

For example, if P = "ca" and Q = "ab", S can be equal to: "ca", "cb", "aa" or "ab". String "aa" has only one distinct letter ('a'), so the answer is 1 (which is minimal among those strings).

Write a function:

class Solution { public int solution(String P, String Q); }

that, given two strings P and Q, both of length N, returns the minimum number of distinct letters of a string S, that can be constructed from P and Q as described above.

Examples:

1. Given P = "abc", Q = "bcd", your function should return 2. All possible strings S that can be constructed are: "abc", "abd", "acc", "acd", "bbc", "bbd", "bcc", "bcd". The minimum number of distinct letters is 2, which be obtained by constructing the following strings: "acc", "bbc", "bbd", "bcc".

2. Given P = "axxz", Q = "yzwy", your function should return 2. String S must consist of at least two distinct letters in this case. We can construct S = "yxxy", where the number of distinct letters is equal to 2, and this is the only optimal solution.

3. Given P = "bacad", Q = "abada", your function should return 1. We can choose the letter 'a' in each position, so S can be equal to "aaaaa".

4. Given P = "amz", Q = "amz", your function should return 3. The input strings are identical, so the only possible S that can be constructed is "amz", and its number of distinct letters is 3.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..200,000];
strings P and Q are both of length N;
strings P and Q are made only of lowercase letters (a−z);
strings P and Q contain a total of at most 20 distinct letters.

*/

func main() {
	P := "adabca"
	Q := "cbdcdb"
	fmt.Println(Solution(P, Q))
}

func Solution(P string, Q string) int {
	// Implement your solution here
	m := make(map[string]int)
	for i := 0; i < len(P); i++ {
		if _, ok := m[string(P[i])+string(Q[i])]; ok {
			continue
		}
		if _, ok := m[string(Q[i])+string(P[i])]; ok {
			continue
		}
		m[string(P[i])+string(Q[i])] = 1
	}
	P1, Q1 := "", ""
	for k := range m {
		P1 = P1 + string(k[0])
		Q1 = Q1 + string(k[1])
	}
	Search(P1, Q1, 0, 0, "")
	return min
}

var min = 20
var mem = make(map[int]string)

func Search(P string, Q string, i int, distinctCount int, S string) {
	if i == len(P) {
		if min > distinctCount {
			min = distinctCount
		}
		return
	}
	if distinctCount >= min {
		return
	}
	if v, ok := mem[i]; ok {
		if v == S {
			return
		} else {
			mem[i] = S
		}
	}

	pCount := distinctCount
	if !strings.ContainsAny(S, tidy(string(P[i]))) {
		pCount++
	}

	Search(P, Q, i+1, pCount, tidy(S+string(P[i])))

	qCount := distinctCount

	if !strings.ContainsAny(S, tidy(string(Q[i]))) {
		qCount++
	}
	Search(P, Q, i+1, qCount, tidy(S+string(Q[i])))
}

func tidy(s string) string {
	m := make(map[rune]int)
	for _, v := range s {
		m[v] = 1
	}

	var slice []rune
	for k := range m {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return string(slice)
}
