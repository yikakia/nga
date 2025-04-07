package main

import (
	"log"

	"github.com/yikakia/nga"
)

func main() {
	c := nga.NewClient(nga.Config{
		BaseUrl:        "https://bbs.nga.cn",
		NgaPassportUid: "1234",
		NgaPassportCid: "4321",
	})

	thread, err := c.Thread("706")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(thread)

}
