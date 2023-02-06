package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PrefixUrl       string
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type MySQL struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var MySQLDatabaseSetting = &MySQL{}

type MongoDB struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var MongoDBDatabaseSetting = &MongoDB{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type WsServer struct {
	Port             string
	Ping             time.Duration
	Pong             time.Duration
	MaxWriteWaitTime time.Duration
	MaxMessageSize   int64
}

var WsServerSetting = &WsServer{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("mysql", MySQLDatabaseSetting)
	mapTo("mongodb", MongoDBDatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("wsServer", WsServerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	WsServerSetting.Ping = (WsServerSetting.Ping * time.Second * 9) / 10
	WsServerSetting.Pong = WsServerSetting.Pong * time.Second
	WsServerSetting.MaxWriteWaitTime = WsServerSetting.MaxWriteWaitTime * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
