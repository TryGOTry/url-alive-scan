/*
* @Author: Try
* @Date:   2021/5/5 9:59
 */
package main

import (
	"flag"
	"url-alive-scan/root"
)

func main() {

	fname := flag.String("f", "", "filename")
	num := flag.Int("s", 5, "线程")
	flag.Parse()
	if *fname != "" {
		root.GoWebScan(*fname, *num, 3)
	}else {
		flag.Usage()
	}
}
