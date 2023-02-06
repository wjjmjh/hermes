package main

import (
	"wjjmjh/hermes/managers"
	"wjjmjh/hermes/pkg/setting"
)

func init() {
	setting.Setup()
}

func main() {
	chatManager := managers.InitialiseManager()
	chatManager.RunWsServer()
	return
}
