package baseEntity

type FxhCoinCode struct {
	TickerId     string `bson:"ticker_id,omitempty"`
	CoinCode     string `bson:"coin_code,omitempty"`
	ExchangeCode string `bson:"exchange_code,omitempty"`
	CoinSymbol   string `bson:"coin_symbol,omitempty"`
	Pair1        string `bson:"pair1,omitempty"`
	Pair2        string `bson:"pair2,omitempty"`
	Title        string `bson:"title,omitempty"`
}
