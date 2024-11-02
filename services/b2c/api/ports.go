package api

import (
	"context"
	"github.com/ercross/zax/services/b2c/data"
)

type Repository interface {
	AuthenticateProduct(ctx context.Context, code string, option data.ScanOption, location data.ScanLocation) (data.ModelProduct, error)
	AuthenticateBatch(ctx context.Context, code string, location data.ScanLocation) (data.ModelProductPackage, error)
	ReportCounterfeit(ctx context.Context, report data.ModelCounterfeitReport) error
	FetchCounterfeitReportsByLocation(ctx context.Context, location data.ScanLocation, sweepRadius int) ([]data.ModelCounterfeitReport, error)
}

type AccountsService interface {
	IsAdmin(ctx context.Context, authToken string) error
}
