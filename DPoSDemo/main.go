package main

import (
	"DPoSDemo/Block"
	"fmt"
	"math/rand"
	"time"
)

var blockChain []Block.Block

// 存放代理人的地址
var delegate = []string{"aaa", "bbb", "ccc", "ddd"}

// 随机调换
func randDelegate() {
	source := rand.NewSource(time.Now().Unix())
	generator := rand.New(source)
	index := generator.Intn(len(delegate))
	index1 := generator.Intn(len(delegate))

	for index1 == index {
		index1 = generator.Intn(len(delegate))
	}

	tempDelegate := delegate[index]
	delegate[index] = delegate[index1]
	delegate[index1] = tempDelegate
}

func main() {

	var firstBlock Block.Block

	blockChain = append(blockChain, firstBlock)

	//通过变量n按顺序让代理人作为矿工
	var n = 0

	var temp = 0

	for {
		randDelegate()

		//每隔3秒产生一个新区快
		time.Sleep(3 * time.Second)

		var nextBlock = Block.GenerateNextBlock(temp, blockChain[temp], delegate[n])

		blockChain = append(blockChain, nextBlock)

		n++

		temp++

		n = n % len(delegate)

		fmt.Println(blockChain)
	}

}
