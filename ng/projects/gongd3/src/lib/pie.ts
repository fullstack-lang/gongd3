// generated code - do not edit

import { PieAPI } from './pie-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Key } from './key'
import { Serie } from './serie'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Pie {

	static GONGSTRUCT_NAME = "Pie"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Heigth: number = 0
	Margin: number = 0

	// insertion point for pointers and slices of pointers declarations
	X?: Key

	Y?: Key

	Set: Array<Serie> = []
}

export function CopyPieToPieAPI(pie: Pie, pieAPI: PieAPI) {

	pieAPI.CreatedAt = pie.CreatedAt
	pieAPI.DeletedAt = pie.DeletedAt
	pieAPI.ID = pie.ID

	// insertion point for basic fields copy operations
	pieAPI.Name = pie.Name
	pieAPI.Width = pie.Width
	pieAPI.Heigth = pie.Heigth
	pieAPI.Margin = pie.Margin

	// insertion point for pointer fields encoding
	pieAPI.PiePointersEncoding.XID.Valid = true
	if (pie.X != undefined) {
		pieAPI.PiePointersEncoding.XID.Int64 = pie.X.ID  
	} else {
		pieAPI.PiePointersEncoding.XID.Int64 = 0 		
	}

	pieAPI.PiePointersEncoding.YID.Valid = true
	if (pie.Y != undefined) {
		pieAPI.PiePointersEncoding.YID.Int64 = pie.Y.ID  
	} else {
		pieAPI.PiePointersEncoding.YID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	pieAPI.PiePointersEncoding.Set = []
	for (let _serie of pie.Set) {
		pieAPI.PiePointersEncoding.Set.push(_serie.ID)
	}

}

// CopyPieAPIToPie update basic, pointers and slice of pointers fields of pie
// from respectively the basic fields and encoded fields of pointers and slices of pointers of pieAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPieAPIToPie(pieAPI: PieAPI, pie: Pie, frontRepo: FrontRepo) {

	pie.CreatedAt = pieAPI.CreatedAt
	pie.DeletedAt = pieAPI.DeletedAt
	pie.ID = pieAPI.ID

	// insertion point for basic fields copy operations
	pie.Name = pieAPI.Name
	pie.Width = pieAPI.Width
	pie.Heigth = pieAPI.Heigth
	pie.Margin = pieAPI.Margin

	// insertion point for pointer fields encoding
	pie.X = frontRepo.map_ID_Key.get(pieAPI.PiePointersEncoding.XID.Int64)
	pie.Y = frontRepo.map_ID_Key.get(pieAPI.PiePointersEncoding.YID.Int64)

	// insertion point for slice of pointers fields encoding
	pie.Set = new Array<Serie>()
	for (let _id of pieAPI.PiePointersEncoding.Set) {
		let _serie = frontRepo.map_ID_Serie.get(_id)
		if (_serie != undefined) {
			pie.Set.push(_serie!)
		}
	}
}
