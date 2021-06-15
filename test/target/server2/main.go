package main

import (
	"io"
	"net/http"
)

const (
	addr = ":9092"
)

// 下游服务器
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, fmt.Sprintf("X-Forwarded-For: %s\n", r.Header["X-Forwarded-For"]))
		//io.WriteString(w, fmt.Sprintf("X-Real-IP: %s\n", r.Header["X-Real-IP"]))
		//io.WriteString(w, fmt.Sprintf("RemoteAddr: %s\n", r.RemoteAddr))

		//w.WriteHeader(http.StatusFound)
		//io.WriteString(w, "RequestURI: "+r.RequestURI)
		io.WriteString(w, "\t下游服务器: "+addr)
	})

	http.ListenAndServe(addr, nil)
}
