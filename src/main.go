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
// 真ん中はすでに空いている状態
func createBingoCard() {
	fmt.Println("ビンゴカードを作成します。")

	numList := createNumList()
	var card [bingoCardSize]int
	for i := 0; i < bingoCardSize; i++ {
		var targetKey int
		for {
			rand.Seed(time.Now().UnixNano())
			targetKey = rand.Intn(maxBingoNum)
			if numList[targetKey] == 0 {
				continue
			} else {
				break
			}
		}
		card[i] = numList[targetKey]
		numList[targetKey] = 0
	}
}

func createNumList() [maxBingoNum]int {
	var numList [maxBingoNum]int
	for i := 1; i <= maxBingoNum; i++ {
		numList[i-1] = i
	}

	return numList
}
