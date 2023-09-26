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
	playground *Playground,
) (barFormCallback *BarFormCallback) {
	barFormCallback = new(BarFormCallback)
	barFormCallback.playground = playground
	barFormCallback.bar = bar

	barFormCallback.CreationMode = (bar == nil)

	return
}

type BarFormCallback struct {
	bar *models.Bar

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (barFormCallback *BarFormCallback) OnSave() {

	log.Println("BarFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	barFormCallback.playground.formStage.Checkout()

	if barFormCallback.bar == nil {
		barFormCallback.bar = new(models.Bar).Stage(barFormCallback.playground.stageOfInterest)
	}
	bar_ := barFormCallback.bar
	_ = bar_

	// get the formGroup
	formGroup := barFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(bar_.Name), formDiv)
		case "X":
			FormDivSelectFieldToField(&(bar_.X), barFormCallback.playground.stageOfInterest, formDiv)
		case "Y":
			FormDivSelectFieldToField(&(bar_.Y), barFormCallback.playground.stageOfInterest, formDiv)
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

	barFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Bar](
		barFormCallback.playground,
	)
	barFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if barFormCallback.CreationMode {
		barFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__BarFormCallback(
				nil,
				barFormCallback.playground,
			),
		}).Stage(barFormCallback.playground.formStage)
		bar := new(models.Bar)
		FillUpForm(bar, newFormGroup, barFormCallback.playground)
		barFormCallback.playground.formStage.Commit()
	}

	fillUpTree(barFormCallback.playground)
}
func __gong__New__KeyFormCallback(
	key *models.Key,
	playground *Playground,
) (keyFormCallback *KeyFormCallback) {
	keyFormCallback = new(KeyFormCallback)
	keyFormCallback.playground = playground
	keyFormCallback.key = key

	keyFormCallback.CreationMode = (key == nil)

	return
}

type KeyFormCallback struct {
	key *models.Key

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (keyFormCallback *KeyFormCallback) OnSave() {

	log.Println("KeyFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	keyFormCallback.playground.formStage.Checkout()

	if keyFormCallback.key == nil {
		keyFormCallback.key = new(models.Key).Stage(keyFormCallback.playground.stageOfInterest)
	}
	key_ := keyFormCallback.key
	_ = key_

	// get the formGroup
	formGroup := keyFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(key_.Name), formDiv)
		}
	}

	keyFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Key](
		keyFormCallback.playground,
	)
	keyFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if keyFormCallback.CreationMode {
		keyFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__KeyFormCallback(
				nil,
				keyFormCallback.playground,
			),
		}).Stage(keyFormCallback.playground.formStage)
		key := new(models.Key)
		FillUpForm(key, newFormGroup, keyFormCallback.playground)
		keyFormCallback.playground.formStage.Commit()
	}

	fillUpTree(keyFormCallback.playground)
}
func __gong__New__PieFormCallback(
	pie *models.Pie,
	playground *Playground,
) (pieFormCallback *PieFormCallback) {
	pieFormCallback = new(PieFormCallback)
	pieFormCallback.playground = playground
	pieFormCallback.pie = pie

	pieFormCallback.CreationMode = (pie == nil)

	return
}

type PieFormCallback struct {
	pie *models.Pie

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (pieFormCallback *PieFormCallback) OnSave() {

	log.Println("PieFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pieFormCallback.playground.formStage.Checkout()

	if pieFormCallback.pie == nil {
		pieFormCallback.pie = new(models.Pie).Stage(pieFormCallback.playground.stageOfInterest)
	}
	pie_ := pieFormCallback.pie
	_ = pie_

	// get the formGroup
	formGroup := pieFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(pie_.Name), formDiv)
		case "X":
			FormDivSelectFieldToField(&(pie_.X), pieFormCallback.playground.stageOfInterest, formDiv)
		case "Y":
			FormDivSelectFieldToField(&(pie_.Y), pieFormCallback.playground.stageOfInterest, formDiv)
		case "Width":
			FormDivBasicFieldToField(&(pie_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(pie_.Heigth), formDiv)
		case "Margin":
			FormDivBasicFieldToField(&(pie_.Margin), formDiv)
		}
	}

	pieFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Pie](
		pieFormCallback.playground,
	)
	pieFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if pieFormCallback.CreationMode {
		pieFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__PieFormCallback(
				nil,
				pieFormCallback.playground,
			),
		}).Stage(pieFormCallback.playground.formStage)
		pie := new(models.Pie)
		FillUpForm(pie, newFormGroup, pieFormCallback.playground)
		pieFormCallback.playground.formStage.Commit()
	}

	fillUpTree(pieFormCallback.playground)
}
func __gong__New__ScatterFormCallback(
	scatter *models.Scatter,
	playground *Playground,
) (scatterFormCallback *ScatterFormCallback) {
	scatterFormCallback = new(ScatterFormCallback)
	scatterFormCallback.playground = playground
	scatterFormCallback.scatter = scatter

	scatterFormCallback.CreationMode = (scatter == nil)

	return
}

type ScatterFormCallback struct {
	scatter *models.Scatter

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (scatterFormCallback *ScatterFormCallback) OnSave() {

	log.Println("ScatterFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	scatterFormCallback.playground.formStage.Checkout()

	if scatterFormCallback.scatter == nil {
		scatterFormCallback.scatter = new(models.Scatter).Stage(scatterFormCallback.playground.stageOfInterest)
	}
	scatter_ := scatterFormCallback.scatter
	_ = scatter_

	// get the formGroup
	formGroup := scatterFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(scatter_.Name), formDiv)
		case "X":
			FormDivSelectFieldToField(&(scatter_.X), scatterFormCallback.playground.stageOfInterest, formDiv)
		case "Y":
			FormDivSelectFieldToField(&(scatter_.Y), scatterFormCallback.playground.stageOfInterest, formDiv)
		case "Text":
			FormDivSelectFieldToField(&(scatter_.Text), scatterFormCallback.playground.stageOfInterest, formDiv)
		case "Width":
			FormDivBasicFieldToField(&(scatter_.Width), formDiv)
		case "Heigth":
			FormDivBasicFieldToField(&(scatter_.Heigth), formDiv)
		case "Margin":
			FormDivBasicFieldToField(&(scatter_.Margin), formDiv)
		}
	}

	scatterFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Scatter](
		scatterFormCallback.playground,
	)
	scatterFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if scatterFormCallback.CreationMode {
		scatterFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ScatterFormCallback(
				nil,
				scatterFormCallback.playground,
			),
		}).Stage(scatterFormCallback.playground.formStage)
		scatter := new(models.Scatter)
		FillUpForm(scatter, newFormGroup, scatterFormCallback.playground)
		scatterFormCallback.playground.formStage.Commit()
	}

	fillUpTree(scatterFormCallback.playground)
}
func __gong__New__SerieFormCallback(
	serie *models.Serie,
	playground *Playground,
) (serieFormCallback *SerieFormCallback) {
	serieFormCallback = new(SerieFormCallback)
	serieFormCallback.playground = playground
	serieFormCallback.serie = serie

	serieFormCallback.CreationMode = (serie == nil)

	return
}

type SerieFormCallback struct {
	serie *models.Serie

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (serieFormCallback *SerieFormCallback) OnSave() {

	log.Println("SerieFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	serieFormCallback.playground.formStage.Checkout()

	if serieFormCallback.serie == nil {
		serieFormCallback.serie = new(models.Serie).Stage(serieFormCallback.playground.stageOfInterest)
	}
	serie_ := serieFormCallback.serie
	_ = serie_

	// get the formGroup
	formGroup := serieFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(serie_.Name), formDiv)
		case "Key":
			FormDivSelectFieldToField(&(serie_.Key), serieFormCallback.playground.stageOfInterest, formDiv)
		case "Bar:Set":
			// we need to retrieve the field owner before the change
			var pastBarOwner *models.Bar
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Bar"
			rf.Fieldname = "Set"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				serieFormCallback.playground.stageOfInterest,
				serieFormCallback.playground.backRepoOfInterest,
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
				for _bar := range *models.GetGongstructInstancesSet[models.Bar](serieFormCallback.playground.stageOfInterest) {

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
				serieFormCallback.playground.stageOfInterest,
				serieFormCallback.playground.backRepoOfInterest,
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
				for _pie := range *models.GetGongstructInstancesSet[models.Pie](serieFormCallback.playground.stageOfInterest) {

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
				serieFormCallback.playground.stageOfInterest,
				serieFormCallback.playground.backRepoOfInterest,
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
				for _scatter := range *models.GetGongstructInstancesSet[models.Scatter](serieFormCallback.playground.stageOfInterest) {

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

	serieFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Serie](
		serieFormCallback.playground,
	)
	serieFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if serieFormCallback.CreationMode {
		serieFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__SerieFormCallback(
				nil,
				serieFormCallback.playground,
			),
		}).Stage(serieFormCallback.playground.formStage)
		serie := new(models.Serie)
		FillUpForm(serie, newFormGroup, serieFormCallback.playground)
		serieFormCallback.playground.formStage.Commit()
	}

	fillUpTree(serieFormCallback.playground)
}
func __gong__New__ValueFormCallback(
	value *models.Value,
	playground *Playground,
) (valueFormCallback *ValueFormCallback) {
	valueFormCallback = new(ValueFormCallback)
	valueFormCallback.playground = playground
	valueFormCallback.value = value

	valueFormCallback.CreationMode = (value == nil)

	return
}

type ValueFormCallback struct {
	value *models.Value

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (valueFormCallback *ValueFormCallback) OnSave() {

	log.Println("ValueFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	valueFormCallback.playground.formStage.Checkout()

	if valueFormCallback.value == nil {
		valueFormCallback.value = new(models.Value).Stage(valueFormCallback.playground.stageOfInterest)
	}
	value_ := valueFormCallback.value
	_ = value_

	// get the formGroup
	formGroup := valueFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
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
				valueFormCallback.playground.stageOfInterest,
				valueFormCallback.playground.backRepoOfInterest,
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
				for _serie := range *models.GetGongstructInstancesSet[models.Serie](valueFormCallback.playground.stageOfInterest) {

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

	valueFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Value](
		valueFormCallback.playground,
	)
	valueFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if valueFormCallback.CreationMode {
		valueFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ValueFormCallback(
				nil,
				valueFormCallback.playground,
			),
		}).Stage(valueFormCallback.playground.formStage)
		value := new(models.Value)
		FillUpForm(value, newFormGroup, valueFormCallback.playground)
		valueFormCallback.playground.formStage.Commit()
	}

	fillUpTree(valueFormCallback.playground)
}
