package save

import (
	"fmt"
	"os"
	"time"
)

func Writefile(wireteString string,fname string) string{
	t := time.Now().Format("2006-01-02")
	filename :=t+"-"+fname+".txt"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte(wireteString+"\n"))
	}
	return filename
}
