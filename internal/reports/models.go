package reports

import (
	"jobreport/internal/reportmodel"
)

type ReportCombainedReference struct {
	ServiceTypes []reportmodel.LookupRef `json:"serviceType"`
	Co           []reportmodel.LookupRef `json:"co"`
}

type JobReportBasicDetails struct {
	reportmodel.Customer
	reportmodel.Request
	TypeOfService reportmodel.LookupRef `json:"typeofservice"`
	CareOf        reportmodel.LookupRef `json:"careof"`
}

type CustomerList struct {
	reportmodel.Customer
	Id    int    `json:"id"`
	Value string `json:"value"`
}
