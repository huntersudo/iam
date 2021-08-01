// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"os"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})

var shutdownHandler chan os.Signal

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
// todo [BP]
func SetupSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler) // panics when called twice
	// 通过 close(onlyOneSignalHandler)来确保 iam-apiserver 组件的代码只调用一次 SetupSignalHandler 函数
	// : close a closed chan will be panic

	shutdownHandler = make(chan os.Signal, 2)

	stop := make(chan struct{})

	signal.Notify(shutdownHandler, shutdownSignals...)
   // 收到一次 SIGINT/ SIGTERM 信号，程序优雅关闭。收到两次 SIGINT/ SIGTERM 信号，程序强制关闭
	go func() {
		<-shutdownHandler
		close(stop)
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

// RequestShutdown emulates a received event that is considered as shutdown signal (SIGTERM/SIGINT)
// This returns whether a handler was notified.
func RequestShutdown() bool {
	if shutdownHandler != nil {
		select {
		case shutdownHandler <- shutdownSignals[0]:
			return true
		default:
		}
	}

	return false
}
