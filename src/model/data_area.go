package model

const (
	Cmd_Station   = "create"
	Cmd_Community = "community"
	Cmd_Street    = "street"
)

type StationData struct {
	Station_Name  string
	Province_Name string
	Province_No   int
	City_Name     string
	City_No       int
	District_Name string
	District_No   int
}
type CommunityData struct {
	Community_Name string
	Station_No     int
}
type StreetData struct {
	Street_Name  string
	Community_No int
}
type DeleteData struct {
	Cmd_Delete string
	Uid        int
}
type UpdateData struct {
	Cmd_Update  string
	Uid         int
	Update_Name string
}
