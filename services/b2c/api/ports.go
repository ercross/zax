package api

import "github.com/ercross/zax/services/b2c/data"

type Repository interface {
	AuthenticateProduct(code string, option data.ScanOption, location data.ScanLocation) (data.ModelProduct, error)
	AuthenticateBatch(code string, location data.ScanLocation) (data.ModelProductPackage, error)
	ReportCounterfeit(report data.ModelCounterfeitReport) error
	FetchCounterfeitReportsByLocation(location data.ScanLocation, sweepRadius int) ([]data.ModelCounterfeitReport, error)
}
