package baseModel

type AmplitudeStage struct {
	Name      string   `json:"name"`
	Begin     float64  `json:"begin"`
	End       float64  `json:"end"`
	Pair1List []string `json:"pair1List"`
}
