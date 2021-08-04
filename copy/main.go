package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gogf/gf/os/gfile"
)

func main() {
	src := "G:\\jinbi1"
	dst := "H:\\Music\\jinbi1"
	//src := "/home/fuqingrong/develop/test/copy/test"
	//dst := "/home/fuqingrong/develop/test/copy/test1"
	list, err := gfile.ScanDirFile(src, "*", true)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range list {
		fmt.Println(v)
		info, _ := gfile.Info(v)
		err = gfile.CopyFile(v, dst+"/"+info.Name())
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(2 * time.Second)
	}
}
