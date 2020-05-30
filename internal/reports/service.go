package reports

import (
	"context"
	"errors"
)

type Service interface {
	initializePage(ctx context.Context) (ReportCombainedReference, error)
	generateReport(ctx context.Context, reportEntity Report)
}
type reportService struct {
	database *Database
}

func NewReportService(d *Database) Service {
	s := &reportService{
		database: d,
	}

	return s
}

func (s *reportService) initializePage(ctx context.Context) (ReportCombainedReference, error) {
	var combainedRef ReportCombainedReference
	serviceTypes, err := s.database.initializeReport(ctx, 2)
	if err != nil {
		return combainedRef, errors.New("unable to Fetech doc metadata in treferencecode")
	}

	co, err := s.database.initializeReport(ctx, 6)
	if err != nil {
		return combainedRef, errors.New("unable to Fetech doc metadata in treferencecode")
	}

	reportCombained := ReportCombainedReference{
		ServiceTypes: serviceTypes,
		Co:           co,
	}

	return reportCombained, nil

}
func (s *reportService) generateReport(ctx context.Context, reportEntity Report) {

}
