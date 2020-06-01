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
	CusId     int    `json:"cusid"`
	FirstName string `json:"cusfirstname"`
	LastName  string `json:"cuslastname"`
	MobileNo  int    `json:"mobno"`
	Address   string `json:"address"`
	Location  string `json:"location"`
	Route     string `json:"route"`
	CusType   int    `json:"custype"`
}

type Request struct {
	RequestId            int       `json:"requestid"`
	CusId                int       `json:"cusid"`
	RequestDate          time.Time `json:"requestdate"`
	Complaint            string    `json:"complaint"`
	TypeOfService        int       `json:"typeofservice"`
	Section              string    `json:"section"`
	OtherItem            string    `json:"otheritem"`
	MaterialsUsed        string    `json:"materialsused"`
	SparePartsUsed       string    `json:"sparepartsused"`
	WorkStatus           int       `json:"workstatus"`
	PendingReason        string    `json:"pendingreason"`
	UdDetails            string    `json:"uddetails"`
	BrokerId             int       `json:"brokerid"`
	ItemsTakenfromClient string    `json:"itemstakenfromclient"`
	DeliveredItems       string    `json:"deliveredItems"`
	DeliveryDate         time.Time `json:"deliverydate"`
	DeliveredBy          int       `json:"deliverydate"`
	Verify               string    `json:"verify"`
	VerifiedBy           int       `json:"verifiedby"`
	TechnicianId         string    `json:"technicianid"`
	WorkDetails          string    `json:"workdetails"`
	WorkStartDate        time.Time `json:"workstartdate"`
	WorkEndDate          time.Time `json:"workenddate"`
	CancellationDate     time.Time `json:"cancellationdate"`
	CancellationReason   string    `json:"cancellattionreason"`
	CustomerApproval     string    `json:"customerapproval"`
	OverTimeReason       string    `json:"overtimereason"`
	OldReqId             int       `json:"oldreqid"`
	CareOf               int       `json:"careof"`
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
	RefCode            string `json:"value"`
	RefCodeDescription string `json:"refcodedescription"`
}
