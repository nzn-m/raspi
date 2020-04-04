package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	// GPIOを制御するためのメモリと接続。エラーならエラー表示してプログラム終了。
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// プログラム終了時にrpio制御を手放す
	defer rpio.Close()
	// ピン配置の設定
	pin := rpio.Pin(2)
	// ピンの出力設定
	pin.Output()
	// 初期状態は消灯
	pin.High()
	// GPIOを10回点滅 (20÷2)
	for i := 0; i < 20; i++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
