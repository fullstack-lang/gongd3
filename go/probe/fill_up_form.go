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
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Bar:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("X", instanceWithInferedType.X, formGroup, probe)
		AssociationFieldToForm("Y", instanceWithInferedType.Y, formGroup, probe)
		AssociationSliceToForm("Set", instanceWithInferedType, &instanceWithInferedType.Set, formGroup, probe)
		BasicFieldtoForm("AutoDomainX", instanceWithInferedType.AutoDomainX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XMin", instanceWithInferedType.XMin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XMax", instanceWithInferedType.XMax, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AutoDomainY", instanceWithInferedType.AutoDomainY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YMin", instanceWithInferedType.YMin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YMax", instanceWithInferedType.YMax, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YLabelPresent", instanceWithInferedType.YLabelPresent, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("YLabelOffset", instanceWithInferedType.YLabelOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XLabelPresent", instanceWithInferedType.XLabelPresent, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("XLabelOffset", instanceWithInferedType.XLabelOffset, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Margin", instanceWithInferedType.Margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Pie:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("X", instanceWithInferedType.X, formGroup, probe)
		AssociationFieldToForm("Y", instanceWithInferedType.Y, formGroup, probe)
		AssociationSliceToForm("Set", instanceWithInferedType, &instanceWithInferedType.Set, formGroup, probe)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Margin", instanceWithInferedType.Margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Scatter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("X", instanceWithInferedType.X, formGroup, probe)
		AssociationFieldToForm("Y", instanceWithInferedType.Y, formGroup, probe)
		AssociationFieldToForm("Text", instanceWithInferedType.Text, formGroup, probe)
		AssociationSliceToForm("Set", instanceWithInferedType, &instanceWithInferedType.Set, formGroup, probe)
		BasicFieldtoForm("Width", instanceWithInferedType.Width, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Heigth", instanceWithInferedType.Heigth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Margin", instanceWithInferedType.Margin, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Serie:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Key", instanceWithInferedType.Key, formGroup, probe)
		AssociationSliceToForm("Values", instanceWithInferedType, &instanceWithInferedType.Values, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Bar"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Bar),
					"Set",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Bar, *models.Serie](
					nil,
					"Set",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Pie"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Pie),
					"Set",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Pie, *models.Serie](
					nil,
					"Set",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Scatter"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Scatter),
					"Set",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Scatter, *models.Serie](
					nil,
					"Set",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Value:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Serie"
			rf.Fieldname = "Values"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Serie),
					"Values",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Serie, *models.Value](
					nil,
					"Values",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
