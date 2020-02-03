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
	nature    string
	street    string
	street_no string
	address   string
	number    string
}
type UnitData struct {
	name                  string
	house_id              int
	license_number        int
	identification_number int
	picture               string
	kind                  string
	scale                 int
	tel                   int
	bank_name             string
	bank_account          string
	comment               string
}
type PersonnelData struct {
	name              string
	occupation        string
	card_no           string
	picture           string
	sex               int
	nation            string
	birthday          string
	address           string
	sign_organization string
	limited_date      string
	history           string
}
