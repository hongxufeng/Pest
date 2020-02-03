package model

type Settings struct {
	Style       string
	TableID     string
	ConfigFile  string
	HasCheckbox bool
	RowList     string
	Condition   string
	Page        int
	Rows        int
	ColPage     int
	Order       string
}

type LoginData struct {
	Username string
	Password string
}
type RegisterData struct {
	Username string
	Password string
	Nickname string
	Power    int
}
type CRUDSettings struct {
	Cmd        string
	ConfigFile string
	TableID    string
	QueryKey   string
	QueryValue string
}
type HouseData struct {
	Nature    string
	Street    string
	Street_No string
	Address   string
	Number    string
}
type UnitData struct {
	Name                  string
	House_ID              int
	License_Number        int
	Identification_Number int
	Picture               string
	Kind                  string
	Scale                 int
	Tel                   int
	Bank_Name             string
	Bank_Account          string
	Comment               string
}
type PersonnelData struct {
	Name              string
	Occupation        string
	Card_No           string
	Picture           string
	Sex               int
	Nation            string
	Birthday          string
	Address           string
	Sign_Organization string
	Limited_Date      string
	History           string
}
type HousePersonnelData struct {
	House_ID        int
	Personnel_ID    int
	Holder_Flag     int
	Relation_Holder string
}
type UnitPersonnelData struct {
	Unit_ID      int
	Personnel_ID int
	Position     string
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
	Hospitalized_Flag int
	Temperature       float64
	Touch_People      string
}
