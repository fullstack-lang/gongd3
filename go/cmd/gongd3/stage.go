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
	__Key__000000_Framework2 := (&models.Key{Name: `Framework2`}).Stage()
	__Key__000001_Released := (&models.Key{Name: `Released`}).Stage()
	__Key__000002_Stars2 := (&models.Key{Name: `Stars2`}).Stage()

	// Declarations of staged instances of Serie
	__Serie__000000_FrameworkSerie := (&models.Serie{Name: `FrameworkSerie`}).Stage()
	__Serie__000001_Released := (&models.Serie{Name: `Released`}).Stage()
	__Serie__000002_Stars := (&models.Serie{Name: `Stars`}).Stage()

	// Declarations of staged instances of Value
	__Value__000000_150793 := (&models.Value{Name: `150793`}).Stage()
	__Value__000001_166443 := (&models.Value{Name: `166443`}).Stage()
	__Value__000002_2010 := (&models.Value{Name: `2010`}).Stage()
	__Value__000003_2011 := (&models.Value{Name: `2011`}).Stage()
	__Value__000004_2013 := (&models.Value{Name: `2013`}).Stage()
	__Value__000005_2014 := (&models.Value{Name: `2014`}).Stage()
	__Value__000006_2016 := (&models.Value{Name: `2016`}).Stage()
	__Value__000007_21471 := (&models.Value{Name: `21471`}).Stage()
	__Value__000008_27647 := (&models.Value{Name: `27647`}).Stage()
	__Value__000009_62342 := (&models.Value{Name: `62342`}).Stage()
	__Value__000010_Angular := (&models.Value{Name: `Angular`}).Stage()
	__Value__000011_Backbone := (&models.Value{Name: `Backbone`}).Stage()
	__Value__000012_Ember := (&models.Value{Name: `Ember`}).Stage()
	__Value__000013_React := (&models.Value{Name: `React`}).Stage()
	__Value__000014_Vue := (&models.Value{Name: `Vue`}).Stage()

	// Setup of values

	// Bar values setup
	__Bar__000000_Stars_per_Framework.Name = `Stars per Framework`

	// Key values setup
	__Key__000000_Framework2.Name = `Framework2`

	// Key values setup
	__Key__000001_Released.Name = `Released`

	// Key values setup
	__Key__000002_Stars2.Name = `Stars2`

	// Serie values setup
	__Serie__000000_FrameworkSerie.Name = `FrameworkSerie`

	// Serie values setup
	__Serie__000001_Released.Name = `Released`

	// Serie values setup
	__Serie__000002_Stars.Name = `Stars`

	// Value values setup
	__Value__000000_150793.Name = `150793`

	// Value values setup
	__Value__000001_166443.Name = `166443`

	// Value values setup
	__Value__000002_2010.Name = `2010`

	// Value values setup
	__Value__000003_2011.Name = `2011`

	// Value values setup
	__Value__000004_2013.Name = `2013`

	// Value values setup
	__Value__000005_2014.Name = `2014`

	// Value values setup
	__Value__000006_2016.Name = `2016`

	// Value values setup
	__Value__000007_21471.Name = `21471`

	// Value values setup
	__Value__000008_27647.Name = `27647`

	// Value values setup
	__Value__000009_62342.Name = `62342`

	// Value values setup
	__Value__000010_Angular.Name = `Angular`

	// Value values setup
	__Value__000011_Backbone.Name = `Backbone`

	// Value values setup
	__Value__000012_Ember.Name = `Ember`

	// Value values setup
	__Value__000013_React.Name = `React`

	// Value values setup
	__Value__000014_Vue.Name = `Vue`

	// Setup of pointers
	__Bar__000000_Stars_per_Framework.X = __Key__000000_Framework2
	__Bar__000000_Stars_per_Framework.Y = __Key__000002_Stars2
	__Bar__000000_Stars_per_Framework.Set = append(__Bar__000000_Stars_per_Framework.Set, __Serie__000002_Stars)
	__Bar__000000_Stars_per_Framework.Set = append(__Bar__000000_Stars_per_Framework.Set, __Serie__000001_Released)
	__Bar__000000_Stars_per_Framework.Set = append(__Bar__000000_Stars_per_Framework.Set, __Serie__000000_FrameworkSerie)
	__Serie__000000_FrameworkSerie.Key = __Key__000000_Framework2
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000014_Vue)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000013_React)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000010_Angular)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000011_Backbone)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000012_Ember)
	__Serie__000001_Released.Key = __Key__000001_Released
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000005_2014)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000004_2013)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000006_2016)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000002_2010)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000003_2011)
	__Serie__000002_Stars.Key = __Key__000002_Stars2
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000001_166443)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000000_150793)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000009_62342)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000008_27647)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000007_21471)
}


