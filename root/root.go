/*
* @Author: Try
* @Date:   2021/5/5 11:18
 */
package root

import (
	"github.com/gookit/color"
	"sync"
	"time"
	"url-alive-scan/getip"
	"url-alive-scan/golimit"
	"url-alive-scan/save"
	"url-alive-scan/scan"
)

func GoWebScan(fname string, num int, timeout int64) {
	dicall := getip.Getip(fname)
	color.Red.Println("[Info] 一个简单的url存活检测工具")
	color.Red.Println("[Info] www.nctry.com")
	color.Red.Println("[Info] 开始扫描中.当前线程:", num)
	color.Red.Println("---------------------------------------")
	//fmt.Println(dicall)
	g := golimit.NewG(num) //设置线程数量
	wg := &sync.WaitGroup{}
	beg := time.Now()
	for i := 0; i < len(dicall); i++ {
		wg.Add(1)
		task := dicall[i]
		g.Run(func() {
			respBody, err := scan.Goscan(task, timeout)
			if err != nil {
				//color.Warn.Println("目标访问错误，可能被ban了！")
				wg.Done()
				return
			}
			if respBody.StatusCode == 200 {
				color.Info.Println("[200] ", respBody.Url+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
				save.Writefile(respBody.Url+" || "+respBody.Title, fname)
			} else if respBody.StatusCode == 403 {
				color.Warn.Println("[403] ", respBody.Url+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
				//writefile.Write(url, "[403] "+respBody.Res+"\n")
				save.Writefile(respBody.Url+" || "+respBody.Title, fname)
			} else if respBody.StatusCode == 302 {
				color.Warn.Println("[302] ", respBody.Url+"   [len]", respBody.Bodylen, "   [title]", respBody.Title, "   [server]", respBody.Server)
				//writefile.Write(url, "[302] "+respBody.Res+"\n")
				save.Writefile(respBody.Url+" || "+respBody.Title, fname)
			}
			wg.Done()
		})
	}
	wg.Wait()
	a := save.Writefile("", fname)
	color.Red.Printf("[info] 扫描完成！当前用时: %fs\n", time.Now().Sub(beg).Seconds())
	color.Red.Println("[info] 保存文件名:",a)
}
