package ebarimt

import (
	"github.com/lambda-platform/ebarimt/posapi"
)

var PosAPI *posapi.PosAPI

func init() {
	api, err := posapi.NewPosAPI("./libPosAPI.so")
	if err != nil {
		panic(err)
	}
	PosAPI = api
	defer PosAPI.Close()
}
