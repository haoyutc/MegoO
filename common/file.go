package common

import (
	"os"
	"path/filepath"
)

var (
	AppPath  string
	WorkPath string
)

// FileExists 文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// GetAppPath 获取进程文件所在目录
func GetAppPath() string {
	if len(AppPath) == 0 {
		appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		AppPath = appPath
	}
	return AppPath
}

// GetWorkPath 获取运行main.go所在目录
func GetWorkPath() string {
	if len(WorkPath) == 0 {
		workPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		WorkPath = workPath
	}
	return WorkPath
}

// GetLogPath 获取日志目录
func GetLogPath(module string) string {
	appPath := GetAppPath()
	workPath := GetWorkPath()

	logPath := filepath.Join(workPath, module, "logger")
	if !FileExists(logPath) {
		logPath = filepath.Join(appPath, "logger")
		if !FileExists(logPath) {
			return ""
		}
	}
	return logPath
}
