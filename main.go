package main

import (
	"fmt"
	"log"

	"github.com/rikioy/xmlydownloader"
)

func main() {
	err := download(31909017)
	if err != nil {
		log.Fatalln(err)
	}
}

func download(albumID int32) (err error) {
	info, err := xmlydownloader.GetAllAudioInfo(31909017)
	//info, err := xmlydownloader.GetTrackList(31909017, 1, true)
	if err != nil {
		return
	}
	for _, v := range info {
		fmt.Printf("%d, %s, %t, %t\n", v.TrackID, v.Title, v.IsFree, v.IsPaid)
		/*
			tInfo, err := xmlydownloader.GetTrackInfoByMobile(v.TrackID, "")
			if err != nil {
				log.Fatalln(err)
			}
			grab.Get("./"+tInfo.Title+".mp3", tInfo.DownloadURL)
			time.Sleep(5 * time.Second)
		*/
	}
	return
}
