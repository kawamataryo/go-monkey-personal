package repl

import (
	"bufio"
	"fmt"
	"github.com/kawamataryo/go-monkey/lexer"
	"github.com/kawamataryo/go-monkey/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// プロンプトを表示
		fmt.Print(PROMPT)
		// 標準入力を待機
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// 入力文字をlineに代入
		line := scanner.Text()

		// 字句解析機へ
		l := lexer.New(line)

		// 軸解解析機で検証
		// 初期化でtokにl.nextToken()の取得結果を代入
		// 条件式でeofを判定
		// 増減式でtok = l.nextToken()を実行
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
