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
	err = req.ParseEncodeUrl(false, "Nature", &data.Nature, "Street", &data.Street, "Street_No", &data.Street_No, "Address", &data.Address, "Number", &data.Number, "Comment", &data.Comment, "Creator_Card_No", &data.Creator_Card_No)
	if err != nil {
		return
	}

	if data.Nature == "" || data.Street == "" || data.Address == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := pest.CreateHouse(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_CreateUnit(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.UnitData
	err = req.ParseEncodeUrl(false, "Name", &data.Name, "House_ID", &data.House_ID, "License_Number", &data.License_Number, "Identification_Number", &data.Identification_Number, "Picture", &data.Picture, "Kind", &data.Kind, "Scale", &data.Scale, "Tel", &data.Tel, "Bank_Name", &data.Bank_Name, "Bank_Account", &data.Bank_Account, "Comment", &data.Comment, "Creator_Card_No", &data.Creator_Card_No)
	if err != nil {
		return
	}

	if data.Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := pest.CreateUnit(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_CreatePersonnel(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.PersonnelData
	err = req.ParseEncodeUrl(false, "Name", &data.Name, "Occupation", &data.Occupation, "Card_No", &data.Card_No, "Card_Picture_Front", &data.Card_Picture_Front, "Card_Picture_Back", &data.Card_Picture_Back, "Face_Picture", &data.Face_Picture, "Sex", &data.Sex, "Nation", &data.Nation, "Birthday", &data.Birthday, "Address", &data.Address, "Sign_Organization", &data.Sign_Organization, "Limited_Date", &data.Limited_Date, "History", &data.History, "Phone", &data.Phone, "Remark", &data.Remark)
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

	res, err := pest.CreatePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_UpdateHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.HouseData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Nature", &data.Nature, "Street", &data.Street, "Street_No", &data.Street_No, "Address", &data.Address, "Number", &data.Number, "Comment", &data.Comment)
	if err != nil {
		return
	}

	if data.Nature == "" || data.Street == "" || data.Address == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := pest.UpdateHouse(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_UpdateUnit(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.UnitData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Name", &data.Name, "License_Number", &data.License_Number, "Identification_Number", &data.Identification_Number, "Picture", &data.Picture, "Kind", &data.Kind, "Scale", &data.Scale, "Tel", &data.Tel, "Bank_Name", &data.Bank_Name, "Bank_Account", &data.Bank_Account, "Comment", &data.Comment)
	if err != nil {
		return
	}

	if data.Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := pest.UpdateUnit(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_UpdatePersonnel(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.PersonnelData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Name", &data.Name, "Occupation", &data.Occupation, "Card_No", &data.Card_No, "Card_Picture_Front", &data.Card_Picture_Front, "Card_Picture_Back", &data.Card_Picture_Back, "Face_Picture", &data.Face_Picture, "Sex", &data.Sex, "Nation", &data.Nation, "Birthday", &data.Birthday, "Address", &data.Address, "Sign_Organization", &data.Sign_Organization, "Limited_Date", &data.Limited_Date, "History", &data.History, "Phone", &data.Phone, "Remark", &data.Remark)
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
	res, err := pest.UpdatePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var uid int
	err = req.ParseEncodeUrl(false, "Uid", &uid)
	if err != nil {
		return
	}

	res, err := pest.DeleteHouse(uid)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteUnit(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var uid int
	err = req.ParseEncodeUrl(false, "Uid", &uid)
	if err != nil {
		return
	}
	res, err := pest.DeleteUnit(uid)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeletePersonnel(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var uid int
	err = req.ParseEncodeUrl(false, "Uid", &uid)
	if err != nil {
		return
	}
	res, err := pest.DeletePersonnel(uid)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_AddRelationPersonnelHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.HousePersonnelData
	err = req.ParseEncodeUrl(false, "House_ID", &data.House_ID, "Personnel_ID", &data.Personnel_ID, "Role", &data.Role, "Relation_Holder", &data.Relation_Holder, "Relation_Together", &data.Relation_Together)
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
func (module *PestModule) Base_UpdateRelationPersonnelHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.HousePersonnelData
	err = req.ParseEncodeUrl(false, "House_ID", &data.House_ID, "Personnel_ID", &data.Personnel_ID, "Role", &data.Role, "Relation_Holder", &data.Relation_Holder, "Relation_Together", &data.Relation_Together)
	if err != nil {
		return
	}
	res, err := pest.UpdateRelationHousePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_UpdateRelationPersonnelUnit(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.UnitPersonnelData
	err = req.ParseEncodeUrl(false, "Unit_ID", &data.Unit_ID, "Personnel_ID", &data.Personnel_ID, "Position", &data.Position)
	if err != nil {
		return
	}
	res, err := pest.UpdateRelationUnitPersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteRelationPersonnelHouse(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.HousePersonnelData
	err = req.ParseEncodeUrl(false, "House_ID", &data.House_ID, "Personnel_ID", &data.Personnel_ID)
	if err != nil {
		return
	}
	res, err := pest.DeleteRelationHousePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteRelationPersonnelUnit(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.UnitPersonnelData
	err = req.ParseEncodeUrl(false, "Unit_ID", &data.Unit_ID, "Personnel_ID", &data.Personnel_ID)
	if err != nil {
		return
	}
	res, err := pest.DeleteRelationUnitPersonnel(&data)
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
	err = req.ParseEncodeUrl(false, "Personnel_ID", &data.Personnel_ID, "Time", &data.Time, "Symptom", &data.Symptom, "Hospitalized_Flag", &data.Hospitalized_Flag, "Temperature", &data.Temperature, "Touch_People", &data.Touch_People)
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
func (module *PestModule) Base_UpdateTouch(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.TouchData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Way", &data.Way, "Time", &data.Time, "Place", &data.Place, "Touch_Number", &data.Touch_Number, "Touch_People", &data.Touch_People)
	if err != nil {
		return
	}

	res, err := pest.UpdateTouch(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_UpdateDailyReport(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.DailyReportData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Time", &data.Time, "Symptom", &data.Symptom, "Hospitalized_Flag", &data.Hospitalized_Flag, "Temperature", &data.Temperature, "Touch_People", &data.Touch_People)
	if err != nil {
		return
	}

	res, err := pest.UpdateDailyReport(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteTouch(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var uid int
	err = req.ParseEncodeUrl(false, "Uid", &uid)
	if err != nil {
		return
	}

	res, err := pest.DeleteTouch(uid)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteDailyReport(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var uid int
	err = req.ParseEncodeUrl(false, "Uid", &uid)
	if err != nil {
		return
	}

	res, err := pest.DeleteDailyReport(uid)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_CreateStructure(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.StructureData
	err = req.ParseEncodeUrl(false, "Parent_ID", &data.Parent_ID, "Unit_ID", &data.Unit_ID, "Structure_Name", &data.Structure_Name, "Structure_Comment", &data.Structure_Comment)
	if err != nil {
		return
	}

	if data.Structure_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := pest.CreateStructure(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_UpdateStructure(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.StructureData
	err = req.ParseEncodeUrl(false, "Uid", &data.Uid, "Structure_Name", &data.Structure_Name, "Structure_Comment", &data.Structure_Comment)
	if err != nil {
		return
	}

	if data.Structure_Name == "" {
		err = service.NewError(service.ERR_MISSING_VALUE, "参数不能为空哦！")
		return
	}

	res, err := pest.UpdateStructure(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteStructure(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var uid int
	err = req.ParseEncodeUrl(false, "Uid", &uid)
	if err != nil {
		return
	}
	res, err := pest.DeleteStructure(uid)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_AddRelationPersonnelStructure(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.StructurePersonnelData
	err = req.ParseEncodeUrl(false, "Structure_ID", &data.Structure_ID, "Personnel_ID", &data.Personnel_ID, "Structure_Position", &data.Structure_Position)
	if err != nil {
		return
	}
	res, err := pest.AddRelationStructurePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_UpdateRelationPersonnelStructure(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.StructurePersonnelData
	err = req.ParseEncodeUrl(false, "Structure_ID", &data.Structure_ID, "Personnel_ID", &data.Personnel_ID, "Structure_Position", &data.Structure_Position)
	if err != nil {
		return
	}
	res, err := pest.UpdateRelationStructurePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
func (module *PestModule) Base_DeleteRelationPersonnelStructure(req *service.HttpRequest, result map[string]interface{}) (err error) {
	var data model.StructurePersonnelData
	err = req.ParseEncodeUrl(false, "Structure_ID", &data.Structure_ID, "Personnel_ID", &data.Personnel_ID)
	if err != nil {
		return
	}
	res, err := pest.DeleteRelationStructurePersonnel(&data)
	if err != nil {
		return
	}
	result["res"] = res
	return
}
