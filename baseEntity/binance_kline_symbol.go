package baseEntity

import (
	"github.com/juntao3000/fxh-spider-go-common/baseEnums"
	"time"
)

type BinanceKlineSymbol struct {
	ExchangeCode   baseEnums.ExchangeCode `bson:"exchange_code"`
	ExchangeSymbol baseEnums.ExchangeCode `bson:"exchange_symbol"`
	CoinCode       baseEnums.ExchangeCode `bson:"coin_code"`
	ContractType   baseEnums.ContractType `bson:"contract_type"`
	SymbolType     baseEnums.SymbolType   `bson:"symbol_type"`
	Pair1          string                 `bson:"pair1"`
	Pair2          string                 `bson:"pair2"`
	UpdateTime     time.Time              `bson:"update_time"`
}
