package ccmd

import (
	"github.com/liasece/micserver/conf"
	"github.com/liasece/micserver/roc"
)

var (
	ConfMongoDB     conf.ConfigKey = "mongodb"
	ConfDataBase    conf.ConfigKey = "database"
	ConfDBUserInfos conf.ConfigKey = "userinfos_collection"
)

var (
	ROCTypePlayer roc.ObjType = "Player"
)
