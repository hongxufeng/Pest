package db

import (
	"model"
	"time"
)

func GetUidbyNo(name string) (uid uint32, err error) {
	uid = 0
	queryStr := "SELECT uid FROM personnel_list WHERE card_no=? limit 0,1"
	result, err := MysqlMain.Query(queryStr, name)
	if err != nil {
		return
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&uid)
	}
	return
}
func GetCheckPersonnel(param string, name string) (bool, error) {
	var count int
	query := "SELECT Count(*) FROM personnel_list WHERE " + param + "=?"
	result, e := MysqlMain.Query(query, name)
	if e != nil {
		return true, e
	}
	defer result.Close()
	if result.Next() {
		e = result.Scan(&count)
	} else {
		return true, nil
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func CheckPersonnelHouse(param1 int, param2 int) (bool, error) {
	var count int
	query := "SELECT Count(*) FROM relation_house_personnel WHERE house_id=? and personnel_id=?"
	result, e := MysqlMain.Query(query, param1, param2)
	if e != nil {
		return true, e
	}
	defer result.Close()
	if result.Next() {
		e = result.Scan(&count)
	} else {
		return true, nil
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func CheckPersonnelUnit(param1 int, param2 int) (bool, error) {
	var count int
	query := "SELECT Count(*) FROM relation_unit_personnel WHERE unit_id=? and personnel_id=?"
	result, e := MysqlMain.Query(query, param1, param2)
	if e != nil {
		return true, e
	}
	defer result.Close()
	if result.Next() {
		e = result.Scan(&count)
	} else {
		return true, nil
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func CheckPersonnelStructure(param1 int, param2 int) (bool, error) {
	var count int
	query := "SELECT Count(*) FROM relation_structure_personnel WHERE structure_id=? and personnel_id=?"
	result, e := MysqlMain.Query(query, param1, param2)
	if e != nil {
		return true, e
	}
	defer result.Close()
	if result.Next() {
		e = result.Scan(&count)
	} else {
		return true, nil
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func ExecCreateHouse(data *model.HouseData) (uid int64, err error) {
	query := "INSERT INTO house_list (nature,street,street_no,address,number,comment,creator_card_no,create_time) VALUES (?,?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Nature, data.Street, data.Street_No, data.Address, data.Number, data.Comment, data.Creator_Card_No, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecCreateUnit(data *model.UnitData) (uid int64, err error) {
	query := "INSERT INTO unit_list (name,house_id,license_number,identification_number,picture,kind,scale,tel,bank_name,bank_account,comment,creator_card_no,create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Name, data.House_ID, data.License_Number, data.Identification_Number, data.Picture, data.Kind, data.Scale, data.Tel, data.Bank_Name, data.Bank_Account, data.Comment, data.Creator_Card_No, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecCreatePersonnel(data *model.PersonnelData) (uid int64, err error) {
	query := "INSERT INTO personnel_list (name,occupation,card_no,phone,home,card_picture_front,card_picture_back,face_picture,sex,nation,birthday,address,sign_organization,limited_date,history,remark,create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Name, data.Occupation, data.Card_No, data.Phone, data.Home, data.Card_Picture_Front, data.Card_Picture_Back, data.Face_Picture, data.Sex, data.Nation, data.Birthday, data.Address, data.Sign_Organization, data.Limited_Date, data.History, data.Remark, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecUpdateHouse(data *model.HouseData) (err error) {
	query := "update house_list set nature=?,street=?,street_no=?,address=?,number=?,comment=? where uid=?"
	err = Exec(query, data.Nature, data.Street, data.Street_No, data.Address, data.Number, data.Comment, data.Uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecUpdateUnit(data *model.UnitData) (err error) {
	query := "update unit_list set name=?,license_number=?,identification_number=?,picture=?,kind=?,scale=?,tel=?,bank_name=?,bank_account=?,comment=? where uid=?"
	err = Exec(query, data.Name, data.License_Number, data.Identification_Number, data.Picture, data.Kind, data.Scale, data.Tel, data.Bank_Name, data.Bank_Account, data.Comment, data.Uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecUpdatePersonnel(data *model.PersonnelData) (err error) {
	query := "update personnel_list set name=?,occupation=?,card_no=?,phone=?,home=?,card_picture_front=?,card_picture_back=?,face_picture=?,sex=?,nation=?,birthday=?,address=?,sign_organization=?,limited_date=?,history=?,remark=? where uid=?"
	err = Exec(query, data.Name, data.Occupation, data.Card_No, data.Phone, data.Home, data.Card_Picture_Front, data.Card_Picture_Back, data.Face_Picture, data.Sex, data.Nation, data.Birthday, data.Address, data.Sign_Organization, data.Limited_Date, data.History, data.Remark, data.Uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecAddRelationHousePersonnel(data *model.HousePersonnelData) (uid int64, err error) {
	query := "INSERT INTO relation_house_personnel (house_id,personnel_id,role,relation_holder,relation_together,create_time) VALUES (?,?,?,?,?,?)"
	uid, err = Insert(query, data.House_ID, data.Personnel_ID, data.Role, data.Relation_Holder, data.Relation_Together, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecAddRelationUnitPersonnel(data *model.UnitPersonnelData) (uid int64, err error) {
	query := "INSERT INTO relation_unit_personnel (unit_id,personnel_id,position,create_time) VALUES (?,?,?,?)"
	uid, err = Insert(query, data.Unit_ID, data.Personnel_ID, data.Position, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecUpdateRelationHousePersonnel(data *model.HousePersonnelData) (err error) {
	query := "Update relation_house_personnel set role=?,relation_holder=?,relation_together=? where house_id=? and personnel_id=?"
	err = Exec(query, data.Role, data.Relation_Holder, data.Relation_Together, data.House_ID, data.Personnel_ID)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecUpdateRelationUnitPersonnel(data *model.UnitPersonnelData) (err error) {
	query := "Update relation_unit_personnel set position=? where unit_id=? and personnel_id=?"
	err = Exec(query, data.Position, data.Unit_ID, data.Personnel_ID)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecDeleteRelationHousePersonnel(data *model.HousePersonnelData) (err error) {
	query := "delete from relation_house_personnel where house_id=? and personnel_id=?"
	err = Exec(query, data.House_ID, data.Personnel_ID)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecDeleteRelationUnitPersonnel(data *model.UnitPersonnelData) (err error) {
	query := "delete from relation_unit_personnel where unit_id=? and personnel_id=?"
	err = Exec(query, data.Unit_ID, data.Personnel_ID)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecAddTouch(data *model.TouchData) (uid int64, err error) {
	query := "INSERT INTO touch_list (personnel_id,way,time,place,touch_number,touch_people,create_time) VALUES (?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Personnel_ID, data.Way, data.Time, data.Place, data.Touch_Number, data.Touch_People, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Touch_History)
	if err != nil {
		return 0, err
	}
	return
}
func ExecAddDailyReport(data *model.DailyReportData) (uid int64, err error) {
	query := "INSERT INTO daily_report_list (personnel_id,symptom,hospitalized_flag,temperature,time,touch_people,create_time) VALUES (?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Personnel_ID, data.Symptom, data.Hospitalized_Flag, data.Temperature, data.Time, data.Touch_People, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Daily_Report)
	if err != nil {
		return 0, err
	}
	return
}
func ExecUpdateTouch(data *model.TouchData) (err error) {
	query := "update touch_list set way=?,time=?,place=?,touch_number=?,touch_people=? where uid=?"
	err = Exec(query, data.Way, data.Time, data.Place, data.Touch_Number, data.Touch_People, data.Uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Touch_History)
	return
}
func ExecUpdateDailyReport(data *model.DailyReportData) (err error) {
	query := "update daily_report_list set symptom=?,hospitalized_flag=?,temperature=?,time=?,touch_people=? where uid=?"
	err = Exec(query, data.Symptom, data.Hospitalized_Flag, data.Temperature, data.Time, data.Touch_People, data.Uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Daily_Report)
	return
}
func ExecDeleteHouse(uid int) (err error) {
	query := "delete house_list,relation_house_personnel,unit_list,relation_unit_personnel from house_list left join relation_house_personnel on relation_house_personnel.house_id=house_list.uid left join unit_list on unit_list.house_id=house_list.uid left join relation_unit_personnel on relation_unit_personnel.unit_id=unit_list.uid where house_list.uid=?"
	err = Exec(query, uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecDeleteUnit(uid int) (err error) {
	query := "delete unit_list,relation_unit_personnel,structure_list,relation_structure_personnel from unit_list left join relation_unit_personnel on relation_unit_personnel.unit_id=unit_list.uid left join structure_list on structure_list.unit_id=unit_list.uid left join relation_structure_personnel on relation_structure_personnel.structure_id=structure_list.uid where unit_list.uid=?"
	err = Exec(query, uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecDeletePersonnel(uid int) (err error) {
	query := "delete personnel_list,relation_house_personnel,relation_unit_personnel,relation_structure_personnel from personnel_list left join relation_house_personnel on relation_house_personnel.personnel_id=personnel_list.uid left join relation_unit_personnel on relation_unit_personnel.personnel_id=personnel_list.uid left join relation_structure_personnel on relation_structure_personnel.personnel_id=personnel_list.uid where personnel_list.uid=?"
	err = Exec(query, uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecDeleteTouch(uid int) (err error) {
	query := "delete from touch_list where uid=?"
	err = Exec(query, uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Touch_History)
	return
}
func ExecDeleteDailyReport(uid int) (err error) {
	query := "delete from daily_report_list where uid=?"
	err = Exec(query, uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Daily_Report)
	return
}
func ExecCreateStructure(data *model.StructureData) (uid int64, err error) {
	query := "INSERT INTO structure_list (parent_id,unit_id,structure_name,structure_comment,create_time) VALUES (?,?,?,?,?)"
	uid, err = Insert(query, data.Parent_ID, data.Unit_ID, data.Structure_Name, data.Structure_Comment, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecUpdateStructure(data *model.StructureData) (err error) {
	query := "update structure_list set structure_name=?,structure_comment=? where uid=?"
	err = Exec(query, data.Structure_Name, data.Structure_Comment, data.Uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecDeleteStructure(uid int) (err error) {
	query := "delete structure_list,relation_structure_personnel from structure_list left join relation_structure_personnel on relation_structure_personnel.structure_id=structure_list.uid where structure_list.uid=?"
	err = Exec(query, uid)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecAddRelationStructurePersonnel(data *model.StructurePersonnelData) (uid int64, err error) {
	query := "INSERT INTO relation_structure_personnel (structure_id,personnel_id,structure_position,create_time) VALUES (?,?,?,?)"
	uid, err = Insert(query, data.Structure_ID, data.Personnel_ID, data.Structure_Position, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecUpdateRelationStructurePersonnel(data *model.StructurePersonnelData) (err error) {
	query := "Update relation_structure_personnel set structure_position=? where structure_id=? and personnel_id=?"
	err = Exec(query, data.Structure_Position, data.Structure_ID, data.Personnel_ID)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
func ExecDeleteRelationStructurePersonnel(data *model.StructurePersonnelData) (err error) {
	query := "delete from relation_structure_personnel where structure_id=? and personnel_id=?"
	err = Exec(query, data.Structure_ID, data.Personnel_ID)
	if err != nil {
		return
	}
	err = DelTableCache(model.XML_Table_Pest)
	return
}
