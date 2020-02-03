package pest

import (
	"datahelper/pest"
	"model"
	"utils/config"
	"utils/function"
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
	err = req.ParseEncodeUrl(false, "Nature", &data.Nature, "Street", &data.Street, "Street_No", &data.Street_No, "Address", &data.Address, "Number", &data.Number)
	if err != nil {
		return
	}

	if data.Nature == "" || data.Street == "" || data.Street_No == "" || data.Address == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	houseCreate, err := pest.CreateHouse(&data)
	if err != nil {
		return
	}
	result["res"] = houseCreate
	return
}
func (module *PestModule) Base_CreateUnit(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.UnitData
	err = req.ParseEncodeUrl(false, "Name", &data.Name, "House_ID", &data.House_ID, "License_Number", &data.License_Number, "Identification_Number", &data.Identification_Number, "Picture", &data.Picture, "Kind", &data.Kind, "Scale", &data.Scale, "Tel", &data.Tel, "Bank_Name", &data.Bank_Name, "Bank_Account", &data.Bank_Account, "Comment", &data.Comment)
	if err != nil {
		return
	}

	if data.Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	unitCreate, err := pest.CreateUnit(&data)
	if err != nil {
		return
	}
	result["res"] = unitCreate
	return
}
func (module *PestModule) Base_CreatePersonnel(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.PersonnelData
	err = req.ParseEncodeUrl(false, "Name", &data.Name, "Occupation", &data.Occupation, "Card_No", &data.Card_No, "Picture", &data.Picture, "Sex", &data.Sex, "Nation", &data.Nation, "Birthday", &data.Birthday, "Address", &data.Address, "Sign_Organization", &data.Sign_Organization, "Limited_Date", &data.Limited_Date, "History", &data.History)
	if err != nil {
		return
	}
	if !function.ValidateCardNo(data.Card_No) {
		err = service.NewError(service.ERR_MISSING_VALUE, "身份证格式哦！")
		return
	}
	if data.Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	personnelCreate, err := pest.CreatePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = personnelCreate
	return
}
func (module *PestModule) Base_AddRelationPersonnelHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.HousePersonnelData
	err = req.ParseEncodeUrl(false, "House_ID", &data.House_ID, "Personnel_ID", &data.Personnel_ID, "Holder_Flag", &data.Holder_Flag, "Relation_Holder", &data.Relation_Holder)
	if err != nil {
		return
	}
	res, err := pest.AddRelationHousePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_AddRelationPersonnelUnit(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.UnitPersonnelData
	err = req.ParseEncodeUrl(false, "Unit_ID", &data.Unit_ID, "Personnel_ID", &data.Personnel_ID, "Position", &data.Position)
	if err != nil {
		return
	}
	res, err := pest.AddRelationUnitPersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_AddTouch(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.TouchData
	err = req.ParseEncodeUrl(false, "Personnel_ID", &data.Personnel_ID, "Way", &data.Way, "Time", &data.Time, "Place", &data.Place, "Touch_Number", &data.Touch_Number, "Touch_People", &data.Touch_People)
	if err != nil {
		return
	}

	res, err := pest.AddTouch(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_AddDailyReport(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.DailyReportData
	err = req.ParseEncodeUrl(false, "Personnel_ID", &data.Personnel_ID, "Symptom", &data.Symptom, "Hospitalized_Flag", &data.Hospitalized_Flag, "Temperature", &data.Temperature, "Touch_People", &data.Touch_People)
	if err != nil {
		return
	}

	res, err := pest.AddDailyReport(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
