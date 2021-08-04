package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gogf/gf/os/gfile"
)

func main() {
	Check(1)
	src := os.Args[1]
	if !gfile.IsDir(src) {
		log.Fatalf("%s not dir\n", src)
	}
	list, err := gfile.ScanDirFile(src, "*", true)
	if err != nil {
		log.Fatalf("scan %s failed, err=%v\n", src, err)
	}

	for _, v := range list {
		info, _ := gfile.Info(v)
		split := strings.Split(info.Name(), ".")
		var newName string
		for i := 0; i < 3-len(split[0]); i++ {
			newName = newName + "0"
		}
		split[0] = newName + split[0]
		newName = strings.Join(split, ".")
		dst := src + "/output/"
		if !gfile.Exists(dst) {
			gfile.Mkdir(dst)
		}
		fmt.Printf("copy: %s to %s\n", src+info.Name(), dst+newName)
		gfile.Copy(src+"/"+info.Name(), dst+newName)
	}
}

func Check(argsNum int) {
	if len(os.Args)-1 != argsNum {
		log.Fatalln("args wrong")
	}
}
