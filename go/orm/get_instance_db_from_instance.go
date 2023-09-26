// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongd3/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | BarDB | KeyDB | PieDB | ScatterDB | SerieDB | ValueDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Bar:
		barInstance := any(concreteInstance).(*models.Bar)
		ret2 := backRepo.BackRepoBar.GetBarDBFromBarPtr(barInstance)
		ret = any(ret2).(*T2)
	case *models.Key:
		keyInstance := any(concreteInstance).(*models.Key)
		ret2 := backRepo.BackRepoKey.GetKeyDBFromKeyPtr(keyInstance)
		ret = any(ret2).(*T2)
	case *models.Pie:
		pieInstance := any(concreteInstance).(*models.Pie)
		ret2 := backRepo.BackRepoPie.GetPieDBFromPiePtr(pieInstance)
		ret = any(ret2).(*T2)
	case *models.Scatter:
		scatterInstance := any(concreteInstance).(*models.Scatter)
		ret2 := backRepo.BackRepoScatter.GetScatterDBFromScatterPtr(scatterInstance)
		ret = any(ret2).(*T2)
	case *models.Serie:
		serieInstance := any(concreteInstance).(*models.Serie)
		ret2 := backRepo.BackRepoSerie.GetSerieDBFromSeriePtr(serieInstance)
		ret = any(ret2).(*T2)
	case *models.Value:
		valueInstance := any(concreteInstance).(*models.Value)
		ret2 := backRepo.BackRepoValue.GetValueDBFromValuePtr(valueInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pie:
		tmp := GetInstanceDBFromInstance[models.Pie, PieDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scatter:
		tmp := GetInstanceDBFromInstance[models.Scatter, ScatterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Serie:
		tmp := GetInstanceDBFromInstance[models.Serie, SerieDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Value:
		tmp := GetInstanceDBFromInstance[models.Value, ValueDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pie:
		tmp := GetInstanceDBFromInstance[models.Pie, PieDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scatter:
		tmp := GetInstanceDBFromInstance[models.Scatter, ScatterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Serie:
		tmp := GetInstanceDBFromInstance[models.Serie, SerieDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Value:
		tmp := GetInstanceDBFromInstance[models.Value, ValueDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
