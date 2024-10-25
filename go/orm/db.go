// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongd3/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	barDBs map[uint]*BarDB

	nextIDBarDB uint

	keyDBs map[uint]*KeyDB

	nextIDKeyDB uint

	pieDBs map[uint]*PieDB

	nextIDPieDB uint

	scatterDBs map[uint]*ScatterDB

	nextIDScatterDB uint

	serieDBs map[uint]*SerieDB

	nextIDSerieDB uint

	valueDBs map[uint]*ValueDB

	nextIDValueDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		barDBs: make(map[uint]*BarDB),

		keyDBs: make(map[uint]*KeyDB),

		pieDBs: make(map[uint]*PieDB),

		scatterDBs: make(map[uint]*ScatterDB),

		serieDBs: make(map[uint]*SerieDB),

		valueDBs: make(map[uint]*ValueDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *BarDB:
		db.nextIDBarDB++
		v.ID = db.nextIDBarDB
		db.barDBs[v.ID] = v
	case *KeyDB:
		db.nextIDKeyDB++
		v.ID = db.nextIDKeyDB
		db.keyDBs[v.ID] = v
	case *PieDB:
		db.nextIDPieDB++
		v.ID = db.nextIDPieDB
		db.pieDBs[v.ID] = v
	case *ScatterDB:
		db.nextIDScatterDB++
		v.ID = db.nextIDScatterDB
		db.scatterDBs[v.ID] = v
	case *SerieDB:
		db.nextIDSerieDB++
		v.ID = db.nextIDSerieDB
		db.serieDBs[v.ID] = v
	case *ValueDB:
		db.nextIDValueDB++
		v.ID = db.nextIDValueDB
		db.valueDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BarDB:
		delete(db.barDBs, v.ID)
	case *KeyDB:
		delete(db.keyDBs, v.ID)
	case *PieDB:
		delete(db.pieDBs, v.ID)
	case *ScatterDB:
		delete(db.scatterDBs, v.ID)
	case *SerieDB:
		delete(db.serieDBs, v.ID)
	case *ValueDB:
		delete(db.valueDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BarDB:
		db.barDBs[v.ID] = v
		return db, nil
	case *KeyDB:
		db.keyDBs[v.ID] = v
		return db, nil
	case *PieDB:
		db.pieDBs[v.ID] = v
		return db, nil
	case *ScatterDB:
		db.scatterDBs[v.ID] = v
		return db, nil
	case *SerieDB:
		db.serieDBs[v.ID] = v
		return db, nil
	case *ValueDB:
		db.valueDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BarDB:
		if existing, ok := db.barDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongd3/go, record not found")
		}
	case *KeyDB:
		if existing, ok := db.keyDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongd3/go, record not found")
		}
	case *PieDB:
		if existing, ok := db.pieDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongd3/go, record not found")
		}
	case *ScatterDB:
		if existing, ok := db.scatterDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongd3/go, record not found")
		}
	case *SerieDB:
		if existing, ok := db.serieDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongd3/go, record not found")
		}
	case *ValueDB:
		if existing, ok := db.valueDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("github.com/fullstack-lang/gongd3/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]BarDB:
        *ptr = make([]BarDB, 0, len(db.barDBs))
        for _, v := range db.barDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]KeyDB:
        *ptr = make([]KeyDB, 0, len(db.keyDBs))
        for _, v := range db.keyDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]PieDB:
        *ptr = make([]PieDB, 0, len(db.pieDBs))
        for _, v := range db.pieDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ScatterDB:
        *ptr = make([]ScatterDB, 0, len(db.scatterDBs))
        for _, v := range db.scatterDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]SerieDB:
        *ptr = make([]SerieDB, 0, len(db.serieDBs))
        for _, v := range db.serieDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ValueDB:
        *ptr = make([]ValueDB, 0, len(db.valueDBs))
        for _, v := range db.valueDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
    default:
        return nil, errors.New("github.com/fullstack-lang/gongd3/go, Find: unsupported type")
    }
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *BarDB:
		tmp, ok := db.barDBs[uint(i)]

		barDB, _ := instanceDB.(*BarDB)
		*barDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *KeyDB:
		tmp, ok := db.keyDBs[uint(i)]

		keyDB, _ := instanceDB.(*KeyDB)
		*keyDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *PieDB:
		tmp, ok := db.pieDBs[uint(i)]

		pieDB, _ := instanceDB.(*PieDB)
		*pieDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ScatterDB:
		tmp, ok := db.scatterDBs[uint(i)]

		scatterDB, _ := instanceDB.(*ScatterDB)
		*scatterDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *SerieDB:
		tmp, ok := db.serieDBs[uint(i)]

		serieDB, _ := instanceDB.(*SerieDB)
		*serieDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ValueDB:
		tmp, ok := db.valueDBs[uint(i)]

		valueDB, _ := instanceDB.(*ValueDB)
		*valueDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongd3/go, Unkown type")
	}
	
	return db, nil
}

