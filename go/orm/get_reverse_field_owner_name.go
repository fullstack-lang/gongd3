// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongd3/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Pie:
		tmp := GetInstanceDBFromInstance[models.Pie, PieDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Scatter:
		tmp := GetInstanceDBFromInstance[models.Scatter, ScatterDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Serie:
		tmp := GetInstanceDBFromInstance[models.Serie, SerieDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			case "Set":
				if _bar, ok := stage.Bar_Set_reverseMap[inst]; ok {
					res = _bar.Name
				}
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			case "Set":
				if _pie, ok := stage.Pie_Set_reverseMap[inst]; ok {
					res = _pie.Name
				}
			}
		case "Scatter":
			switch reverseField.Fieldname {
			case "Set":
				if _scatter, ok := stage.Scatter_Set_reverseMap[inst]; ok {
					res = _scatter.Name
				}
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Value:
		tmp := GetInstanceDBFromInstance[models.Value, ValueDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			case "Values":
				if _serie, ok := stage.Serie_Values_reverseMap[inst]; ok {
					res = _serie.Name
				}
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Pie:
		tmp := GetInstanceDBFromInstance[models.Pie, PieDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Scatter:
		tmp := GetInstanceDBFromInstance[models.Scatter, ScatterDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Serie:
		tmp := GetInstanceDBFromInstance[models.Serie, SerieDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			case "Set":
				res = stage.Bar_Set_reverseMap[inst]
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			case "Set":
				res = stage.Pie_Set_reverseMap[inst]
			}
		case "Scatter":
			switch reverseField.Fieldname {
			case "Set":
				res = stage.Scatter_Set_reverseMap[inst]
			}
		case "Serie":
			switch reverseField.Fieldname {
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	case *models.Value:
		tmp := GetInstanceDBFromInstance[models.Value, ValueDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			}
		case "Scatter":
			switch reverseField.Fieldname {
			}
		case "Serie":
			switch reverseField.Fieldname {
			case "Values":
				res = stage.Serie_Values_reverseMap[inst]
			}
		case "Value":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
