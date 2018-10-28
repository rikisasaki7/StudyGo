package main
import(
	"fmt"
	"time"
)

func task1(){
	// 重い処理を想定
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finished!")
}

func task2(){
	fmt.Println("task2 finished!")
}

func main(){
	// go routine
    fmt.Println("========go routineのテスト=======")
	// goをつけて並行処理にする
	go task1()
	go task2()

	// goroutineが終わる前にmain関数自体が終わるため、待ち時間をつける
	time.Sleep(time.Second * 3)
}