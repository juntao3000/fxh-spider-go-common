package baseEnums

type SymbolType int

const (
	SymbolTypeUnknown  SymbolType = 0
	SymbolTypeSpot     SymbolType = 1
	SymbolTypeFutures  SymbolType = 2
	SymbolTypeDelivery SymbolType = 3
)

func (t SymbolType) String() string {
	switch t {
	case SymbolTypeSpot:
		return "spot"

	case SymbolTypeFutures:
		return "futures"

	case SymbolTypeDelivery:
		return "delivery"

	default:
		return ""
	}
}
