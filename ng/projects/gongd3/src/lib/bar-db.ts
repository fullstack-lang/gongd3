// insertion point for imports
import { KeyDB } from './key-db'
import { SerieDB } from './serie-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class BarDB {
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

	// insertion point for other declarations
	X?: KeyDB
	XID: NullInt64 = new NullInt64 // if pointer is null, X.ID = 0

	Y?: KeyDB
	YID: NullInt64 = new NullInt64 // if pointer is null, Y.ID = 0

	Set?: Array<SerieDB>
}
