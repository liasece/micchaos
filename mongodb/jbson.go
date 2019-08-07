package mongodb

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBsonProxyJsonByObj(obj interface{}) (bson.M, error) {
	// 取 json
	jsonb, errJson := json.Marshal(obj)
	if errJson != nil {
		return nil, errJson
	}
	// 解析中转对象
	var proxyObj interface{}
	errJsonU := json.Unmarshal(jsonb, &proxyObj)
	if errJsonU != nil {
		return nil, errJsonU
	}
	return proxyObj.(bson.M), nil
}

func GetObjProxyJsonByBson(bdata bson.M, obj *interface{}) error {
	// 取 json
	jsonb, errJson := json.Marshal(bdata)
	if errJson != nil {
		return errJson
	}
	// 解析中转对象
	errJsonU := json.Unmarshal(jsonb, obj)
	if errJsonU != nil {
		return errJsonU
	}
	return nil
}
