package util

import (
	"log"

	"github.com/acmestack/gorm-plus/gplus"
	"github.com/bwmarrin/snowflake"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var node *snowflake.Node
func getId() string{

	node, _ = snowflake.NewNode(1)
  return node.Generate().String()
}


var gormDb *gorm.DB
var gormDbMap = make(map[string]*gorm.DB)

func DbChange(tag string){
 // gplus的连接的数据库
 gormDb := gormDbMap[tag]
 gplus.Init(gormDb)
}
func DbInit(tag string,url string) {
  var err error
  gormDb, err = gorm.Open(mysql.Open(url+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
    NamingStrategy: schema.NamingStrategy{
      TablePrefix: "",   // 数据库表前缀
      SingularTable: true, // 不用给表名加复数
      NoLowerCase: false, // 要不要把表名全小写 .false 默认,true 转小写
    },
    // Logger: logger.Discard, // 不输出日志
  })
  if err != nil {
    log.Println(err)
  }
  gormDbMap[tag] = gormDb 
}