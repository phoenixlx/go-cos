package main

import (
	"context"
	//"net/url"
	"os"

	"net/http"

	"github.com/lewzylu/go-cos"
	"github.com/lewzylu/go-cos/debug"
)

func main() {
	u, _ := url.Parse("https://test-1253846586.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	name := "test/hello.txt"
	_, err := c.Object.Head(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
}
