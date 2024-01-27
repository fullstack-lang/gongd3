// generated code - do not edit

import { PieDB } from './pie-db'
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

export function CopyPieToPieDB(pie: Pie, pieDB: PieDB) {

	pieDB.CreatedAt = pie.CreatedAt
	pieDB.DeletedAt = pie.DeletedAt
	pieDB.ID = pie.ID

	// insertion point for basic fields copy operations
	pieDB.Name = pie.Name
	pieDB.Width = pie.Width
	pieDB.Heigth = pie.Heigth
	pieDB.Margin = pie.Margin

	// insertion point for pointer fields encoding
	pieDB.PiePointersEncoding.XID.Valid = true
	if (pie.X != undefined) {
		pieDB.PiePointersEncoding.XID.Int64 = pie.X.ID  
	} else {
		pieDB.PiePointersEncoding.XID.Int64 = 0 		
	}

	pieDB.PiePointersEncoding.YID.Valid = true
	if (pie.Y != undefined) {
		pieDB.PiePointersEncoding.YID.Int64 = pie.Y.ID  
	} else {
		pieDB.PiePointersEncoding.YID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	pieDB.PiePointersEncoding.Set = []
	for (let _serie of pie.Set) {
		pieDB.PiePointersEncoding.Set.push(_serie.ID)
	}

}

// CopyPieDBToPie update basic, pointers and slice of pointers fields of pie
// from respectively the basic fields and encoded fields of pointers and slices of pointers of pieDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPieDBToPie(pieDB: PieDB, pie: Pie, frontRepo: FrontRepo) {

	pie.CreatedAt = pieDB.CreatedAt
	pie.DeletedAt = pieDB.DeletedAt
	pie.ID = pieDB.ID

	// insertion point for basic fields copy operations
	pie.Name = pieDB.Name
	pie.Width = pieDB.Width
	pie.Heigth = pieDB.Heigth
	pie.Margin = pieDB.Margin

	// insertion point for pointer fields encoding
	pie.X = frontRepo.map_ID_Key.get(pieDB.PiePointersEncoding.XID.Int64)
	pie.Y = frontRepo.map_ID_Key.get(pieDB.PiePointersEncoding.YID.Int64)

	// insertion point for slice of pointers fields encoding
	pie.Set = new Array<Serie>()
	for (let _id of pieDB.PiePointersEncoding.Set) {
		let _serie = frontRepo.map_ID_Serie.get(_id)
		if (_serie != undefined) {
			pie.Set.push(_serie!)
		}
	}
}
