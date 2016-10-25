package main

import (
	"fmt"
	"github.com/cgao/change-date-format-tweets/conf"
)

func main() {
	conf.ParseConfig("./conf/cfg.json")

	t1 := time.Now()
	fmt.Println("The programming is runing, do NOT close the WINDOW.")
	handle.Handle()
	elapsed := time.Since(t1)
	fmt.Println("app elapsed:", elapsed)
}