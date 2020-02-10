package area

import (
	"datahelper/db"
	"model"
)

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
	if err == nil {
		res["uid"] = uid
		res["msg"] = "街道生成成功！"
	}
	return
}
