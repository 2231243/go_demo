package util

import (
	"demo/inter"

	"fmt"
	"net/url"
	"strconv"
)

func Query(c, m map[string]interface{}) string {

	u := url.Values{}
	for k, v := range c {
		if k == "key" {
			continue
		}
		m[k] = v
	}

	for k, v := range m {
		if _, ok := v.(int); ok {
			v = strconv.Itoa(v.(int))
		}
		u.Set(k, v.(string))
	}
	return u.Encode()
}

func GetRtmpHlsHdl(cc inter.Cloud) (result inter.Result) {
	rtmpChan := make(chan map[string]interface{})
	hdlChan := make(chan map[string]interface{})
	hlsChan := make(chan map[string]interface{})

	go cc.Hdl(cc, hdlChan)
	go cc.Rtmp(cc, rtmpChan)
	go cc.Hls(cc, hlsChan)
	var (
		rtmp = make(map[string]interface{})
		hdl  = make(map[string]interface{})
		hls  = make(map[string]interface{})
	)
	rtmp = <-rtmpChan
	hdl = <-hdlChan
	hls = <-hlsChan
	result = inter.Result{
		Rtmp: rtmp["rtmp"].(string),
		Hdl:  hdl["hdl"].(string),
		Hls:  hls["hls"].(string),
		Id:   cc.GetId(),
	}
	if r, ok := rtmp["rtmp_format"]; ok {
		result.RtmpFormat = r.(string)
	}
	if h, ok := hdl["hdl_format"]; ok {
		result.HdlFormat = h.(string)
	}
	return
}

func getCloudData(c inter.Cloud) (int, string, string, map[string]string, string) {
	var (
		id    int               = c.GetId()
		q     string            = c.Query()
		appid string            = c.GetAppid()
		urls  map[string]string = c.GetUrl()
		path  string            = c.GetPath()
	)
	return id, q, appid, urls, path
}

func GetPullLinkByType(pull_type string, c inter.Cloud, ch chan map[string]interface{}) {
	switch pull_type {
	case "rtmp":
		RtmpFormat(c, ch)
		break
	case "hdl":
		HdlFormat(c, ch)
		break
	case "hls":
		HlsFormat(c, ch)
		break
	}
}

func RtmpFormat(c inter.Cloud, rtmp chan map[string]interface{}) {
	makeRtmp := make(map[string]interface{})
	id, q, appid, urls, path := getCloudData(c)

	makeRtmp["rtmp"] = fmt.Sprintf("rtmp://%s/%s?%s/%s",
		urls["rtmp"], appid, q, path)
	if id == 2 {
		makeRtmp["rtmp_format"] = "rtmp://%s/" +
			fmt.Sprintf("%s?vhost=%s&%s/%s",
				appid, urls["rtmp"], q, path)
	}
	rtmp <- makeRtmp
	close(rtmp)
}

func HdlFormat(c inter.Cloud, hdl chan map[string]interface{}) {
	makeHdl := make(map[string]interface{})
	id, q, appid, urls, path := getCloudData(c)

	makeHdl["hdl"] = fmt.Sprintf("http://%s:8086/%s/%s.fiv?%s",
		urls["hdl"], appid, path, q)
	if id == 2 {
		makeHdl["hdl_format"] = "http://%s:8086/" +
			fmt.Sprintf("%s/%s.fiv?vhost=%s&%s", appid, path, urls["hdl"], q)
	}
	hdl <- makeHdl
	close(hdl)
}

func HlsFormat(c inter.Cloud, hls chan map[string]interface{}) {
	makeHls := make(map[string]interface{})
	_, q, appid, urls, path := getCloudData(c)
	makeHls["hls"] = fmt.Sprintf("//%s/%s/%s/index.m3u8?%s",
		urls["hls"], appid, path, q)

	hls <- makeHls
	close(hls)
}
