package utils

import (
	"fmt"
	"time"
)

/**
 * Copyright (C), 2019-2020
 * FileName: request
 * Author:   LinGuanHong
 * Date:     2020/12/11 10:49 上午
 * Description:
 */

func RetryReq(retryTime int, delayTime time.Duration, reqFunc func() (interface{}, error)) (interface{}, error) {
	var (
		ret interface{}
		err error
	)
	for i := 0; i <= retryTime; i++ {
		if ret, err = reqFunc(); err != nil {
			fmt.Println("retry err:", err.Error())
			time.Sleep(delayTime)
			continue
		}
		break
	}
	return ret, nil
}
