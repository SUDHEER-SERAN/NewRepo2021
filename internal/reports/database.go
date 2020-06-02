package reports

import (
	"context"
	"errors"
	"jobreport/internal/database"
	"jobreport/internal/reportmodel"

	"github.com/sirupsen/logrus"
)

type Database struct {
	database *database.PostgresDatabase
}

func NewReportDatabase(database *database.PostgresDatabase) *Database {
	return &Database{
		database: database,
	}
}

func (d *Database) getReferenceListById(ctx context.Context, id int) ([]reportmodel.LookupRef, error) {
	var listOfRef []reportmodel.LookupRef
	rows, err := d.database.Conn.Query(ctx, "select refcodeid,refcode from treferencecode where refid=$1", id)
	if err != nil {
		logrus.WithError(err).Warn("unable to select  doc in treferencecode")
		return listOfRef, errors.New("unable to select  doc in in treferencecode")
	}

	defer rows.Close()

	for rows.Next() {
		typeRef := reportmodel.LookupRef{}
		err := rows.Scan(&typeRef.RefCodeId, &typeRef.RefCode)

		if err != nil {
			logrus.WithError(err).Warn("unable to select  doc in treferencecode")
			return listOfRef, errors.New("unable to select  doc in treferencecode")
		}
		listOfRef = append(listOfRef, typeRef)
	}

	return listOfRef, nil
}

func (d *Database) generateReport(ctx context.Context, reportEntity JobReportBasicDetails) error {
	conn := d.database.Conn

	tx, err := conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		logrus.WithError(err).Warn("unable to Begin the Transaction in generateReport")
		return errors.New("unable to Fetech doc metadata in treferencecode")
	}
	var cusId int
	if err := tx.QueryRow(ctx, `select custid from tcustomer where mobileno=$1`, reportEntity.MobileNo).Scan(&cusId); err != nil {
		logrus.WithError(err).Warn("unable to Select the record in tcustomer generateReport")
	}
	if cusId == 0 {
		if err := tx.QueryRow(ctx, `insert into tcustomer(firstname,lastname,mobileno,address,location,route,custtype) values($1,$2,$3,$4,$5,$6,$7) returning custid`, reportEntity.FirstName, reportEntity.LastName, reportEntity.MobileNo, reportEntity.Address, reportEntity.Location, reportEntity.Route, reportEntity.CusType).Scan(&cusId); err != nil {
			logrus.WithError(err).Warn("unable to Insert the record in tcustomer generateReport")
			return errors.New("unable to Fetech doc metadata in treferencecode")
		}
	}
	if _, err := tx.Exec(ctx, `insert into trequest(custid,typeofservice,requestdate,careof,otheritem) values($1,$2,$3,$4,$5)`, cusId, reportEntity.TypeOfService.RefCodeId, reportEntity.RequestDate, reportEntity.CareOf.RefCodeId, reportEntity.OtherItem); err != nil {
		logrus.WithError(err).Warn("unable to Insert the record in tcustomer generateReport")
		return errors.New("unable to Fetech doc metadata in treferencecode")
	}
	tx.Commit(ctx)

	return nil

}
func (d *Database) getReferenceList(ctx context.Context, id int, searchKey string) ([]reportmodel.LookupRef, error) {
	var listOfRef []reportmodel.LookupRef
	rows, err := d.database.Conn.Query(ctx, "select refcodeid,refcode from treferencecode where refid=$1 and lower(refcode) like  lower($2) || '%' limit 20", id, searchKey)
	if err != nil {
		logrus.WithError(err).Warn("unable to select  doc in treferencecode")
		return listOfRef, errors.New("unable to select  doc in in treferencecode")
	}

	defer rows.Close()

	for rows.Next() {
		typeRef := reportmodel.LookupRef{}
		err := rows.Scan(&typeRef.RefCodeId, &typeRef.RefCode)

		if err != nil {
			logrus.WithError(err).Warn("unable to select  doc in treferencecode")
			return listOfRef, errors.New("unable to select  doc in treferencecode")
		}
		listOfRef = append(listOfRef, typeRef)
	}

	return listOfRef, nil
}
func (d *Database) getCustomerList(ctx context.Context, searchKey string) ([]CustomerList, error) {
	var listOfRef []CustomerList
	rows, err := d.database.Conn.Query(ctx, "select custid,firstname,lastname from tcustomer where lower(firstname) like  lower($1) || '%' limit 20", searchKey)
	if err != nil {
		logrus.WithError(err).Warn("unable to select  doc in treferencecode")
		return listOfRef, errors.New("unable to select  doc in in treferencecode")
	}

	defer rows.Close()

	for rows.Next() {
		typeRef := CustomerList{}
		err := rows.Scan(&typeRef.CusId, &typeRef.FirstName, &typeRef.LastName)
		typeRef.Value = typeRef.FirstName + " " + typeRef.LastName
		typeRef.Id = typeRef.CusId
		if err != nil {
			logrus.WithError(err).Warn("unable to select  doc in tcustomer")
			return listOfRef, errors.New("unable to select  doc in tcustomer")
		}
		listOfRef = append(listOfRef, typeRef)
	}

	return listOfRef, nil
}

func (d *Database) getReports(ctx context.Context) ([]JobReportBasicDetails, error) {
	// rows, err := d.database.Conn.Query(ctx, "SELECT req.requestid requestid,getcodedescbyid(req.typeofservice) typeofservice,req.requestdate requestdate,cus.firstname firstname from trequest req inner join tcustomer cus on cus.custid=req.custid")
	// if err != nil {
	// 	logrus.WithError(err).Warn("unable to select  doc in treferencecode")
	// 	return []JobReportBasicDetails{}, errors.New("unable to select  doc in in treferencecode")
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	typeRef := JobReportBasicDetails{}
	// 	err := rows.Scan(&typeRef.RequestId, &typeRef.FirstName, &typeRef.LastName)
	// 	typeRef.Value = typeRef.FirstName + " " + typeRef.LastName
	// 	typeRef.Id = typeRef.CusId
	// 	if err != nil {
	// 		logrus.WithError(err).Warn("unable to select  doc in tcustomer")
	// 		return listOfRef, errors.New("unable to select  doc in tcustomer")
	// 	}
	// 	listOfRef = append(listOfRef, typeRef)
	// }

	return nil, nil
}
