package reports

import "time"

type ReportListReference struct {
	Value string `json:"value"`
	Id    int    `json:"id"`
}

type ReportCombainedReference struct {
	ServiceTypes []ReportListReference `json:"serviceType"`
	Co           []ReportListReference `json:"co"`
}

type Report struct {
	EnquiryDate    time.Time           `json:"enquiryDate"`
	CustomerName   string              `json:customerName`
	MobNo          int                 `json:mobNo`
	Address        string              `json:address`
	Location       string              `json:location`
	TypeOfService  ReportListReference `json:typeOfService`
	Co             ReportListReference `json:co`
	TypeOfCustomer int                 `json:typeOfCustomer`
}
