package models

type Bar struct {

	// to be used by the client to identify which bar to display
	Name string

	// data to display
	X   *Key
	Y   *Key
	Set []*Serie

	// Domain for X & Y
	AutoDomainX bool // computes automaticaly XMin and XMax
	XMin        float64
	XMax        float64

	AutoDomainY bool // computes automaticaly YMin and YMax
	YMin        float64
	YMax        float64

	YLabelPresent bool
	YLabelOffset  float64

	XLabelPresent bool
	XLabelOffset  float64

	// display settings
	Width  float64
	Heigth float64
	Margin float64
}
