// insertion point for imports
import { KeyDB } from './key-db'
import { SerieDB } from './serie-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ScatterDB {

	static GONGSTRUCT_NAME = "Scatter"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Heigth: number = 0
	Margin: number = 0

	// insertion point for other decls

	ScatterPointersEncoding: ScatterPointersEncoding = new ScatterPointersEncoding
}

export class ScatterPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	XID: NullInt64 = new NullInt64 // if pointer is null, X.ID = 0

	YID: NullInt64 = new NullInt64 // if pointer is null, Y.ID = 0

	TextID: NullInt64 = new NullInt64 // if pointer is null, Text.ID = 0

	Set: number[] = []
}
