package mongodb

import (
	"encoding/json"
	"fmt"
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
	return bson.M(proxyObj.(map[string]interface{})), nil
}

func GetObjProxyJsonByBson(bdata bson.M, obj interface{}) error {
	if bdata == nil || obj == nil {
		return fmt.Errorf("data or obj is nil")
	}
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
