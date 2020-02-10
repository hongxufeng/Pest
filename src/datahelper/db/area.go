package db

import (
	"model"
	"time"
)

func ExecAddStation(data *model.StationData) (uid int64, err error) {
	query := "INSERT INTO station_list (station_name,province_no,province_name,city_no,city_name,district_no,district_name,create_time) VALUES (?,?,?,?,?,?,?,?)"
	//fmt.Println(time.Now())
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
	//fmt.Println(time.Now())
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
	//fmt.Println(time.Now())
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
