// generated code - do not edit

import { KeyAPI } from './key-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Key {

	static GONGSTRUCT_NAME = "Key"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyKeyToKeyAPI(key: Key, keyAPI: KeyAPI) {

	keyAPI.CreatedAt = key.CreatedAt
	keyAPI.DeletedAt = key.DeletedAt
	keyAPI.ID = key.ID

	// insertion point for basic fields copy operations
	keyAPI.Name = key.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyKeyAPIToKey update basic, pointers and slice of pointers fields of key
// from respectively the basic fields and encoded fields of pointers and slices of pointers of keyAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyKeyAPIToKey(keyAPI: KeyAPI, key: Key, frontRepo: FrontRepo) {

	key.CreatedAt = keyAPI.CreatedAt
	key.DeletedAt = keyAPI.DeletedAt
	key.ID = keyAPI.ID

	// insertion point for basic fields copy operations
	key.Name = keyAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
