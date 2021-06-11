/*
* @Author: Try
* @Date:   2021/5/5 11:45
* 处理c段
 */
package getip

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Getip(fname string) []string { //将ip转换为c段数组
	var s []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("[info] 加载失败！")
		return nil
	}
	reader := bufio.NewReader(file)
	for {
		url, err := reader.ReadString('\n') //注意是字符
		str1 := strings.Replace(url, "\n", "", -1)
		str := strings.Replace(str1, "\r", "", -1)
		if err == io.EOF {
			file.Close()
		}
		if err != nil {
			break
		}
		s = append(s, str)
		//fmt.Println(str)
	}
	return s
}
