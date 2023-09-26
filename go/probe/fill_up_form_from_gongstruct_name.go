// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongd3/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Bar":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Bar Form",
			OnSave: __gong__New__BarFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		bar := new(models.Bar)
		FillUpForm(bar, formGroup, playground)
	case "Key":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Key Form",
			OnSave: __gong__New__KeyFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		key := new(models.Key)
		FillUpForm(key, formGroup, playground)
	case "Pie":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Pie Form",
			OnSave: __gong__New__PieFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		pie := new(models.Pie)
		FillUpForm(pie, formGroup, playground)
	case "Scatter":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Scatter Form",
			OnSave: __gong__New__ScatterFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		scatter := new(models.Scatter)
		FillUpForm(scatter, formGroup, playground)
	case "Serie":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Serie Form",
			OnSave: __gong__New__SerieFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		serie := new(models.Serie)
		FillUpForm(serie, formGroup, playground)
	case "Value":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Value Form",
			OnSave: __gong__New__ValueFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		value := new(models.Value)
		FillUpForm(value, formGroup, playground)
	}
	formStage.Commit()
}
