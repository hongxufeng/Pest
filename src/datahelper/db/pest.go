package db

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
