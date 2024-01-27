// generated code - do not edit

import { ValueDB } from './value-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Value {

	static GONGSTRUCT_NAME = "Value"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyValueToValueDB(value: Value, valueDB: ValueDB) {

	valueDB.CreatedAt = value.CreatedAt
	valueDB.DeletedAt = value.DeletedAt
	valueDB.ID = value.ID

	// insertion point for basic fields copy operations
	valueDB.Name = value.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyValueDBToValue update basic, pointers and slice of pointers fields of value
// from respectively the basic fields and encoded fields of pointers and slices of pointers of valueDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyValueDBToValue(valueDB: ValueDB, value: Value, frontRepo: FrontRepo) {

	value.CreatedAt = valueDB.CreatedAt
	value.DeletedAt = valueDB.DeletedAt
	value.ID = valueDB.ID

	// insertion point for basic fields copy operations
	value.Name = valueDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
