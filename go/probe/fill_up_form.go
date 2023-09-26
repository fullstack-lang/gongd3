// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongd3/go/models"
	"github.com/fullstack-lang/gongd3/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Bar:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("X", instanceWithInferedType.X, formGroup, playground)
		AssociationFieldToForm("Y", instanceWithInferedType.Y, formGroup, playground)
		AssociationSliceToForm("Set", instanceWithInferedType, &instanceWithInferedType.Set, formGroup, playground)
		BasicFieldtoForm("AutoDomainX", instanceWithInferedType.AutoDomainX, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("XMin", instanceWithInferedType.XMin, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("XMax", instanceWithInferedType.XMax, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("AutoDomainY", instanceWithInferedType.AutoDomainY, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("YMin", instanceWithInferedType.YMin, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("YMax", instanceWithInferedType.YMax, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("YLabelPresent", instanceWithInferedType.YLabelPresent, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("YLabelOffset", instanceWithInferedType.YLabelOffset, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("XLabelPresent", instanceWithInferedType.XLabelPresent, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("XLabelOffset", instanceWithInferedType.XLabelOffset, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Margin", instanceWithInferedType.Margin, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Pie:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("X", instanceWithInferedType.X, formGroup, playground)
		AssociationFieldToForm("Y", instanceWithInferedType.Y, formGroup, playground)
		AssociationSliceToForm("Set", instanceWithInferedType, &instanceWithInferedType.Set, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Margin", instanceWithInferedType.Margin, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Scatter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("X", instanceWithInferedType.X, formGroup, playground)
		AssociationFieldToForm("Y", instanceWithInferedType.Y, formGroup, playground)
		AssociationFieldToForm("Text", instanceWithInferedType.Text, formGroup, playground)
		AssociationSliceToForm("Set", instanceWithInferedType, &instanceWithInferedType.Set, formGroup, playground)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Margin", instanceWithInferedType.Margin, instanceWithInferedType, playground.formStage, formGroup, false)

	case *models.Serie:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationFieldToForm("Key", instanceWithInferedType.Key, formGroup, playground)
		AssociationSliceToForm("Values", instanceWithInferedType, &instanceWithInferedType.Values, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Bar"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Bar),
					"Set",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Bar, *models.Serie](
					nil,
					"Set",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Pie"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Pie),
					"Set",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Pie, *models.Serie](
					nil,
					"Set",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Scatter"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Scatter),
					"Set",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Scatter, *models.Serie](
					nil,
					"Set",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Value:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Serie"
			rf.Fieldname = "Values"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Serie),
					"Values",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Serie, *models.Value](
					nil,
					"Values",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
