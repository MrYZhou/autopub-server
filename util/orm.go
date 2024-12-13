package util

import (
	"log"
	"os"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/bwmarrin/snowflake"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var node *snowflake.Node

func init() {
	// 加载 .env 文件
	godotenv.Load()
	dbUrl := os.Getenv("dbUrl")
	DbInit("default", dbUrl)
	DbChange("default")
	node, _ = snowflake.NewNode(1)
}

func GetId() string {
	return node.Generate().String()
}

var gormDb *gorm.DB
var gormDbMap = make(map[string]*gorm.DB)

/*
从gormDbMap中切换数据库

tag 库名key
*/
func DbChange(tag string) {
	// gplus的连接的数据库
	gormDb := gormDbMap[tag]
	gplus.Init(gormDb)
}

/*
初始化库存在gormDbMap中

tag 库名key ,url数据库地址
*/
func DbInit(tag string, url string) {
	var err error
	gormDb, err = gorm.Open(mysql.Open(url+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 数据库表前缀
			SingularTable: true,  // 不用给表名加复数
			NoLowerCase:   false, // 要不要把表名全小写 .false 默认,true 转小写
		},
		// Logger: logger.Discard, // 不输出日志
	})
	if err != nil {
		log.Println(err)
	}
	gormDbMap[tag] = gormDb
}
