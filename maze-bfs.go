/*
 * 幅優先探索(Breadth First Search)を採用した迷路探索アルゴリズム
 * <要件>
 * 左右（X方向に+-1）と上下（Y方向に+-1）のマスに移動可能
 * スタートからゴールまでの順路全てを表示
 * 一度訪れたマスに再度訪れることは不可能
 * スタートはx=0, y=0のマス、ゴールはx=MAX_X, y=MAX_Yのマスとする
 * x, yともに最大値より大きいマスおよび最小値より小さいマスへの進行は不可とする
 */

package main

import "fmt"

type Location struct {
	X int
	Y int
}

// 現在座標と現在座標に至るまでの経路の組み合わせ
type State struct {
	currentLocation Location
	course          []Location
}

// 進行可能なマス
const FLAT = 0

// 進行不可能なマス
const WALL = 1

// X軸の最小値
const MIN_X = 0

// Y軸の最小値
const MIN_Y = 0

// X軸の最大値
const MAX_X = 4

// Y軸の最大値
const MAX_Y = 4

// 迷路のフィールド
var field = [][]int{
	{FLAT, FLAT, FLAT, FLAT, FLAT},
	{FLAT, WALL, FLAT, WALL, FLAT},
	{FLAT, WALL, FLAT, WALL, FLAT},
	{FLAT, FLAT, FLAT, WALL, WALL},
	{FLAT, WALL, FLAT, FLAT, FLAT},
}

// 現在座標と移動先座標の差（0: X軸, 1: Y軸）
var diffs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func main() {
	start := Location{MIN_X, MIN_Y}
	visit(
		append(
			[]State{},
			State{
				currentLocation: start,
				course:          append([]Location{}, start),
			},
		),
	)
}

/*
 * 迷路探索を行う再帰関数
 * @param current Location 現在地の座標
 * @param course []Location 現在地に至るまでに辿ってきた座標
 */
func visit(queue []State) {
	// デキュー処理
	currentState := queue[0]
	queue = queue[1:]

	// 現在地がゴールである場合はゴールまでの順路を表示し、それ以外の場合は現在地から移動可能な座標をエンキューする
	if currentState.currentLocation.X == MAX_X && currentState.currentLocation.Y == MAX_Y {
		fmt.Println("<course>")
		for _, location := range currentState.course {
			fmt.Println("X:", location.X, "Y:", location.Y)
		}
	} else {
		for _, diff := range diffs {
			// 移動先の座標
			nextLocation := Location{currentState.currentLocation.X + diff[0], currentState.currentLocation.Y + diff[1]}
			// 移動先の座標へ移動可能な場合は移動先および、移動先までの経路をエンキューする
			if isMoveable(nextLocation, currentState.course) {
				queue = append(
					queue,
					State{
						currentLocation: nextLocation,
						course:          append(currentState.course, nextLocation),
					},
				)
			}
		}
	}

	if len(queue) > 0 {
		visit(queue)
	}
}

/*
 * 指定のマスに対して移動可能か判定する関数
 * @param target Location 判定対象のマス
 * @param course []Location 現在地に至るまでに辿ってきた座標
 * @return bool （true：移動可, false：移動不可）
 */
func isMoveable(target Location, course []Location) bool {
	// 迷路のフィールド外への移動は不可
	if target.X < MIN_X || target.X > MAX_X || target.Y < MIN_Y || target.Y > MAX_Y {
		return false
	}
	// 指定されたマスが進行不可能の場合は移動不可
	if field[target.Y][target.X] == WALL {
		return false
	}
	// 指定されたマスが既に辿ったマスである場合は移動不可
	for _, location := range course {
		if target == location {
			return false
		}
	}
	// 上記の条件全てに当てはまらない場合は移動可
	return true
}
