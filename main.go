package main

import (
	"goreq/handler"
	"strings"

	"github.com/imroc/req"
	"github.com/wonderivan/logger"
)

func checkUrlContent(url string, tz string) bool {
	// only url is required, others are optional.
	r, err := req.Get(url)
	if err != nil {
		logger.Error("请求出错:", url, err)
		return false
	}
	data := r.String()
	return strings.Contains(data, tz)
}

func main() {

	handler.OpenSqlite()

	logger.Info("运行完毕")
}
