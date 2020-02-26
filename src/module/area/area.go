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
	qrurl string
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
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.StationData
	err = req.ParseEncodeUrl(false, "Station_Name", &data.Station_Name, "Province_Name", &data.Province_Name, "Province_No", &data.Province_No, "City_Name", &data.City_Name, "City_No", &data.City_No, "District_Name", &data.District_Name, "District_No", &data.District_No, "Station_Head", &data.Station_Head, "Station_Phone", &data.Station_Phone)
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
func (module *AreaModule) User_AddOffice(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.OfficeData
	err = req.ParseEncodeUrl(false, "Office_Name", &data.Office_Name, "Station_No", &data.Station_No, "Office_Head", &data.Office_Head, "Office_Phone", &data.Office_Phone)
	if err != nil {
		return
	}

	if data.Office_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.AddOffice(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_AddCommunity(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.CommunityData
	err = req.ParseEncodeUrl(false, "Community_Name", &data.Community_Name, "Office_No", &data.Office_No, "Community_Head", &data.Community_Head, "Community_Phone", &data.Community_Phone)
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
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.StreetData
	err = req.ParseEncodeUrl(false, "Street_Name", &data.Street_Name, "Community_No", &data.Community_No, "Street_Head", &data.Street_Head, "Street_Phone", &data.Street_Phone, "Street_Property_Name", &data.Street_Property_Name, "Street_Property_Phone", &data.Street_Property_Phone)
	if err != nil {
		return
	}

	if data.Street_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.AddStreet(module.qrurl, &data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_DeleteArea(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.DeleteData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Cmd_Delete", &data.Cmd_Delete)
	if err != nil {
		return
	}
	res, err := area.DeleteArea(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_UpdateStation(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.StationData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Station_Name", &data.Station_Name, "Station_Head", &data.Station_Head, "Station_Phone", &data.Station_Phone)
	if err != nil {
		return
	}

	if data.Station_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.UpdateStation(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_UpdateOffice(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.OfficeData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Office_Name", &data.Office_Name, "Office_Head", &data.Office_Head, "Office_Phone", &data.Office_Phone)
	if err != nil {
		return
	}

	if data.Office_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.UpdateOffice(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_UpdateCommunity(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.CommunityData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Community_Name", &data.Community_Name, "Community_Head", &data.Community_Head, "Community_Phone", &data.Community_Phone)
	if err != nil {
		return
	}

	if data.Community_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.UpdateCommunity(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *AreaModule) User_UpdateStreet(req *service.HttpRequest, result map[string]interface{}) (err error) {
	if req.Power > 0 {
		err = service.NewError(service.ERR_POWER_DENIED, "所处用户权限不足！")
		return
	}
	var data model.StreetData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Street_Name", &data.Street_Name, "Street_Head", &data.Street_Head, "Street_Phone", &data.Street_Phone, "Street_Property_Name", &data.Street_Property_Name, "Street_Property_Phone", &data.Street_Property_Phone)
	if err != nil {
		return
	}

	if data.Street_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := area.UpdateStreet(module.qrurl, &data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
