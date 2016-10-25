// 将tweets的json格式数据导入到MySQL
// 2016年10月24日11:33:27

package main

import (
	"fmt"
	"github.com/cgao/mv-tweets-data/conf"
	"github.com/cgao/mv-tweets-data/handle"
	"time"
)

func main() {
	conf.ParseConfig("./conf/cfg.json")

	t1 := time.Now()
	fmt.Println("正在从json格式文件导入到MySQL，请勿关闭..")
	handle.Handle()
	elapsed := time.Since(t1)
	fmt.Println("app elapsed:", elapsed)
}
