package main

import (
	"errors"
	"log"
)

func main() {
	msg := "ok"
	body := "demo"
	channel := make(chan string)
	go func(b string) {
		if err := DataProcess(b); err != nil {
			channel <- err.Error()
		}
		channel <- "处理成功"
	}(body)
	msg = <-channel
	log.Println("MSG", msg)

	body = ""

	go func(b string) {
		if err := DataProcess(b); err != nil {
			channel <- err.Error()
		}
		channel <- "处理成功"
	}(body)
	msg = <-channel
	log.Println("MSG", msg)

}
func DataProcess(b string) error {
	if b == "" {
		return nil

	}
	return errors.New(b)
}
