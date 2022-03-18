package handler

import (
	"time"

	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"github.com/wonderivan/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Content struct {
	ID      int
	Content string
}

func OpenSqlite() {
	db, err := gorm.Open(sqlite.Open("test.db3"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("数据库打开失败")
	}

	var contents []Content
	db.Raw("SELECT ID,content From Content").Scan(&contents)
	logger.Info("数据个数: ", len(contents))

	//遍历数组

	for _, value := range contents {
		newcontent := Rewrite(value.Content)
		if len(newcontent) > 1 {
			var cc Content
			db.Raw("UPDATE Content SET content= ? WHERE ID = ? RETURNING ID, content", newcontent, value.ID).Scan(&cc)
			logger.Info("更新ID:", cc.ID)
		}
		logger.Info("休息5秒...")
		time.Sleep(5 * time.Second)
	}

}

func Rewrite(txt string) string {

	header := req.Header{
		"Accept":        "application/x-www-form-urlencoded; charset=UTF-8",
		"Authorization": "C2220F5C64B94EA29349756B2AECF1E1",
	}

	param := req.Param{
		"txt": txt,
	}

	r, err := req.Post("http://apis.5118.com/wyc/rewrite", header, param)
	if err != nil {
		logger.Error("请求出错:", err)
		return ""
	}

	data := gjson.Get(r.String(), "data").String()
	logger.Info("重写结果: ", data)
	if len(data) == 0 {
		logger.Warn("没有获取到结果", r.String())
	}

	return data

}
