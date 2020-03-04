package imports

import (
	"utils/config"
	"utils/service"

	"github.com/hongxufeng/fileLogger"
	"github.com/tealeg/xlsx"
)

var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger

type ImportsModule struct {
	level service.LEVEL
}

func (module *ImportsModule) Init(conf *config.Config) error {
	module.level = service.SetEnvironment(conf.Environment)
	Info = fileLogger.NewDefaultLogger(conf.LogDir, "Imports_Info.log")
	Error = fileLogger.NewDefaultLogger(conf.LogDir, "Imports_Error.log")
	Info.SetPrefix("[Imports]")
	Error.SetPrefix("[Imports]")
	return nil
}
func (module *ImportsModule) User_ImportHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	bytes, filename, err := req.PostFileInfo("file")
	//get extname from fname
	if err != nil {
		err = service.NewError(service.ERR_INVALID_PARAM, err.Error(), "上传错误")
		return
	}
	xlFile, err := xlsx.OpenBinary(bytes)
	if err != nil {
		err = service.NewError(service.ERR_INVALID_PARAM, err.Error(), "文件不能解析！")
		return
	}
	Info.Info(filename + "开始导入。。。。。。。。")
	res, err := AddExcelData(xlFile)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
