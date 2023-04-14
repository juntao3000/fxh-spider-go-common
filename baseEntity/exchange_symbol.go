package baseEntity

import (
	"fmt"
	"github.com/juntao3000/fxh-spider-go-common/baseEnums"
	"strings"
	"time"
)

type ExchangeSymbol struct {
	ExchangeCode baseEnums.ExchangeCode `bson:"exchange_code"`
	SymbolType   baseEnums.SymbolType   `bson:"symbol_type"`
	ContractType baseEnums.ContractType `bson:"contract_type"`
	Symbol       string                 `bson:"symbol"`
	Pair         string                 `bson:"pair,omitempty"`
	ContractSize float64                `bson:"contract_size,omitempty"`
	Pair1        string                 `bson:"pair1"`
	Pair2        string                 `bson:"pair2"`
	Status       string                 `bson:"status,omitempty"`
	CoinCode     string                 `bson:"coin_code,omitempty"`
	UpdateTime   time.Time              `bson:"update_time"`
}

func (s *ExchangeSymbol) TickerId() string {
	// ticker_id = ({exchange_code}_{pair1}_{contract_type}_{pair2}).lower()

	contractType := s.ContractType.String()

	tickerId := ""
	if len(contractType) == 0 || s.ContractType == baseEnums.ContractTypeSpot {
		tickerId = fmt.Sprintf("%s_%s_%s", s.ExchangeCode, s.Pair1, s.Pair2)
	} else {
		tickerId = fmt.Sprintf("%s_%s_%s_%s", s.ExchangeCode, s.Pair1, contractType, s.Pair2)
	}

	return strings.ToLower(tickerId)
}
