package initialization

import (
	"CSForum/config"
	model_database "CSForum/model/database"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	init_mysql()
	get_mysql_tables(DB)
}

func init_mysql() {
	m := config.CS_forum_db
	dsn := m.Get_dsn()
	Config := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    false, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	var err error
	if DB, err = gorm.Open(mysql.New(Config)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(m.Max_idle_conns)
		sqlDB.SetMaxOpenConns(m.Max_open_conns)
		sqlDB.SetConnMaxLifetime(m.Max_lifetime)
	}
}

func get_mysql_tables(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(
		new(model_database.User),
	)
	if err != nil {
		fmt.Println("register mysql table failed:" + err.Error())
		panic("register mysql table failed:" + err.Error())
	}
}