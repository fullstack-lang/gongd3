// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongsvg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Animate_sql sql.NullBool
var dummy_Animate_time time.Duration
var dummy_Animate_sort sort.Float64Slice

// AnimateAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model animateAPI
type AnimateAPI struct {
	gorm.Model

	models.Animate

	// encoding of pointers
	AnimatePointersEnconding
}

// AnimatePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type AnimatePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Circle{}.Animations []*Animate
	Circle_AnimationsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Circle_AnimationsDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Ellipse{}.Animates []*Animate
	Ellipse_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Ellipse_AnimatesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Line{}.Animates []*Animate
	Line_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Line_AnimatesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field LinkAnchoredText{}.Animates []*Animate
	LinkAnchoredText_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	LinkAnchoredText_AnimatesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Path{}.Animates []*Animate
	Path_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Path_AnimatesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Polygone{}.Animates []*Animate
	Polygone_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Polygone_AnimatesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Polyline{}.Animates []*Animate
	Polyline_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Polyline_AnimatesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Rect{}.Animations []*Animate
	Rect_AnimationsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Rect_AnimationsDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field RectAnchoredText{}.Animates []*Animate
	RectAnchoredText_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	RectAnchoredText_AnimatesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Text{}.Animates []*Animate
	Text_AnimatesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Text_AnimatesDBID_Index sql.NullInt64
}

// AnimateDB describes a animate in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model animateDB
type AnimateDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field animateDB.Name
	Name_Data sql.NullString

	// Declation for basic field animateDB.AttributeName
	AttributeName_Data sql.NullString

	// Declation for basic field animateDB.Values
	Values_Data sql.NullString

	// Declation for basic field animateDB.Dur
	Dur_Data sql.NullString

	// Declation for basic field animateDB.RepeatCount
	RepeatCount_Data sql.NullString
	// encoding of pointers
	AnimatePointersEnconding
}

// AnimateDBs arrays animateDBs
// swagger:response animateDBsResponse
type AnimateDBs []AnimateDB

// AnimateDBResponse provides response
// swagger:response animateDBResponse
type AnimateDBResponse struct {
	AnimateDB
}

// AnimateWOP is a Animate without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type AnimateWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	AttributeName string `xlsx:"2"`

	Values string `xlsx:"3"`

	Dur string `xlsx:"4"`

	RepeatCount string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var Animate_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"AttributeName",
	"Values",
	"Dur",
	"RepeatCount",
}

type BackRepoAnimateStruct struct {
	// stores AnimateDB according to their gorm ID
	Map_AnimateDBID_AnimateDB map[uint]*AnimateDB

	// stores AnimateDB ID according to Animate address
	Map_AnimatePtr_AnimateDBID map[*models.Animate]uint

	// stores Animate according to their gorm ID
	Map_AnimateDBID_AnimatePtr map[uint]*models.Animate

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoAnimate *BackRepoAnimateStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoAnimate.stage
	return
}

func (backRepoAnimate *BackRepoAnimateStruct) GetDB() *gorm.DB {
	return backRepoAnimate.db
}

// GetAnimateDBFromAnimatePtr is a handy function to access the back repo instance from the stage instance
func (backRepoAnimate *BackRepoAnimateStruct) GetAnimateDBFromAnimatePtr(animate *models.Animate) (animateDB *AnimateDB) {
	id := backRepoAnimate.Map_AnimatePtr_AnimateDBID[animate]
	animateDB = backRepoAnimate.Map_AnimateDBID_AnimateDB[id]
	return
}

// BackRepoAnimate.CommitPhaseOne commits all staged instances of Animate to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAnimate *BackRepoAnimateStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for animate := range stage.Animates {
		backRepoAnimate.CommitPhaseOneInstance(animate)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, animate := range backRepoAnimate.Map_AnimateDBID_AnimatePtr {
		if _, ok := stage.Animates[animate]; !ok {
			backRepoAnimate.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoAnimate.CommitDeleteInstance commits deletion of Animate to the BackRepo
func (backRepoAnimate *BackRepoAnimateStruct) CommitDeleteInstance(id uint) (Error error) {

	animate := backRepoAnimate.Map_AnimateDBID_AnimatePtr[id]

	// animate is not staged anymore, remove animateDB
	animateDB := backRepoAnimate.Map_AnimateDBID_AnimateDB[id]
	query := backRepoAnimate.db.Unscoped().Delete(&animateDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoAnimate.Map_AnimatePtr_AnimateDBID, animate)
	delete(backRepoAnimate.Map_AnimateDBID_AnimatePtr, id)
	delete(backRepoAnimate.Map_AnimateDBID_AnimateDB, id)

	return
}

// BackRepoAnimate.CommitPhaseOneInstance commits animate staged instances of Animate to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoAnimate *BackRepoAnimateStruct) CommitPhaseOneInstance(animate *models.Animate) (Error error) {

	// check if the animate is not commited yet
	if _, ok := backRepoAnimate.Map_AnimatePtr_AnimateDBID[animate]; ok {
		return
	}

	// initiate animate
	var animateDB AnimateDB
	animateDB.CopyBasicFieldsFromAnimate(animate)

	query := backRepoAnimate.db.Create(&animateDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoAnimate.Map_AnimatePtr_AnimateDBID[animate] = animateDB.ID
	backRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID] = animate
	backRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB.ID] = &animateDB

	return
}

// BackRepoAnimate.CommitPhaseTwo commits all staged instances of Animate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnimate *BackRepoAnimateStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, animate := range backRepoAnimate.Map_AnimateDBID_AnimatePtr {
		backRepoAnimate.CommitPhaseTwoInstance(backRepo, idx, animate)
	}

	return
}

// BackRepoAnimate.CommitPhaseTwoInstance commits {{structname }} of models.Animate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnimate *BackRepoAnimateStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, animate *models.Animate) (Error error) {

	// fetch matching animateDB
	if animateDB, ok := backRepoAnimate.Map_AnimateDBID_AnimateDB[idx]; ok {

		animateDB.CopyBasicFieldsFromAnimate(animate)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoAnimate.db.Save(&animateDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Animate intance %s", animate.Name))
		return err
	}

	return
}

// BackRepoAnimate.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoAnimate *BackRepoAnimateStruct) CheckoutPhaseOne() (Error error) {

	animateDBArray := make([]AnimateDB, 0)
	query := backRepoAnimate.db.Find(&animateDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	animateInstancesToBeRemovedFromTheStage := make(map[*models.Animate]any)
	for key, value := range backRepoAnimate.stage.Animates {
		animateInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, animateDB := range animateDBArray {
		backRepoAnimate.CheckoutPhaseOneInstance(&animateDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		animate, ok := backRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]
		if ok {
			delete(animateInstancesToBeRemovedFromTheStage, animate)
		}
	}

	// remove from stage and back repo's 3 maps all animates that are not in the checkout
	for animate := range animateInstancesToBeRemovedFromTheStage {
		animate.Unstage(backRepoAnimate.GetStage())

		// remove instance from the back repo 3 maps
		animateID := backRepoAnimate.Map_AnimatePtr_AnimateDBID[animate]
		delete(backRepoAnimate.Map_AnimatePtr_AnimateDBID, animate)
		delete(backRepoAnimate.Map_AnimateDBID_AnimateDB, animateID)
		delete(backRepoAnimate.Map_AnimateDBID_AnimatePtr, animateID)
	}

	return
}

// CheckoutPhaseOneInstance takes a animateDB that has been found in the DB, updates the backRepo and stages the
// models version of the animateDB
func (backRepoAnimate *BackRepoAnimateStruct) CheckoutPhaseOneInstance(animateDB *AnimateDB) (Error error) {

	animate, ok := backRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]
	if !ok {
		animate = new(models.Animate)

		backRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID] = animate
		backRepoAnimate.Map_AnimatePtr_AnimateDBID[animate] = animateDB.ID

		// append model store with the new element
		animate.Name = animateDB.Name_Data.String
		animate.Stage(backRepoAnimate.GetStage())
	}
	animateDB.CopyBasicFieldsToAnimate(animate)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	animate.Stage(backRepoAnimate.GetStage())

	// preserve pointer to animateDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_AnimateDBID_AnimateDB)[animateDB hold variable pointers
	animateDB_Data := *animateDB
	preservedPtrToAnimate := &animateDB_Data
	backRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB.ID] = preservedPtrToAnimate

	return
}

// BackRepoAnimate.CheckoutPhaseTwo Checkouts all staged instances of Animate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnimate *BackRepoAnimateStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, animateDB := range backRepoAnimate.Map_AnimateDBID_AnimateDB {
		backRepoAnimate.CheckoutPhaseTwoInstance(backRepo, animateDB)
	}
	return
}

// BackRepoAnimate.CheckoutPhaseTwoInstance Checkouts staged instances of Animate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnimate *BackRepoAnimateStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, animateDB *AnimateDB) (Error error) {

	animate := backRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB.ID]
	_ = animate // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitAnimate allows commit of a single animate (if already staged)
func (backRepo *BackRepoStruct) CommitAnimate(animate *models.Animate) {
	backRepo.BackRepoAnimate.CommitPhaseOneInstance(animate)
	if id, ok := backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID[animate]; ok {
		backRepo.BackRepoAnimate.CommitPhaseTwoInstance(backRepo, id, animate)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitAnimate allows checkout of a single animate (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutAnimate(animate *models.Animate) {
	// check if the animate is staged
	if _, ok := backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID[animate]; ok {

		if id, ok := backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID[animate]; ok {
			var animateDB AnimateDB
			animateDB.ID = id

			if err := backRepo.BackRepoAnimate.db.First(&animateDB, id).Error; err != nil {
				log.Panicln("CheckoutAnimate : Problem with getting object with id:", id)
			}
			backRepo.BackRepoAnimate.CheckoutPhaseOneInstance(&animateDB)
			backRepo.BackRepoAnimate.CheckoutPhaseTwoInstance(backRepo, &animateDB)
		}
	}
}

// CopyBasicFieldsFromAnimate
func (animateDB *AnimateDB) CopyBasicFieldsFromAnimate(animate *models.Animate) {
	// insertion point for fields commit

	animateDB.Name_Data.String = animate.Name
	animateDB.Name_Data.Valid = true

	animateDB.AttributeName_Data.String = animate.AttributeName
	animateDB.AttributeName_Data.Valid = true

	animateDB.Values_Data.String = animate.Values
	animateDB.Values_Data.Valid = true

	animateDB.Dur_Data.String = animate.Dur
	animateDB.Dur_Data.Valid = true

	animateDB.RepeatCount_Data.String = animate.RepeatCount
	animateDB.RepeatCount_Data.Valid = true
}

// CopyBasicFieldsFromAnimateWOP
func (animateDB *AnimateDB) CopyBasicFieldsFromAnimateWOP(animate *AnimateWOP) {
	// insertion point for fields commit

	animateDB.Name_Data.String = animate.Name
	animateDB.Name_Data.Valid = true

	animateDB.AttributeName_Data.String = animate.AttributeName
	animateDB.AttributeName_Data.Valid = true

	animateDB.Values_Data.String = animate.Values
	animateDB.Values_Data.Valid = true

	animateDB.Dur_Data.String = animate.Dur
	animateDB.Dur_Data.Valid = true

	animateDB.RepeatCount_Data.String = animate.RepeatCount
	animateDB.RepeatCount_Data.Valid = true
}

// CopyBasicFieldsToAnimate
func (animateDB *AnimateDB) CopyBasicFieldsToAnimate(animate *models.Animate) {
	// insertion point for checkout of basic fields (back repo to stage)
	animate.Name = animateDB.Name_Data.String
	animate.AttributeName = animateDB.AttributeName_Data.String
	animate.Values = animateDB.Values_Data.String
	animate.Dur = animateDB.Dur_Data.String
	animate.RepeatCount = animateDB.RepeatCount_Data.String
}

// CopyBasicFieldsToAnimateWOP
func (animateDB *AnimateDB) CopyBasicFieldsToAnimateWOP(animate *AnimateWOP) {
	animate.ID = int(animateDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	animate.Name = animateDB.Name_Data.String
	animate.AttributeName = animateDB.AttributeName_Data.String
	animate.Values = animateDB.Values_Data.String
	animate.Dur = animateDB.Dur_Data.String
	animate.RepeatCount = animateDB.RepeatCount_Data.String
}

// Backup generates a json file from a slice of all AnimateDB instances in the backrepo
func (backRepoAnimate *BackRepoAnimateStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "AnimateDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AnimateDB, 0)
	for _, animateDB := range backRepoAnimate.Map_AnimateDBID_AnimateDB {
		forBackup = append(forBackup, animateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Animate ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Animate file", err.Error())
	}
}

// Backup generates a json file from a slice of all AnimateDB instances in the backrepo
func (backRepoAnimate *BackRepoAnimateStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*AnimateDB, 0)
	for _, animateDB := range backRepoAnimate.Map_AnimateDBID_AnimateDB {
		forBackup = append(forBackup, animateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Animate")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Animate_Fields, -1)
	for _, animateDB := range forBackup {

		var animateWOP AnimateWOP
		animateDB.CopyBasicFieldsToAnimateWOP(&animateWOP)

		row := sh.AddRow()
		row.WriteStruct(&animateWOP, -1)
	}
}

// RestoreXL from the "Animate" sheet all AnimateDB instances
func (backRepoAnimate *BackRepoAnimateStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoAnimateid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Animate"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoAnimate.rowVisitorAnimate)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoAnimate *BackRepoAnimateStruct) rowVisitorAnimate(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var animateWOP AnimateWOP
		row.ReadStruct(&animateWOP)

		// add the unmarshalled struct to the stage
		animateDB := new(AnimateDB)
		animateDB.CopyBasicFieldsFromAnimateWOP(&animateWOP)

		animateDB_ID_atBackupTime := animateDB.ID
		animateDB.ID = 0
		query := backRepoAnimate.db.Create(animateDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB.ID] = animateDB
		BackRepoAnimateid_atBckpTime_newID[animateDB_ID_atBackupTime] = animateDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "AnimateDB.json" in dirPath that stores an array
// of AnimateDB and stores it in the database
// the map BackRepoAnimateid_atBckpTime_newID is updated accordingly
func (backRepoAnimate *BackRepoAnimateStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoAnimateid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "AnimateDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Animate file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*AnimateDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_AnimateDBID_AnimateDB
	for _, animateDB := range forRestore {

		animateDB_ID_atBackupTime := animateDB.ID
		animateDB.ID = 0
		query := backRepoAnimate.db.Create(animateDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB.ID] = animateDB
		BackRepoAnimateid_atBckpTime_newID[animateDB_ID_atBackupTime] = animateDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Animate file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Animate>id_atBckpTime_newID
// to compute new index
func (backRepoAnimate *BackRepoAnimateStruct) RestorePhaseTwo() {

	for _, animateDB := range backRepoAnimate.Map_AnimateDBID_AnimateDB {

		// next line of code is to avert unused variable compilation error
		_ = animateDB

		// insertion point for reindexing pointers encoding
		// This reindex animate.Animations
		if animateDB.Circle_AnimationsDBID.Int64 != 0 {
			animateDB.Circle_AnimationsDBID.Int64 =
				int64(BackRepoCircleid_atBckpTime_newID[uint(animateDB.Circle_AnimationsDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.Ellipse_AnimatesDBID.Int64 != 0 {
			animateDB.Ellipse_AnimatesDBID.Int64 =
				int64(BackRepoEllipseid_atBckpTime_newID[uint(animateDB.Ellipse_AnimatesDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.Line_AnimatesDBID.Int64 != 0 {
			animateDB.Line_AnimatesDBID.Int64 =
				int64(BackRepoLineid_atBckpTime_newID[uint(animateDB.Line_AnimatesDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.LinkAnchoredText_AnimatesDBID.Int64 != 0 {
			animateDB.LinkAnchoredText_AnimatesDBID.Int64 =
				int64(BackRepoLinkAnchoredTextid_atBckpTime_newID[uint(animateDB.LinkAnchoredText_AnimatesDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.Path_AnimatesDBID.Int64 != 0 {
			animateDB.Path_AnimatesDBID.Int64 =
				int64(BackRepoPathid_atBckpTime_newID[uint(animateDB.Path_AnimatesDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.Polygone_AnimatesDBID.Int64 != 0 {
			animateDB.Polygone_AnimatesDBID.Int64 =
				int64(BackRepoPolygoneid_atBckpTime_newID[uint(animateDB.Polygone_AnimatesDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.Polyline_AnimatesDBID.Int64 != 0 {
			animateDB.Polyline_AnimatesDBID.Int64 =
				int64(BackRepoPolylineid_atBckpTime_newID[uint(animateDB.Polyline_AnimatesDBID.Int64)])
		}

		// This reindex animate.Animations
		if animateDB.Rect_AnimationsDBID.Int64 != 0 {
			animateDB.Rect_AnimationsDBID.Int64 =
				int64(BackRepoRectid_atBckpTime_newID[uint(animateDB.Rect_AnimationsDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.RectAnchoredText_AnimatesDBID.Int64 != 0 {
			animateDB.RectAnchoredText_AnimatesDBID.Int64 =
				int64(BackRepoRectAnchoredTextid_atBckpTime_newID[uint(animateDB.RectAnchoredText_AnimatesDBID.Int64)])
		}

		// This reindex animate.Animates
		if animateDB.Text_AnimatesDBID.Int64 != 0 {
			animateDB.Text_AnimatesDBID.Int64 =
				int64(BackRepoTextid_atBckpTime_newID[uint(animateDB.Text_AnimatesDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoAnimate.db.Model(animateDB).Updates(*animateDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// BackRepoAnimate.ResetReversePointers commits all staged instances of Animate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoAnimate *BackRepoAnimateStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, animate := range backRepoAnimate.Map_AnimateDBID_AnimatePtr {
		backRepoAnimate.ResetReversePointersInstance(backRepo, idx, animate)
	}

	return
}

func (backRepoAnimate *BackRepoAnimateStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.Animate) (Error error) {

	// fetch matching animateDB
	if animateDB, ok := backRepoAnimate.Map_AnimateDBID_AnimateDB[idx]; ok {
		_ = animateDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		if animateDB.Circle_AnimationsDBID.Int64 != 0 {
			animateDB.Circle_AnimationsDBID.Int64 = 0
			animateDB.Circle_AnimationsDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.Ellipse_AnimatesDBID.Int64 != 0 {
			animateDB.Ellipse_AnimatesDBID.Int64 = 0
			animateDB.Ellipse_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.Line_AnimatesDBID.Int64 != 0 {
			animateDB.Line_AnimatesDBID.Int64 = 0
			animateDB.Line_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.LinkAnchoredText_AnimatesDBID.Int64 != 0 {
			animateDB.LinkAnchoredText_AnimatesDBID.Int64 = 0
			animateDB.LinkAnchoredText_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.Path_AnimatesDBID.Int64 != 0 {
			animateDB.Path_AnimatesDBID.Int64 = 0
			animateDB.Path_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.Polygone_AnimatesDBID.Int64 != 0 {
			animateDB.Polygone_AnimatesDBID.Int64 = 0
			animateDB.Polygone_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.Polyline_AnimatesDBID.Int64 != 0 {
			animateDB.Polyline_AnimatesDBID.Int64 = 0
			animateDB.Polyline_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.Rect_AnimationsDBID.Int64 != 0 {
			animateDB.Rect_AnimationsDBID.Int64 = 0
			animateDB.Rect_AnimationsDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.RectAnchoredText_AnimatesDBID.Int64 != 0 {
			animateDB.RectAnchoredText_AnimatesDBID.Int64 = 0
			animateDB.RectAnchoredText_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		if animateDB.Text_AnimatesDBID.Int64 != 0 {
			animateDB.Text_AnimatesDBID.Int64 = 0
			animateDB.Text_AnimatesDBID.Valid = true

			// save the reset
			if q := backRepoAnimate.db.Save(animateDB); q.Error != nil {
				return q.Error
			}
		}
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoAnimateid_atBckpTime_newID map[uint]uint