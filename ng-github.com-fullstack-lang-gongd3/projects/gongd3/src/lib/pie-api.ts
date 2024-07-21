// insertion point for imports
import { KeyAPI } from './key-api'
import { SerieAPI } from './serie-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PieAPI {

	static GONGSTRUCT_NAME = "Pie"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Width: number = 0
	Heigth: number = 0
	Margin: number = 0

	// insertion point for other decls

	PiePointersEncoding: PiePointersEncoding = new PiePointersEncoding
}

export class PiePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	XID: NullInt64 = new NullInt64 // if pointer is null, X.ID = 0

	YID: NullInt64 = new NullInt64 // if pointer is null, Y.ID = 0

	Set: number[] = []
}
