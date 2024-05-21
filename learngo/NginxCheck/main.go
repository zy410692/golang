package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/blinkbean/dingtalk"
)

func task() {
	//https://oapi.dingtalk.com/robot/send?access_token=5d66630c6039402b225c8c33abcd48cdedb70d0e0de0c4c8f1bf944a831fdf5f
	var dingToken = "5d66630c6039402b225c8c33abcd48cdedb70d0e0de0c4c8f1bf944a831fdf5f"
	cli := dingtalk.InitDingTalkWithSecret(dingToken, "SEC65e6bfdfd2b2e4a999b9c33cc4d70232ca5b5eafa29b14b188c7e9ea2b20c213")

	cmd := exec.Command("ls", "-l")  //调用ls -l命令
	out, err := cmd.CombinedOutput() //获取命令的输出
	if err != nil {
		log.Println(err)
		cli.SendTextMessage(err.Error())
	}
	fmt.Println(string(out))
	if len(out) > 0 {
		cli.SendTextMessage("nginx -t 状态正常")
	}
}

func main() {
	ticker := time.NewTicker(1 * time.Hour)

	go func() {
		for range ticker.C {
			task() // 每小时执行一次任务
		}
	}()

	// 使用信号量阻塞，否则主进程结束，goroutine也会结束
	select {}

}
