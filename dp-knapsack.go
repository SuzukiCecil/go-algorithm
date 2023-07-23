/*
 * 動的計画法を用いたナップザック問題の解法アルゴリズム
 * 価値と重量が定義されている品物が複数と最大積載重量が定義されているナップザックから、ナップザックに積載可能な品物の価値の合計値の最大を算出するプログラム
 */
package main

import "fmt"

// ナップザックに積載する品物
type Item struct {
	Value  int // 品物の価値
	Weight int // 品物の重量
}

// ナップザックの最大積載重量
const CAPACITY = 30

func main() {
	items := initItems()
	dp := initDp(items)
	for i := 1; i <= len(items); i++ {
		for j := 1; j <= CAPACITY; j++ {
			if j >= items[i-1].Weight {
				dp[i][j] = items[i-1].Value
			} else {
				dp[i][j] = 0
			}

			if dp[i-1][j] > dp[i][j] {
				dp[i][j] = dp[i-1][j]
			}
			if j > items[i-1].Weight && dp[i-1][j-items[i-1].Weight]+items[i-1].Value > dp[i][j] {
				dp[i][j] = dp[i-1][j-items[i-1].Weight] + items[i-1].Value
			}
		}
	}

	fmt.Println(dp[len(items)][CAPACITY])
}

/*
 * 品物を定義し返す関数
 */
func initItems() []Item {
	return []Item{
		Item{Value: 10, Weight: 5},
		Item{Value: 8, Weight: 3},
		Item{Value: 11, Weight: 6},
		Item{Value: 8, Weight: 4},
		Item{Value: 15, Weight: 8},
		Item{Value: 3, Weight: 1},
		Item{Value: 4, Weight: 2},
		Item{Value: 7, Weight: 4},
		Item{Value: 5, Weight: 2},
		Item{Value: 12, Weight: 6},
	}
}

/*
 * 動的計画法に用いるメモ領域を返す関数
 * メモ領域は(品物数+1)*(ナップザックの最大積載重量+1)とする
 */
func initDp(items []Item) [][]int {
	dp := make([][]int, len(items)+1)
	for i := 0; i <= len(items); i++ {
		dp[i] = make([]int, CAPACITY+1)
	}
	return dp
}
