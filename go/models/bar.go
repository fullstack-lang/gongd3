package models

type Bar struct {

	// to be used by the client to identify which bar to display
	Name string

	// data to display
	X   *Key
	Y   *Key
	Set []*Serie

	// display settings
	Width  float64
	Heigth float64
	Margin float64
}
