package main

import (
	"demo/cloud"
	"encoding/json"
	"fmt"
)

func main() {
	urls := make(map[string]string)
	urls["rtmp"] = "rtmp.www.baidu.com"
	urls["hdl"] = "hdl.www.baidu.com"
	urls["hls"] = "hls.www.baidu.com"
	params := make(map[string]interface{})
	params["key"] = "456"
	params["t"] = 5

	qn := cloud.NewQNCloud(2,
		"127.0.0.1",
		"2019071010001",
		"2de8992d3f3b73717d55edcc16d230c607585764",
		urls,
		params)

	b, _ := json.Marshal(qn.Link())

	fmt.Println(string(b))

}
