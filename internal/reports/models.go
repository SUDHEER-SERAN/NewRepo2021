package reports

type ReportListReference struct{
	Value string `json:"value"`
	Id int `json:"id"`
}

type ReportCombainedReference struct{
	ServiceTypes []ReportListReference `json:"serviceType"`
	Co []ReportListReference `json:"co"`
}