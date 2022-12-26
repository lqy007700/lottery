package conf

type DBConfig struct {
	Host      string
	Port      int32
	User      string
	Pwd       string
	Databases string
}

var DBMasterList = []DBConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "root",
		Databases: "lottery",
	},
}

var DBMaster = DBMasterList[0]
