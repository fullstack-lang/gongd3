// generated code - do not edit

import { ScatterDB } from './scatter-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Key } from './key'
import { Serie } from './serie'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Scatter {

	static GONGSTRUCT_NAME = "Scatter"

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

	Text?: Key

	Set: Array<Serie> = []
}

export function CopyScatterToScatterDB(scatter: Scatter, scatterDB: ScatterDB) {

	scatterDB.CreatedAt = scatter.CreatedAt
	scatterDB.DeletedAt = scatter.DeletedAt
	scatterDB.ID = scatter.ID

	// insertion point for basic fields copy operations
	scatterDB.Name = scatter.Name
	scatterDB.Width = scatter.Width
	scatterDB.Heigth = scatter.Heigth
	scatterDB.Margin = scatter.Margin

	// insertion point for pointer fields encoding
	scatterDB.ScatterPointersEncoding.XID.Valid = true
	if (scatter.X != undefined) {
		scatterDB.ScatterPointersEncoding.XID.Int64 = scatter.X.ID  
	} else {
		scatterDB.ScatterPointersEncoding.XID.Int64 = 0 		
	}

	scatterDB.ScatterPointersEncoding.YID.Valid = true
	if (scatter.Y != undefined) {
		scatterDB.ScatterPointersEncoding.YID.Int64 = scatter.Y.ID  
	} else {
		scatterDB.ScatterPointersEncoding.YID.Int64 = 0 		
	}

	scatterDB.ScatterPointersEncoding.TextID.Valid = true
	if (scatter.Text != undefined) {
		scatterDB.ScatterPointersEncoding.TextID.Int64 = scatter.Text.ID  
	} else {
		scatterDB.ScatterPointersEncoding.TextID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	scatterDB.ScatterPointersEncoding.Set = []
	for (let _serie of scatter.Set) {
		scatterDB.ScatterPointersEncoding.Set.push(_serie.ID)
	}

}

// CopyScatterDBToScatter update basic, pointers and slice of pointers fields of scatter
// from respectively the basic fields and encoded fields of pointers and slices of pointers of scatterDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyScatterDBToScatter(scatterDB: ScatterDB, scatter: Scatter, frontRepo: FrontRepo) {

	scatter.CreatedAt = scatterDB.CreatedAt
	scatter.DeletedAt = scatterDB.DeletedAt
	scatter.ID = scatterDB.ID

	// insertion point for basic fields copy operations
	scatter.Name = scatterDB.Name
	scatter.Width = scatterDB.Width
	scatter.Heigth = scatterDB.Heigth
	scatter.Margin = scatterDB.Margin

	// insertion point for pointer fields encoding
	scatter.X = frontRepo.map_ID_Key.get(scatterDB.ScatterPointersEncoding.XID.Int64)
	scatter.Y = frontRepo.map_ID_Key.get(scatterDB.ScatterPointersEncoding.YID.Int64)
	scatter.Text = frontRepo.map_ID_Key.get(scatterDB.ScatterPointersEncoding.TextID.Int64)

	// insertion point for slice of pointers fields encoding
	scatter.Set = new Array<Serie>()
	for (let _id of scatterDB.ScatterPointersEncoding.Set) {
		let _serie = frontRepo.map_ID_Serie.get(_id)
		if (_serie != undefined) {
			scatter.Set.push(_serie!)
		}
	}
}
