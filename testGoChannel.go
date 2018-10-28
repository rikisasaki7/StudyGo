package main

import(
	"fmt"
	"time"
)

/*
 それぞれのタスクからメインの方で使いたい。
 タスクでreturnなどをしたいところだが、goroutineではそれができない。
 channelはデータの受け渡しをするパイプの様なイメージ。
   例えば、task1 の方で channel に対してデータを渡して、main の方でそのデータを取り出す。
 channel はスライスなどと同じ参照型のデータ
   makeで作成可能
 */

 func task1(result chan string){
	 // 重い処理を想定
	 time.Sleep(time.Second * 2)
	 fmt.Println("task1 finished!")

	 result <- "task1 result"
 }

 func task2() {
    fmt.Println("task2 finished!")
}

func main(){
	// chan型で受け渡すデータの型はstring
	result := make(chan string)

	// 並行処理
	go task1(result)
	go task2()

	// resultの中に何も入っていなければ入ってくるまで待つ仕様
	fmt.Println(<-result)

    // goroutineが終わる前に main 関数自体が終わる為、待ち時間をつける。
	time.Sleep(time.Second * 3)
}