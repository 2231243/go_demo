package cloud

import (
	"crypto/md5"
	"demo/base"
	"demo/inter"
	"demo/util"
	"fmt"

	"strconv"
	"strings"
	"time"
)

type QNCloud struct {
	base.Cloud
	sercet_path string
	Rtime       int64
}

func (c *QNCloud) md5str(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func (c *QNCloud) getPath(proto string) string {
	switch proto {
	case "rtmp":
		return "/" + c.Appid
	case "hls":
		return "/" + c.Appid + "/" + c.Path + "/index.m3u8"
	case "hdl":
		return "/" + c.Appid + "/" + c.Path + ".flv"
	}
	return ""
}
func (c *QNCloud) strTohex() string {
	t := c.Params["t"]
	if _, ok := t.(int); !ok {
		t, _ = strconv.Atoi(t.(string))
	}
	s := strconv.FormatInt(int64(int(c.Rtime)+t.(int)*60), 16)
	return strings.ToLower(fmt.Sprintf("%v", s))
}

func (c *QNCloud) Sign() string {
	s := c.Params["key"].(string) + c.sercet_path + c.strTohex()
	hash_str := c.md5str(s)
	return strings.ToLower(hash_str)
}

func (c *QNCloud) GetPullLink(cc inter.Cloud) (result inter.Result) {
	result = util.GetRtmpHlsHdl(c)
	return
}

func (c *QNCloud) Link() inter.Result {
	link := c.GetPullLink(c)
	return link
}

func (c *QNCloud) Query() string {
	t := c.strTohex()
	m := make(map[string]interface{})
	m["t"] = t
	m["l_t"] = t
	m["id"] = c.Id
	m["ip"] = c.Ip
	m["sign"] = c.Sign()
	m["sercet_path"] = c.sercet_path
	return util.Query(c.Params, m)
}

func (c *QNCloud) Rtmp(cc inter.Cloud, rtmp chan map[string]interface{}) {
	c.sercet_path = c.getPath("rtmp")
	c.Cloud.Rtmp(cc, rtmp)
	return
}

func (c *QNCloud) Hdl(cc inter.Cloud, hdl chan map[string]interface{}) {
	c.sercet_path = c.getPath("hdl")
	c.Cloud.Hdl(cc, hdl)
	return
}

func (c *QNCloud) Hls(cc inter.Cloud, hls chan map[string]interface{}) {
	c.sercet_path = c.getPath("hls")
	c.Cloud.Hls(cc, hls)
	return
}

func NewQNCloud(id int, ip, path, appid string, url map[string]string, params map[string]interface{}) *QNCloud {
	qn := &QNCloud{
		Cloud: base.Cloud{
			Id:     id,
			Ip:     ip,
			Path:   path,
			Appid:  appid,
			Url:    url,
			Params: params,
		},
		Rtime: time.Now().Unix(),
	}
	return qn
}
