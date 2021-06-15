package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

const (
	targetURL  = "http://127.0.0.1:8002/base"
	reverseURL = ":9090"
)

var targets = []*url.URL{
	{
		Scheme: "http",
		Host:   "localhost:9091",
	},
	{
		Scheme: "http",
		Host:   "localhost:9092",
	},
}

// 代理服务器
func main() {
	//URL, err := url.Parse(targetURL)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//proxy := httputil.NewSingleHostReverseProxy(URL)
	proxy := &httputil.ReverseProxy{}
	// 自定义director
	proxy.Director = func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}

	// modify resp
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
