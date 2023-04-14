package baseUtil

import (
	"github.com/juntao3000/fxh-spider-go-common/baseEnums"
	"strings"
)

func ParseSymbolTypeByPair2(pair2 string) baseEnums.SymbolType {
	pair2 = strings.TrimSpace(pair2)

	if strings.EqualFold(pair2, "USDT") {
		return baseEnums.SymbolTypeFutures
	}

	if strings.EqualFold(pair2, "USD") {
		return baseEnums.SymbolTypeDelivery
	}

	return baseEnums.SymbolTypeUnknown
}

func ParseSymbolType(symbolType string) baseEnums.SymbolType {
	symbolType = strings.TrimSpace(symbolType)

	if strings.EqualFold(symbolType, "futures") {
		return baseEnums.SymbolTypeFutures
	}

	if strings.EqualFold(symbolType, "delivery") {
		return baseEnums.SymbolTypeDelivery
	}

	if strings.EqualFold(symbolType, "spot") {
		return baseEnums.SymbolTypeSpot
	}

	return baseEnums.SymbolTypeUnknown
}

func ParseContractType(contractType string) baseEnums.ContractType {
	contractType = strings.TrimSpace(contractType)

	if strings.EqualFold(contractType, "spot") {
		return baseEnums.ContractTypeSpot
	}

	if strings.EqualFold(contractType, "swap") {
		return baseEnums.ContractTypeSwap
	}

	if strings.EqualFold(contractType, "cw") {
		return baseEnums.ContractTypeCurrentWeek
	}

	if strings.EqualFold(contractType, "nw") {
		return baseEnums.ContractTypeNextWeek
	}

	if strings.EqualFold(contractType, "cq") {
		return baseEnums.ContractTypeCurrentQuarter
	}

	if strings.EqualFold(contractType, "nq") {
		return baseEnums.ContractTypeNextQuarter
	}

	return baseEnums.ContractTypeUnknown
}
