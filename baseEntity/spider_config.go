package baseEntity

type SymbolConfig struct {
	Key   string `bson:"key"`
	Value any    `bson:"value"`
}
