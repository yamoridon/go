package main

import (
	"fmt"
	"gobook/ch5/links"
	"log"
	"os"
)

// breadthFirst は worklink 内の個々の項目に対して f を呼び出します。
// f から返された全ての項目は worklist へ追加されます。
// f は、それぞれの項目に対して高々一度しか呼び出されません。
func breathFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// コマンドライン引数から開始して、
	// ウェブを幅優先でクロールする
	breathFirst(crawl, os.Args[1:])
}
