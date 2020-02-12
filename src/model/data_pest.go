package model

type HouseData struct {
	Nature          string
	Street          string
	Street_No       int
	Address         string
	Number          int
	Comment         string
	Creator_Card_No string
}
type UnitData struct {
	Name                  string
	House_ID              int
	License_Number        string
	Identification_Number string
	Picture               string
	Kind                  int
	Scale                 int
	Tel                   string
	Bank_Name             string
	Bank_Account          string
	Comment               string
}
type PersonnelData struct {
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
}
type HousePersonnelData struct {
	House_ID          int
	Personnel_ID      int
	Role              int
	Relation_Holder   int
	Relation_Together int
}
type UnitPersonnelData struct {
	Unit_ID      int
	Personnel_ID int
	Position     int
}
type TouchData struct {
	Personnel_ID int
	Way          string
	Time         string
	Place        string
	Touch_Number int
	Touch_People string
}
type DailyReportData struct {
	Personnel_ID      int
	Symptom           string
	Time              string
	Hospitalized_Flag int
	Temperature       float64
	Touch_People      string
}
type IsolationData struct {
	House_ID           int
	Flag               int
	Reason             string
	Start_Time         string
	End_Time           string
	Real_Compelte_Time string
	Unit_ID            string
}
