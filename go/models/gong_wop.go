// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Bar_WOP struct {
	// insertion point
	Name string
	AutoDomainX bool
	XMin float64
	XMax float64
	AutoDomainY bool
	YMin float64
	YMax float64
	YLabelPresent bool
	YLabelOffset float64
	XLabelPresent bool
	XLabelOffset float64
	Width float64
	Heigth float64
	Margin float64
}

func (from *Bar) CopyBasicFields(to *Bar) {
	// insertion point
	to.Name = from.Name
	to.AutoDomainX = from.AutoDomainX
	to.XMin = from.XMin
	to.XMax = from.XMax
	to.AutoDomainY = from.AutoDomainY
	to.YMin = from.YMin
	to.YMax = from.YMax
	to.YLabelPresent = from.YLabelPresent
	to.YLabelOffset = from.YLabelOffset
	to.XLabelPresent = from.XLabelPresent
	to.XLabelOffset = from.XLabelOffset
	to.Width = from.Width
	to.Heigth = from.Heigth
	to.Margin = from.Margin
}

type Key_WOP struct {
	// insertion point
	Name string
}

func (from *Key) CopyBasicFields(to *Key) {
	// insertion point
	to.Name = from.Name
}

type Pie_WOP struct {
	// insertion point
	Name string
	Width float64
	Heigth float64
	Margin float64
}

func (from *Pie) CopyBasicFields(to *Pie) {
	// insertion point
	to.Name = from.Name
	to.Width = from.Width
	to.Heigth = from.Heigth
	to.Margin = from.Margin
}

type Scatter_WOP struct {
	// insertion point
	Name string
	Width float64
	Heigth float64
	Margin float64
}

func (from *Scatter) CopyBasicFields(to *Scatter) {
	// insertion point
	to.Name = from.Name
	to.Width = from.Width
	to.Heigth = from.Heigth
	to.Margin = from.Margin
}

type Serie_WOP struct {
	// insertion point
	Name string
}

func (from *Serie) CopyBasicFields(to *Serie) {
	// insertion point
	to.Name = from.Name
}

type Value_WOP struct {
	// insertion point
	Name string
}

func (from *Value) CopyBasicFields(to *Value) {
	// insertion point
	to.Name = from.Name
}

