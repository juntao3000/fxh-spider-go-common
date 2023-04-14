package baseEntity

import (
	"time"
)

type Kline struct {
	ExchangeSymbol string    `bson:"exchange_symbol"`
	CoinCode       string    `bson:"coin_code"`
	Pair1          string    `bson:"pair1"`
	Pair2          string    `bson:"pair2"`
	Timestamp      time.Time `bson:"timestamp"`
	Open           float64   `bson:"open"`
	High           float64   `bson:"high"`
	Low            float64   `bson:"low"`
	Close          float64   `bson:"close"`
	Volume         float64   `bson:"volume"`
	Amount         float64   `bson:"amount"`
}
