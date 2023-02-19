package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Email  Email  `yaml:"email" json:"email" mapstructure:"email"`
	System System `yaml:"system" json:"system" mapstructure:"system"`
}

type System struct {
	Port int `mapstructure:"port" json:"port" yaml:"port"`
}
type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // 收件人:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 发件人  你自己要发邮件的邮箱
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 昵称    发件人昵称 通常为自己的邮箱
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`        // 是否SSL   是否开启SSL
}

var EmailConf *Config

func Init() error {

	// main.go文件的绝对路径
	mainDirectory, _ := os.Getwd()
	path := mainDirectory + "/config.yml"

	// 读文件的时候必须给一个文件的名称或者地址
	viper.SetConfigFile(path)
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	// 把读取到的配置信息反序列化到 EmailConf 变量中
	if err = viper.Unmarshal(EmailConf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(EmailConf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return nil
}
