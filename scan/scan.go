package scan

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/encoding/gcharset"
	"regexp"
	"strings"
	"time"
)

var (
	title = `<title>([\s\S]+?)</title>`
)

type Webinfo struct {
	StatusCode int
	Title      string
	Server     string
	Powered    string
	Body       string
	Url        string //成功的结果
	Bodylen    int    //返回包长度
}

func Goscan(url1 string, timeout int64) (Webinfo, error) {
	var Web Webinfo
	client := resty.New().SetTimeout(time.Duration(timeout) * time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}) //忽略https证书错误，设置超时时间
	client.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	//client.Header.Set("Accept-Charset", "gb2312")
	//client.Header.Add("Accept-Charset", "UTF-8")
	resp, err := client.R().EnableTrace().Get(url1) //开始请求扫描
	if err != nil {
		//log.Println(err)
		return Web, err
	}

	str := resp.Body()
	body := string(str)
	//body = url.QueryEscape(body)
	//fmt.Println(body)
	if strings.Contains(body, "<title>") {
		re1 := regexp.MustCompile(title) //正则取标题
		titlename := re1.FindAllStringSubmatch(body, 1)
		if len(titlename) > 0 {
			if strings.Contains(body, "UTF-8") == false {
				srcCharset := "GB2312"
				dstCharset := "UTF-8"
				str, _ := gcharset.Convert(dstCharset, srcCharset, titlename[0][1])
				//fmt.Println(str)
				Web.Title = string(str)
			}
			Web.Title = titlename[0][1]
		}
	}
	Web.Url = url1
	Web.StatusCode = resp.StatusCode()
	Web.Powered = resp.Header().Get("X-Powered-By")
	//Web.Title = titlename[0][1]
	Web.Server = resp.Header().Get("server")
	Web.Body = body
	Web.Bodylen = len(body)
	return Web, nil
}
