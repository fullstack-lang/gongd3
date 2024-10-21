// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Bars           map[*Bar]any
	Bars_mapString map[string]*Bar

	// insertion point for slice of pointers maps
	Bar_Set_reverseMap map[*Serie]*Bar

	OnAfterBarCreateCallback OnAfterCreateInterface[Bar]
	OnAfterBarUpdateCallback OnAfterUpdateInterface[Bar]
	OnAfterBarDeleteCallback OnAfterDeleteInterface[Bar]
	OnAfterBarReadCallback   OnAfterReadInterface[Bar]

	Keys           map[*Key]any
	Keys_mapString map[string]*Key

	// insertion point for slice of pointers maps
	OnAfterKeyCreateCallback OnAfterCreateInterface[Key]
	OnAfterKeyUpdateCallback OnAfterUpdateInterface[Key]
	OnAfterKeyDeleteCallback OnAfterDeleteInterface[Key]
	OnAfterKeyReadCallback   OnAfterReadInterface[Key]

	Pies           map[*Pie]any
	Pies_mapString map[string]*Pie

	// insertion point for slice of pointers maps
	Pie_Set_reverseMap map[*Serie]*Pie

	OnAfterPieCreateCallback OnAfterCreateInterface[Pie]
	OnAfterPieUpdateCallback OnAfterUpdateInterface[Pie]
	OnAfterPieDeleteCallback OnAfterDeleteInterface[Pie]
	OnAfterPieReadCallback   OnAfterReadInterface[Pie]

	Scatters           map[*Scatter]any
	Scatters_mapString map[string]*Scatter

	// insertion point for slice of pointers maps
	Scatter_Set_reverseMap map[*Serie]*Scatter

	OnAfterScatterCreateCallback OnAfterCreateInterface[Scatter]
	OnAfterScatterUpdateCallback OnAfterUpdateInterface[Scatter]
	OnAfterScatterDeleteCallback OnAfterDeleteInterface[Scatter]
	OnAfterScatterReadCallback   OnAfterReadInterface[Scatter]

	Series           map[*Serie]any
	Series_mapString map[string]*Serie

	// insertion point for slice of pointers maps
	Serie_Values_reverseMap map[*Value]*Serie

	OnAfterSerieCreateCallback OnAfterCreateInterface[Serie]
	OnAfterSerieUpdateCallback OnAfterUpdateInterface[Serie]
	OnAfterSerieDeleteCallback OnAfterDeleteInterface[Serie]
	OnAfterSerieReadCallback   OnAfterReadInterface[Serie]

	Values           map[*Value]any
	Values_mapString map[string]*Value

	// insertion point for slice of pointers maps
	OnAfterValueCreateCallback OnAfterCreateInterface[Value]
	OnAfterValueUpdateCallback OnAfterUpdateInterface[Value]
	OnAfterValueDeleteCallback OnAfterDeleteInterface[Value]
	OnAfterValueReadCallback   OnAfterReadInterface[Value]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongd3/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitBar(bar *Bar)
	CheckoutBar(bar *Bar)
	CommitKey(key *Key)
	CheckoutKey(key *Key)
	CommitPie(pie *Pie)
	CheckoutPie(pie *Pie)
	CommitScatter(scatter *Scatter)
	CheckoutScatter(scatter *Scatter)
	CommitSerie(serie *Serie)
	CheckoutSerie(serie *Serie)
	CommitValue(value *Value)
	CheckoutValue(value *Value)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Bars:           make(map[*Bar]any),
		Bars_mapString: make(map[string]*Bar),

		Keys:           make(map[*Key]any),
		Keys_mapString: make(map[string]*Key),

		Pies:           make(map[*Pie]any),
		Pies_mapString: make(map[string]*Pie),

		Scatters:           make(map[*Scatter]any),
		Scatters_mapString: make(map[string]*Scatter),

		Series:           make(map[*Serie]any),
		Series_mapString: make(map[string]*Serie),

		Values:           make(map[*Value]any),
		Values_mapString: make(map[string]*Value),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Bar"] = len(stage.Bars)
	stage.Map_GongStructName_InstancesNb["Key"] = len(stage.Keys)
	stage.Map_GongStructName_InstancesNb["Pie"] = len(stage.Pies)
	stage.Map_GongStructName_InstancesNb["Scatter"] = len(stage.Scatters)
	stage.Map_GongStructName_InstancesNb["Serie"] = len(stage.Series)
	stage.Map_GongStructName_InstancesNb["Value"] = len(stage.Values)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Bar"] = len(stage.Bars)
	stage.Map_GongStructName_InstancesNb["Key"] = len(stage.Keys)
	stage.Map_GongStructName_InstancesNb["Pie"] = len(stage.Pies)
	stage.Map_GongStructName_InstancesNb["Scatter"] = len(stage.Scatters)
	stage.Map_GongStructName_InstancesNb["Serie"] = len(stage.Series)
	stage.Map_GongStructName_InstancesNb["Value"] = len(stage.Values)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts bar to the model stage
func (bar *Bar) Stage(stage *StageStruct) *Bar {
	stage.Bars[bar] = __member
	stage.Bars_mapString[bar.Name] = bar

	return bar
}

// Unstage removes bar off the model stage
func (bar *Bar) Unstage(stage *StageStruct) *Bar {
	delete(stage.Bars, bar)
	delete(stage.Bars_mapString, bar.Name)
	return bar
}

// UnstageVoid removes bar off the model stage
func (bar *Bar) UnstageVoid(stage *StageStruct) {
	delete(stage.Bars, bar)
	delete(stage.Bars_mapString, bar.Name)
}

// commit bar to the back repo (if it is already staged)
func (bar *Bar) Commit(stage *StageStruct) *Bar {
	if _, ok := stage.Bars[bar]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBar(bar)
		}
	}
	return bar
}

func (bar *Bar) CommitVoid(stage *StageStruct) {
	bar.Commit(stage)
}

// Checkout bar to the back repo (if it is already staged)
func (bar *Bar) Checkout(stage *StageStruct) *Bar {
	if _, ok := stage.Bars[bar]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBar(bar)
		}
	}
	return bar
}

// for satisfaction of GongStruct interface
func (bar *Bar) GetName() (res string) {
	return bar.Name
}

// Stage puts key to the model stage
func (key *Key) Stage(stage *StageStruct) *Key {
	stage.Keys[key] = __member
	stage.Keys_mapString[key.Name] = key

	return key
}

// Unstage removes key off the model stage
func (key *Key) Unstage(stage *StageStruct) *Key {
	delete(stage.Keys, key)
	delete(stage.Keys_mapString, key.Name)
	return key
}

// UnstageVoid removes key off the model stage
func (key *Key) UnstageVoid(stage *StageStruct) {
	delete(stage.Keys, key)
	delete(stage.Keys_mapString, key.Name)
}

// commit key to the back repo (if it is already staged)
func (key *Key) Commit(stage *StageStruct) *Key {
	if _, ok := stage.Keys[key]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitKey(key)
		}
	}
	return key
}

func (key *Key) CommitVoid(stage *StageStruct) {
	key.Commit(stage)
}

// Checkout key to the back repo (if it is already staged)
func (key *Key) Checkout(stage *StageStruct) *Key {
	if _, ok := stage.Keys[key]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutKey(key)
		}
	}
	return key
}

// for satisfaction of GongStruct interface
func (key *Key) GetName() (res string) {
	return key.Name
}

// Stage puts pie to the model stage
func (pie *Pie) Stage(stage *StageStruct) *Pie {
	stage.Pies[pie] = __member
	stage.Pies_mapString[pie.Name] = pie

	return pie
}

// Unstage removes pie off the model stage
func (pie *Pie) Unstage(stage *StageStruct) *Pie {
	delete(stage.Pies, pie)
	delete(stage.Pies_mapString, pie.Name)
	return pie
}

// UnstageVoid removes pie off the model stage
func (pie *Pie) UnstageVoid(stage *StageStruct) {
	delete(stage.Pies, pie)
	delete(stage.Pies_mapString, pie.Name)
}

// commit pie to the back repo (if it is already staged)
func (pie *Pie) Commit(stage *StageStruct) *Pie {
	if _, ok := stage.Pies[pie]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPie(pie)
		}
	}
	return pie
}

func (pie *Pie) CommitVoid(stage *StageStruct) {
	pie.Commit(stage)
}

// Checkout pie to the back repo (if it is already staged)
func (pie *Pie) Checkout(stage *StageStruct) *Pie {
	if _, ok := stage.Pies[pie]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPie(pie)
		}
	}
	return pie
}

// for satisfaction of GongStruct interface
func (pie *Pie) GetName() (res string) {
	return pie.Name
}

// Stage puts scatter to the model stage
func (scatter *Scatter) Stage(stage *StageStruct) *Scatter {
	stage.Scatters[scatter] = __member
	stage.Scatters_mapString[scatter.Name] = scatter

	return scatter
}

// Unstage removes scatter off the model stage
func (scatter *Scatter) Unstage(stage *StageStruct) *Scatter {
	delete(stage.Scatters, scatter)
	delete(stage.Scatters_mapString, scatter.Name)
	return scatter
}

// UnstageVoid removes scatter off the model stage
func (scatter *Scatter) UnstageVoid(stage *StageStruct) {
	delete(stage.Scatters, scatter)
	delete(stage.Scatters_mapString, scatter.Name)
}

// commit scatter to the back repo (if it is already staged)
func (scatter *Scatter) Commit(stage *StageStruct) *Scatter {
	if _, ok := stage.Scatters[scatter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitScatter(scatter)
		}
	}
	return scatter
}

func (scatter *Scatter) CommitVoid(stage *StageStruct) {
	scatter.Commit(stage)
}

// Checkout scatter to the back repo (if it is already staged)
func (scatter *Scatter) Checkout(stage *StageStruct) *Scatter {
	if _, ok := stage.Scatters[scatter]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutScatter(scatter)
		}
	}
	return scatter
}

// for satisfaction of GongStruct interface
func (scatter *Scatter) GetName() (res string) {
	return scatter.Name
}

// Stage puts serie to the model stage
func (serie *Serie) Stage(stage *StageStruct) *Serie {
	stage.Series[serie] = __member
	stage.Series_mapString[serie.Name] = serie

	return serie
}

// Unstage removes serie off the model stage
func (serie *Serie) Unstage(stage *StageStruct) *Serie {
	delete(stage.Series, serie)
	delete(stage.Series_mapString, serie.Name)
	return serie
}

// UnstageVoid removes serie off the model stage
func (serie *Serie) UnstageVoid(stage *StageStruct) {
	delete(stage.Series, serie)
	delete(stage.Series_mapString, serie.Name)
}

// commit serie to the back repo (if it is already staged)
func (serie *Serie) Commit(stage *StageStruct) *Serie {
	if _, ok := stage.Series[serie]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSerie(serie)
		}
	}
	return serie
}

func (serie *Serie) CommitVoid(stage *StageStruct) {
	serie.Commit(stage)
}

// Checkout serie to the back repo (if it is already staged)
func (serie *Serie) Checkout(stage *StageStruct) *Serie {
	if _, ok := stage.Series[serie]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSerie(serie)
		}
	}
	return serie
}

// for satisfaction of GongStruct interface
func (serie *Serie) GetName() (res string) {
	return serie.Name
}

// Stage puts value to the model stage
func (value *Value) Stage(stage *StageStruct) *Value {
	stage.Values[value] = __member
	stage.Values_mapString[value.Name] = value

	return value
}

// Unstage removes value off the model stage
func (value *Value) Unstage(stage *StageStruct) *Value {
	delete(stage.Values, value)
	delete(stage.Values_mapString, value.Name)
	return value
}

// UnstageVoid removes value off the model stage
func (value *Value) UnstageVoid(stage *StageStruct) {
	delete(stage.Values, value)
	delete(stage.Values_mapString, value.Name)
}

// commit value to the back repo (if it is already staged)
func (value *Value) Commit(stage *StageStruct) *Value {
	if _, ok := stage.Values[value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitValue(value)
		}
	}
	return value
}

func (value *Value) CommitVoid(stage *StageStruct) {
	value.Commit(stage)
}

// Checkout value to the back repo (if it is already staged)
func (value *Value) Checkout(stage *StageStruct) *Value {
	if _, ok := stage.Values[value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutValue(value)
		}
	}
	return value
}

// for satisfaction of GongStruct interface
func (value *Value) GetName() (res string) {
	return value.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMBar(Bar *Bar)
	CreateORMKey(Key *Key)
	CreateORMPie(Pie *Pie)
	CreateORMScatter(Scatter *Scatter)
	CreateORMSerie(Serie *Serie)
	CreateORMValue(Value *Value)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMBar(Bar *Bar)
	DeleteORMKey(Key *Key)
	DeleteORMPie(Pie *Pie)
	DeleteORMScatter(Scatter *Scatter)
	DeleteORMSerie(Serie *Serie)
	DeleteORMValue(Value *Value)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Bars = make(map[*Bar]any)
	stage.Bars_mapString = make(map[string]*Bar)

	stage.Keys = make(map[*Key]any)
	stage.Keys_mapString = make(map[string]*Key)

	stage.Pies = make(map[*Pie]any)
	stage.Pies_mapString = make(map[string]*Pie)

	stage.Scatters = make(map[*Scatter]any)
	stage.Scatters_mapString = make(map[string]*Scatter)

	stage.Series = make(map[*Serie]any)
	stage.Series_mapString = make(map[string]*Serie)

	stage.Values = make(map[*Value]any)
	stage.Values_mapString = make(map[string]*Value)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Bars = nil
	stage.Bars_mapString = nil

	stage.Keys = nil
	stage.Keys_mapString = nil

	stage.Pies = nil
	stage.Pies_mapString = nil

	stage.Scatters = nil
	stage.Scatters_mapString = nil

	stage.Series = nil
	stage.Series_mapString = nil

	stage.Values = nil
	stage.Values_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for bar := range stage.Bars {
		bar.Unstage(stage)
	}

	for key := range stage.Keys {
		key.Unstage(stage)
	}

	for pie := range stage.Pies {
		pie.Unstage(stage)
	}

	for scatter := range stage.Scatters {
		scatter.Unstage(stage)
	}

	for serie := range stage.Series {
		serie.Unstage(stage)
	}

	for value := range stage.Values {
		value.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Bar]any:
		return any(&stage.Bars).(*Type)
	case map[*Key]any:
		return any(&stage.Keys).(*Type)
	case map[*Pie]any:
		return any(&stage.Pies).(*Type)
	case map[*Scatter]any:
		return any(&stage.Scatters).(*Type)
	case map[*Serie]any:
		return any(&stage.Series).(*Type)
	case map[*Value]any:
		return any(&stage.Values).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Bar:
		return any(&stage.Bars_mapString).(*Type)
	case map[string]*Key:
		return any(&stage.Keys_mapString).(*Type)
	case map[string]*Pie:
		return any(&stage.Pies_mapString).(*Type)
	case map[string]*Scatter:
		return any(&stage.Scatters_mapString).(*Type)
	case map[string]*Serie:
		return any(&stage.Series_mapString).(*Type)
	case map[string]*Value:
		return any(&stage.Values_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Bar:
		return any(&stage.Bars).(*map[*Type]any)
	case Key:
		return any(&stage.Keys).(*map[*Type]any)
	case Pie:
		return any(&stage.Pies).(*map[*Type]any)
	case Scatter:
		return any(&stage.Scatters).(*map[*Type]any)
	case Serie:
		return any(&stage.Series).(*map[*Type]any)
	case Value:
		return any(&stage.Values).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Bar:
		return any(&stage.Bars).(*map[Type]any)
	case *Key:
		return any(&stage.Keys).(*map[Type]any)
	case *Pie:
		return any(&stage.Pies).(*map[Type]any)
	case *Scatter:
		return any(&stage.Scatters).(*map[Type]any)
	case *Serie:
		return any(&stage.Series).(*map[Type]any)
	case *Value:
		return any(&stage.Values).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Bar:
		return any(&stage.Bars_mapString).(*map[string]*Type)
	case Key:
		return any(&stage.Keys_mapString).(*map[string]*Type)
	case Pie:
		return any(&stage.Pies_mapString).(*map[string]*Type)
	case Scatter:
		return any(&stage.Scatters_mapString).(*map[string]*Type)
	case Serie:
		return any(&stage.Series_mapString).(*map[string]*Type)
	case Value:
		return any(&stage.Values_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Bar:
		return any(&Bar{
			// Initialisation of associations
			// field is initialized with an instance of Key with the name of the field
			X: &Key{Name: "X"},
			// field is initialized with an instance of Key with the name of the field
			Y: &Key{Name: "Y"},
			// field is initialized with an instance of Serie with the name of the field
			Set: []*Serie{{Name: "Set"}},
		}).(*Type)
	case Key:
		return any(&Key{
			// Initialisation of associations
		}).(*Type)
	case Pie:
		return any(&Pie{
			// Initialisation of associations
			// field is initialized with an instance of Key with the name of the field
			X: &Key{Name: "X"},
			// field is initialized with an instance of Key with the name of the field
			Y: &Key{Name: "Y"},
			// field is initialized with an instance of Serie with the name of the field
			Set: []*Serie{{Name: "Set"}},
		}).(*Type)
	case Scatter:
		return any(&Scatter{
			// Initialisation of associations
			// field is initialized with an instance of Key with the name of the field
			X: &Key{Name: "X"},
			// field is initialized with an instance of Key with the name of the field
			Y: &Key{Name: "Y"},
			// field is initialized with an instance of Key with the name of the field
			Text: &Key{Name: "Text"},
			// field is initialized with an instance of Serie with the name of the field
			Set: []*Serie{{Name: "Set"}},
		}).(*Type)
	case Serie:
		return any(&Serie{
			// Initialisation of associations
			// field is initialized with an instance of Key with the name of the field
			Key: &Key{Name: "Key"},
			// field is initialized with an instance of Value with the name of the field
			Values: []*Value{{Name: "Values"}},
		}).(*Type)
	case Value:
		return any(&Value{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Bar
	case Bar:
		switch fieldname {
		// insertion point for per direct association field
		case "X":
			res := make(map[*Key][]*Bar)
			for bar := range stage.Bars {
				if bar.X != nil {
					key_ := bar.X
					var bars []*Bar
					_, ok := res[key_]
					if ok {
						bars = res[key_]
					} else {
						bars = make([]*Bar, 0)
					}
					bars = append(bars, bar)
					res[key_] = bars
				}
			}
			return any(res).(map[*End][]*Start)
		case "Y":
			res := make(map[*Key][]*Bar)
			for bar := range stage.Bars {
				if bar.Y != nil {
					key_ := bar.Y
					var bars []*Bar
					_, ok := res[key_]
					if ok {
						bars = res[key_]
					} else {
						bars = make([]*Bar, 0)
					}
					bars = append(bars, bar)
					res[key_] = bars
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Key
	case Key:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pie
	case Pie:
		switch fieldname {
		// insertion point for per direct association field
		case "X":
			res := make(map[*Key][]*Pie)
			for pie := range stage.Pies {
				if pie.X != nil {
					key_ := pie.X
					var pies []*Pie
					_, ok := res[key_]
					if ok {
						pies = res[key_]
					} else {
						pies = make([]*Pie, 0)
					}
					pies = append(pies, pie)
					res[key_] = pies
				}
			}
			return any(res).(map[*End][]*Start)
		case "Y":
			res := make(map[*Key][]*Pie)
			for pie := range stage.Pies {
				if pie.Y != nil {
					key_ := pie.Y
					var pies []*Pie
					_, ok := res[key_]
					if ok {
						pies = res[key_]
					} else {
						pies = make([]*Pie, 0)
					}
					pies = append(pies, pie)
					res[key_] = pies
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Scatter
	case Scatter:
		switch fieldname {
		// insertion point for per direct association field
		case "X":
			res := make(map[*Key][]*Scatter)
			for scatter := range stage.Scatters {
				if scatter.X != nil {
					key_ := scatter.X
					var scatters []*Scatter
					_, ok := res[key_]
					if ok {
						scatters = res[key_]
					} else {
						scatters = make([]*Scatter, 0)
					}
					scatters = append(scatters, scatter)
					res[key_] = scatters
				}
			}
			return any(res).(map[*End][]*Start)
		case "Y":
			res := make(map[*Key][]*Scatter)
			for scatter := range stage.Scatters {
				if scatter.Y != nil {
					key_ := scatter.Y
					var scatters []*Scatter
					_, ok := res[key_]
					if ok {
						scatters = res[key_]
					} else {
						scatters = make([]*Scatter, 0)
					}
					scatters = append(scatters, scatter)
					res[key_] = scatters
				}
			}
			return any(res).(map[*End][]*Start)
		case "Text":
			res := make(map[*Key][]*Scatter)
			for scatter := range stage.Scatters {
				if scatter.Text != nil {
					key_ := scatter.Text
					var scatters []*Scatter
					_, ok := res[key_]
					if ok {
						scatters = res[key_]
					} else {
						scatters = make([]*Scatter, 0)
					}
					scatters = append(scatters, scatter)
					res[key_] = scatters
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Serie
	case Serie:
		switch fieldname {
		// insertion point for per direct association field
		case "Key":
			res := make(map[*Key][]*Serie)
			for serie := range stage.Series {
				if serie.Key != nil {
					key_ := serie.Key
					var series []*Serie
					_, ok := res[key_]
					if ok {
						series = res[key_]
					} else {
						series = make([]*Serie, 0)
					}
					series = append(series, serie)
					res[key_] = series
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Value
	case Value:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Bar
	case Bar:
		switch fieldname {
		// insertion point for per direct association field
		case "Set":
			res := make(map[*Serie]*Bar)
			for bar := range stage.Bars {
				for _, serie_ := range bar.Set {
					res[serie_] = bar
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Key
	case Key:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pie
	case Pie:
		switch fieldname {
		// insertion point for per direct association field
		case "Set":
			res := make(map[*Serie]*Pie)
			for pie := range stage.Pies {
				for _, serie_ := range pie.Set {
					res[serie_] = pie
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Scatter
	case Scatter:
		switch fieldname {
		// insertion point for per direct association field
		case "Set":
			res := make(map[*Serie]*Scatter)
			for scatter := range stage.Scatters {
				for _, serie_ := range scatter.Set {
					res[serie_] = scatter
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Serie
	case Serie:
		switch fieldname {
		// insertion point for per direct association field
		case "Values":
			res := make(map[*Value]*Serie)
			for serie := range stage.Series {
				for _, value_ := range serie.Values {
					res[value_] = serie
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Value
	case Value:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Bar:
		res = "Bar"
	case Key:
		res = "Key"
	case Pie:
		res = "Pie"
	case Scatter:
		res = "Scatter"
	case Serie:
		res = "Serie"
	case Value:
		res = "Value"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Bar:
		res = "Bar"
	case *Key:
		res = "Key"
	case *Pie:
		res = "Pie"
	case *Scatter:
		res = "Scatter"
	case *Serie:
		res = "Serie"
	case *Value:
		res = "Value"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Bar:
		res = []string{"Name", "X", "Y", "Set", "AutoDomainX", "XMin", "XMax", "AutoDomainY", "YMin", "YMax", "YLabelPresent", "YLabelOffset", "XLabelPresent", "XLabelOffset", "Width", "Heigth", "Margin"}
	case Key:
		res = []string{"Name"}
	case Pie:
		res = []string{"Name", "X", "Y", "Set", "Width", "Heigth", "Margin"}
	case Scatter:
		res = []string{"Name", "X", "Y", "Text", "Set", "Width", "Heigth", "Margin"}
	case Serie:
		res = []string{"Name", "Key", "Values"}
	case Value:
		res = []string{"Name"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Bar:
		var rf ReverseField
		_ = rf
	case Key:
		var rf ReverseField
		_ = rf
	case Pie:
		var rf ReverseField
		_ = rf
	case Scatter:
		var rf ReverseField
		_ = rf
	case Serie:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Bar"
		rf.Fieldname = "Set"
		res = append(res, rf)
		rf.GongstructName = "Pie"
		rf.Fieldname = "Set"
		res = append(res, rf)
		rf.GongstructName = "Scatter"
		rf.Fieldname = "Set"
		res = append(res, rf)
	case Value:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Serie"
		rf.Fieldname = "Values"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Bar:
		res = []string{"Name", "X", "Y", "Set", "AutoDomainX", "XMin", "XMax", "AutoDomainY", "YMin", "YMax", "YLabelPresent", "YLabelOffset", "XLabelPresent", "XLabelOffset", "Width", "Heigth", "Margin"}
	case *Key:
		res = []string{"Name"}
	case *Pie:
		res = []string{"Name", "X", "Y", "Set", "Width", "Heigth", "Margin"}
	case *Scatter:
		res = []string{"Name", "X", "Y", "Text", "Set", "Width", "Heigth", "Margin"}
	case *Serie:
		res = []string{"Name", "Key", "Values"}
	case *Value:
		res = []string{"Name"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Bar:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			if inferedInstance.X != nil {
				res = inferedInstance.X.Name
			}
		case "Y":
			if inferedInstance.Y != nil {
				res = inferedInstance.Y.Name
			}
		case "Set":
			for idx, __instance__ := range inferedInstance.Set {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "AutoDomainX":
			res = fmt.Sprintf("%t", inferedInstance.AutoDomainX)
		case "XMin":
			res = fmt.Sprintf("%f", inferedInstance.XMin)
		case "XMax":
			res = fmt.Sprintf("%f", inferedInstance.XMax)
		case "AutoDomainY":
			res = fmt.Sprintf("%t", inferedInstance.AutoDomainY)
		case "YMin":
			res = fmt.Sprintf("%f", inferedInstance.YMin)
		case "YMax":
			res = fmt.Sprintf("%f", inferedInstance.YMax)
		case "YLabelPresent":
			res = fmt.Sprintf("%t", inferedInstance.YLabelPresent)
		case "YLabelOffset":
			res = fmt.Sprintf("%f", inferedInstance.YLabelOffset)
		case "XLabelPresent":
			res = fmt.Sprintf("%t", inferedInstance.XLabelPresent)
		case "XLabelOffset":
			res = fmt.Sprintf("%f", inferedInstance.XLabelOffset)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Heigth":
			res = fmt.Sprintf("%f", inferedInstance.Heigth)
		case "Margin":
			res = fmt.Sprintf("%f", inferedInstance.Margin)
		}
	case *Key:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Pie:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			if inferedInstance.X != nil {
				res = inferedInstance.X.Name
			}
		case "Y":
			if inferedInstance.Y != nil {
				res = inferedInstance.Y.Name
			}
		case "Set":
			for idx, __instance__ := range inferedInstance.Set {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Heigth":
			res = fmt.Sprintf("%f", inferedInstance.Heigth)
		case "Margin":
			res = fmt.Sprintf("%f", inferedInstance.Margin)
		}
	case *Scatter:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			if inferedInstance.X != nil {
				res = inferedInstance.X.Name
			}
		case "Y":
			if inferedInstance.Y != nil {
				res = inferedInstance.Y.Name
			}
		case "Text":
			if inferedInstance.Text != nil {
				res = inferedInstance.Text.Name
			}
		case "Set":
			for idx, __instance__ := range inferedInstance.Set {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Heigth":
			res = fmt.Sprintf("%f", inferedInstance.Heigth)
		case "Margin":
			res = fmt.Sprintf("%f", inferedInstance.Margin)
		}
	case *Serie:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Key":
			if inferedInstance.Key != nil {
				res = inferedInstance.Key.Name
			}
		case "Values":
			for idx, __instance__ := range inferedInstance.Values {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Value:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Bar:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			if inferedInstance.X != nil {
				res = inferedInstance.X.Name
			}
		case "Y":
			if inferedInstance.Y != nil {
				res = inferedInstance.Y.Name
			}
		case "Set":
			for idx, __instance__ := range inferedInstance.Set {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "AutoDomainX":
			res = fmt.Sprintf("%t", inferedInstance.AutoDomainX)
		case "XMin":
			res = fmt.Sprintf("%f", inferedInstance.XMin)
		case "XMax":
			res = fmt.Sprintf("%f", inferedInstance.XMax)
		case "AutoDomainY":
			res = fmt.Sprintf("%t", inferedInstance.AutoDomainY)
		case "YMin":
			res = fmt.Sprintf("%f", inferedInstance.YMin)
		case "YMax":
			res = fmt.Sprintf("%f", inferedInstance.YMax)
		case "YLabelPresent":
			res = fmt.Sprintf("%t", inferedInstance.YLabelPresent)
		case "YLabelOffset":
			res = fmt.Sprintf("%f", inferedInstance.YLabelOffset)
		case "XLabelPresent":
			res = fmt.Sprintf("%t", inferedInstance.XLabelPresent)
		case "XLabelOffset":
			res = fmt.Sprintf("%f", inferedInstance.XLabelOffset)
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Heigth":
			res = fmt.Sprintf("%f", inferedInstance.Heigth)
		case "Margin":
			res = fmt.Sprintf("%f", inferedInstance.Margin)
		}
	case Key:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Pie:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			if inferedInstance.X != nil {
				res = inferedInstance.X.Name
			}
		case "Y":
			if inferedInstance.Y != nil {
				res = inferedInstance.Y.Name
			}
		case "Set":
			for idx, __instance__ := range inferedInstance.Set {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Heigth":
			res = fmt.Sprintf("%f", inferedInstance.Heigth)
		case "Margin":
			res = fmt.Sprintf("%f", inferedInstance.Margin)
		}
	case Scatter:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "X":
			if inferedInstance.X != nil {
				res = inferedInstance.X.Name
			}
		case "Y":
			if inferedInstance.Y != nil {
				res = inferedInstance.Y.Name
			}
		case "Text":
			if inferedInstance.Text != nil {
				res = inferedInstance.Text.Name
			}
		case "Set":
			for idx, __instance__ := range inferedInstance.Set {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Width":
			res = fmt.Sprintf("%f", inferedInstance.Width)
		case "Heigth":
			res = fmt.Sprintf("%f", inferedInstance.Heigth)
		case "Margin":
			res = fmt.Sprintf("%f", inferedInstance.Margin)
		}
	case Serie:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Key":
			if inferedInstance.Key != nil {
				res = inferedInstance.Key.Name
			}
		case "Values":
			for idx, __instance__ := range inferedInstance.Values {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Value:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
