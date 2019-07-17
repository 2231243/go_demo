package base

import (
	"demo/inter"
	"demo/util"
)

type Cloud struct {
	Id     int
	Ip     string
	Url    map[string]string
	Path   string
	Appid  string
	Params map[string]interface{}
}

func (c *Cloud) GetId() int {
	return c.Id
}
func (c *Cloud) GetPath() string {
	return c.Path
}
func (c *Cloud) GetAppid() string {
	return c.Appid
}
func (c *Cloud) GetUrl() map[string]string {
	return c.Url
}

func (c *Cloud) Sign() string {
	return ""
}

func (c *Cloud) GetPullLink(cc inter.Cloud) (result inter.Result) {
	result = util.GetRtmpHlsHdl(cc)
	return
}

func (c *Cloud) Link() inter.Result {
	ni := inter.Result{}
	return ni
}

func (c *Cloud) Query() string {
	return ""
}

func (c *Cloud) Rtmp(cc inter.Cloud, rtmp chan map[string]interface{}) {
	util.GetPullLinkByType("rtmp", cc, rtmp)
}

func (c *Cloud) Hdl(cc inter.Cloud, hdl chan map[string]interface{}) {
	util.GetPullLinkByType("hdl", cc, hdl)
}

func (c *Cloud) Hls(cc inter.Cloud, hls chan map[string]interface{}) {
	util.GetPullLinkByType("hls", cc, hls)
}
