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
	targetURL  = "http://127.0.0.1:8002/base"
	reverseURL = ":8001"
)

// 代理服务器
func main() {
	URL, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(URL)
	proxy.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode == http.StatusFound {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			newPayload := []byte("rewrite resp body: " + string(body))
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(newPayload))
			resp.ContentLength = int64(len(newPayload))
			resp.Header.Set("Content-Length", strconv.Itoa(len(newPayload)))
		}
		return nil
	}
	log.Printf("代理服务器：%s\t", reverseURL)
	log.Fatalln(http.ListenAndServe(reverseURL, proxy))
}
