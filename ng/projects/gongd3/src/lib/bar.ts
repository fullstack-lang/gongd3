// generated code - do not edit

import { BarAPI } from './bar-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Key } from './key'
import { Serie } from './serie'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Bar {

	static GONGSTRUCT_NAME = "Bar"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	AutoDomainX: boolean = false
	XMin: number = 0
	XMax: number = 0
	AutoDomainY: boolean = false
	YMin: number = 0
	YMax: number = 0
	YLabelPresent: boolean = false
	YLabelOffset: number = 0
	XLabelPresent: boolean = false
	XLabelOffset: number = 0
	Width: number = 0
	Heigth: number = 0
	Margin: number = 0

	// insertion point for pointers and slices of pointers declarations
	X?: Key

	Y?: Key

	Set: Array<Serie> = []
}

export function CopyBarToBarAPI(bar: Bar, barAPI: BarAPI) {

	barAPI.CreatedAt = bar.CreatedAt
	barAPI.DeletedAt = bar.DeletedAt
	barAPI.ID = bar.ID

	// insertion point for basic fields copy operations
	barAPI.Name = bar.Name
	barAPI.AutoDomainX = bar.AutoDomainX
	barAPI.XMin = bar.XMin
	barAPI.XMax = bar.XMax
	barAPI.AutoDomainY = bar.AutoDomainY
	barAPI.YMin = bar.YMin
	barAPI.YMax = bar.YMax
	barAPI.YLabelPresent = bar.YLabelPresent
	barAPI.YLabelOffset = bar.YLabelOffset
	barAPI.XLabelPresent = bar.XLabelPresent
	barAPI.XLabelOffset = bar.XLabelOffset
	barAPI.Width = bar.Width
	barAPI.Heigth = bar.Heigth
	barAPI.Margin = bar.Margin

	// insertion point for pointer fields encoding
	barAPI.BarPointersEncoding.XID.Valid = true
	if (bar.X != undefined) {
		barAPI.BarPointersEncoding.XID.Int64 = bar.X.ID  
	} else {
		barAPI.BarPointersEncoding.XID.Int64 = 0 		
	}

	barAPI.BarPointersEncoding.YID.Valid = true
	if (bar.Y != undefined) {
		barAPI.BarPointersEncoding.YID.Int64 = bar.Y.ID  
	} else {
		barAPI.BarPointersEncoding.YID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	barAPI.BarPointersEncoding.Set = []
	for (let _serie of bar.Set) {
		barAPI.BarPointersEncoding.Set.push(_serie.ID)
	}

}

// CopyBarAPIToBar update basic, pointers and slice of pointers fields of bar
// from respectively the basic fields and encoded fields of pointers and slices of pointers of barAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBarAPIToBar(barAPI: BarAPI, bar: Bar, frontRepo: FrontRepo) {

	bar.CreatedAt = barAPI.CreatedAt
	bar.DeletedAt = barAPI.DeletedAt
	bar.ID = barAPI.ID

	// insertion point for basic fields copy operations
	bar.Name = barAPI.Name
	bar.AutoDomainX = barAPI.AutoDomainX
	bar.XMin = barAPI.XMin
	bar.XMax = barAPI.XMax
	bar.AutoDomainY = barAPI.AutoDomainY
	bar.YMin = barAPI.YMin
	bar.YMax = barAPI.YMax
	bar.YLabelPresent = barAPI.YLabelPresent
	bar.YLabelOffset = barAPI.YLabelOffset
	bar.XLabelPresent = barAPI.XLabelPresent
	bar.XLabelOffset = barAPI.XLabelOffset
	bar.Width = barAPI.Width
	bar.Heigth = barAPI.Heigth
	bar.Margin = barAPI.Margin

	// insertion point for pointer fields encoding
	bar.X = frontRepo.map_ID_Key.get(barAPI.BarPointersEncoding.XID.Int64)
	bar.Y = frontRepo.map_ID_Key.get(barAPI.BarPointersEncoding.YID.Int64)

	// insertion point for slice of pointers fields encoding
	bar.Set = new Array<Serie>()
	for (let _id of barAPI.BarPointersEncoding.Set) {
		let _serie = frontRepo.map_ID_Serie.get(_id)
		if (_serie != undefined) {
			bar.Set.push(_serie!)
		}
	}
}
