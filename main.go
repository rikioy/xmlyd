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
	cookie := "341318425&A12BB1C0140ND095F1D43A93EFBA1CBF9CC0B65F13A7D5CAA4FB929503ED2F149606644FC94E139MA66077EE060FFE6_"
	err := download("jb2", "jinbi2", 38624498, true, false, cookie)
	//err := download("jinbi2", 31909017, false, false, cookie)
	if err != nil {
		log.Fatalln(err)
	}
}

func download(prefix, path string, albumID int, dFree, dVip bool, cookie string) (err error) {
	info, err := xmlydownloader.GetAllAudioInfo(albumID)
	if err != nil {
		return
	}
	if !gfile.Exists(path) {
		gfile.Mkdir(path)
	}
	var index = 1
	for _, v := range info {
		fmt.Printf("%d, %s, %t, %t\n", v.TrackID, v.Title, v.IsFree, v.IsPaid)
		//if v.IsFree {
		if true {
			if dFree {
				tInfo, err := xmlydownloader.GetFreeTrackInfo(v.TrackID, "")
				if err != nil {
					log.Fatalln(err)
				}

				prefixStr := fmt.Sprintf("%s.%03d", prefix, index)
				filePath := fmt.Sprintf("./%s/%s-%s.mp3", path, prefixStr, tInfo.Title)
				fmt.Println(filePath)
				grab.Get(filePath, tInfo.DownloadURL)
				time.Sleep(2 * time.Second)
			}
		} else {
			if dVip {
				tInfo, err := xmlydownloader.GetVipAudioInfo(v.TrackID, cookie)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(tInfo.PlayPathAacv164)
				filePath := fmt.Sprintf("./%s/%s.mp3", path, tInfo.Title)
				fmt.Println(filePath)
				grab.Get(filePath, tInfo.PlayPathAacv164)
				time.Sleep(2 * time.Second)
			}
		}
		index++
	}
	return
}
