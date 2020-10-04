package main

import (
	"flag"
)

var cfg = flag.String("f", "suning.yml", "suning spider")

func main() {
	flag.Parse()

	app, err := CreateApp(*cfg)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	//等待结束信号
	app.AwaitSignal()
}
