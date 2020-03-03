package model

type HouseData struct {
	Uid             int64
	Nature          string
	Street          string
	Street_No       int64
	Address         string
	Number          int
	Comment         string
	Creator_Card_No string
}
type UnitData struct {
	Uid                   int64
	Name                  string
	House_ID              int64
	License_Number        string
	Identification_Number string
	Picture               string
	Kind                  int
	Scale                 int
	Tel                   string
	Bank_Name             string
	Bank_Account          string
	Comment               string
	Creator_Card_No       string
}
type PersonnelData struct {
	Uid                int64
	Name               string
	Occupation         string
	Card_No            string
	Card_Picture_Front string
	Card_Picture_Back  string
	Face_Picture       string
	Sex                int
	Nation             string
	Birthday           string
	Address            string
	Sign_Organization  string
	Limited_Date       string
	History            string
	Phone              string
	Remark             string
	Home               string
}
type HousePersonnelData struct {
	House_ID          int64
	Personnel_ID      int64
	Role              int
	Relation_Holder   int
	Relation_Together int
}
type UnitPersonnelData struct {
	Unit_ID      int64
	Personnel_ID int64
	Position     int
}
type TouchData struct {
	Uid          int64
	Personnel_ID int64
	Way          string
	Time         string
	Place        string
	Touch_Number int
	Touch_People string
}
type DailyReportData struct {
	Uid               int64
	Personnel_ID      int64
	Symptom           string
	Time              string
	Hospitalized_Flag int
	Temperature       float64
	Touch_People      string
}
type IsolationData struct {
	Uid                int64
	House_ID           int64
	Flag               int
	Reason             string
	Start_Time         string
	End_Time           string
	Real_Compelte_Time string
	Unit_ID            string
}
type StructureData struct {
	Uid               int64
	Parent_ID         int64
	Unit_ID           int64
	Structure_Name    string
	Structure_Comment string
}
type StructurePersonnelData struct {
	Structure_ID       int64
	Personnel_ID       int64
	Structure_Position string
}
