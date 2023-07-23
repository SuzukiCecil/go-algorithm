/*
 * クイックソートを採用したソートアルゴリズム
 * クイックソートによりソートした配列とソートパッケージによりソートした配列を比較し、ソート結果の成否を表示する
 *
 * ＜出力例＞
 * Before: [88 93 87 41 90 25 80 46 56 52 81 32 29 72 8 58 86 28 1 96 14 91 25 36 38 24 19 98 23 22 3 86 16 67 91 38 92 5 54 55 58 42 74 96 20 19 95 1 11 84 78 72 42 84 90 89 99 28 99 32 96 98 48 67 38 66 35 74 72 66 96 2 86 53 66 60 84 46 36 5 41 27 72 21 11 46 87 27 51 86 62 57 19 13 56 47 18 61 41 58]
 * Sort by quick sort: [1 1 2 3 5 5 8 11 11 13 14 16 18 19 19 19 20 21 22 23 24 25 25 27 27 28 28 29 32 32 35 36 36 38 38 38 41 41 41 42 42 46 46 46 47 48 51 52 53 54 55 56 56 57 58 58 58 60 61 62 66 66 66 67 67 72 72 72 72 74 74 78 80 81 84 84 84 86 86 86 86 87 87 88 89 90 90 91 91 92 93 95 96 96 96 96 98 98 99 99]
 * Sort by sort package: [1 1 2 3 5 5 8 11 11 13 14 16 18 19 19 19 20 21 22 23 24 25 25 27 27 28 28 29 32 32 35 36 36 38 38 38 41 41 41 42 42 46 46 46 47 48 51 52 53 54 55 56 56 57 58 58 58 60 61 62 66 66 66 67 67 72 72 72 72 74 74 78 80 81 84 84 84 86 86 86 86 87 87 88 89 90 90 91 91 92 93 95 96 96 96 96 98 98 99 99]
 * Equals sort result: true
 */
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// ソート前の配列
	list := initList()
	// クイックソートによりソートされた配列
	sortedList := list
	quickSort(0, len(sortedList)-1, &sortedList)

	fmt.Println("Before:", list)
	fmt.Println("Sort by quick sort:", sortedList)

	sort.Ints(list[:])
	fmt.Println("Sort by sort package:", list)

	fmt.Println("Equals sort result:", list == sortedList)
}

/*
 * 乱数によりソート前の配列を生成する関数
 */
func initList() [100]int {
	list := [100]int{}
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(100)
	}
	return list
}

func quickSort(start int, end int, list *[100]int) {
	if end <= start || len(list) < 2 {
		return
	}

	pivot := list[int((start+end)/2)]

	left := start
	right := end

	for {
		for {
			if list[left] >= pivot {
				break
			}
			left++
		}
		for {
			if list[right] <= pivot {
				break
			}
			right--
		}
		if right <= left {
			break
		}

		list[left], list[right] = swap(list[left], list[right])

		left++
		right--
	}

	if start < left-1 {
		quickSort(start, left-1, list)
	}
	if right+1 < end {
		quickSort(right+1, end, list)
	}
}

/*
 * 値の入れ替えを行う関数
 */
func swap(a int, b int) (int, int) {
	return b, a
}
