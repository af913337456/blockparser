package main

import (
	"fmt"
	"github.com/af913337456/blockparser/demo/btc_chain"
	"github.com/af913337456/blockparser/scanner"
	"github.com/af913337456/blockparser/types"
	"time"
)

/**
 * Copyright (C), 2019-2021
 * FileName: main
 * Author:   LinGuanHong
 * Description: 例子
 */

func main() {
	// demo
	_scanner := scanner.NewBlockScanner(scanner.InitBlockScanner{
		Chain: btc_chain.NewBlockCypherAPI(),
		Db:    btc_chain.NewBtcBlockScannerMysqlStorage(nil),
		Log:   new(myBlockParserLogger),
		Control: types.DelayControl{
			RoundDelay: time.Millisecond * 200,
			CatchDelay: time.Millisecond * 400,
		},
		FrontNumber: 6,
	})
	// _ = _scanner.SetStartScannerHeight(10000)
	if err := _scanner.Start(); err != nil {
		panic(err)
	}
	select {}
}

type myBlockParserLogger struct{}

func (l *myBlockParserLogger) Info(args ...interface{}) {
	fmt.Println(args...)
}
func (l *myBlockParserLogger) Error(args ...interface{}) { fmt.Println(args...) }
func (l *myBlockParserLogger) Warn(args ...interface{})  { fmt.Println(args...) }
