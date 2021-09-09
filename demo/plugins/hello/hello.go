package hello

import "github.com/ri0day/rboot"

func init() {
	// 注册脚本
	rboot.RegisterPlugin(`hello`, rboot.Plugin{
		// 脚本处理函数
		Action: func(bot *rboot.Robot, incoming *rboot.Message) []*rboot.Message {
			return rboot.NewMessages("Hello World!")
		},
		Ruleset: map[string]string{`hello`: `^!hello`}, // 脚本规则集
		Usage: map[string]string{
			"!hello": "say hello world",
		},
		Description: `example 'Hello World' script for rboot`,
	})
}
