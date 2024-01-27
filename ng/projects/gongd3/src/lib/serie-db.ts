// insertion point for imports
import { KeyDB } from './key-db'
import { ValueDB } from './value-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SerieDB {

	static GONGSTRUCT_NAME = "Serie"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	SeriePointersEncoding: SeriePointersEncoding = new SeriePointersEncoding
}

export class SeriePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	KeyID: NullInt64 = new NullInt64 // if pointer is null, Key.ID = 0

	Values: number[] = []
}
