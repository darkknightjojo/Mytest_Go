package config

// 配置文件解析
import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Config  = tomlConfig{}
	ViewDir string
	LogDir  string
)

type tomlConfig struct {
	RootDir string
	Port    string
	DBDebug bool
}

type configByOS struct {
	Windows tomlConfig
	//OSX		tomlConfig
}

func init() {
	var configOS configByOS
	if _, err := toml.DecodeFile("config.toml", &configOS); err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		Config = configOS.Windows
	}

	ViewDir = filepath.Join(Config.RootDir, "templates")
	LogDir = filepath.Join(Config.RootDir, "logs")

	_ = os.MkdirAll(LogDir, os.ModePerm)
}
