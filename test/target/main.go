package main

import (
	"net/http"
)

func main() {
	addr := ":2003"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("我是被代理的下游服务器, "))
		w.Write([]byte("请求都被代理到我这里来处理了->"))
		w.Write([]byte(addr))
	})
	http.ListenAndServe(addr, nil)
}
