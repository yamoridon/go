// nonempty はスライス内のアルゴリズムの例です。
package main

import "fmt"

// nonempty は空文字ではない文字列を保持するスライスを返します。
// 基底配列は呼び出し中に修正されます。
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "trhee"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
}
