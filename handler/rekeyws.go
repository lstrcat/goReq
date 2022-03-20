package handler

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/wonderivan/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RandRules() string {
	fmts := []string{"%s_%s安卓版下载",
		"%s_%s官方版下载",
		"%s下载_2022%s下载",
		"2022%s_官方版%s下载",
		"%s_%s手机版下载",
		"%s安卓版_%sAndroid下载",
		"%s苹果版_%sIOS版下载",
		"%s手机版下载_%s官方版",
		"%s2022稳定版_%s最新版下载"}

	b := rand.Intn(len(fmts))
	return fmts[b]
}

func RuleKeyws(kw string) string {
	r := RandRules()

	s := fmt.Sprintf(r, kw, kw)

	return s
}

func loadTxt() []string {
	var lines []string

	fi, err := os.Open("6666.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return lines
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}

type Content struct {
	ID int
}

func OpenSqlite() {
	db, err := gorm.Open(sqlite.Open("rere.db3"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("数据库打开失败")
	}

	var contents []Content
	db.Raw("SELECT ID From Content").Scan(&contents)
	logger.Info("sqlite数据个数: ", len(contents))

	lines := loadTxt()
	logger.Info("txt数据个数: ", len(lines))

	if len(lines) == 0 || len(contents) == 0 {
		logger.Error("没有数据")
		return
	}

	//遍历数组

	txtIdx := 0
	for idx, value := range contents {

		if idx > len(lines) {
			//如果数据库条数大于文本条数 ,文本重读
			txtIdx = 0
		}
		kw1 := lines[txtIdx]
		kw2 := kw1
		kw3 := RuleKeyws(kw1)

		var cc Content
		db.Raw("UPDATE Content SET 标题= ?,关键词= ?,副标题= ? WHERE ID = ? RETURNING ID", kw1, kw2, kw3, value.ID).Scan(&cc)
		logger.Info("更新ID:", cc.ID)

		txtIdx++
	}

}
