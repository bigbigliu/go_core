package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Config 表示配置文件的结构体
type Config struct {
	App    *AppConf    `yaml:"app"`    // App app配置
	Db     *DbConf     `yaml:"db"`     // Db 数据库配置
	Redis  *RedisConf  `yaml:"redis"`  // Redis redis配置
	Logger *LoggerConf `yaml:"logger"` // Logger 日志配置
	Jwt    *JwtConf    `yaml:"jwt"`    // Jwt jwt配置
}

// AppConf app配置
type AppConf struct {
	Host string `yaml:"host"` // Host app启动host
	Port int    `yaml:"port"` // Port app启动port
	Mode string `yaml:"mode"` // Mode gin启动模式
}

// DbConf 数据库配置
type DbConf struct {
	Host               string `yaml:"host"`                 // Host 数据库服务host
	Port               int    `yaml:"port"`                 // Port 数据库服务port
	User               string `yaml:"user"`                 // User 数据库服务用户名
	Password           string `yaml:"password"`             // Password 数据库服务密码
	Name               string `yaml:"name"`                 // Name 数据库名
	MaxIdleConnections string `yaml:"max_idle_connections"` // MaxIdleConnections 设置空闲连接池中连接的最大数量
	MaxOpenConnections string `yaml:"max_open_connections"` // MaxOpenConnections 设置数据库的最大打开连接数
}

// RedisConf redis配置
type RedisConf struct {
	Addr string `yaml:"addr"` // Addr redis服务host
	Port int    `yaml:"port"` // Port redis服务port
	Pwd  string `yaml:"pwd"`  // Pwd redis服务密码
	Db   int    `yaml:"db"`   // Db redis服务数据库
}

// LoggerConf 日志配置
type LoggerConf struct {
	Path  string `yaml:"path"`  // Path 日志保存
	Level string `yaml:"level"` // Level 日志级别
}

// JwtConf jwt配置
type JwtConf struct {
	Secret  string `yaml:"secret"`  // Secret jwt密钥
	Timeout int    `yaml:"timeout"` // Timeout token过期时间 单位 秒/s
}

// SingletonConfig 是Config的唯一实例
var singletonConfig *Config

// LoadConfig 从文件加载配置信息
func loadConfig(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		os.Exit(-1)
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		fmt.Println("配置文件反序列化失败: ", err)
		os.Exit(-1)
	}

	singletonConfig = &config
	fmt.Println("配置文件加载成功")
}

// GetConfig 获取SingletonConfig实例
func GetConfig() *Config {
	return singletonConfig
}

// InitConfig 配置文件初始化方法
func init() {
	configFile := os.Getenv("CONFIG")
	if configFile == "" {
		configFile = "./config/settings.yml"
	}
	loadConfig(configFile)
}
