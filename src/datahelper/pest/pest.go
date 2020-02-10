package pest

import (
	"datahelper/db"
	"model"
	"utils/service"
)

func CreateHouse(data *model.HouseData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecCreateHouse(data)
	if err == nil {
		res["house_id"] = uid
		res["msg"] = "房屋生成成功！"
	}
	return
}
func CreateUnit(data *model.UnitData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecCreateUnit(data)
	if err == nil {
		res["unit_id"] = uid
		res["msg"] = "单位生成成功！"
	}
	return
}
func CreatePersonnel(data *model.PersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	isexist, _ := db.GetCheckPersonnel("card_no", data.Card_No)
	if isexist {
		err = service.NewError(service.ERR_INVALID_PARAM, "身份证信息已存在！")
		return
	}
	uid, err := db.ExecCreatePersonnel(data)
	if err == nil {
		res["personnel_id"] = uid
		res["msg"] = "人员生成成功！"
	}
	return
}
func AddRelationHousePersonnel(data *model.HousePersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	isexist, _ := db.CheckPersonnelHouse(data.House_ID, data.Personnel_ID)
	if isexist {
		err = service.NewError(service.ERR_INVALID_PARAM, "关系已经存在！")
		return
	}
	uid, err := db.ExecAddRelationHousePersonnel(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "人员房屋关系添加成功！"
	}
	return
}
func AddRelationUnitPersonnel(data *model.UnitPersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	isexist, _ := db.CheckPersonnelUnit(data.Unit_ID, data.Personnel_ID)
	if isexist {
		err = service.NewError(service.ERR_INVALID_PARAM, "关系已经存在！")
		return
	}
	uid, err := db.ExecAddRelationUnitPersonnel(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "人员单位关系添加成功！"
	}
	return
}
func AddTouch(data *model.TouchData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddTouch(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "其他接触史添加成功！"
	}
	return
}
func AddDailyReport(data *model.DailyReportData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddDailyReport(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "每日上报添加成功！"
	}
	return
}
