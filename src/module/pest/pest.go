package pest

import (
	"model"
	"utils/config"
	"utils/service"

	"github.com/hongxufeng/fileLogger"
)

var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger

type PestModule struct {
	level service.LEVEL
}

func (module *PestModule) Init(conf *config.Config) error {
	module.level = service.SetEnvironment(conf.Environment)
	Info = fileLogger.NewDefaultLogger(conf.LogDir, "Pest_Info.log")
	Error = fileLogger.NewDefaultLogger(conf.LogDir, "Pest_Error.log")
	Info.SetPrefix("[pest] ")
	Error.SetPrefix("[pest] ")
	return nil
}

func (module *PestModule) Base_CreateHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.HouseData
	err = req.ParseEncodeUrl(false, "nature", &data.nature, "street", &data.street, "street_no", &data.street_no, "address", &data.address, "number", &data.number)
	if err != nil {
		return
	}

	if data.nature == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "任务内容不能为空哦！")
		return
	}

	pestCreate, err := pest.Createpest(&data)
	if err != nil {
		return
	}
	result["res"] = pestCreate
	return
}
