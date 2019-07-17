package inter

type Result struct {
	Rtmp       string `json:"rtmp"`
	Hdl        string `json:"hdl"`
	Hls        string `json:"hls"`
	Id         int    `json:"id"`
	RtmpFormat string `json:"rtmp_format"`
	HdlFormat  string `json:"hdl_format"`
}

type Cloud interface {
	GetId() int
	GetPath() string
	GetAppid() string
	GetUrl() map[string]string
	GetPullLink(c Cloud) (result Result)

	Sign() string
	Link() Result
	Query() string
	Rtmp(c Cloud, rtmp chan map[string]interface{})
	Hdl(c Cloud, hdl chan map[string]interface{})
	Hls(c Cloud, hls chan map[string]interface{})
}
