package baseModel

import "time"

type BaseConfig struct {
	FxhWebConfigBaseUrl string `json:"fxhWebConfigBaseUrl"`

	BrowserUseFireFox               bool     `json:"browserUseFireFox"`
	BrowserTimeoutSecond            int      `json:"browserTimeoutSecond"`
	BrowserUserAgent                string   `json:"browserUserAgent"`
	BrowserInitJs                   string   `json:"browserInitJs"`
	BrowserArgs                     []string `json:"browserArgs"`
	BrowserAbortRequestResourceType []string `json:"browserAbortRequestResourceType"`

	// http proxy
	HttpProxyHost string `json:"httpProxyHost"`
	HttpProxyPort int    `json:"httpProxyPort"`
	// socks5 Proxy
	Socks5ProxyHost string `json:"socks5ProxyHost"`
	Socks5ProxyPort int    `json:"socks5ProxyPort"`

	// hk http proxy
	HkHttpProxyHost string `json:"hkHttpProxyHost"`
	HkHttpProxyPort int    `json:"hkHttpProxyPort"`
	// hk socks5 Proxy
	HkSocks5ProxyHost string `json:"hkSocks5ProxyHost"`
	HkSocks5ProxyPort int    `json:"hkSocks5ProxyPort"`

	// ssr socks5 Proxy
	SsrSocks5ProxyHost string `json:"ssrSocks5ProxyHost"`
	SsrSocks5ProxyPort int    `json:"ssrSocks5ProxyPort"`

	MysqlConnectionString string `json:"mysqlConnectionString"`

	// mongodb
	DexMongoConnectionString string `json:"dexMongoConnectionString"`
	MongoConnectionString    string `json:"mongoConnectionString"`
	MongoBatchSize           int    `json:"mongoBatchSize"`
	MongoSaveIntervalSecond  int    `json:"mongoSaveIntervalSecond"`

	// dex redis
	DexRedisHost     string `json:"dexRedisHost"`
	DexRedisPort     int    `json:"dexRedisPort"`
	DexRedisPassword string `json:"dexRedisPassword"`
	DexRedisDatabase int    `json:"dexRedisDatabase"`

	// redis
	RedisHost     string `json:"redisHost"`
	RedisPort     int    `json:"redisPort"`
	RedisPassword string `json:"redisPassword"`
	RedisDatabase int    `json:"redisDatabase"`

	// oss
	OssEndPoint   string `json:"ossEndPoint"`
	OssBucketName string `json:"ossBucketName"`
	OssAccessId   string `json:"ossAccessId"`
	OssAccessKey  string `json:"ossAccessKey"`

	// clickhouse
	ClickHouseProtocol          int    `json:"clickHouseProtocol"` //Native=0,Http=1
	ClickHouseHost              string `json:"clickHouseHost"`
	ClickHousePort              int    `json:"clickHousePort"`
	ClickHouseUser              string `json:"clickHouseUser"`
	ClickHousePassword          string `json:"clickHousePassword"`
	ClickHouseDatabase          string `json:"clickHouseDatabase"`
	ClickHouseMaxOpenConnection int    `json:"clickHouseMaxOpenConnection"` //default=5

	// chain data
	ChainDataMongoConnectionString string `json:"chainDataMongoConnectionString"`
	ChainDataRedisHost             string `json:"chainDataRedisHost"`
	ChainDataRedisPort             int    `json:"chainDataRedisPort"`
	ChainDataRedisPassword         string `json:"chainDataRedisPassword"`
	ChainDataRedisDatabase         int    `json:"chainDataRedisDatabase"`
}

func (c *BaseConfig) GetMongoSaveDuration() time.Duration {
	return time.Second * time.Duration(c.MongoSaveIntervalSecond)
}
