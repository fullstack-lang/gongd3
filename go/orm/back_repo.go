// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongd3/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoBar BackRepoBarStruct

	BackRepoKey BackRepoKeyStruct

	BackRepoPie BackRepoPieStruct

	BackRepoScatter BackRepoScatterStruct

	BackRepoSerie BackRepoSerieStruct

	BackRepoValue BackRepoValueStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&BarDB{},
		&KeyDB{},
		&PieDB{},
		&ScatterDB{},
		&SerieDB{},
		&ValueDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoBar = BackRepoBarStruct{
		Map_BarDBID_BarPtr: make(map[uint]*models.Bar, 0),
		Map_BarDBID_BarDB:  make(map[uint]*BarDB, 0),
		Map_BarPtr_BarDBID: make(map[*models.Bar]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKey = BackRepoKeyStruct{
		Map_KeyDBID_KeyPtr: make(map[uint]*models.Key, 0),
		Map_KeyDBID_KeyDB:  make(map[uint]*KeyDB, 0),
		Map_KeyPtr_KeyDBID: make(map[*models.Key]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPie = BackRepoPieStruct{
		Map_PieDBID_PiePtr: make(map[uint]*models.Pie, 0),
		Map_PieDBID_PieDB:  make(map[uint]*PieDB, 0),
		Map_PiePtr_PieDBID: make(map[*models.Pie]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScatter = BackRepoScatterStruct{
		Map_ScatterDBID_ScatterPtr: make(map[uint]*models.Scatter, 0),
		Map_ScatterDBID_ScatterDB:  make(map[uint]*ScatterDB, 0),
		Map_ScatterPtr_ScatterDBID: make(map[*models.Scatter]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSerie = BackRepoSerieStruct{
		Map_SerieDBID_SeriePtr: make(map[uint]*models.Serie, 0),
		Map_SerieDBID_SerieDB:  make(map[uint]*SerieDB, 0),
		Map_SeriePtr_SerieDBID: make(map[*models.Serie]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoValue = BackRepoValueStruct{
		Map_ValueDBID_ValuePtr: make(map[uint]*models.Value, 0),
		Map_ValueDBID_ValueDB:  make(map[uint]*ValueDB, 0),
		Map_ValuePtr_ValueDBID: make(map[*models.Value]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()
	
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoBar.CommitPhaseOne(stage)
	backRepo.BackRepoKey.CommitPhaseOne(stage)
	backRepo.BackRepoPie.CommitPhaseOne(stage)
	backRepo.BackRepoScatter.CommitPhaseOne(stage)
	backRepo.BackRepoSerie.CommitPhaseOne(stage)
	backRepo.BackRepoValue.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoBar.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKey.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPie.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScatter.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSerie.CommitPhaseTwo(backRepo)
	backRepo.BackRepoValue.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoBar.CheckoutPhaseOne()
	backRepo.BackRepoKey.CheckoutPhaseOne()
	backRepo.BackRepoPie.CheckoutPhaseOne()
	backRepo.BackRepoScatter.CheckoutPhaseOne()
	backRepo.BackRepoSerie.CheckoutPhaseOne()
	backRepo.BackRepoValue.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoBar.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKey.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPie.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScatter.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSerie.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoValue.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoBar.Backup(dirPath)
	backRepo.BackRepoKey.Backup(dirPath)
	backRepo.BackRepoPie.Backup(dirPath)
	backRepo.BackRepoScatter.Backup(dirPath)
	backRepo.BackRepoSerie.Backup(dirPath)
	backRepo.BackRepoValue.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoBar.BackupXL(file)
	backRepo.BackRepoKey.BackupXL(file)
	backRepo.BackRepoPie.BackupXL(file)
	backRepo.BackRepoScatter.BackupXL(file)
	backRepo.BackRepoSerie.BackupXL(file)
	backRepo.BackRepoValue.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoBar.RestorePhaseOne(dirPath)
	backRepo.BackRepoKey.RestorePhaseOne(dirPath)
	backRepo.BackRepoPie.RestorePhaseOne(dirPath)
	backRepo.BackRepoScatter.RestorePhaseOne(dirPath)
	backRepo.BackRepoSerie.RestorePhaseOne(dirPath)
	backRepo.BackRepoValue.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoBar.RestorePhaseTwo()
	backRepo.BackRepoKey.RestorePhaseTwo()
	backRepo.BackRepoPie.RestorePhaseTwo()
	backRepo.BackRepoScatter.RestorePhaseTwo()
	backRepo.BackRepoSerie.RestorePhaseTwo()
	backRepo.BackRepoValue.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoBar.RestoreXLPhaseOne(file)
	backRepo.BackRepoKey.RestoreXLPhaseOne(file)
	backRepo.BackRepoPie.RestoreXLPhaseOne(file)
	backRepo.BackRepoScatter.RestoreXLPhaseOne(file)
	backRepo.BackRepoSerie.RestoreXLPhaseOne(file)
	backRepo.BackRepoValue.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	activeChannels := make([]chan int, 0)

	for _, ch := range backRepoStruct.subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			activeChannels = append(activeChannels, ch)
		default:
			// Assume channel is no longer active; don't add to activeChannels
			log.Println("Channel no longer active")
			close(ch)
		}
	}
	backRepoStruct.subscribers = activeChannels
}
