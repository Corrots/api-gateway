package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

const (
	reverseURL = "http://127.0.0.1:2003/base"
	proxyURL   = "127.0.0.1:2002"
)

func main() {
	URL, err := url.Parse(reverseURL)
	if err != nil {
		panic(err)
	}
	proxyHandler := httputil.NewSingleHostReverseProxy(URL)
	proxyHandler.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode == 200 {
			oldPayload, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			// 追加内容
			newPayload := []byte("StatusCode err: " + string(oldPayload))
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(newPayload))
			size := int64(len(newPayload))
			resp.ContentLength = size
			resp.Header.Set("Content-Length", strconv.FormatInt(size, 10))
		}
		return nil
	}
	log.Println("我是代理服务器：" + proxyURL)
	log.Fatalln(http.ListenAndServe(proxyURL, proxyHandler))
}
