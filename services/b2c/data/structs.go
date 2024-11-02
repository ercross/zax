package data

type ScanOption string

const (
	ScanOptionTrackerChip      ScanOption = "TrackerChip"
	ScanOptionQRCode           ScanOption = "QRCode"
	ScanOptionAlphanumericCode ScanOption = "AlphanumericCode"
)

type DaoProduct struct{}
type ModelProduct struct{}
type DaoCounterfeitReport struct{}
type ModelCounterfeitReport struct{}
type DaoProductPackage struct{}
type ModelProductPackage struct{}
type ScanLocation struct{}
