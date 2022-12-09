package main

import (
	"time"

	"gongd3/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage models.StageStruct
var ___dummy__Time time.Time

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = stageInjection
// }

// stageInjection will stage objects of database "stage"
func stageInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Bar
	__Bar__000000_Stars_per_Framework := (&models.Bar{Name: `Stars per Framework`}).Stage()

	// Declarations of staged instances of Key
	__Key__000000_Framework := (&models.Key{Name: `Framework`}).Stage()
	__Key__000001_Released := (&models.Key{Name: `Released`}).Stage()
	__Key__000002_Stars := (&models.Key{Name: `Stars`}).Stage()

	// Declarations of staged instances of Serie
	__Serie__000000_Framework := (&models.Serie{Name: `Framework`}).Stage()
	__Serie__000001_Released := (&models.Serie{Name: `Released`}).Stage()
	__Serie__000002_Stars := (&models.Serie{Name: `Stars`}).Stage()

	// Declarations of staged instances of Value
	__Value__000000_150793 := (&models.Value{Name: `150793`}).Stage()
	__Value__000001_166443 := (&models.Value{Name: `166443`}).Stage()
	__Value__000002_21471 := (&models.Value{Name: `21471`}).Stage()
	__Value__000003_27647 := (&models.Value{Name: `27647`}).Stage()
	__Value__000004_62342 := (&models.Value{Name: `62342`}).Stage()
	__Value__000005_Angular := (&models.Value{Name: `Angular`}).Stage()
	__Value__000006_Backbone := (&models.Value{Name: `Backbone`}).Stage()
	__Value__000007_Ember := (&models.Value{Name: `Ember`}).Stage()
	__Value__000008_React := (&models.Value{Name: `React`}).Stage()
	__Value__000009_Vue := (&models.Value{Name: `Vue`}).Stage()

	// Setup of values

	// Bar values setup
	__Bar__000000_Stars_per_Framework.Name = `Stars per Framework`

	// Key values setup
	__Key__000000_Framework.Name = `Framework`

	// Key values setup
	__Key__000001_Released.Name = `Released`

	// Key values setup
	__Key__000002_Stars.Name = `Stars`

	// Serie values setup
	__Serie__000000_Framework.Name = `Framework`

	// Serie values setup
	__Serie__000001_Released.Name = `Released`

	// Serie values setup
	__Serie__000002_Stars.Name = `Stars`

	// Value values setup
	__Value__000000_150793.Name = `150793`

	// Value values setup
	__Value__000001_166443.Name = `166443`

	// Value values setup
	__Value__000002_21471.Name = `21471`

	// Value values setup
	__Value__000003_27647.Name = `27647`

	// Value values setup
	__Value__000004_62342.Name = `62342`

	// Value values setup
	__Value__000005_Angular.Name = `Angular`

	// Value values setup
	__Value__000006_Backbone.Name = `Backbone`

	// Value values setup
	__Value__000007_Ember.Name = `Ember`

	// Value values setup
	__Value__000008_React.Name = `React`

	// Value values setup
	__Value__000009_Vue.Name = `Vue`

	// Setup of pointers
	__Bar__000000_Stars_per_Framework.X = __Key__000000_Framework
	__Bar__000000_Stars_per_Framework.Y = __Key__000002_Stars
	__Bar__000000_Stars_per_Framework.Set = append(__Bar__000000_Stars_per_Framework.Set, __Serie__000002_Stars)
	__Bar__000000_Stars_per_Framework.Set = append(__Bar__000000_Stars_per_Framework.Set, __Serie__000001_Released)
	__Bar__000000_Stars_per_Framework.Set = append(__Bar__000000_Stars_per_Framework.Set, __Serie__000000_Framework)
	__Serie__000000_Framework.Key = __Key__000000_Framework
	__Serie__000000_Framework.Values = append(__Serie__000000_Framework.Values, __Value__000007_Ember)
	__Serie__000000_Framework.Values = append(__Serie__000000_Framework.Values, __Value__000008_React)
	__Serie__000000_Framework.Values = append(__Serie__000000_Framework.Values, __Value__000006_Backbone)
	__Serie__000000_Framework.Values = append(__Serie__000000_Framework.Values, __Value__000009_Vue)
	__Serie__000000_Framework.Values = append(__Serie__000000_Framework.Values, __Value__000005_Angular)
	__Serie__000001_Released.Key = __Key__000001_Released
	__Serie__000002_Stars.Key = __Key__000002_Stars
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000000_150793)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000004_62342)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000001_166443)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000003_27647)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000002_21471)
}


