package baseModel

type AmplitudeInfo struct {
	Time         int64   `json:"Time"`
	Exchange     string  `json:"Exchange"`
	Pair1        string  `json:"Pair1"`
	Pair2        string  `json:"Pair2"`
	Title        string  `json:"Title"`
	Price        float64 `json:"Price"`
	Open24H      float64 `json:"Open_24h"`
	High24H      float64 `json:"High_24h"`
	Low24H       float64 `json:"Low_24h"`
	Volume24H    float64 `json:"Volume_24h"`
	Amplitude24h float64 `json:"Amplitude_24h"`
}
