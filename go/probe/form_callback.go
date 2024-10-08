// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongd3/go/models"
	"github.com/fullstack-lang/gongd3/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__BarFormCallback(
	bar *models.Bar,
	probe *Probe,
	formGroup *table.FormGroup,
) (barFormCallback *BarFormCallback) {
	barFormCallback = new(BarFormCallback)
	barFormCallback.probe = probe
	barFormCallback.bar = bar
	barFormCallback.formGroup = formGroup

	barFormCallback.CreationMode = (bar == nil)

	return
}

type BarFormCallback struct {
	bar *models.Bar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (barFormCallback *BarFormCallback) OnSave() {

	log.Println("BarFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	barFormCallback.probe.formStage.Checkout()

	if barFormCallback.bar == nil {
		barFormCallback.bar = new(models.Bar).Stage(barFormCallback.probe.stageOfInterest)
	}
	bar_ := barFormCallback.bar
	_ = bar_

	for _, formDiv := range barFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bar_.Name), formDiv)
		case "X":
			FormDivSelectFieldToField(&(bar_.X), barFormCallback.probe.stageOfInterest, formDiv)
		case "Y":
			FormDivSelectFieldToField(&(bar_.Y), barFormCallback.probe.stageOfInterest, formDiv)
		case "AutoDomainX":
			FormDivBasicFieldToField(&(bar_.AutoDomainX), formDiv)
		case "XMin":
			FormDivBasicFieldToField(&(bar_.XMin), formDiv)
		case "XMax":
			FormDivBasicFieldToField(&(bar_.XMax), formDiv)
		case "AutoDomainY":
			FormDivBasicFieldToField(&(bar_.AutoDomainY), formDiv)
		case "YMin":
			FormDivBasicFieldToField(&(bar_.YMin), formDiv)
		case "YMax":
			FormDivBasicFieldToField(&(bar_.YMax), formDiv)
		case "YLabelPresent":
			FormDivBasicFieldToField(&(bar_.YLabelPresent), formDiv)
		case "YLabelOffset":
			FormDivBasicFieldToField(&(bar_.YLabelOffset), formDiv)
		case "XLabelPresent":
			FormDivBasicFieldToField(&(bar_.XLabelPresent), formDiv)
		case "XLabelOffset":
			FormDivBasicFieldToField(&(bar_.XLabelOffset), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(bar_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(bar_.Heigth), formDiv)
		case "Margin":
			FormDivBasicFieldToField(&(bar_.Margin), formDiv)
		}
	}

	// manage the suppress operation
	if barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		bar_.Unstage(barFormCallback.probe.stageOfInterest)
	}

	barFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Bar](
		barFormCallback.probe,
	)
	barFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if barFormCallback.CreationMode || barFormCallback.formGroup.HasSuppressButtonBeenPressed {
		barFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(barFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__BarFormCallback(
			nil,
			barFormCallback.probe,
			newFormGroup,
		)
		bar := new(models.Bar)
		FillUpForm(bar, newFormGroup, barFormCallback.probe)
		barFormCallback.probe.formStage.Commit()
	}

	fillUpTree(barFormCallback.probe)
}
func __gong__New__KeyFormCallback(
	key *models.Key,
	probe *Probe,
	formGroup *table.FormGroup,
) (keyFormCallback *KeyFormCallback) {
	keyFormCallback = new(KeyFormCallback)
	keyFormCallback.probe = probe
	keyFormCallback.key = key
	keyFormCallback.formGroup = formGroup

	keyFormCallback.CreationMode = (key == nil)

	return
}

type KeyFormCallback struct {
	key *models.Key

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (keyFormCallback *KeyFormCallback) OnSave() {

	log.Println("KeyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	keyFormCallback.probe.formStage.Checkout()

	if keyFormCallback.key == nil {
		keyFormCallback.key = new(models.Key).Stage(keyFormCallback.probe.stageOfInterest)
	}
	key_ := keyFormCallback.key
	_ = key_

	for _, formDiv := range keyFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(key_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		key_.Unstage(keyFormCallback.probe.stageOfInterest)
	}

	keyFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Key](
		keyFormCallback.probe,
	)
	keyFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if keyFormCallback.CreationMode || keyFormCallback.formGroup.HasSuppressButtonBeenPressed {
		keyFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(keyFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__KeyFormCallback(
			nil,
			keyFormCallback.probe,
			newFormGroup,
		)
		key := new(models.Key)
		FillUpForm(key, newFormGroup, keyFormCallback.probe)
		keyFormCallback.probe.formStage.Commit()
	}

	fillUpTree(keyFormCallback.probe)
}
func __gong__New__PieFormCallback(
	pie *models.Pie,
	probe *Probe,
	formGroup *table.FormGroup,
) (pieFormCallback *PieFormCallback) {
	pieFormCallback = new(PieFormCallback)
	pieFormCallback.probe = probe
	pieFormCallback.pie = pie
	pieFormCallback.formGroup = formGroup

	pieFormCallback.CreationMode = (pie == nil)

	return
}

type PieFormCallback struct {
	pie *models.Pie

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (pieFormCallback *PieFormCallback) OnSave() {

	log.Println("PieFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pieFormCallback.probe.formStage.Checkout()

	if pieFormCallback.pie == nil {
		pieFormCallback.pie = new(models.Pie).Stage(pieFormCallback.probe.stageOfInterest)
	}
	pie_ := pieFormCallback.pie
	_ = pie_

	for _, formDiv := range pieFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pie_.Name), formDiv)
		case "X":
			FormDivSelectFieldToField(&(pie_.X), pieFormCallback.probe.stageOfInterest, formDiv)
		case "Y":
			FormDivSelectFieldToField(&(pie_.Y), pieFormCallback.probe.stageOfInterest, formDiv)
		case "Width":
			FormDivBasicFieldToField(&(pie_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(pie_.Heigth), formDiv)
		case "Margin":
			FormDivBasicFieldToField(&(pie_.Margin), formDiv)
		}
	}

	// manage the suppress operation
	if pieFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pie_.Unstage(pieFormCallback.probe.stageOfInterest)
	}

	pieFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Pie](
		pieFormCallback.probe,
	)
	pieFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if pieFormCallback.CreationMode || pieFormCallback.formGroup.HasSuppressButtonBeenPressed {
		pieFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(pieFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PieFormCallback(
			nil,
			pieFormCallback.probe,
			newFormGroup,
		)
		pie := new(models.Pie)
		FillUpForm(pie, newFormGroup, pieFormCallback.probe)
		pieFormCallback.probe.formStage.Commit()
	}

	fillUpTree(pieFormCallback.probe)
}
func __gong__New__ScatterFormCallback(
	scatter *models.Scatter,
	probe *Probe,
	formGroup *table.FormGroup,
) (scatterFormCallback *ScatterFormCallback) {
	scatterFormCallback = new(ScatterFormCallback)
	scatterFormCallback.probe = probe
	scatterFormCallback.scatter = scatter
	scatterFormCallback.formGroup = formGroup

	scatterFormCallback.CreationMode = (scatter == nil)

	return
}

type ScatterFormCallback struct {
	scatter *models.Scatter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (scatterFormCallback *ScatterFormCallback) OnSave() {

	log.Println("ScatterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	scatterFormCallback.probe.formStage.Checkout()

	if scatterFormCallback.scatter == nil {
		scatterFormCallback.scatter = new(models.Scatter).Stage(scatterFormCallback.probe.stageOfInterest)
	}
	scatter_ := scatterFormCallback.scatter
	_ = scatter_

	for _, formDiv := range scatterFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(scatter_.Name), formDiv)
		case "X":
			FormDivSelectFieldToField(&(scatter_.X), scatterFormCallback.probe.stageOfInterest, formDiv)
		case "Y":
			FormDivSelectFieldToField(&(scatter_.Y), scatterFormCallback.probe.stageOfInterest, formDiv)
		case "Text":
			FormDivSelectFieldToField(&(scatter_.Text), scatterFormCallback.probe.stageOfInterest, formDiv)
		case "Width":
			FormDivBasicFieldToField(&(scatter_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(scatter_.Heigth), formDiv)
		case "Margin":
			FormDivBasicFieldToField(&(scatter_.Margin), formDiv)
		}
	}

	// manage the suppress operation
	if scatterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scatter_.Unstage(scatterFormCallback.probe.stageOfInterest)
	}

	scatterFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Scatter](
		scatterFormCallback.probe,
	)
	scatterFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if scatterFormCallback.CreationMode || scatterFormCallback.formGroup.HasSuppressButtonBeenPressed {
		scatterFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(scatterFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ScatterFormCallback(
			nil,
			scatterFormCallback.probe,
			newFormGroup,
		)
		scatter := new(models.Scatter)
		FillUpForm(scatter, newFormGroup, scatterFormCallback.probe)
		scatterFormCallback.probe.formStage.Commit()
	}

	fillUpTree(scatterFormCallback.probe)
}
func __gong__New__SerieFormCallback(
	serie *models.Serie,
	probe *Probe,
	formGroup *table.FormGroup,
) (serieFormCallback *SerieFormCallback) {
	serieFormCallback = new(SerieFormCallback)
	serieFormCallback.probe = probe
	serieFormCallback.serie = serie
	serieFormCallback.formGroup = formGroup

	serieFormCallback.CreationMode = (serie == nil)

	return
}

type SerieFormCallback struct {
	serie *models.Serie

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (serieFormCallback *SerieFormCallback) OnSave() {

	log.Println("SerieFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	serieFormCallback.probe.formStage.Checkout()

	if serieFormCallback.serie == nil {
		serieFormCallback.serie = new(models.Serie).Stage(serieFormCallback.probe.stageOfInterest)
	}
	serie_ := serieFormCallback.serie
	_ = serie_

	for _, formDiv := range serieFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(serie_.Name), formDiv)
		case "Key":
			FormDivSelectFieldToField(&(serie_.Key), serieFormCallback.probe.stageOfInterest, formDiv)
		case "Bar:Set":
			// we need to retrieve the field owner before the change
			var pastBarOwner *models.Bar
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Bar"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				serieFormCallback.probe.stageOfInterest,
				serieFormCallback.probe.backRepoOfInterest,
				serie_,
				&rf)

			if reverseFieldOwner != nil {
				pastBarOwner = reverseFieldOwner.(*models.Bar)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastBarOwner != nil {
					idx := slices.Index(pastBarOwner.Set, serie_)
					pastBarOwner.Set = slices.Delete(pastBarOwner.Set, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _bar := range *models.GetGongstructInstancesSet[models.Bar](serieFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _bar.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newBarOwner := _bar // we have a match
						if pastBarOwner != nil {
							if newBarOwner != pastBarOwner {
								idx := slices.Index(pastBarOwner.Set, serie_)
								pastBarOwner.Set = slices.Delete(pastBarOwner.Set, idx, idx+1)
								newBarOwner.Set = append(newBarOwner.Set, serie_)
							}
						} else {
							newBarOwner.Set = append(newBarOwner.Set, serie_)
						}
					}
				}
			}
		case "Pie:Set":
			// we need to retrieve the field owner before the change
			var pastPieOwner *models.Pie
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Pie"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				serieFormCallback.probe.stageOfInterest,
				serieFormCallback.probe.backRepoOfInterest,
				serie_,
				&rf)

			if reverseFieldOwner != nil {
				pastPieOwner = reverseFieldOwner.(*models.Pie)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastPieOwner != nil {
					idx := slices.Index(pastPieOwner.Set, serie_)
					pastPieOwner.Set = slices.Delete(pastPieOwner.Set, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _pie := range *models.GetGongstructInstancesSet[models.Pie](serieFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _pie.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newPieOwner := _pie // we have a match
						if pastPieOwner != nil {
							if newPieOwner != pastPieOwner {
								idx := slices.Index(pastPieOwner.Set, serie_)
								pastPieOwner.Set = slices.Delete(pastPieOwner.Set, idx, idx+1)
								newPieOwner.Set = append(newPieOwner.Set, serie_)
							}
						} else {
							newPieOwner.Set = append(newPieOwner.Set, serie_)
						}
					}
				}
			}
		case "Scatter:Set":
			// we need to retrieve the field owner before the change
			var pastScatterOwner *models.Scatter
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Scatter"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				serieFormCallback.probe.stageOfInterest,
				serieFormCallback.probe.backRepoOfInterest,
				serie_,
				&rf)

			if reverseFieldOwner != nil {
				pastScatterOwner = reverseFieldOwner.(*models.Scatter)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastScatterOwner != nil {
					idx := slices.Index(pastScatterOwner.Set, serie_)
					pastScatterOwner.Set = slices.Delete(pastScatterOwner.Set, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _scatter := range *models.GetGongstructInstancesSet[models.Scatter](serieFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _scatter.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newScatterOwner := _scatter // we have a match
						if pastScatterOwner != nil {
							if newScatterOwner != pastScatterOwner {
								idx := slices.Index(pastScatterOwner.Set, serie_)
								pastScatterOwner.Set = slices.Delete(pastScatterOwner.Set, idx, idx+1)
								newScatterOwner.Set = append(newScatterOwner.Set, serie_)
							}
						} else {
							newScatterOwner.Set = append(newScatterOwner.Set, serie_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if serieFormCallback.formGroup.HasSuppressButtonBeenPressed {
		serie_.Unstage(serieFormCallback.probe.stageOfInterest)
	}

	serieFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Serie](
		serieFormCallback.probe,
	)
	serieFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if serieFormCallback.CreationMode || serieFormCallback.formGroup.HasSuppressButtonBeenPressed {
		serieFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(serieFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__SerieFormCallback(
			nil,
			serieFormCallback.probe,
			newFormGroup,
		)
		serie := new(models.Serie)
		FillUpForm(serie, newFormGroup, serieFormCallback.probe)
		serieFormCallback.probe.formStage.Commit()
	}

	fillUpTree(serieFormCallback.probe)
}
func __gong__New__ValueFormCallback(
	value *models.Value,
	probe *Probe,
	formGroup *table.FormGroup,
) (valueFormCallback *ValueFormCallback) {
	valueFormCallback = new(ValueFormCallback)
	valueFormCallback.probe = probe
	valueFormCallback.value = value
	valueFormCallback.formGroup = formGroup

	valueFormCallback.CreationMode = (value == nil)

	return
}

type ValueFormCallback struct {
	value *models.Value

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (valueFormCallback *ValueFormCallback) OnSave() {

	log.Println("ValueFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	valueFormCallback.probe.formStage.Checkout()

	if valueFormCallback.value == nil {
		valueFormCallback.value = new(models.Value).Stage(valueFormCallback.probe.stageOfInterest)
	}
	value_ := valueFormCallback.value
	_ = value_

	for _, formDiv := range valueFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(value_.Name), formDiv)
		case "Serie:Values":
			// we need to retrieve the field owner before the change
			var pastSerieOwner *models.Serie
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Serie"
			rf.Fieldname = "Values"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				valueFormCallback.probe.stageOfInterest,
				valueFormCallback.probe.backRepoOfInterest,
				value_,
				&rf)

			if reverseFieldOwner != nil {
				pastSerieOwner = reverseFieldOwner.(*models.Serie)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSerieOwner != nil {
					idx := slices.Index(pastSerieOwner.Values, value_)
					pastSerieOwner.Values = slices.Delete(pastSerieOwner.Values, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _serie := range *models.GetGongstructInstancesSet[models.Serie](valueFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _serie.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSerieOwner := _serie // we have a match
						if pastSerieOwner != nil {
							if newSerieOwner != pastSerieOwner {
								idx := slices.Index(pastSerieOwner.Values, value_)
								pastSerieOwner.Values = slices.Delete(pastSerieOwner.Values, idx, idx+1)
								newSerieOwner.Values = append(newSerieOwner.Values, value_)
							}
						} else {
							newSerieOwner.Values = append(newSerieOwner.Values, value_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		value_.Unstage(valueFormCallback.probe.stageOfInterest)
	}

	valueFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Value](
		valueFormCallback.probe,
	)
	valueFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if valueFormCallback.CreationMode || valueFormCallback.formGroup.HasSuppressButtonBeenPressed {
		valueFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(valueFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ValueFormCallback(
			nil,
			valueFormCallback.probe,
			newFormGroup,
		)
		value := new(models.Value)
		FillUpForm(value, newFormGroup, valueFormCallback.probe)
		valueFormCallback.probe.formStage.Commit()
	}

	fillUpTree(valueFormCallback.probe)
}
