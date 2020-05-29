package reports

import (
	"context"
	"errors"
	"jobreport/internal/database"

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

func (d *Database) initializeReport(ctx context.Context, id int) ([]ReportListReference, error) {
	var listOfRef []ReportListReference
	rows, err := d.database.Conn.Query(ctx, "select refcodeid,refcode from treferencecode where refid=$1", id)
	if err != nil {
		logrus.WithError(err).Warn("unable to select  doc in treferencecode")
		return listOfRef, errors.New("unable to select  doc in in treferencecode")
	}

	defer rows.Close()

	for rows.Next() {
		typeRef := ReportListReference{}
		err := rows.Scan(&typeRef.Id, &typeRef.Value)

		if err != nil {
			logrus.WithError(err).Warn("unable to select  doc in treferencecode")
			return listOfRef, errors.New("unable to select  doc in treferencecode")
		}
		listOfRef = append(listOfRef, typeRef)
	}

	return listOfRef, nil
}
