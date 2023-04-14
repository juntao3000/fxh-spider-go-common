package baseEnums

type ContractType int

const (
	ContractTypeUnknown        ContractType = 0
	ContractTypeSpot           ContractType = 1
	ContractTypeSwap           ContractType = 2
	ContractTypeCurrentWeek    ContractType = 3
	ContractTypeNextWeek       ContractType = 4
	ContractTypeCurrentQuarter ContractType = 5
	ContractTypeNextQuarter    ContractType = 6
)

func (t ContractType) String() string {
	switch t {
	case ContractTypeSpot:
		return "spot"

	case ContractTypeSwap:
		return "swap"

	case ContractTypeCurrentWeek:
		return "cw"

	case ContractTypeNextWeek:
		return "nw"

	case ContractTypeCurrentQuarter:
		return "cq"

	case ContractTypeNextQuarter:
		return "nq"

	default:
		return ""
	}
}
