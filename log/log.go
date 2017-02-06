// Copyright 2017 ashan.org. All rights reserved.
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"time"
)

const (
	lerror = 4
	lwarn  = 2
	llog   = 1
)

// 日志级别，默认为0，均不打印
//
// 日志级别包含三个级别，均使用数字表示
//
// 1:打印log日志
//
// 2:打印warn警告
//
// 4:打印error错误
//
// 三个数字可以互相组合，使用方式类似Linux文件权限操作
//
// LOG_LEVEL = 1 只打印log日志
//
// LOG_LEVEL = 2 只打印warn警告
//
// LOG_LEVEL = 4 只打印error错误
//
// LOG_LEVEL = 3 打印 log日志 和 warn警告
//
// LOG_LEVEL = 5 打印 log日志 和 error错误
//
// LOG_LEVEL = 6 打印 warn警告 和 error错误
//
// LOG_LEVEL = 7 全部打印
var LOG_LEVEL = 0

// 调试打印开启选项，开发阶段可设置为true
var DEBUG = false

func log_Getime() string {
	times := time.Now().Unix()
	tm := time.Unix(times, 0)
	return tm.Format("02/01/2006 15:04:05 PM")
}

// 打印错误日志
func E(a ...interface{}) {
	if LOG_LEVEL&lerror == lerror {
		fmt.Println("\x1b[31mERROR: [", log_Getime(), "]", a, "\x1b[0m")
	}
}

// 打印警告日志
func W(a ...interface{}) {
	if LOG_LEVEL&lwarn == lwarn {
		fmt.Println("\x1b[33mWARN : [", log_Getime(), "]", a, "\x1b[0m")
	}
}

// 打印Log日志
func L(a ...interface{}) {
	if LOG_LEVEL&llog == llog {
		fmt.Println("\x1b[32mLOG  : [", log_Getime(), "]", a, "\x1b[0m")
	}
}

// 打印debug调试信息
func D(a ...interface{}) {
	if DEBUG == true {
		fmt.Println("\x1b[35mDEBUG: [", log_Getime(), "]", a, "\x1b[0m")
	}
}
