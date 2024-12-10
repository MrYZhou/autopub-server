package main

import (
	"testing"

	"log"
	"time"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var node *snowflake.Node

func getId() string {

	node, _ = snowflake.NewNode(1)
	return node.Generate().String()
}

type Users struct {
	ID        int64 
	Username  string
	Password  string
	Address   string
	Age       int
	Phone     string
	Score     int
	Dept      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Config struct {
	Id   string   
	Name string
}

var gormDb *gorm.DB

func TestGorm(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",    // 数据库表前缀
			SingularTable: true,  // 不用给表名加复数
			NoLowerCase:   false, // 要不要把表名全小写 .false 默认,true 转小写
		},
	})
	if err != nil {
		log.Println(err)
	}
	// 初始化gplus
	gplus.Init(gormDb)

	users, resultDb := gplus.SelectList[Users](nil)
	log.Println("error:", resultDb.Error)
	log.Println("RowsAffected:", resultDb.RowsAffected)
	for _, user := range users {
		log.Println("user:", user)
	}

	config, _ := gplus.SelectById[Config](1)

	log.Println("config:", config)

}

