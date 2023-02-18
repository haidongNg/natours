package setting

import (
	"time"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// Setup initialize the configuration instance
func Setup() {
	AppSetting.JwtSecret = "233"
	AppSetting.PageSize = 10
	AppSetting.PrefixUrl = "http://127.0.0.1:8000"

	AppSetting.RuntimeRootPath = "runtime/"

	AppSetting.LogSavePath = "logs/"
	AppSetting.LogSaveName = "log"
	AppSetting.LogFileExt = "log"
	AppSetting.TimeFormat = "20060102"

	ServerSetting.RunMode = "debug"
	ServerSetting.HttpPort = 8000
	ServerSetting.ReadTimeout = 60
	ServerSetting.WriteTimeout = 60

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}
