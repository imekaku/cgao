package main

import (
	"fmt"
	"github.com/cgao/change-date-format-tweets/conf"
	"github.com/cgao/change-date-format-tweets/handle"
	"time"
)

func main() {
	conf.ParseConfig("./conf/cfg.json")

	t1 := time.Now()
	fmt.Println("The programming is runing, do NOT close the window.")
	handle.Handle()
	elapsed := time.Since(t1)
	fmt.Println("app elapsed:", elapsed)
}
