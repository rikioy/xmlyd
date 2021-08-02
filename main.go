package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cavaliercoder/grab"
	"github.com/gogf/gf/os/gfile"
	"github.com/rikioy/xmlydownloader"
)

func main() {
	err := download("jinbi1", 31909017)
	if err != nil {
		log.Fatalln(err)
	}
}

func download(path string, albumID int32) (err error) {
	info, err := xmlydownloader.GetAllAudioInfo(31909017)
	//info, err := xmlydownloader.GetTrackList(31909017, 1, true)
	if err != nil {
		return
	}
	if !gfile.Exists(path) {
		gfile.Mkdir(path)
	}
	for _, v := range info {
		fmt.Printf("%d, %s, %t, %t\n", v.TrackID, v.Title, v.IsFree, v.IsPaid)
		if v.IsFree {
			tInfo, err := xmlydownloader.GetFreeTrackInfo(v.TrackID, "")
			if err != nil {
				log.Fatalln(err)
			}
			filePath := fmt.Sprintf("./%s/%s.mp3", path, tInfo.Title)
			fmt.Println(filePath)
			grab.Get(filePath, tInfo.DownloadURL)
			time.Sleep(2 * time.Second)
		}
	}
	return
}
