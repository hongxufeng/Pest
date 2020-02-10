package area

import (
	"datahelper/area"
	"model"
	"utils/config"
	"utils/service"

	"github.com/hongxufeng/fileLogger"
)

var Info *fileLogger.FileLogger
var Error *fileLogger.FileLogger

type AreaModule struct {
	level service.LEVEL
}

func (module *AreaModule) Init(conf *config.Config) error {
	module.level = service.SetEnvironment(conf.Environment)
	Info = fileLogger.NewDefaultLogger(conf.LogDir, "Area_Info.log")
	Error = fileLogger.NewDefaultLogger(conf.LogDir, "Area_Error.log")
	Info.SetPrefix("[area] ")
	Error.SetPrefix("[area] ")
	return nil
}

func (module *AreaModule) User_AddStation(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.StationData
	err = req.ParseEncodeUrl(false, "Station_Name", &data.Station_Name, "Province_Name", &data.Province_Name, "Province_No", &data.Province_No, "City_Name", &data.City_Name, "City_No", &data.City_No, "District_Name", &data.District_Name, "District_No", &data.District_No)
	if err != nil {
		return
	}

	if data.Station_Name == "" || data.Province_Name == "" || data.City_Name == "" || data.District_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.AddStation(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_AddCommunity(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.CommunityData
	err = req.ParseEncodeUrl(false, "Community_Name", &data.Community_Name, "Station_No", &data.Station_No)
	if err != nil {
		return
	}

	if data.Community_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.AddCommunity(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_AddStreet(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.StreetData
	err = req.ParseEncodeUrl(false, "Street_Name", &data.Street_Name, "Community_No", &data.Community_No)
	if err != nil {
		return
	}

	if data.Street_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.AddStreet(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
