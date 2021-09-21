// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package apiserver does all of the work necessary to create a iam APIServer.
package apiserver

import (
	"github.com/marmotedu/iam/internal/apiserver/config"
	"github.com/marmotedu/iam/internal/apiserver/options"
	"github.com/marmotedu/iam/pkg/app"
	"github.com/marmotedu/iam/pkg/log"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more iam-apiserver information at:
    https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-apiserver.md`

// NewApp creates a App object with default parameters.
func NewApp(basename string) *app.App {
	// 创建带有默认值的 Options 类型变量 opts
	// 最终在 App 框架中，被来自于命令行参数或配置文件的配置（也可能是二者 Merge 后的配置）所填充，opts 变量中各个字段的值会用来创建应用配置
	opts := options.NewOptions()
	application := app.NewApp("IAM API Server",
		basename,
		app.WithOptions(opts),  // todo  Options implement CliOptions
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush() // todo

        // todo 2 通过CreateConfigFromOptions函数来构建应用配置：
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
