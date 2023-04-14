package baseConst

import "github.com/juntao3000/fxh-spider-go-common/baseCommon"

const (
	BaseConfigFilePath = "FXH_BASE_CONFIG_FILE_PATH"
	Proxy3ConfigPath   = "FXH_PROXY3_CONFIG_PATH"
)

func GetConfigDir() string {
	if baseCommon.IsDebug() {
		return "d:\\git\\rx\\fxh_spider\\conf"
	}

	return "/etc/spider"
}

func GetDefaultBaseConfigFilePath() string {
	if baseCommon.IsDebug() {
		return "d:\\git\\rx\\fxh_spider\\conf\\go-config\\base.yaml"
	}

	return "/etc/spider/go-config/base.yaml"
}

func GetDefaultProxy3ConfigDir() string {
	if baseCommon.IsDebug() {
		return "d:\\git\\rx\\fxh_spider\\conf\\proxy\\3proxy"
	}

	return "/etc/spider/proxy/3proxy"
}
