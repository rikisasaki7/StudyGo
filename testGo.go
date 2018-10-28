package main

import (
	"fmt"
	"errors"
)

func main(){
	var msg string
	// 定義パターンのテスト
	fmt.Println("========定義パターンのテスト=======")
	msg = "hello world"
	fmt.Println(msg)

	var msg2 string = "hello world2"
	fmt.Println(msg2)

	var msg3 = "hello world3"
	fmt.Println(msg3)

	var a, b int
	a,b = 1, 3
	fmt.Printf("a: %d, b: %d\n", a,b) // 書式は一緒

	msg4 := "hellohwllo"
	fmt.Println(msg4)

	var(
		c int
		d string
	)
	c, d = 20, "hhhhh"
	fmt.Printf("c: %d, d: %s\n", c, d)

	fmt.Println("========変数定義時のテスト=======")
	var in int
	var fl float64
	var st string
	var bo bool

	fmt.Printf("a:%d, b:%f, c:%s, d:%t\n ", in, fl, st, bo)

	fmt.Println("========定数のテスト=======")
	const(
		sun = iota // 0
		mon
		tue
	)
	fmt.Printf("sun:%d, mon:%d, tue:%d\n ", sun, mon, tue)
	const(
		sunsun = "SUN"
		monmon = "MON"
		tuetue = "TUE"
	)
	fmt.Printf("sun:%s, mon:%s, tue:%s\n ", sunsun, monmon, tuetue)

	fmt.Println("========ポインタのテスト=======")
	poin := 5
	var pa *int // int型を格納する領域のアドレスを格納する準備
	pa = &poin // &poin = poinのアドレス

	// poinの領域にあるデータの値 = *poin
	fmt.Println(pa) // ポインタのアドレスが出力される
	fmt.Println(*pa) // paポインタに格納されている値（&poin）が出力される
	// どうやらポインタは構造体とセットで語られることが多いようだ。

	fmt.Println("========関数のテスト=======")
	sayHi()
	sayName("Bob")
	fmt.Println(getMessage("Bob x"))
	fmt.Println(getHelloMessage("Gemcook"))
	fmt.Println(swap(10, 9))

	// 関数もデータ型なので、変数に代入可能
	// その際は関数名はいらない
	f := swap // これでどうやらswap関数を示せるようだ。
	f2 := func(a int, b int)(int, int){ // こういう風に無名関数も渡せるようだ
		return a+5, a+10
	}
	fmt.Println(f(3, 8))
	fmt.Println(f2(10, 10))

	fmt.Println("========配列のテスト=======")
	// 宣言と代入を分ける
	var aArray[5]int //a[0] - a[4] -> 初期値は0のようだ。それぞれのデータ型で決まっていそうだ
	aArray[1] = 1
	aArray[2] = 2
	aArray[3] = 3
	aArray[4] = 4
	fmt.Print("aArray: ")
	fmt.Println(aArray)

	// 宣言とともに初期化する
	bArray := [3]int{1, 3, 5}
	fmt.Print("bArray: ")
	fmt.Println(bArray)

	// 配列の個数が未定の場合
	cArray := [...]int{2, 4, 6, 8, 10, 0}
	fmt.Print("cArray: ")
	fmt.Println(cArray)

	fmt.Println("========スライスのテスト=======")
	aSlice := [5]int{2, 10, 8, 15, 4} // 添え字は０からみたいだ
	bSlice := aSlice[2:4] // [8, 15] // この場合、指定した数字はどちらも含むようだ
	cSlice := aSlice[2:] // [8, 15, 4] // この場合、指定した数字は含むようだ
	dSlice := aSlice[:4] // [2, 10, 8, 15] // この場合は、指定した数字は含まないようだ
	eSlice := aSlice[:] // [2, 10, 8, 15, 4] // こうしたら全部
	fmt.Print("bSlice: ")
	fmt.Println(bSlice)
	fmt.Print("cSlice: ")
	fmt.Println(cSlice)
	fmt.Print("dSlice: ")
	fmt.Println(dSlice)
	fmt.Print("eSlice: ")
	fmt.Println(eSlice)

	// スライスの長さを取得
	fmt.Printf("cSlice len: %d\n", len(cSlice)) // 3が出力される
	// スライスの最大容量を取得
	fmt.Printf("cSlice capacity: %d\n", cap(cSlice))

	// make関数でスライスを作成
	s1 := make([]int, 3) // [0, 0, 0]

	// スライス作成と同時に初期化
	s2 := []int{1, 3, 5}
	fmt.Print("s1: ")
	fmt.Println(s1)
	fmt.Print("s2: ")
	fmt.Println(s2)

	// appendでスライスの末尾に要素を追加
	s3 := append(s2, 8, 2, 10)
	fmt.Print("s3: ")
	fmt.Println(s3)

	// copyでスライスをコピー
	// 返り値は要素数
	t := make([]int, len(s3))
	s4 := copy(t, s3)
	fmt.Print("s4: ")
	fmt.Println(s4)

	fmt.Println("========mapのテスト=======")
	// mapの宣言
	m := make(map[string]string)
	// mapに値を設定
	m["key"] = "value"
	fmt.Print("m(map): ")
	fmt.Println(m)

	// 値を指定しながら宣言する
	m2 := map[string]int{"key2": 200, "key3": 300}
	fmt.Print("m2(map): ")
	fmt.Println(m2)

	// キーの存在を調べる
	v, ok := m["key"] // どうやらvalueも一緒にとれるようだ
	fmt.Printf("v: %s, ok: %t\n", v, ok)

	// 要素を削除する
	delete(m2, "key3")
	fmt.Print("m2(map) after delete: ")
	fmt.Println(m2)


	if err := doError(); err != nil {
		fmt.Println("pointerで例外が発生しました。")
	}
}

func pointer() int {
	return 1
}

func doError() error {
	return errors.New("error!!")
}

// 引数無関数
func sayHi() {
	fmt.Println("sayHi called: hi!!")
}

// 引数あり関数
func sayName(name string){
	fmt.Printf("my name is %s!!!\n", name)
}

// return 関数
func getMessage(name string) string {
	msg := "hi! my name is " + name
	return msg
}

// 名前付きreturn関数
// 関数内で使った変数名を返す
func getHelloMessage(name string)(msg string){
	msg = "Hello " + name
	return
}

// 複数の返り値
func swap(a int, b int)(int, int){ // 複数のときはタプルのようだ
	fmt.Println("swap called!!!!!!!!!!!!!!!")
	return b, a
}