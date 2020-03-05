package model

const (
	Cmd_Station   = "station"
	Cmd_Office    = "office"
	Cmd_Community = "community"
	Cmd_Street    = "street"
)

type StationData struct {
	Uid           int64
	Station_Name  string
	Province_Name string
	Province_No   int
	City_Name     string
	City_No       int
	District_Name string
	District_No   int
	Station_Head  string
	Station_Phone string
}
type OfficeData struct {
	Uid          int64
	Office_Name  string
	Station_No   int64
	Office_Head  string
	Office_Phone string
}
type CommunityData struct {
	Uid             int64
	Community_Name  string
	Office_No       int64
	Community_Head  string
	Community_Phone string
}
type StreetData struct {
	Uid                   int64
	Street_Name           string
	Community_No          int64
	Street_Head           string
	Street_Phone          string
	Street_Property_Name  string
	Street_Property_Phone string
}
type DeleteData struct {
	Cmd_Delete string
	Uid        int64
}
