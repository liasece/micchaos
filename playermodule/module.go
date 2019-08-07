package playermodule

import (
	"github.com/liasece/micserver/module"
	"mongodb"
)

type PlayerModule struct {
	mongodb_micworld *mongodb.MongoClient
	module.BaseModule
}

func (this *PlayerModule) AfterInitModule() {
	this.BaseModule.AfterInitModule()
	this.Info("PlayerModule init finish %s", this.GetModuleID())

	var err error
	this.mongodb_micworld, err = mongodb.NewClient("mongodb://localhost:27017")
	if err != nil {
		this.Error("mongodb.NewClient err: %s", err.Error())
	} else {
		this.Info("mongodb.NewClient scesse")
	}
}
