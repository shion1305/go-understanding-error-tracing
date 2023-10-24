package main

import (
	"fmt"
	"github.com/pkg/errors" // pkg/errors パッケージをインポート
)

func main() {
	err := doSomething()
	if err != nil {
		// %+v とすることで、スタックトレースを含めた詳細なエラー情報を出力できます
		fmt.Printf("%+v\n", err)
	}
}

func doSomething() error {
	if err := doSomething1(); err != nil {
		// スタックトレース情報をエラーに追加します
		return errors.Wrap(err, "new error") // fmt.Errorf の代わりに errors.Wrap を使用
	}
	return nil
}

func doSomething1() error {
	return errors.New("test error") // 通常のエラー生成
}
