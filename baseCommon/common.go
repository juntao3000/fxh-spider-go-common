package baseCommon

import (
	"context"
	"github.com/juntao3000/fxh-spider-go-common/baseModel"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	SysExitChan                 = make(chan os.Signal, 1)
	GlobalErrChan               = make(chan error, 1)
	GlobalCtx, GlobalCancelFunc = context.WithCancel(context.Background())

	BaseConfig *baseModel.BaseConfig
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	signal.Notify(SysExitChan, syscall.SIGINT, syscall.SIGTERM)

	BaseConfig = &baseModel.BaseConfig{}
	err := GetBaseConfig(BaseConfig)
	if err != nil {
		logrus.Errorf("get base config failed:%s", err.Error())
	}

	hostIpEnv := os.Getenv("HOST_IP")
	hostIpEnv = strings.TrimSpace(hostIpEnv)
	if len(hostIpEnv) > 0 {
		logrus.Debugf("replace proxyHost with HOST_IP(env),old HttpProxyHost=%s,old Socks5ProxyHost=%s,HOST_IP=%s", BaseConfig.HttpProxyHost, BaseConfig.Socks5ProxyHost, hostIpEnv)

		BaseConfig.HttpProxyHost = hostIpEnv
		BaseConfig.Socks5ProxyHost = hostIpEnv
	}
}

func Dispose() {
}
