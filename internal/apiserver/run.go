// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apiserver

import "github.com/marmotedu/iam/internal/apiserver/config"

// Run runs the specified APIServer. This should never exit.
func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}
   //todo  调用PrepareRun方法，进行 HTTP/GRPC 服务器启动前的准备。
   // 在准备函数中，我们可以做各种初始化操作，例如初始化数据库，安装业务相关的 Gin 中间件、RESTful API 路由等。

   //todo  完成 HTTP/GRPC 服务器启动前的准备之后，调用Run方法启动 HTTP/GRPC 服务。在Run方法中，分别启动了 GRPC 和 HTTP 服务。
	return server.PrepareRun().Run()
}
