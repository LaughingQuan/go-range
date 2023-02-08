package sqltargets

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"xmirror.cn/iast/goat/config"
	"xmirror.cn/iast/goat/util"
)

const DbFile string = "tempDatabase.db"
const SqlFile string = "db.sql"

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		initConn()
	}
	return db
}

type Vulnerability struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Severity int    `json:"level"`
	Desc     string `json:"desc"`
}

func (Vulnerability) TableName() string {
	return "vulnerability"
}

var (
	sampleItems = []Vulnerability{
		{
			ID:       1,
			Name:     "不安全的加密算法",
			Severity: 1,
			Desc:     "应用使用了不安全的或加密强度弱的加密算法，使被加密数据有可能被攻击者破解。",
		},
		{
			ID:       2,
			Name:     "违反信任边界规则",
			Severity: 2,
			Desc:     "应用将未处理的外部数据被存到可信区域中，如Session中。",
		},
		{
			ID:       3,
			Name:     "XStream反序列化",
			Severity: 3,
			Desc:     "XStream可以将对象序列化成XML或将XML反序列化为对象。在使用XStream进行反序列化时，如果程序在对外部数据反序列化时，没有校验，会导致反序列化漏洞。",
		},
		{
			ID:       4,
			Name:     "Headers敏感信息泄露",
			Severity: 0,
			Desc:     "网站请求响应中Headers的字段会显示网站web中间件或开发框架版本信息。",
		},
		{
			ID:       5,
			Name:     "服务端模板注入漏洞",
			Severity: 3,
			Desc:     "服务端模板注入和常见Web注入的成因一样，也是服务端接收了用户的输入，将其作为 Web 应用模板内容的一部分，在进行目标编译渲染的过程中，执行了用户插入的恶意内容，因而可能导致了敏感信息泄露、代码执行、GetShell等问题，其影响范围主要取决于模版引擎的复杂性。",
		},
	}
)

func initConn() {
	var err error
	mysqlConf, err := config.GetConfMysql()
	if err != nil {
		log.Printf("failed to get the database config %s", err)
		return
	}
	dnf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)
	db, err = gorm.Open(mysql.Open(dnf), &gorm.Config{})
	if err == nil {
		return
	}
	sqlPath := filepath.Join(util.StartDir(), DbFile)
	db, err = gorm.Open(sqlite.Open(sqlPath), &gorm.Config{})
	if err != nil {
		log.Printf("failed to link to the database %s", err)
		return
	}
	backExecTable()
}

func backExecTable() {
	err := db.AutoMigrate(&Vulnerability{})
	if err != nil {
		log.Println(err)
		return
	}

	tx := db.Model(&Vulnerability{}).Create(sampleItems)
	if tx.Error != nil {
		log.Println(err)
		return
	}
}

func DBCloseAndRemove() error {
	exPath := util.StartDir()
	sqlPath, _ := filepath.Abs(filepath.Join(exPath, DbFile))
	os.Remove(sqlPath)
	return nil
}
