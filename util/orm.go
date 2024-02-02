package util

import (
	"log"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/bwmarrin/snowflake"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var node *snowflake.Node
func getId() string{

	node, _ = snowflake.NewNode(1)
  return node.Generate().String()
}


var gormDb *gorm.DB
func DbInit() {
  dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
  var err error
  gormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
  })
  if err != nil {
    log.Println(err)
  }

  // 初始化gplus
  gplus.Init(gormDb)
}