package baseEnums

type ExchangeCode string

const (
	ExchangeCodeUnknown ExchangeCode = ""
	ExchangeCodeBinance ExchangeCode = "binance"
	ExchangeCodeHuobi   ExchangeCode = "huobipro"
	ExchangeCodeOkex    ExchangeCode = "okex"
	ExchangeCodeGateio  ExchangeCode = "gate-io"
	ExchangeCodeBitmex  ExchangeCode = "bitmex"
	ExchangeCodeAax     ExchangeCode = "aax"
)
