package ebarimt

import "github.com/lambda-platform/ebarimt/posapi"

var PosAPI *posapi.PosAPI

func init() {
	api, err := posapi.NewPosAPI("/home/khankhulgen/web/ebarimt/rest/6787556.so")
	if err != nil {
		panic(err)
	}
	PosAPI = api
	defer PosAPI.Close()
}
