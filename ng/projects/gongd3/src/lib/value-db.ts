// insertion point for imports
import { SerieDB } from './serie-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ValueDB {

	static GONGSTRUCT_NAME = "Value"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Serie_ValuesDBID: NullInt64 = new NullInt64
	Serie_ValuesDBID_Index: NullInt64  = new NullInt64 // store the index of the value instance in Serie.Values
	Serie_Values_reverse?: SerieDB 

}
