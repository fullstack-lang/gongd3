// generated code - do not edit

import { SerieAPI } from './serie-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Key } from './key'
import { Value } from './value'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Serie {

	static GONGSTRUCT_NAME = "Serie"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Key?: Key

	Values: Array<Value> = []
}

export function CopySerieToSerieAPI(serie: Serie, serieAPI: SerieAPI) {

	serieAPI.CreatedAt = serie.CreatedAt
	serieAPI.DeletedAt = serie.DeletedAt
	serieAPI.ID = serie.ID

	// insertion point for basic fields copy operations
	serieAPI.Name = serie.Name

	// insertion point for pointer fields encoding
	serieAPI.SeriePointersEncoding.KeyID.Valid = true
	if (serie.Key != undefined) {
		serieAPI.SeriePointersEncoding.KeyID.Int64 = serie.Key.ID  
	} else {
		serieAPI.SeriePointersEncoding.KeyID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	serieAPI.SeriePointersEncoding.Values = []
	for (let _value of serie.Values) {
		serieAPI.SeriePointersEncoding.Values.push(_value.ID)
	}

}

// CopySerieAPIToSerie update basic, pointers and slice of pointers fields of serie
// from respectively the basic fields and encoded fields of pointers and slices of pointers of serieAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopySerieAPIToSerie(serieAPI: SerieAPI, serie: Serie, frontRepo: FrontRepo) {

	serie.CreatedAt = serieAPI.CreatedAt
	serie.DeletedAt = serieAPI.DeletedAt
	serie.ID = serieAPI.ID

	// insertion point for basic fields copy operations
	serie.Name = serieAPI.Name

	// insertion point for pointer fields encoding
	serie.Key = frontRepo.map_ID_Key.get(serieAPI.SeriePointersEncoding.KeyID.Int64)

	// insertion point for slice of pointers fields encoding
	serie.Values = new Array<Value>()
	for (let _id of serieAPI.SeriePointersEncoding.Values) {
		let _value = frontRepo.map_ID_Value.get(_id)
		if (_value != undefined) {
			serie.Values.push(_value!)
		}
	}
}
