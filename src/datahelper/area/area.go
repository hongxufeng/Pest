package area

import (
	"datahelper/db"
	"model"
	"os"
	"path"
	"utils/function"
	"utils/service"

	"github.com/skip2/go-qrcode"
)

func GenerateQrCode(url string, filename string) error {
	path := path.Dir(filename)
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return qrcode.WriteFile(url, qrcode.Medium, 256, filename)
}
func AddStation(data *model.StationData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddStation(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "派出所生成成功！"
	}
	return
}
func AddOffice(data *model.OfficeData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddOffice(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "办事处生成成功！"
	}
	return
}
func AddCommunity(data *model.CommunityData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddCommunity(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "社区生成成功！"
	}
	return
}
func AddStreet(qrurl string, data *model.StreetData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddStreet(data)
	if err != nil {
		return
	}
	_, province_name, city_name, district_name, station_name, office_name, community_name, street_name, _, err := db.GetStreetbyID(uid)
	if err != nil {
		return
	}
	url := qrurl + "?street_no=" + function.Int64ToString(uid)
	objname := function.MakePath("qr", province_name, city_name, district_name, station_name, office_name, community_name) + "/" + street_name + ".png"
	filename := "web/" + objname
	err = GenerateQrCode(url, filename)
	if err != nil {
		return
	}
	err = db.ExecUpdateStreetQr(uid, objname)
	if err == nil {
		res["uid"] = uid
		res["qr"] = objname
		res["msg"] = "小区生成成功！"
	}
	return
}
func DeleteArea(data *model.DeleteData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	switch data.Cmd_Delete {
	case model.Cmd_Station:
		err = db.ExecDeleteStation(data.Uid)
		break
	case model.Cmd_Office:
		err = db.ExecDeleteOffice(data.Uid)
		break
	case model.Cmd_Community:
		err = db.ExecDeleteCommunity(data.Uid)
		break
	case model.Cmd_Street:
		err = db.ExecDeleteStreet(data.Uid)
		break
	default:
		err = service.NewError(service.ERR_INVALID_PARAM, "输入cmd参数错误！")
	}
	if err == nil {
		res["deletestatus"] = 1
		res["msg"] = "删除成功！"
	}
	return
}
func UpdateStation(data *model.StationData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateStation(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "派出所更新成功！"
	}
	return
}
func UpdateOffice(data *model.OfficeData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateOffice(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "办事处更新成功！"
	}
	return
}
func UpdateCommunity(data *model.CommunityData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateCommunity(data)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "社区更新成功！"
	}
	return
}
func UpdateStreet(qrurl string, data *model.StreetData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	err = db.ExecUpdateStreet(data)
	if err != nil {
		return
	}
	err = UpdateQrCode(data.Uid, qrurl)
	if err == nil {
		res["updatestatus"] = 1
		res["msg"] = "小区更新成功！"
	}
	return
}
func UpdateQrCode(uid int64, qrurl string) (err error) {
	_, province_name, city_name, district_name, station_name, office_name, community_name, street_name, street_qrcode, err := db.GetStreetbyID(uid)
	if err != nil {
		return
	}
	url := qrurl + "?street_no=" + function.Int64ToString(uid)
	objname := function.MakePath("qr", province_name, city_name, district_name, station_name, office_name, community_name) + "/" + street_name + ".png"
	filename := "web/" + objname
	if objname != street_qrcode {
		_, e := os.Stat(filename)
		if e == nil {
			err = os.Remove(filename)
			if err != nil {
				return
			}
		}
	}
	err = GenerateQrCode(url, filename)
	if err != nil {
		return
	}
	err = db.ExecUpdateStreetQr(uid, objname)
	return
}
