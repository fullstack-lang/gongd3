// generated code - do not edit

import { SerieDB } from './serie-db'
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

export function CopySerieToSerieDB(serie: Serie, serieDB: SerieDB) {

	serieDB.CreatedAt = serie.CreatedAt
	serieDB.DeletedAt = serie.DeletedAt
	serieDB.ID = serie.ID
	
	// insertion point for basic fields copy operations
	serieDB.Name = serie.Name

	// insertion point for pointer fields encoding
    serieDB.SeriePointersEncoding.KeyID.Valid = true
	if (serie.Key != undefined) {
		serieDB.SeriePointersEncoding.KeyID.Int64 = serie.Key.ID  
    } else {
		serieDB.SeriePointersEncoding.KeyID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
	serieDB.SeriePointersEncoding.Values = []
    for (let _value of serie.Values) {
		serieDB.SeriePointersEncoding.Values.push(_value.ID)
    }
	
}

export function CopySerieDBToSerie(serieDB: SerieDB, serie: Serie, frontRepo: FrontRepo) {

	serie.CreatedAt = serieDB.CreatedAt
	serie.DeletedAt = serieDB.DeletedAt
	serie.ID = serieDB.ID
	
	// insertion point for basic fields copy operations
	serie.Name = serieDB.Name

	// insertion point for pointer fields encoding
	serie.Key = frontRepo.Keys.get(serieDB.SeriePointersEncoding.KeyID.Int64)

	// insertion point for slice of pointers fields encoding
	serie.Values = new Array<Value>()
	for (let _id of serieDB.SeriePointersEncoding.Values) {
	  let _value = frontRepo.Values.get(_id)
	  if (_value != undefined) {
		serie.Values.push(_value!)
	  }
	}
}
