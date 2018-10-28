package main

import (
    "fmt"
)

// 構造体を定義
// 構造体：複数の値を意味のあるまとまりとして新しい型を定義する
type user struct {
    // フィールド
    name string 
    score int
}

func main(){

    fmt.Println("========構造体のテスト=======")
    // newでuser構造体分の領域を確保して、初期化して、そのポインタを返す
    u1 := new(user)
    // ポインタが帰ってきているので下記の書き方でもok
    (*u1).name = "testUser"
    u1.name = "testtest"
    u1.score = 200
    fmt.Println(u1)

    fmt.Println("========メソッドのテスト=======")
    // golang ではオブジェクト志向プログラミングで出てくるクラスやその中で定義されるメソッドはサポートしていない。
    // golang のメソッドは構造体などのデータ型に紐付いた関数になっていて、データ型の定義の中ではなく、データ型とは別に書く
    // 構造体とメソッド(関数)の紐付けには レシーバ を使う
    met := user{name:"testmethod", score: 300}
    met.hit() // 引数で構造体を指定するとかじゃあないみたいだ
    met.show()

    fmt.Println("========インタフェースのテスト=======")
    greeters := []greeter{japanese{}, american{}}
    for _, greeter := range greeters {
        greeter.greet()
        show(greeter)
    }
}

// レシーバーで構造体とメソッドを紐づける
func(u user) show(){
    fmt.Printf("name: %s, socre: %d\n", u.name, u.score)
}

// デフォルトではuserには価渡し（コピー）でメソッドに渡される
// *をつける事で参照渡し（ポインタを渡す）になる
func(u *user) hit(){
    u.score++
}

type greeter interface {
    greet()
}
type japanese struct{}
type american struct{}
func(ja japanese) greet(){
    fmt.Println("こんにちは")
}
func(us american) greet(){
    fmt.Println("Hello")
}

// 空のインタフェース型で引数を受け取る
func show(t interface{}){
    // 型アサーション（キャストっぽい）
    // ２つの値が返る
    _, ok := t.(japanese)
    // okを使って条件分岐
    if ok {
        fmt.Println("i am japanese")
    } else {
        fmt.Println("i am not japanese")
    }

    // 型switch
    switch t.(type){
    case japanese:
        fmt.Println("僕は日本人なんだな")
    default:
        fmt.Println("僕は日本人じゃないんだな")
    }
}