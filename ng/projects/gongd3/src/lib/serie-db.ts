// insertion point for imports
import { KeyDB } from './key-db'
import { ValueDB } from './value-db'
import { BarDB } from './bar-db'
import { PieDB } from './pie-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SerieDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Key?: KeyDB
	KeyID: NullInt64 = new NullInt64 // if pointer is null, Key.ID = 0

	Values?: Array<ValueDB>
	Bar_SetDBID: NullInt64 = new NullInt64
	Bar_SetDBID_Index: NullInt64  = new NullInt64 // store the index of the serie instance in Bar.Set
	Bar_Set_reverse?: BarDB 

	Pie_SetDBID: NullInt64 = new NullInt64
	Pie_SetDBID_Index: NullInt64  = new NullInt64 // store the index of the serie instance in Pie.Set
	Pie_Set_reverse?: PieDB 

}
