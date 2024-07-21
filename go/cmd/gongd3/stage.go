package main

import (
	"time"

	"github.com/fullstack-lang/gongd3/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage models.StageStruct
var ___dummy__Time time.Time

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = _
// }

// _ will stage objects of database "stage"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Bar
	__Bar__000000_Bar_Stars_per_Framework := (&models.Bar{Name: `Bar Stars per Framework`}).Stage(stage)
	__Bar__000001_Random_X_Y := (&models.Bar{Name: `Random X & Y`}).Stage(stage)

	// Declarations of staged instances of Key
	__Key__000000_Framework_not_hardcoded := (&models.Key{Name: `Framework not hardcoded`}).Stage(stage)
	__Key__000001_Released_not_hardcoded := (&models.Key{Name: `Released not hardcoded`}).Stage(stage)
	__Key__000002_Stars_not_hardcoded := (&models.Key{Name: `Stars not hardcoded`}).Stage(stage)
	__Key__000003_X := (&models.Key{Name: `X`}).Stage(stage)
	__Key__000004_Y := (&models.Key{Name: `Y`}).Stage(stage)

	// Declarations of staged instances of Pie
	__Pie__000000_Pie_Stars_per_Framework := (&models.Pie{Name: `Pie Stars per Framework`}).Stage(stage)

	// Declarations of staged instances of Scatter
	__Scatter__000000_Scatter_Stars_per_Framework := (&models.Scatter{Name: `Scatter Stars per Framework`}).Stage(stage)

	// Declarations of staged instances of Serie
	__Serie__000000_FrameworkSerie := (&models.Serie{Name: `FrameworkSerie`}).Stage(stage)
	__Serie__000001_Released := (&models.Serie{Name: `Released`}).Stage(stage)
	__Serie__000002_Stars := (&models.Serie{Name: `Stars`}).Stage(stage)
	__Serie__000003_XSerie := (&models.Serie{Name: `XSerie`}).Stage(stage)
	__Serie__000004_YSerie := (&models.Serie{Name: `YSerie`}).Stage(stage)

	// Declarations of staged instances of Value
	__Value__000000_14_2 := (&models.Value{Name: `14.2`}).Stage(stage)
	__Value__000001_150 := (&models.Value{Name: `150`}).Stage(stage)
	__Value__000002_150793 := (&models.Value{Name: `150793`}).Stage(stage)
	__Value__000003_160 := (&models.Value{Name: `160`}).Stage(stage)
	__Value__000004_166443 := (&models.Value{Name: `166443`}).Stage(stage)
	__Value__000005_170 := (&models.Value{Name: `170`}).Stage(stage)
	__Value__000006_180 := (&models.Value{Name: `180`}).Stage(stage)
	__Value__000007_2010 := (&models.Value{Name: `2010`}).Stage(stage)
	__Value__000008_2011 := (&models.Value{Name: `2011`}).Stage(stage)
	__Value__000009_2013 := (&models.Value{Name: `2013`}).Stage(stage)
	__Value__000010_2014 := (&models.Value{Name: `2014`}).Stage(stage)
	__Value__000011_2016 := (&models.Value{Name: `2016`}).Stage(stage)
	__Value__000012_21471 := (&models.Value{Name: `21471`}).Stage(stage)
	__Value__000013_22_4 := (&models.Value{Name: `22.4`}).Stage(stage)
	__Value__000014_25_2 := (&models.Value{Name: `25.2`}).Stage(stage)
	__Value__000015_26_3 := (&models.Value{Name: `26.3`}).Stage(stage)
	__Value__000016_27647 := (&models.Value{Name: `27647`}).Stage(stage)
	__Value__000017_62342 := (&models.Value{Name: `62342`}).Stage(stage)
	__Value__000018_Angular := (&models.Value{Name: `Angular`}).Stage(stage)
	__Value__000019_Backbone := (&models.Value{Name: `Backbone`}).Stage(stage)
	__Value__000020_Ember := (&models.Value{Name: `Ember`}).Stage(stage)
	__Value__000021_React := (&models.Value{Name: `React`}).Stage(stage)
	__Value__000022_Vue := (&models.Value{Name: `Vue`}).Stage(stage)

	// Setup of values

	// Bar values setup
	__Bar__000000_Bar_Stars_per_Framework.Name = `Bar Stars per Framework`
	__Bar__000000_Bar_Stars_per_Framework.AutoDomainX = false
	__Bar__000000_Bar_Stars_per_Framework.XMin = 0.000000
	__Bar__000000_Bar_Stars_per_Framework.XMax = 0.000000
	__Bar__000000_Bar_Stars_per_Framework.AutoDomainY = false
	__Bar__000000_Bar_Stars_per_Framework.YMin = 0.000000
	__Bar__000000_Bar_Stars_per_Framework.YMax = 180000.000000
	__Bar__000000_Bar_Stars_per_Framework.YLabelPresent = false
	__Bar__000000_Bar_Stars_per_Framework.YLabelOffset = 0.000000
	__Bar__000000_Bar_Stars_per_Framework.XLabelPresent = false
	__Bar__000000_Bar_Stars_per_Framework.XLabelOffset = 0.000000
	__Bar__000000_Bar_Stars_per_Framework.Width = 750.000000
	__Bar__000000_Bar_Stars_per_Framework.Heigth = 500.000000
	__Bar__000000_Bar_Stars_per_Framework.Margin = 100.000000

	// Bar values setup
	__Bar__000001_Random_X_Y.Name = `Random X & Y`
	__Bar__000001_Random_X_Y.AutoDomainX = false
	__Bar__000001_Random_X_Y.XMin = 100.000000
	__Bar__000001_Random_X_Y.XMax = 230.000000
	__Bar__000001_Random_X_Y.AutoDomainY = false
	__Bar__000001_Random_X_Y.YMin = 15.000000
	__Bar__000001_Random_X_Y.YMax = 35.000000
	__Bar__000001_Random_X_Y.YLabelPresent = false
	__Bar__000001_Random_X_Y.YLabelOffset = 0.000000
	__Bar__000001_Random_X_Y.XLabelPresent = false
	__Bar__000001_Random_X_Y.XLabelOffset = 0.000000
	__Bar__000001_Random_X_Y.Width = 750.000000
	__Bar__000001_Random_X_Y.Heigth = 500.000000
	__Bar__000001_Random_X_Y.Margin = 50.000000

	// Key values setup
	__Key__000000_Framework_not_hardcoded.Name = `Framework not hardcoded`

	// Key values setup
	__Key__000001_Released_not_hardcoded.Name = `Released not hardcoded`

	// Key values setup
	__Key__000002_Stars_not_hardcoded.Name = `Stars not hardcoded`

	// Key values setup
	__Key__000003_X.Name = `X`

	// Key values setup
	__Key__000004_Y.Name = `Y`

	// Pie values setup
	__Pie__000000_Pie_Stars_per_Framework.Name = `Pie Stars per Framework`
	__Pie__000000_Pie_Stars_per_Framework.Width = 750.000000
	__Pie__000000_Pie_Stars_per_Framework.Heigth = 600.000000
	__Pie__000000_Pie_Stars_per_Framework.Margin = 50.000000

	// Scatter values setup
	__Scatter__000000_Scatter_Stars_per_Framework.Name = `Scatter Stars per Framework`
	__Scatter__000000_Scatter_Stars_per_Framework.Width = 750.000000
	__Scatter__000000_Scatter_Stars_per_Framework.Heigth = 400.000000
	__Scatter__000000_Scatter_Stars_per_Framework.Margin = 50.000000

	// Serie values setup
	__Serie__000000_FrameworkSerie.Name = `FrameworkSerie`

	// Serie values setup
	__Serie__000001_Released.Name = `Released`

	// Serie values setup
	__Serie__000002_Stars.Name = `Stars`

	// Serie values setup
	__Serie__000003_XSerie.Name = `XSerie`

	// Serie values setup
	__Serie__000004_YSerie.Name = `YSerie`

	// Value values setup
	__Value__000000_14_2.Name = `14.2`

	// Value values setup
	__Value__000001_150.Name = `150`

	// Value values setup
	__Value__000002_150793.Name = `150793`

	// Value values setup
	__Value__000003_160.Name = `160`

	// Value values setup
	__Value__000004_166443.Name = `166443`

	// Value values setup
	__Value__000005_170.Name = `170`

	// Value values setup
	__Value__000006_180.Name = `180`

	// Value values setup
	__Value__000007_2010.Name = `2010`

	// Value values setup
	__Value__000008_2011.Name = `2011`

	// Value values setup
	__Value__000009_2013.Name = `2013`

	// Value values setup
	__Value__000010_2014.Name = `2014`

	// Value values setup
	__Value__000011_2016.Name = `2016`

	// Value values setup
	__Value__000012_21471.Name = `21471`

	// Value values setup
	__Value__000013_22_4.Name = `22.4`

	// Value values setup
	__Value__000014_25_2.Name = `25.2`

	// Value values setup
	__Value__000015_26_3.Name = `26.3`

	// Value values setup
	__Value__000016_27647.Name = `27647`

	// Value values setup
	__Value__000017_62342.Name = `62342`

	// Value values setup
	__Value__000018_Angular.Name = `Angular`

	// Value values setup
	__Value__000019_Backbone.Name = `Backbone`

	// Value values setup
	__Value__000020_Ember.Name = `Ember`

	// Value values setup
	__Value__000021_React.Name = `React`

	// Value values setup
	__Value__000022_Vue.Name = `Vue`

	// Setup of pointers
	__Bar__000000_Bar_Stars_per_Framework.X = __Key__000000_Framework_not_hardcoded
	__Bar__000000_Bar_Stars_per_Framework.Y = __Key__000002_Stars_not_hardcoded
	__Bar__000000_Bar_Stars_per_Framework.Set = append(__Bar__000000_Bar_Stars_per_Framework.Set, __Serie__000002_Stars)
	__Bar__000000_Bar_Stars_per_Framework.Set = append(__Bar__000000_Bar_Stars_per_Framework.Set, __Serie__000001_Released)
	__Bar__000000_Bar_Stars_per_Framework.Set = append(__Bar__000000_Bar_Stars_per_Framework.Set, __Serie__000000_FrameworkSerie)
	__Bar__000001_Random_X_Y.X = __Key__000003_X
	__Bar__000001_Random_X_Y.Y = __Key__000004_Y
	__Bar__000001_Random_X_Y.Set = append(__Bar__000001_Random_X_Y.Set, __Serie__000003_XSerie)
	__Bar__000001_Random_X_Y.Set = append(__Bar__000001_Random_X_Y.Set, __Serie__000004_YSerie)
	__Pie__000000_Pie_Stars_per_Framework.X = __Key__000000_Framework_not_hardcoded
	__Pie__000000_Pie_Stars_per_Framework.Y = __Key__000002_Stars_not_hardcoded
	__Pie__000000_Pie_Stars_per_Framework.Set = append(__Pie__000000_Pie_Stars_per_Framework.Set, __Serie__000001_Released)
	__Pie__000000_Pie_Stars_per_Framework.Set = append(__Pie__000000_Pie_Stars_per_Framework.Set, __Serie__000002_Stars)
	__Pie__000000_Pie_Stars_per_Framework.Set = append(__Pie__000000_Pie_Stars_per_Framework.Set, __Serie__000000_FrameworkSerie)
	__Scatter__000000_Scatter_Stars_per_Framework.X = __Key__000001_Released_not_hardcoded
	__Scatter__000000_Scatter_Stars_per_Framework.Y = __Key__000002_Stars_not_hardcoded
	__Scatter__000000_Scatter_Stars_per_Framework.Text = __Key__000000_Framework_not_hardcoded
	__Scatter__000000_Scatter_Stars_per_Framework.Set = append(__Scatter__000000_Scatter_Stars_per_Framework.Set, __Serie__000002_Stars)
	__Scatter__000000_Scatter_Stars_per_Framework.Set = append(__Scatter__000000_Scatter_Stars_per_Framework.Set, __Serie__000000_FrameworkSerie)
	__Scatter__000000_Scatter_Stars_per_Framework.Set = append(__Scatter__000000_Scatter_Stars_per_Framework.Set, __Serie__000001_Released)
	__Serie__000000_FrameworkSerie.Key = __Key__000000_Framework_not_hardcoded
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000022_Vue)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000021_React)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000018_Angular)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000019_Backbone)
	__Serie__000000_FrameworkSerie.Values = append(__Serie__000000_FrameworkSerie.Values, __Value__000020_Ember)
	__Serie__000001_Released.Key = __Key__000001_Released_not_hardcoded
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000010_2014)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000009_2013)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000011_2016)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000007_2010)
	__Serie__000001_Released.Values = append(__Serie__000001_Released.Values, __Value__000008_2011)
	__Serie__000002_Stars.Key = __Key__000002_Stars_not_hardcoded
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000004_166443)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000002_150793)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000017_62342)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000016_27647)
	__Serie__000002_Stars.Values = append(__Serie__000002_Stars.Values, __Value__000012_21471)
	__Serie__000003_XSerie.Key = __Key__000003_X
	__Serie__000003_XSerie.Values = append(__Serie__000003_XSerie.Values, __Value__000001_150)
	__Serie__000003_XSerie.Values = append(__Serie__000003_XSerie.Values, __Value__000005_170)
	__Serie__000003_XSerie.Values = append(__Serie__000003_XSerie.Values, __Value__000003_160)
	__Serie__000003_XSerie.Values = append(__Serie__000003_XSerie.Values, __Value__000006_180)
	__Serie__000004_YSerie.Key = __Key__000004_Y
	__Serie__000004_YSerie.Values = append(__Serie__000004_YSerie.Values, __Value__000014_25_2)
	__Serie__000004_YSerie.Values = append(__Serie__000004_YSerie.Values, __Value__000015_26_3)
	__Serie__000004_YSerie.Values = append(__Serie__000004_YSerie.Values, __Value__000013_22_4)
	__Serie__000004_YSerie.Values = append(__Serie__000004_YSerie.Values, __Value__000000_14_2)
}
