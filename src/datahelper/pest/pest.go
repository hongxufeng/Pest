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
func UpdateHouse(data *model.HouseData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateHouse(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "房屋更新成功！"
	}
	return
}
func UpdateUnit(data *model.UnitData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateUnit(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "单位更新成功！"
	}
	return
}
func UpdatePersonnel(data *model.PersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdatePersonnel(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "人员更新成功！"
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
func UpdateRelationHousePersonnel(data *model.HousePersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateRelationHousePersonnel(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "人员房屋关系更新成功！"
	}
	return
}
func UpdateRelationUnitPersonnel(data *model.UnitPersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateRelationUnitPersonnel(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "人员单位关系更新成功！"
	}
	return
}
func DeleteRelationHousePersonnel(data *model.HousePersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteRelationHousePersonnel(data)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "人员房屋关系删除成功！"
	}
	return
}
func DeleteRelationUnitPersonnel(data *model.UnitPersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteRelationUnitPersonnel(data)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "人员单位关系删除成功！"
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
func UpdateTouch(data *model.TouchData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateTouch(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "其他接触史更新成功！"
	}
	return
}
func UpdateDailyReport(data *model.DailyReportData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateDailyReport(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "每日上报更新成功！"
	}
	return
}
func DeleteHouse(uid int) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteHouse(uid)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "房屋删除成功！"
	}
	return
}
func DeleteUnit(uid int) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteUnit(uid)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "单位删除成功！"
	}
	return
}
func DeletePersonnel(uid int) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeletePersonnel(uid)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "人员删除成功！"
	}
	return
}
func DeleteTouch(uid int) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteTouch(uid)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "其他接触史删除成功！"
	}
	return
}
func DeleteDailyReport(uid int) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteDailyReport(uid)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "每日上报删除成功！"
	}
	return
}
func CreateStructure(data *model.StructureData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecCreateStructure(data)
	if err == nil {
		res["structure_id"] = uid
		res["msg"] = "结构生成成功！"
	}
	return
}
func UpdateStructure(data *model.StructureData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateStructure(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "结构更新成功！"
	}
	return
}
func DeleteStructure(uid int) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteStructure(uid)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "结构删除成功！"
	}
	return
}
func AddRelationStructurePersonnel(data *model.StructurePersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	isexist, _ := db.CheckPersonnelStructure(data.Structure_ID, data.Personnel_ID)
	if isexist {
		err = service.NewError(service.ERR_INVALID_PARAM, "关系已经存在！")
		return
	}
	uid, err := db.ExecAddRelationStructurePersonnel(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "人员结构关系添加成功！"
	}
	return
}
func UpdateRelationStructurePersonnel(data *model.StructurePersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateRelationStructurePersonnel(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "人员结构关系更新成功！"
	}
	return
}
func DeleteRelationStructurePersonnel(data *model.StructurePersonnelData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecDeleteRelationStructurePersonnel(data)
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "人员结构关系删除成功！"
	}
	return
}
