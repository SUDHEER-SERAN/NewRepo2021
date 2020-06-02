package reports

import (
	"context"
	"errors"
	"jobreport/internal/reportmodel"
)

type Service interface {
	initializePage(ctx context.Context) (ReportCombainedReference, error)
	generateReport(ctx context.Context, reportEntity JobReportBasicDetails) error
	getReports(ctx context.Context) error
	getjrList(ctx context.Context, id int, searchKey string) ([]reportmodel.LookupRef, error)
	getCustomerList(ctx context.Context, searchKey string) ([]CustomerList, error)
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
	serviceTypes, err := s.database.getReferenceListById(ctx, 2)
	if err != nil {
		return combainedRef, errors.New("unable to Fetech doc metadata in treferencecode")
	}

	co, err := s.database.getReferenceListById(ctx, 6)
	if err != nil {
		return combainedRef, errors.New("unable to Fetech doc metadata in treferencecode")
	}

	reportCombained := ReportCombainedReference{
		ServiceTypes: serviceTypes,
		Co:           co,
	}

	return reportCombained, nil

}
func (s *reportService) generateReport(ctx context.Context, reportEntity JobReportBasicDetails) error {

	if err := s.database.generateReport(ctx, reportEntity); err != nil {
		return errors.New("unable to Create the report")
	}
	return nil
}

func (s *reportService) getReports(ctx context.Context) error {

	// if reports, err := s.database.getReports(ctx); err != nil {
	// 	return errors.New("unable to Fetch the reports")
	// }
	return nil
}

func (s *reportService) getjrList(ctx context.Context, id int, searchKey string) ([]reportmodel.LookupRef, error) {

	list, err := s.database.getReferenceList(ctx, id, searchKey)
	if err != nil {
		return nil, errors.New("unable to Fetch the list")
	}
	if list == nil {
		var refList = []reportmodel.LookupRef{}
		return refList, nil
	}
	return list, nil

}
func (s *reportService) getCustomerList(ctx context.Context, searchKey string) ([]CustomerList, error) {

	list, err := s.database.getCustomerList(ctx, searchKey)
	if err != nil {
		return nil, errors.New("unable to Fetch the list")
	}
	if list == nil {
		var refList = []CustomerList{}
		return refList, nil
	}
	return list, nil

}
