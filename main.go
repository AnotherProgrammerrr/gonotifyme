package main

import (
	"log"
	"os"

	"gopkg.in/toast.v1"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	var imageDir string = path + "/image.png"

	notification := toast.Notification{
		AppID:   "Yamaha",
		Title:   "omg, hi!!",
		Message: "hi!!",
		Icon:    imageDir,
		Actions: []toast.Action{
			{"protocol", "hi!!", ""},
			{"protocol", "bye.", ""},
		},
	}

	Noterr := notification.Push()
	if Noterr != nil {
		log.Fatalln(err)
	}

}
