package baseService

import (
	"context"
	"fmt"
	"github.com/juntao3000/fxh-spider-go-common/baseCommon"
	"github.com/juntao3000/fxh-spider-go-common/baseConst"
	"github.com/juntao3000/fxh-spider-go-common/baseEntity"
	"github.com/juntao3000/fxh-spider-go-common/baseEnums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

var (
	MongoClient *mongo.Client
)

func DisposeMongo() {
	if MongoClient != nil {
		_ = MongoClient.Disconnect(context.Background())
	}
}

func InitMongo(isDex bool) error {
	if MongoClient != nil {
		return nil
	}

	var err error
	if isDex {
		MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(baseCommon.BaseConfig.DexMongoConnectionString))
	} else {
		MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(baseCommon.BaseConfig.MongoConnectionString))
	}

	if err != nil {
		return err
	}

	return MongoClient.Ping(context.Background(), nil)
}

func ListExchangeSymbol(exchangeCode baseEnums.ExchangeCode, symbolType baseEnums.SymbolType,
	contractType baseEnums.ContractType, status string) ([]*baseEntity.ExchangeSymbol, error) {
	coll := MongoClient.
		Database(baseConst.MongoDbNameSpider).
		Collection(baseConst.MongoCollNameSpiderExchangeSymbol)

	filter := bson.D{
		{"exchange_code", exchangeCode},
		{"symbol_type", symbolType},
	}
	if contractType != baseEnums.ContractTypeUnknown {
		filter = append(filter, bson.E{Key: "contract_type", Value: contractType})
	}
	if len(status) != 0 {
		filter = append(filter, bson.E{Key: "status", Value: status})
	}

	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("get data from db error:%v", err)
	}
	defer func() {
		_ = cursor.Close(context.Background())
	}()

	list := make([]*baseEntity.ExchangeSymbol, 0)
	for cursor.Next(context.Background()) {
		var info baseEntity.ExchangeSymbol
		if err := cursor.Decode(&info); err != nil {
			return nil, fmt.Errorf("decode data from db error:%v", err)
		}
		list = append(list, &info)
	}

	return list, nil
}

func ListExchangeSymbolName(exchangeCode baseEnums.ExchangeCode, symbolType baseEnums.SymbolType,
	contractType baseEnums.ContractType, status string) ([]string, error) {
	coll := MongoClient.
		Database(baseConst.MongoDbNameSpider).
		Collection(baseConst.MongoCollNameSpiderExchangeSymbol)

	filter := bson.D{
		{"exchange_code", exchangeCode},
		{"symbol_type", symbolType},
	}
	if contractType != baseEnums.ContractTypeUnknown {
		filter = append(filter, bson.E{Key: "contract_type", Value: contractType})
	}
	if len(status) != 0 {
		filter = append(filter, bson.E{Key: "status", Value: status})
	}

	opts := options.Find().SetProjection(bson.D{{"_id", 0}, {"symbol", 1}}) //.SetLimit(1)
	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("get data from db error:%v", err)
	}
	defer func() {
		_ = cursor.Close(context.Background())
	}()

	list := make([]struct {
		Symbol string `bson:"symbol"`
	}, 0)
	err = cursor.All(context.Background(), &list)
	if err != nil {
		return nil, fmt.Errorf("decode data from db error:%v", err)
	}

	strList := make([]string, 0)
	for _, item := range list {
		strList = append(strList, item.Symbol)
	}

	return strList, nil
}

func GetGetCoinCodeByPair1(pair1 string) (string, error) {
	if len(pair1) == 0 {
		return "", fmt.Errorf("invalid pair1")
	}

	coll := MongoClient.
		Database(baseConst.MongoDbNameSpider).
		Collection(baseConst.MongoCollNameSpiderFxhCoinCode)

	filterArr := bson.A{
		bson.D{{"coin_symbol", pair1}},
	}
	lower := strings.ToLower(pair1)
	if pair1 != lower {
		filterArr = append(filterArr, bson.D{{"coin_symbol", lower}})
	}
	upper := strings.ToUpper(pair1)
	if pair1 != upper {
		filterArr = append(filterArr, bson.D{{"coin_symbol", upper}})
	}
	filter := bson.D{{"$or", filterArr}}
	//
	//filter := bson.M{"$regex": primitive.Regex{Pattern: pair1, Options: "i"}}

	opts := options.Find().SetProjection(bson.D{{"_id", 0}, {"coin_code", 1}}) //.SetLimit(1)
	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		return "", fmt.Errorf("get data from db error:%v", err)
	}
	defer func() {
		_ = cursor.Close(context.Background())
	}()

	var coinCode string
	for cursor.Next(context.Background()) {
		var info struct {
			CoinCode string `bson:"coin_code"`
		}

		if err := cursor.Decode(&info); err != nil {
			return "", fmt.Errorf("decode data from db error:%v", err)
		}

		if len(info.CoinCode) != 0 {
			coinCode = info.CoinCode
			break
		}
	}

	return coinCode, nil
}

func GetUsdStableCoin() ([]string, error) {
	coll := MongoClient.
		Database(baseConst.MongoDbNameSpider).
		Collection(baseConst.MongoCollNameSpiderConfig)

	filter := bson.D{{"key", "usd_stable_coin_list"}}
	result := coll.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, fmt.Errorf("get data from db error:%v", result.Err())
	}
	var cfg struct {
		Key               string   `bson:"key"`
		UsdStableCoinList []string `bson:"value"`
	}
	if err := result.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode data from db error:%v", err)
	}

	return cfg.UsdStableCoinList, nil
}

func GetBscPublicHttpNodeList(getPro bool) ([]string, error) {
	const bscPublicHttpNodeConfigKey = "bscPublicHttpNode"
	const bscPublicHttpNodeProConfigKey = "bscPublicHttpNodePro"

	var key string
	if getPro {
		key = bscPublicHttpNodeProConfigKey
	} else {
		key = bscPublicHttpNodeConfigKey
	}

	coll := MongoClient.
		Database(baseConst.MongoDbNameSpider).
		Collection(baseConst.MongoCollNameSpiderConfig)

	filter := bson.D{{"key", key}}
	result := coll.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return nil, fmt.Errorf("get data from db error:%v", result.Err())
	}
	var cfg struct {
		Key  string   `bson:"key"`
		List []string `bson:"value"`
	}
	if err := result.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode data from db error:%v", err)
	}

	return cfg.List, nil
}
