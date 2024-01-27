// generated code - do not edit

import { KeyDB } from './key-db'
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

export function CopyKeyToKeyDB(key: Key, keyDB: KeyDB) {

	keyDB.CreatedAt = key.CreatedAt
	keyDB.DeletedAt = key.DeletedAt
	keyDB.ID = key.ID

	// insertion point for basic fields copy operations
	keyDB.Name = key.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyKeyDBToKey update basic, pointers and slice of pointers fields of key
// from respectively the basic fields and encoded fields of pointers and slices of pointers of keyDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyKeyDBToKey(keyDB: KeyDB, key: Key, frontRepo: FrontRepo) {

	key.CreatedAt = keyDB.CreatedAt
	key.DeletedAt = keyDB.DeletedAt
	key.ID = keyDB.ID

	// insertion point for basic fields copy operations
	key.Name = keyDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
