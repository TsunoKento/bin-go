package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxBingoNum   int = 50
	bingoCardSize int = 25
)

func main() {
	createBingoCard()
}

// 5*5で作成する
// 数字は1~50の間で重複なしのランダム
// 真ん中はすでに空いている(0)状態
func createBingoCard() {
	fmt.Println("ビンゴカードを作成します。")

	numList := createNumList(maxBingoNum)
	var card [bingoCardSize]int
	for i := 0; i < bingoCardSize; i++ {
		if i == bingoCardSize/2 {
			continue
		}
		var randomKey int
		for {
			rand.Seed(time.Now().UnixNano())
			randomKey = rand.Intn(maxBingoNum)
			if numList[randomKey] == 0 {
				continue
			} else {
				break
			}
		}
		card[i] = numList[randomKey]
		numList[randomKey] = 0
	}
}

func createNumList(maxNum int) []int {
	numList := make([]int, maxNum)
	for i := 1; i <= maxNum; i++ {
		numList[i-1] = i
	}

	return numList
}
