package reportmodel

import (
	"time"
)

type EmployeeDetails struct {
	EmpId     int    `json:"empid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	MobileNo  int    `json:"mobno"`
	Address   string `json:"address"`
	Role      int    `json:"role"`
}

type Users struct {
	EmpId    int    `json:"empid"`
	UserId   int    `json:"userid"`
	Password string `json:"password"`
	UserName string `json:"username"`
}

type Customer struct {
	CusId     int    `json:"cusid,omitempty"`
	FirstName string `json:"cusfirstname,omitempty"`
	LastName  string `json:"cuslastname,omitempty"`
	MobileNo  int    `json:"mobno,omitempty"`
	Address   string `json:"address,omitempty"`
	Location  string `json:"location,omitempty"`
	Route     string `json:"route,omitempty"`
	CusType   int    `json:"custype,omitempty"`
}

type Request struct {
	RequestId            int       `json:"requestid,omitempty"`
	CusId                int       `json:"cusid,omitempty"`
	RequestDate          time.Time `json:"requestdate,omitempty"`
	Complaint            string    `json:"complaint,omitempty"`
	TypeOfService        int       `json:"typeofservice,omitempty"`
	Section              string    `json:"section,omitempty"`
	OtherItem            string    `json:"otheritem,omitempty"`
	MaterialsUsed        string    `json:"materialsused,omitempty"`
	SparePartsUsed       string    `json:"sparepartsused,omitempty"`
	WorkStatus           int       `json:"workstatus,omitempty"`
	PendingReason        string    `json:"pendingreason,omitempty"`
	UdDetails            string    `json:"uddetails,omitempty"`
	BrokerId             int       `json:"brokerid,omitempty"`
	ItemsTakenfromClient string    `json:"itemstakenfromclient,omitempty"`
	DeliveredItems       string    `json:"deliveredItems,omitempty"`
	DeliveryDate         time.Time `json:"deliverydate,omitempty"`
	DeliveredBy          int       `json:"deliverydate,omitempty"`
	Verify               string    `json:"verify,omitempty"`
	VerifiedBy           int       `json:"verifiedby,omitempty"`
	TechnicianId         string    `json:"technicianid,omitempty"`
	WorkDetails          string    `json:"workdetails,omitempty"`
	WorkStartDate        time.Time `json:"workstartdate,omitempty"`
	WorkEndDate          time.Time `json:"workenddate,omitempty"`
	CancellationDate     time.Time `json:"cancellationdate,omitempty"`
	CancellationReason   string    `json:"cancellattionreason,omitempty"`
	CustomerApproval     string    `json:"customerapproval,omitempty"`
	OverTimeReason       string    `json:"overtimereason,omitempty"`
	OldReqId             int       `json:"oldreqid,omitempty"`
	CareOf               int       `json:"careof,omitempty"`
}

type BrokerDetails struct {
	BrokerId  int    `json:"brokerid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	MobNo     int    `json:"mobno"`
}

type Payment struct {
	PayId            int `json:"payid"`
	RequestId        int `json:"requestid"`
	ModeofPayment    int `json:"modeofpayment"`
	EstimationAmount int `json:"estimationamount"`
	AgreedAmount     int `json:"agreedamount"`
	ActualAmount     int `json:"actualamount"`
	AdvanceAmount    int `json:"advanceamount"`
	PaymentSatus     int `json:"paymentstatus"`
}

type Charges struct {
	RequestId        int `json:"requestid"`
	InspectionCharge int `json:inceptioncharge`
	SparePartsAmount int `json:"sparepartsamount"`
	AdditionToolRent int `json:"additiontoolrent"`
	TransportCharges int `json:"transportcharges"`
	LatheworkCharge  int `json:"latheworkcharge"`
	VendorCost       int `json:"vendorcost"`
	BrokerCharge     int `json:"brokercharge"`
	Misc             int `json:"misc"`
	TotalAmount      int `json:"totalamount"`
}

type LookupRef struct {
	RefCodeId          int    `json:"id"`
	RefId              int    `json:"refid"`
	RefCode            string `json:"value,"`
	RefCodeDescription string `json:"refcodedescription"`
}
