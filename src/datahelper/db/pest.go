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
func ExecCreateHouse(data *model.HouseData) (uid int64, err error) {
	query := "INSERT INTO house_list (nature,street,street_no,address,number,create_time) VALUES (?,?,?,?,?,?)"
	//fmt.Println(time.Now())
	uid, err = Insert(query, data.Nature, data.Street, data.Street_No, data.Address, data.Number, time.Now().UnixNano())
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
	query := "INSERT INTO unit_list (name,house_id,license_number,identification_number,picture,kind,scale,tel,bank_name,bank_account,comment,create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	//fmt.Println(time.Now())
	uid, err = Insert(query, data.Name, data.House_ID, data.License_Number, data.Identification_Number, data.Picture, data.Kind, data.Scale, data.Tel, data.Bank_Name, data.Bank_Account, data.Comment, time.Now().UnixNano())
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
	query := "INSERT INTO personnel_list (name,occupation,card_no,picture,sex,nation,birthday,address,sign_organization,limited_date,history,create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	//fmt.Println(time.Now())
	uid, err = Insert(query, data.Name, data.Occupation, data.Card_No, data.Picture, data.Sex, data.Nation, data.Birthday, data.Address, data.Sign_Organization, data.Limited_Date, data.History, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Pest)
	if err != nil {
		return 0, err
	}
	return
}
func ExecAddRelationHousePersonnel(data *model.HousePersonnelData) (uid int64, err error) {
	query := "INSERT INTO relation_house_personnel (house_id,personnel_id,holder_flag,relation_holder,create_time) VALUES (?,?,?,?,?)"
	//fmt.Println(time.Now())
	uid, err = Insert(query, data.House_ID, data.Personnel_ID, data.Holder_Flag, data.Relation_Holder, time.Now().UnixNano())
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
	//fmt.Println(time.Now())
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
func ExecAddTouch(data *model.TouchData) (uid int64, err error) {
	query := "INSERT INTO touch_list (personnel_id,way,time,place,touch_number,touch_people,create_time) VALUES (?,?,?,?,?,?,?)"
	//fmt.Println(time.Now())
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
	query := "INSERT INTO daily_report_list (personnel_id,symptom,hospitalized_flag,temperature,touch_people,create_time) VALUES (?,?,?,?,?,?)"
	//fmt.Println(time.Now())
	uid, err = Insert(query, data.Personnel_ID, data.Symptom, data.Hospitalized_Flag, data.Temperature, data.Touch_People, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Daily_Report)
	if err != nil {
		return 0, err
	}
	return
}
