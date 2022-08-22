package config

const (
	ServiceName = ""
	ServicePort = ":8080"
	LogName     = "platform.logger"
	Module      = "MegoO"
)

type ConfYaml struct {
	CommYaml
}

type CommYaml struct {
	DB  *Database `json:"db"`
	Log *Log      `json:"log"`
}

// Database config
type Database struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Database    string `json:"database"`
	Debug       bool   `json:"debug"`
	MaxOpenConn int    `json:"maxOpenConn"`
	MaxIdleConn int    `json:"maxIdleConn"`
	MaxLifetime int    `json:"maxLifetime"`
}

type Log struct {
	Level             string `json:"level"`
	Out               string `json:"out"`
	Format            string `json:"format"`
	LogPath           string `json:"logPath"`
	LogName           string `json:"logName"`
	MaxAge            int    `json:"maxAge"`
	RotationTime      int    `json:"rotationTime"`
	LogSize           int    `json:"logSize"`
	ReportCaller      bool   `json:"reportCaller"`
	DbUser            string `json:"username"`
	DbPassword        string `json:"password"`
	DbHost            string `json:"host"`
	MaxOpenConn       int    `json:"maxOpenConn"`
	MaxIdleConn       int    `json:"maxIdleConn"`
	MaxLifetime       int    `json:"maxLifetime"`
	TableNameUseMonth bool   `json:"tableNameUseMonth"`
}
