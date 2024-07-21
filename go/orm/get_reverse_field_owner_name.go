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
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pie:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scatter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Serie:
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			case "Set":
				if _bar, ok := stage.Bar_Set_reverseMap[inst]; ok {
					res = _bar.Name
				}
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
		}

	case *models.Value:
		switch reverseField.GongstructName {
		// insertion point
		case "Serie":
			switch reverseField.Fieldname {
			case "Values":
				if _serie, ok := stage.Serie_Values_reverseMap[inst]; ok {
					res = _serie.Name
				}
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
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pie:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scatter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Serie:
		switch reverseField.GongstructName {
		// insertion point
		case "Bar":
			switch reverseField.Fieldname {
			case "Set":
				res = stage.Bar_Set_reverseMap[inst]
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
		}

	case *models.Value:
		switch reverseField.GongstructName {
		// insertion point
		case "Serie":
			switch reverseField.Fieldname {
			case "Values":
				res = stage.Serie_Values_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
