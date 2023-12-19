package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type dsn struct {
	MysqlDSN          string
	RedisDSN          string
	RabbitDSN         string
	EtcdDSN           string
	UserDB            string
	SignDB            string
	ChooseDB          string
	OrderDB           string
	RabbitVhost       string
	UserNameDB        string
	PassWordDB        string
	UserNameRabbit    string
	PassWordRabbit    string
	BaseService       string
	SignService       string
	ChooseService     string
	ApiServiceName    string
	BaseServiceName   string
	SignServiceName   string
	ChooseServiceName string
}

type Slice struct {
	Mod   int64
	Slice map[string][]string
}

type SnowFlow struct {
	BeginStamp    int64
	WorkId        int64
	TimeStampBits int
	WorkIdBits    int
	SequenceBits  int
}

type Cache struct {
	User     int
	Group    int
	Prize    int
	Activity int
	Sign     int
	Order    int
}

type Config struct {
	DSN         *dsn
	UserSlice   *Slice
	GroupSlice  *Slice
	SignSlice   *Slice
	ChooseSlice *Slice
	SnowFlow    *SnowFlow
	Cache       *Cache
	JwtSecret   string
}

var GlobalConfig *Config

func NewConfig() *Config {
	var conf *Config
	config := viper.New()
	config.SetConfigFile("conf/conf.json")
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("找不到配置文件")
		} else {
			panic(err)
		}
	}
	conf = &Config{
		DSN: &dsn{
			MysqlDSN:          config.GetString("dsn.mysqlDSN"),
			RedisDSN:          config.GetString("dsn.redisDSN"),
			RabbitDSN:         config.GetString("dsn.rabbitDSN"),
			EtcdDSN:           config.GetString("dsn.etcdDSN"),
			UserDB:            config.GetString("dsn.userDB"),
			SignDB:            config.GetString("dsn.signDB"),
			ChooseDB:          config.GetString("dsn.chooseDB"),
			OrderDB:           config.GetString("dsn.orderDB"),
			RabbitVhost:       config.GetString("dsn.rabbitVhost"),
			UserNameDB:        config.GetString("dsn.userNameDB"),
			PassWordDB:        config.GetString("dsn.passWordDB"),
			UserNameRabbit:    config.GetString("dsn.userNameRabbit"),
			PassWordRabbit:    config.GetString("dsn.passwordRabbit"),
			BaseService:       config.GetString("dsn.baseService"),
			SignService:       config.GetString("dsn.signService"),
			ChooseService:     config.GetString("dsn.chooseService"),
			ApiServiceName:    config.GetString("dsn.apiServiceName"),
			BaseServiceName:   config.GetString("dsn.baseServiceName"),
			SignServiceName:   config.GetString("dsn.signServiceName"),
			ChooseServiceName: config.GetString("dsn.chooseServiceName"),
		},
		UserSlice: &Slice{
			Mod:   config.GetInt64("userSlice.mod"),
			Slice: config.GetStringMapStringSlice("userSlice.slice"),
		},
		GroupSlice: &Slice{
			Mod:   config.GetInt64("groupSlice.mod"),
			Slice: config.GetStringMapStringSlice("groupSlice.slice"),
		},
		SignSlice: &Slice{
			Mod:   config.GetInt64("signSlice.mod"),
			Slice: config.GetStringMapStringSlice("signSlice.slice"),
		},
		ChooseSlice: &Slice{
			Mod:   config.GetInt64("chooseSlice.mod"),
			Slice: config.GetStringMapStringSlice("chooseSlice.slice"),
		},
		SnowFlow: &SnowFlow{
			BeginStamp:    config.GetInt64("snowflow.beginStamp"),
			WorkId:        config.GetInt64("snowflow.workId"),
			TimeStampBits: config.GetInt("snowflow.timeStampBits"),
			WorkIdBits:    config.GetInt("snowflow.workIdBits"),
			SequenceBits:  config.GetInt("snowflow.sequenceBits"),
		},
		Cache: &Cache{
			User:     config.GetInt("cache.user"),
			Group:    config.GetInt("cache.group"),
			Prize:    config.GetInt("cache.prize"),
			Activity: config.GetInt("cache.activity"),
			Sign:     config.GetInt("cache.sign"),
			Order:    config.GetInt("cache.order"),
		},
		JwtSecret: config.GetString("jwtSecret"),
	}
	config.WatchConfig()
	config.OnConfigChange(func(in fsnotify.Event) {
		if err := config.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				panic("找不到配置文件")
			} else {
				panic(err)
			}
		}
		conf = &Config{
			DSN: &dsn{
				MysqlDSN:          config.GetString("dsn.mysqlDSN"),
				RedisDSN:          config.GetString("dsn.redisDSN"),
				RabbitDSN:         config.GetString("dsn.rabbitDSN"),
				EtcdDSN:           config.GetString("dsn.etcdDSN"),
				UserDB:            config.GetString("dsn.userDB"),
				SignDB:            config.GetString("dsn.signDB"),
				ChooseDB:          config.GetString("dsn.chooseDB"),
				OrderDB:           config.GetString("dsn.orderDB"),
				RabbitVhost:       config.GetString("dsn.rabbitVhost"),
				UserNameDB:        config.GetString("dsn.userNameDB"),
				PassWordDB:        config.GetString("dsn.passWordDB"),
				BaseService:       config.GetString("dsn.baseService"),
				SignService:       config.GetString("dsn.signService"),
				ChooseService:     config.GetString("dsn.chooseService"),
				ApiServiceName:    config.GetString("dsn.apiServiceName"),
				BaseServiceName:   config.GetString("dsn.baseServiceName"),
				SignServiceName:   config.GetString("dsn.signServiceName"),
				ChooseServiceName: config.GetString("dsn.chooseServiceName"),
			},
			UserSlice: &Slice{
				Mod:   config.GetInt64("userSlice.mod"),
				Slice: config.GetStringMapStringSlice("userSlice.slice"),
			},
			GroupSlice: &Slice{
				Mod:   config.GetInt64("groupSlice.mod"),
				Slice: config.GetStringMapStringSlice("groupSlice.slice"),
			},
			SignSlice: &Slice{
				Mod:   config.GetInt64("signSlice.mod"),
				Slice: config.GetStringMapStringSlice("signSlice.slice"),
			},
			SnowFlow: &SnowFlow{
				BeginStamp:    config.GetInt64("snowflow.beginStamp"),
				WorkId:        config.GetInt64("snowflow.workId"),
				TimeStampBits: config.GetInt("snowflow.timeStampBits"),
				WorkIdBits:    config.GetInt("snowflow.workIdBits"),
				SequenceBits:  config.GetInt("snowflow.sequenceBits"),
			},
			Cache: &Cache{
				User:     config.GetInt("cache.user"),
				Group:    config.GetInt("cache.group"),
				Prize:    config.GetInt("cache.prize"),
				Activity: config.GetInt("cache.activity"),
				Sign:     config.GetInt("cache.sign"),
				Order:    config.GetInt("cache.order"),
			},
			JwtSecret: config.GetString("jwtSecret"),
		}
	})
	return conf
}

func init() {
	if GlobalConfig == nil {
		GlobalConfig = NewConfig()
	}
}
