package area

import (
	"datahelper/db"
	"model"
	"os"
	"path"
	"utils/function"

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
func AddCommunity(data *model.CommunityData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddCommunity(data)
	if err == nil {
		res["uid"] = uid
		res["msg"] = "社区生成成功！"
	}
	return
}
func AddStreet(data *model.StreetData) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	uid, err := db.ExecAddStreet(data)
	if err != nil {
		return
	}
	_, province_name, city_name, district_name, station_name, community_name, street_name, err := db.GetStreetbyID(uid)
	if err != nil {
		return
	}
	url := "http://47.92.86.80:8081/index.html?street_no=" + function.Int64ToString(uid)
	objname := function.MakePath("qr", province_name, city_name, district_name, station_name, community_name) + "/" + street_name + ".png"
	filename := "web/" + objname
	err = GenerateQrCode(url, filename)
	if err != nil {
		return
	}
	err = db.ExecUpdateStreetQr(objname)
	if err == nil {
		res["uid"] = uid
		res["qr"] = objname
		res["msg"] = "街道生成成功！"
	}
	return
}
