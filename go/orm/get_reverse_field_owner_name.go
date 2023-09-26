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
				if tmp != nil && tmp.Bar_SetDBID.Int64 != 0 {
					id := uint(tmp.Bar_SetDBID.Int64)
					reservePointerTarget := backRepo.BackRepoBar.Map_BarDBID_BarPtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			case "Set":
				if tmp != nil && tmp.Pie_SetDBID.Int64 != 0 {
					id := uint(tmp.Pie_SetDBID.Int64)
					reservePointerTarget := backRepo.BackRepoPie.Map_PieDBID_PiePtr[id]
					res = reservePointerTarget.Name
				}
			}
		case "Scatter":
			switch reverseField.Fieldname {
			case "Set":
				if tmp != nil && tmp.Scatter_SetDBID.Int64 != 0 {
					id := uint(tmp.Scatter_SetDBID.Int64)
					reservePointerTarget := backRepo.BackRepoScatter.Map_ScatterDBID_ScatterPtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.Serie_ValuesDBID.Int64 != 0 {
					id := uint(tmp.Serie_ValuesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoSerie.Map_SerieDBID_SeriePtr[id]
					res = reservePointerTarget.Name
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
				if tmp != nil && tmp.Bar_SetDBID.Int64 != 0 {
					id := uint(tmp.Bar_SetDBID.Int64)
					reservePointerTarget := backRepo.BackRepoBar.Map_BarDBID_BarPtr[id]
					res = reservePointerTarget
				}
			}
		case "Key":
			switch reverseField.Fieldname {
			}
		case "Pie":
			switch reverseField.Fieldname {
			case "Set":
				if tmp != nil && tmp.Pie_SetDBID.Int64 != 0 {
					id := uint(tmp.Pie_SetDBID.Int64)
					reservePointerTarget := backRepo.BackRepoPie.Map_PieDBID_PiePtr[id]
					res = reservePointerTarget
				}
			}
		case "Scatter":
			switch reverseField.Fieldname {
			case "Set":
				if tmp != nil && tmp.Scatter_SetDBID.Int64 != 0 {
					id := uint(tmp.Scatter_SetDBID.Int64)
					reservePointerTarget := backRepo.BackRepoScatter.Map_ScatterDBID_ScatterPtr[id]
					res = reservePointerTarget
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
				if tmp != nil && tmp.Serie_ValuesDBID.Int64 != 0 {
					id := uint(tmp.Serie_ValuesDBID.Int64)
					reservePointerTarget := backRepo.BackRepoSerie.Map_SerieDBID_SeriePtr[id]
					res = reservePointerTarget
				}
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
