package imports

import (
	"datahelper/db"
	"errors"
	"model"
	"utils/function"

	"github.com/tealeg/xlsx"
)

type ExcelInfo struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
	I string
	J string
	K string
	L string
	M string
	N string
	O string
	P string
}

type DataInfo struct {
	House_ID           int64
	Personnel_ID       int64
	Street_No          int64
	Street_Name        string
	Addrees            string
	Nature             string
	Personnel_Name     string
	Personnel_Sex      int
	Personnel_Card_No  string
	Personnel_Birthday string
	Personnel_Home     string
	Personnel_Address  string
	Personnel_Phone    string
	Relation_Role      int
	Relation_Holder    int
	Relation_Together  int
}

func AddExcelData(xlFile *xlsx.File) (res map[string]interface{}, err error) {
	res = make(map[string]interface{}, 0)
	listlistExcelInfo, err := ReadExcelData(xlFile)
	if err != nil {
		return
	}
	skip, houseaffect, personnelaffect := ExecAddExcelData(listlistExcelInfo)
	res["msg"] = "导入数据成功！"
	res["skip"] = skip
	res["houseaffect"] = houseaffect
	res["personnelaffect"] = personnelaffect
	return
}
func ReadExcelData(xlFile *xlsx.File) (listExcelInfo []ExcelInfo, err error) {
	for _, sheet := range xlFile.Sheets {
		info := ExcelInfo{}
		for _, row := range sheet.Rows {
			var strs []string
			for _, cell := range row.Cells {
				text := cell.String()
				strs = append(strs, text)
			}
			info.A = strs[0]
			info.B = strs[1]
			info.C = strs[2]
			info.D = strs[3]
			info.E = strs[4]
			info.F = strs[5]
			info.G = strs[6]
			info.H = strs[7]
			info.I = strs[8]
			info.J = strs[9]
			info.K = strs[10]
			info.L = strs[11]
			info.M = strs[12]
			info.N = strs[13]
			info.O = strs[14]
			info.P = strs[15]
			listExcelInfo = append(listExcelInfo, info)
		}
	}
	defer func() {

		if r := recover(); r != nil {
			//check exactly what the panic was and create error.
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Excel表格格式有误！")
			}
		}
	}()
	return
}
func ExecAddExcelData(listExcelInfo []ExcelInfo) (skip int, houseaffect int, personnelaffect int) {
	for k, i := range listExcelInfo {
		data := DataInfo{}
		if i.A == "区县" && i.B == "街道办事处" || i.C == "社区" || i.D == "小区名称" {
			continue
		}
		streetid, name, e1 := db.GetStreetbyParam(i.A, i.B, i.C, i.D)
		if e1 != nil || streetid == 0 {
			skip++
			Info.Info("GetStreetbyParam,已跳过第%d条数据,区县%s街道办事处%s社区%s小区名称%s", k, i.A, i.B, i.C, i.D)
			continue
		}
		data.Street_Name = name
		data.Street_No = streetid
		data.Addrees = i.E + "号楼" + i.F + "单元" + i.G
		houseid, e2 := db.GetHousebyParam(data.Street_No, data.Street_Name, data.Addrees)
		data.House_ID = houseid
		if e2 != nil {
			Info.Info("GetHousebyParam,出现异常,异常信息为:%s", e2.Error())
			continue
		} else if houseid == 0 {
			if i.H == "" {
				data.Nature = "闲置"
			} else if i.P == "是" {
				data.Nature = "出租"
			} else {
				data.Nature = "自住"
			}
			house := &model.HouseData{0, data.Nature, data.Street_Name, data.Street_No, data.Addrees, 0, "", ""}
			houseid, e3 := db.ExecCreateHouse(house)
			data.House_ID = houseid
			if e3 != nil {
				Info.Info("ExecCreateHouse,创建房屋出现异常,异常信息为:%s", e3.Error())
				continue
			}
			houseaffect++
		}
		data.Personnel_Card_No = i.K
		if !function.ValidateCardNo(data.Personnel_Card_No) {
			Info.Info("ValidateCardNo,第%d条数据,身份证%s不合法", k, i.K)
			continue
		}
		personnelid, e4 := db.GetPersonnelbyNo(data.Personnel_Card_No)
		data.Personnel_ID = personnelid
		if e4 != nil {
			Info.Info("GetPersonnelbyNo,出现异常,异常信息为:%s", e4.Error())
			continue
		} else if personnelid == 0 {
			data.Personnel_Name = i.H
			if i.I == "女" {
				data.Personnel_Sex = 2
			} else {
				data.Personnel_Sex = 1
			}
			data.Personnel_Birthday = function.Substr(i.K, 6, 4) + "-" + function.Substr(i.K, 10, 2) + "-" + function.Substr(i.K, 12, 2)
			data.Personnel_Home = i.L
			data.Personnel_Address = i.M
			data.Personnel_Phone = i.O
			personnel := &model.PersonnelData{0, data.Personnel_Name, "", data.Personnel_Card_No, "", "", "", data.Personnel_Sex, "", data.Personnel_Birthday, data.Personnel_Address, "", "", "", data.Personnel_Phone, "", data.Personnel_Home}
			personnelid, e5 := db.ExecCreatePersonnel(personnel)
			data.Personnel_ID = personnelid
			if e5 != nil {
				Info.Info("ExecCreatePersonnel,创建人员出现异常,异常信息为:%s", e5.Error())
				continue
			}
			personnelaffect++
		}
		isexist, e6 := db.CheckPersonnelHouse(data.House_ID, data.Personnel_ID)
		if e6 != nil {
			Info.Info("CheckPersonnelHouse,出现异常,异常信息为:%s", e6.Error())
			continue
		} else if !isexist {
			if i.P == "是" {
				data.Relation_Role = 3
			} else {
				data.Relation_Role = 2
				data.Relation_Together = 4
			}
			housepersonnel := &model.HousePersonnelData{data.House_ID, data.Personnel_ID, data.Relation_Role, data.Relation_Holder, data.Relation_Together}
			_, e7 := db.ExecAddRelationHousePersonnel(housepersonnel)
			if e7 != nil {
				Info.Info("ExecCreatePersonnel,创建房屋人员关系出现异常,异常信息为:%s", e7.Error())
				continue
			}
		}
	}
	return
}
