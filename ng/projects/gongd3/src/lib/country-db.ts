// insertion point for imports
import { HelloDB } from './hello-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CountryDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Hello?: HelloDB
	HelloID: NullInt64 = new NullInt64 // if pointer is null, Hello.ID = 0

}
