# 字句解析
ソースコードからトークン列への変換を字句解析という。
プログラミングコードから空白、コメントなどを取り除き意味を持つ文字列の配列を作り出す。


# Goの文法

### `:=`
変数の宣言と代入を一度にやる。関数の内部でしか使えない。
Short variable declarations
https://tour.golang.org/basics/10

### `&型名`
どうやら型のポインタを抽出するらしい..

### `type 名前 struct {}`
構造体を作る。任意の型をまとめたもの。

### `func (型名) 関数名() {}`
構造体のメソッドを定義する？

### `len`
letは名前空間が関数内に限定されるぽい。JSのletと同じ


# Goのテスト
t.Fatal()を使う。これが呼ばれるとテストが落ちる。正常系だと呼ばれない。なるほど。。全然考え方が違う
https://golang.org/pkg/testing/

# Goのパッケージのパス
好きなところに置いていいわけではない..!!
https://hodalog.com/golang-standard-project-layout/
