// insertion point for imports
import { KeyAPI } from './key-api'
import { SerieAPI } from './serie-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BarAPI {

	static GONGSTRUCT_NAME = "Bar"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	AutoDomainX: boolean = false
	XMin: number = 0
	XMax: number = 0
	AutoDomainY: boolean = false
	YMin: number = 0
	YMax: number = 0
	YLabelPresent: boolean = false
	YLabelOffset: number = 0
	XLabelPresent: boolean = false
	XLabelOffset: number = 0
	Width: number = 0
	Heigth: number = 0
	Margin: number = 0

	// insertion point for other decls

	BarPointersEncoding: BarPointersEncoding = new BarPointersEncoding
}

export class BarPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	XID: NullInt64 = new NullInt64 // if pointer is null, X.ID = 0

	YID: NullInt64 = new NullInt64 // if pointer is null, Y.ID = 0

	Set: number[] = []
}
