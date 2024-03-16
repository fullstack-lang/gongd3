// generated code - do not edit

import { ScatterAPI } from './scatter-api'
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

export function CopyScatterToScatterAPI(scatter: Scatter, scatterAPI: ScatterAPI) {

	scatterAPI.CreatedAt = scatter.CreatedAt
	scatterAPI.DeletedAt = scatter.DeletedAt
	scatterAPI.ID = scatter.ID

	// insertion point for basic fields copy operations
	scatterAPI.Name = scatter.Name
	scatterAPI.Width = scatter.Width
	scatterAPI.Heigth = scatter.Heigth
	scatterAPI.Margin = scatter.Margin

	// insertion point for pointer fields encoding
	scatterAPI.ScatterPointersEncoding.XID.Valid = true
	if (scatter.X != undefined) {
		scatterAPI.ScatterPointersEncoding.XID.Int64 = scatter.X.ID  
	} else {
		scatterAPI.ScatterPointersEncoding.XID.Int64 = 0 		
	}

	scatterAPI.ScatterPointersEncoding.YID.Valid = true
	if (scatter.Y != undefined) {
		scatterAPI.ScatterPointersEncoding.YID.Int64 = scatter.Y.ID  
	} else {
		scatterAPI.ScatterPointersEncoding.YID.Int64 = 0 		
	}

	scatterAPI.ScatterPointersEncoding.TextID.Valid = true
	if (scatter.Text != undefined) {
		scatterAPI.ScatterPointersEncoding.TextID.Int64 = scatter.Text.ID  
	} else {
		scatterAPI.ScatterPointersEncoding.TextID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	scatterAPI.ScatterPointersEncoding.Set = []
	for (let _serie of scatter.Set) {
		scatterAPI.ScatterPointersEncoding.Set.push(_serie.ID)
	}

}

// CopyScatterAPIToScatter update basic, pointers and slice of pointers fields of scatter
// from respectively the basic fields and encoded fields of pointers and slices of pointers of scatterAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyScatterAPIToScatter(scatterAPI: ScatterAPI, scatter: Scatter, frontRepo: FrontRepo) {

	scatter.CreatedAt = scatterAPI.CreatedAt
	scatter.DeletedAt = scatterAPI.DeletedAt
	scatter.ID = scatterAPI.ID

	// insertion point for basic fields copy operations
	scatter.Name = scatterAPI.Name
	scatter.Width = scatterAPI.Width
	scatter.Heigth = scatterAPI.Heigth
	scatter.Margin = scatterAPI.Margin

	// insertion point for pointer fields encoding
	scatter.X = frontRepo.map_ID_Key.get(scatterAPI.ScatterPointersEncoding.XID.Int64)
	scatter.Y = frontRepo.map_ID_Key.get(scatterAPI.ScatterPointersEncoding.YID.Int64)
	scatter.Text = frontRepo.map_ID_Key.get(scatterAPI.ScatterPointersEncoding.TextID.Int64)

	// insertion point for slice of pointers fields encoding
	scatter.Set = new Array<Serie>()
	for (let _id of scatterAPI.ScatterPointersEncoding.Set) {
		let _serie = frontRepo.map_ID_Serie.get(_id)
		if (_serie != undefined) {
			scatter.Set.push(_serie!)
		}
	}
}
