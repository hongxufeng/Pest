package db

import (
	"model"
	"time"
)

func GetStreetbyID(uid int64) (name string, province_name string, city_name string, district_name string, station_name string, office_name string, community_name string, street_name string, street_qrcode string, err error) {
	queryStr := "SELECT name,province_name,city_name,district_name,station_name,office_name,community_name,street_name,COALESCE (street_qrcode,'') FROM area_list WHERE uid=? limit 0,1"
	result, err := MysqlMain.Query(queryStr, uid)
	if err != nil {
		return
	}
	defer result.Close()
	if result.Next() {
		err = result.Scan(&name, &province_name, &city_name, &district_name, &station_name, &office_name, &community_name, &street_name, &street_qrcode)
	}
	return
}
func ExecAddStation(data *model.StationData) (uid int64, err error) {
	query := "INSERT INTO station_list (station_name,province_no,province_name,city_no,city_name,district_no,district_name,station_head,station_phone,create_time) VALUES (?,?,?,?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Station_Name, data.Province_No, data.Province_Name, data.City_No, data.City_Name, data.District_No, data.District_Name, data.Station_Head, data.Station_Phone, time.Now().UnixNano())
	if err != nil {
		return 0, err
	}
	err = DelTableCache(model.XML_Table_Area)
	if err != nil {
		return 0, err
	}
	return
}
func ExecAddOffice(data *model.OfficeData) (uid int64, err error) {
	query := "INSERT INTO office_list (office_name,station_no,office_head,office_phone,create_time) VALUES (?,?,?,?,?)"
	uid, err = Insert(query, data.Office_Name, data.Station_No, data.Office_Head, data.Office_Phone, time.Now().UnixNano())
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
	query := "INSERT INTO community_list (community_name,office_no,community_head,community_phone,create_time) VALUES (?,?,?,?,?)"
	uid, err = Insert(query, data.Community_Name, data.Office_No, data.Community_Head, data.Community_Phone, time.Now().UnixNano())
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
	query := "INSERT INTO street_list (street_name,community_no,street_head,street_phone,street_property_name,street_property_phone,create_time) VALUES (?,?,?,?,?,?,?)"
	uid, err = Insert(query, data.Street_Name, data.Community_No, data.Street_Head, data.Street_Phone, data.Street_Property_Name, data.Street_Property_Phone, time.Now().UnixNano())
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
	query := "delete station_list,office_list,community_list,street_list from station_list left join office_list on office_list.station_no=station_list.uid LEFT JOIN community_list ON community_list.office_no=station_list.uid LEFT JOIN street_list ON street_list.community_no=community_list.uid where station_list.uid=?"
	err := Exec(query, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecDeleteOffice(uid int) error {
	query := "delete street_list,community_list,office_list from office_list left join community_list on community.office_no=office_list.uid left join street_list on street_list.community_no=community_list.uid where office_list.uid=?"
	err := Exec(query, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecDeleteCommunity(uid int) error {
	query := "delete street_list,community_list from community_list left join street_list on street_list.community_no=community_list.uid where community_list.uid=?"
	err := Exec(query, uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecUpdateStreet(data *model.StreetData) error {
	query := "update street_list set street_name=?,street_head=?,street_phone=?,street_property_name=?,street_property_phone=? where uid=?"
	err := Exec(query, data.Street_Name, data.Street_Head, data.Street_Phone, data.Street_Property_Name, data.Street_Property_Phone, data.Uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecUpdateOffice(data *model.OfficeData) error {
	query := "update office_list set office_name=?,office_head=?,office_phone=? where uid=?"
	err := Exec(query, data.Office_Name, data.Office_Head, data.Office_Phone, data.Uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecUpdateCommunity(data *model.CommunityData) error {
	query := "update community_list set community_name=?,community_head=?,community_phone=? where uid=?"
	err := Exec(query, data.Community_Name, data.Community_Head, data.Community_Phone, data.Uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
func ExecUpdateStation(data *model.StationData) error {
	query := "update station_list set station_name=?,station_head=?,station_phone=? where uid=?"
	err := Exec(query, data.Station_Name, data.Station_Head, data.Station_Phone, data.Uid)
	if err != nil {
		return err
	}
	return DelTableCache(model.XML_Table_Area)
}
