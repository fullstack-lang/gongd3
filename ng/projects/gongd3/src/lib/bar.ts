// generated code - do not edit

import { BarDB } from './bar-db'
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

export function CopyBarToBarDB(bar: Bar, barDB: BarDB) {

	barDB.CreatedAt = bar.CreatedAt
	barDB.DeletedAt = bar.DeletedAt
	barDB.ID = bar.ID
	
	// insertion point for basic fields copy operations
	barDB.Name = bar.Name
	barDB.AutoDomainX = bar.AutoDomainX
	barDB.XMin = bar.XMin
	barDB.XMax = bar.XMax
	barDB.AutoDomainY = bar.AutoDomainY
	barDB.YMin = bar.YMin
	barDB.YMax = bar.YMax
	barDB.YLabelPresent = bar.YLabelPresent
	barDB.YLabelOffset = bar.YLabelOffset
	barDB.XLabelPresent = bar.XLabelPresent
	barDB.XLabelOffset = bar.XLabelOffset
	barDB.Width = bar.Width
	barDB.Heigth = bar.Heigth
	barDB.Margin = bar.Margin

	// insertion point for pointer fields encoding
    barDB.BarPointersEncoding.XID.Valid = true
	if (bar.X != undefined) {
		barDB.BarPointersEncoding.XID.Int64 = bar.X.ID  
    } else {
		barDB.BarPointersEncoding.XID.Int64 = 0 		
	}

    barDB.BarPointersEncoding.YID.Valid = true
	if (bar.Y != undefined) {
		barDB.BarPointersEncoding.YID.Int64 = bar.Y.ID  
    } else {
		barDB.BarPointersEncoding.YID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	barDB.BarPointersEncoding.Set = []
    for (let _serie of bar.Set) {
		barDB.BarPointersEncoding.Set.push(_serie.ID)
    }
	
}

export function CopyBarDBToBar(barDB: BarDB, bar: Bar, frontRepo: FrontRepo) {

	bar.CreatedAt = barDB.CreatedAt
	bar.DeletedAt = barDB.DeletedAt
	bar.ID = barDB.ID
	
	// insertion point for basic fields copy operations
	bar.Name = barDB.Name
	bar.AutoDomainX = barDB.AutoDomainX
	bar.XMin = barDB.XMin
	bar.XMax = barDB.XMax
	bar.AutoDomainY = barDB.AutoDomainY
	bar.YMin = barDB.YMin
	bar.YMax = barDB.YMax
	bar.YLabelPresent = barDB.YLabelPresent
	bar.YLabelOffset = barDB.YLabelOffset
	bar.XLabelPresent = barDB.XLabelPresent
	bar.XLabelOffset = barDB.XLabelOffset
	bar.Width = barDB.Width
	bar.Heigth = barDB.Heigth
	bar.Margin = barDB.Margin

	// insertion point for pointer fields encoding
	bar.X = frontRepo.Keys.get(barDB.BarPointersEncoding.XID.Int64)
	bar.Y = frontRepo.Keys.get(barDB.BarPointersEncoding.YID.Int64)

	// insertion point for slice of pointers fields encoding
	bar.Set = new Array<Serie>()
	for (let _id of barDB.BarPointersEncoding.Set) {
	  let _serie = frontRepo.Series.get(_id)
	  if (_serie != undefined) {
		bar.Set.push(_serie!)
	  }
	}
}
