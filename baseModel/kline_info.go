package baseModel

import "github.com/juntao3000/fxh-spider-go-common/baseEntity"

type KlineInfo struct {
	Period string
	Kline  *baseEntity.Kline
}
