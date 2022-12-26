package conf

type RedisConfig struct {
	Host string
	Port int32
	User string
	Pwd  string
}

var RedisCacheList = []RedisConfig{
	{
		Host: "127.0.0.1",
		Port: 6379,
		User: "",
		Pwd:  "",
	},
}

var RedisCache = RedisCacheList[0]
