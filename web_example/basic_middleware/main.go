// advanced-middleware.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging
// 全てのリクエストについてパスと処理時間をログに出力する。
func Logging() Middleware {

	// 新しい Middleware を作る。
	return func(f http.HandlerFunc) http.HandlerFunc {

		// http.HandlerFunc を定義する。
		return func(w http.ResponseWriter, r *http.Request) {

			// ミドルウェアとしての処理を行う。
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// チェインの次のミドルウェア/HTTP ハンドラを呼び出す。
			f(w, r)
		}
	}
}

// Method
// 特定の HTTP メソッドから呼ばれたことを保証する。それ以外の場合は 400 Bad Request を返す。
func Method(m string) Middleware {

	// 新しい Middleware を作る。
	return func(f http.HandlerFunc) http.HandlerFunc {

		// http.HandlerFunc を定義する。
		return func(w http.ResponseWriter, r *http.Request) {

			// ミドルウェアとしての処理を行う。
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// チェインの次のミドルウェア/HTTP ハンドラを呼び出す。
			f(w, r)
		}
	}
}

// Chain
// http.HandlerFunc に対し全てのミドルウェアを適用する。
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}
