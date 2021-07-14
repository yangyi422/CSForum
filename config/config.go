package config

import (
	"time"
)

// 数据库连接信息
type Db_config struct {
	IP             string            `json:"ip" yaml:"ip"`
	PORT           string            `json:"port" yaml:"port"`
	Config         map[string]string `json:"config" yaml:"config"`
	Db_name        string            `json:"db_name" yaml:"db_name"`
	User_name      string            `json:"user_name" yaml:"user_name"`
	Password       string            `json:"password" yaml:"password"`
	Max_idle_conns int               `json:"max_idle_conns" yaml:"max_idle_conns"`
	Max_open_conns int               `json:"max_open_conns" yaml:"max_open_conns"`
	Max_lifetime   time.Duration     `json:"Max_lifetime" yaml:"Max_lifetime"`
	Db_type        string            `json:"db_type" yaml:"db_type"`
	Log_type       string            `json:"log_type" yaml:"log_type"`
}

// 数据库信息
var (
	MYSQL_PASSWORD string = "123456"
	MYSQL_IP       string = "47.106.215.114"
	MYSQL_PORT     string = "3306"
	snk_db_config         = map[string]string{
		"charset":          "utf8",
		"maxAllowedPacket": "1073741824",
		"multiStatements":  "true",
		"parseTime":        "true",
		"loc":              "Local",
	}
	CS_forum_db = &Db_config{
		IP:             MYSQL_IP,
		PORT:           MYSQL_PORT,
		Config:         snk_db_config,
		Db_name:        "CSForum",
		User_name:      "root",
		Password:       MYSQL_PASSWORD,
		Max_idle_conns: 21,
		Max_open_conns: 20,
		Max_lifetime:   120 * time.Second,
	}
)

// 获取数据库连接
func (m *Db_config) Get_dsn() string {
	var settings string
	if len(m.Config) > 0 {
		for k, v := range m.Config {
			setting := k + "=" + v + "&"
			settings += setting
		}
		settings = settings[:len(settings)-1]
	}
	if len(m.User_name) == 0 && len(m.Password) == 0 {
		return "tcp://" + m.IP + ":" + m.PORT + "?database=" + m.Db_name + "&" + settings
	} else {
		return m.User_name + ":" + m.Password + "@tcp(" + m.IP + ":" + m.PORT + ")/" + m.Db_name + "?" + settings
	}
}
