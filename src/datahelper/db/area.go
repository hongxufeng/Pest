package db

import (
	"model"
	"time"
)

func GetStreetbyID(uid int64) (name string, province_name string, city_name string, district_name string, station_name string, community_name string, street_name string, street_qrcode string, err error) {
	queryStr := "SELECT name,province_name,city_name,district_name,station_name,community_name,street_name,COALESCE (street_qrcode,'') FROM area_list WHERE uid=? limit 0,1"
	result, err := MysqlMain.Query(queryStr, uid)
	if err != nil {
		return
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&name, &province_name, &city_name, &district_name, &station_name, &community_name, &street_name, &street_qrcode)
	}
	return
}
func ExecAddStation(data *model.StationData) (uid int64, err error) {
	query := "INSERT INTO station_list (station_name,province_no,province_name,city_no,city_name,district_no,district_name,create_time) VALUES (?,?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Station_Name, data.Province_No, data.Province_Name, data.City_No, data.City_Name, data.District_No, data.District_Name, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Area)
	if err != nil {
		return 0, err
	}
	return
}
func ExecAddCommunity(data *model.CommunityData) (uid int64, err error) {
	query := "INSERT INTO community_list (community_name,station_no,create_time) VALUES (?,?,?)"
	uid, err = Insert(query, data.Community_Name, data.Station_No, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Area)
	if err != nil {
		return 0, err
	}
	return
}
func ExecAddStreet(data *model.StreetData) (uid int64, err error) {
	query := "INSERT INTO street_list (street_name,community_no,create_time) VALUES (?,?,?)"
	uid, err = Insert(query, data.Street_Name, data.Community_No, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Area)
	if err != nil {
		return 0, err
	}
	return
}
func ExecUpdateStreetQr(uid int64, qrfile string) error {
	query := "update street_list set street_qrcode=? where uid=?"
	err := Exec(query, qrfile, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecDeleteStreet(uid int) error {
	query := "delete from street_list where uid=?"
	err := Exec(query, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecDeleteStation(uid int) error {
	query := "delete station_list,community_list,street_list from station_list,community_list,street_list where station_list.uid=? and community_list.station_no=station_list.uid and street_list.community_no=community_list.uid"
	err := Exec(query, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecDeleteCommunity(uid int) error {
	query := "delete street_list,community_list from street_list,community_list where community_list.uid=? and community_list.uid=street_list.community_no"
	err := Exec(query, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecUpdateStreet(data *model.UpdateData) error {
	query := "update street_list set street_name=? where uid=?"
	err := Exec(query, data.Update_Name, data.Uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecUpdateCommunity(data *model.UpdateData) error {
	query := "update community_list set community_name=? where uid=?"
	err := Exec(query, data.Update_Name, data.Uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecUpdateStation(data *model.UpdateData) error {
	query := "update station_list set station_name=? where uid=?"
	err := Exec(query, data.Update_Name, data.Uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
