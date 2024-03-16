// generated code - do not edit

import { ValueAPI } from './value-api'
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

export function CopyValueToValueAPI(value: Value, valueAPI: ValueAPI) {

	valueAPI.CreatedAt = value.CreatedAt
	valueAPI.DeletedAt = value.DeletedAt
	valueAPI.ID = value.ID

	// insertion point for basic fields copy operations
	valueAPI.Name = value.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyValueAPIToValue update basic, pointers and slice of pointers fields of value
// from respectively the basic fields and encoded fields of pointers and slices of pointers of valueAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyValueAPIToValue(valueAPI: ValueAPI, value: Value, frontRepo: FrontRepo) {

	value.CreatedAt = valueAPI.CreatedAt
	value.DeletedAt = valueAPI.DeletedAt
	value.ID = valueAPI.ID

	// insertion point for basic fields copy operations
	value.Name = valueAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
