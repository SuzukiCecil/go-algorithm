/*
 * 動的計画法を用いた最長増加部分列（LIS）の解法アルゴリズム
 * 数列A_0, A_1, ..., A_nから0<=i_0<i_1<...<i_k<nかつa_i_0<a_i_1<...<a_i_k を満たす増加部分列のうち、最長の増加部分列の長さを算出する
 * 例えば数列が[63 99 29 29 80 6 41 13 19 33] の場合、[6 13 19 33]が最長増加部分列となり、長さは4となる
 * ここでは数列の長さを10^5, 数列の各要素の最大値を10^9とし、数列は乱数により生成するものとする
 */
package main

import (
	"fmt"
	"math/rand"
)

// 数列Aの長さ
const N = 100_000

// 数列の各要素の最大値
const MAX = 1_000_000_000

func main() {
	// 数列A
	sequence := initSequence()
	// 動的計画法に用いるメモ領域
	dp := [N]int{}
	// 最長増加部分列の長さ
	max := 0
	for i := 0; i < N; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if sequence[i] > sequence[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	fmt.Println(max)
}

/*
 * 数列Aを生成する関数
 */
func initSequence() [N]int {
	sequence := [N]int{}
	for i := 0; i < N; i++ {
		sequence[i] = rand.Intn(MAX)
	}
	return sequence
}
